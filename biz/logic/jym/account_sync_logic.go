package jym

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
	"github.com/thoas/go-funk"
	"regexp"
	"strings"
	"time"
)

// 账号同步逻辑
type AccountSyncContext struct {
	ctx context.Context

	req   *vo.AccountSearchReq
	err   error
	funcs []func()

	//****中间变量****
	//查询到的所有商品列表信息
	GoodsInfoList []vo.GetSgzGameZoneItemListRespResultGoodsInfo
	//查询到的所有商品详情信息
	GoodsInfoItemMap map[string]*vo.AccountItemInfo
}

func NewAccountSyncContext(ctx context.Context, req *vo.AccountSearchReq) *AccountSyncContext {
	runCtx := &AccountSyncContext{
		ctx:              ctx,
		req:              req,
		GoodsInfoList:    make([]vo.GetSgzGameZoneItemListRespResultGoodsInfo, 0),
		GoodsInfoItemMap: make(map[string]*vo.AccountItemInfo, 0),
	}
	runCtx.funcs = []func(){
		//1.查询符合条件的账号列表
		runCtx.searchAccountList,
	}
	return runCtx
}

func (runCtx *AccountSyncContext) Process() error {
	for _, f := range runCtx.funcs {
		f()
		if runCtx.err != nil {
			hlog.CtxErrorf(runCtx.ctx, runCtx.err.Error())
			return runCtx.err
		}
	}
	return nil
}

//计算一共多少页
func (runCtx *AccountSyncContext) getTotalPage() int64 {
	req := runCtx.buildGetSgzGameZoneItemListReq(cast.ToInt64(1))

	httpResp, err := util.HttpGet(runCtx.ctx, consts.Url_GetSgzGameZoneItemList, nil, util.StructToMap(req))
	if err != nil {
		hlog.CtxErrorf(runCtx.ctx, "url:%s , HttpGet err:%v", consts.Url_GetSgzGameZoneItemList, err)
		runCtx.err = err
		return 0
	}
	resp := &vo.GetSgzGameZoneItemListResp{}
	err = json.Unmarshal([]byte(httpResp), resp)
	if err != nil {
		//一般是限流导致，重试
		hlog.CtxErrorf(runCtx.ctx, "Unmarshal err:%v", err)
		return 0
	}
	if !resp.Success {
		hlog.CtxErrorf(runCtx.ctx, "resp result is failed")
		runCtx.err = errors.New(fmt.Sprintf("url:%s , resp result is failed", consts.Url_GetSgzGameZoneItemList))
		return 0
	}
	runCtx.GoodsInfoList = append(runCtx.GoodsInfoList, resp.Result.GoodsList...)
	return resp.Result.TotalCnt
}
func (runCtx *AccountSyncContext) searchAccountList() {
	//页数统计
	totalSize := runCtx.getTotalPage()
	totalPageSize := int(totalSize / 30)
	hlog.CtxInfof(runCtx.ctx, "总数：%d ,按每页30个，页面数量估算：%d", totalSize, totalPageSize)

	for i := 0; i < totalPageSize; i++ {

	}

	//翻页查询
	for i := 0; i < totalPageSize+10; i++ {
		pageNo := i + 1
		hlog.CtxInfof(runCtx.ctx, "翻页查询第%d页", pageNo)
		req := runCtx.buildGetSgzGameZoneItemListReq(cast.ToInt64(pageNo))

		httpResp, err := util.HttpGet(runCtx.ctx, consts.Url_GetSgzGameZoneItemList, nil, util.StructToMap(req))
		if err != nil {
			//重试
			hlog.CtxErrorf(runCtx.ctx, "url:%s , HttpGet err:%v", consts.Url_GetSgzGameZoneItemList, err)
			i--
			continue
		}
		resp := &vo.GetSgzGameZoneItemListResp{}
		err = json.Unmarshal([]byte(httpResp), resp)
		if err != nil {
			//一般是限流导致，重试
			hlog.CtxErrorf(runCtx.ctx, "Unmarshal err:%v", err)
			i--
			continue
		}
		if !resp.Success {
			hlog.CtxErrorf(runCtx.ctx, "resp result is failed")
			runCtx.err = errors.New(fmt.Sprintf("url:%s , resp result is failed", consts.Url_GetSgzGameZoneItemList))
			return
		}
		runCtx.GoodsInfoList = append(runCtx.GoodsInfoList, resp.Result.GoodsList...)
		hlog.CtxInfof(runCtx.ctx, "翻页查询第%d页 , 商品总数：%d", pageNo, resp.Result.TotalCnt)
		//当前页面不足15个，不需要再翻页了，交易猫默认一页15条
		if len(resp.Result.GoodsList) == 0 {
			hlog.CtxInfof(runCtx.ctx, "翻页查询第%d页，本页商品数量：%d，不需要继续翻页", i+1, len(resp.Result.GoodsList))
			break
		}

		//2.获取账号详情信息
		runCtx.searchAccountDetail()
		//3.保存数据
		runCtx.saveAccountInfo()
	}
}

func (runCtx *AccountSyncContext) buildGetSgzGameZoneItemListReq(pageNo int64) *vo.GetSgzGameZoneItemListReq {
	req := &vo.GetSgzGameZoneItemListReq{
		GameId:           consts.GameId,
		Fcid:             consts.FcId,
		OsId:             consts.OsId,
		Cid:              consts.CId,
		PlatformId:       consts.PlatformId,
		Sort:             consts.Sort,
		StdCatId:         consts.StdCatId,
		JymCatId:         consts.JymCatId,
		FilterLowQuality: consts.FilterLowQuality,
		Keyword:          runCtx.req.Keyword,
		Page:             pageNo,
	}
	//指定特技等
	extConditions := &vo.ExtConditions{}
	if len(runCtx.req.MustSpecialTech) > 0 {
		extConditions.EquipSkill = runCtx.buildSpecialTech()
	}
	//指定英雄
	if len(runCtx.req.DefiniteHeros) > 0 {
		extConditions.Hero = runCtx.buildHeros()
	}
	//指定红度
	if runCtx.req.DefiniteStage != "" {
		extConditions.Stage = runCtx.req.DefiniteStage
	}
	//指定总红度
	if runCtx.req.DefiniteTotalStage != "" {
		extConditions.StageSum = runCtx.req.DefiniteTotalStage
	}
	//可跨服，公示
	if runCtx.req.CrossServerAndPublic {
		extConditions.CrossServerAndPublic = "1,2"
	}
	//五星武将数量
	if runCtx.req.FiveStarHeroNum != "" {
		extConditions.FiveStarHeroNum = runCtx.req.FiveStarHeroNum
	}
	//S战法数量
	if runCtx.req.SskillNum != "" {
		extConditions.SskillNum = runCtx.req.SskillNum
	}

	req.ExtConditions = util.ToJsonString(runCtx.ctx, extConditions)
	//价格范围
	if len(strings.Trim(runCtx.req.PriceRange, " ")) > 0 {
		req.PriceRange = runCtx.req.PriceRange
	} else {
		req.PriceRange = util.ToJsonString(runCtx.ctx, []string{})
	}
	return req
}

func (runCtx *AccountSyncContext) buildHeros() string {
	heroIds := ""
	for i, heroId := range runCtx.req.DefiniteHeros {
		heroIds += heroId
		if i < len(runCtx.req.DefiniteHeros) {
			heroIds += ","
		}
	}
	return heroIds
}

func (runCtx *AccountSyncContext) buildSpecialTech() string {
	skillIds := ""
	for i, skillId := range runCtx.req.DefiniteSkill {
		skillIds += skillId
		if i < len(runCtx.req.DefiniteHeros) {
			skillIds += ","
		}
	}
	return skillIds
}

func (runCtx *AccountSyncContext) searchAccountDetail() {
	for i, goodsItem := range runCtx.GoodsInfoList {
		hlog.CtxInfof(runCtx.ctx, "商品查询进度：%d/%d", i+1, len(runCtx.GoodsInfoList))
		//防止被限流
		time.Sleep(500 * time.Millisecond)

		err := retry.Do(func() error {
			httpRes, err := util.HttpGet(runCtx.ctx, goodsItem.DetailUrl, nil, nil)
			if err != nil {
				hlog.CtxErrorf(runCtx.ctx, "url:%s , HttpGet err:%v", goodsItem.DetailUrl, err)
				return err
			}
			//fmt.Printf("httpRes:\n%s", httpRes)
			reg, err := regexp.Compile("window.__INITIAL_STATE__ =.*\\n")
			if err != nil {
				hlog.CtxErrorf(runCtx.ctx, "url:%s , regexp err:%v", goodsItem.DetailUrl, err)
				return err
			}
			jsonStr := reg.FindString(httpRes)
			jsonStr = strings.ReplaceAll(jsonStr, "window.__INITIAL_STATE__ =", "")
			//fmt.Printf("jsonStr:\n%s", jsonStr)

			data := &vo.AccountItemInfo{}
			err = json.Unmarshal([]byte(jsonStr), data)
			if err != nil {
				//hlog.CtxErrorf(runCtx.ctx, "json unmarshal err:%v ,\njsonStr:%s\nurl:%s\n:resp:%s", err, jsonStr, goodsItem.DetailUrl, httpRes)
				hlog.CtxErrorf(runCtx.ctx, "json unmarshal err:%v ,\njsonStr:%s", err, jsonStr)
				return err
			}

			//整理商品详情结果
			runCtx.GoodsInfoItemMap[goodsItem.DetailUrl] = data
			return nil
		}, retry.Attempts(3), retry.Delay(1*time.Second))
		if err != nil {
			hlog.CtxErrorf(runCtx.ctx, "url:%s, retry err:%v", goodsItem.DetailUrl, err)
		}
	}
}

func (runCtx *AccountSyncContext) saveAccountInfo() {
	goods := make([]*po.JymGoods, 0)
	for url, itemInfo := range runCtx.GoodsInfoItemMap {
		goods = append(goods, &po.JymGoods{
			GoodsUrl:    url,
			GoodsDetail: util.ToJsonString(runCtx.ctx, itemInfo),
		})
	}
	goodsNew := funk.Chunk(goods, 100)

	for _, goodsArr := range goodsNew.([][]*po.JymGoods) {
		err := mysql.NewJymGoods().BatchSaveJymGoods(runCtx.ctx, goodsArr)
		if err != nil {
			hlog.CtxErrorf(runCtx.ctx, "saveAccountInfo BatchSaveJymGoods err:%v", err)
			runCtx.err = err
			return
		}
	}
	runCtx.GoodsInfoList = make([]vo.GetSgzGameZoneItemListRespResultGoodsInfo, 0)
}
