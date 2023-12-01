package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
	"regexp"
	"strings"
	"time"
)

// 账号查找逻辑
type AccountSearchContext struct {
	ctx context.Context

	req   *vo.GetSgzGameZoneItemListReq
	err   error
	funcs []func()

	//中间变量
	GoodsInfoList []vo.GetSgzGameZoneItemListRespResultGoodsInfo
}

func NewAccountSearchContext(ctx context.Context, req *vo.GetSgzGameZoneItemListReq) *AccountSearchContext {
	runCtx := &AccountSearchContext{
		ctx:           ctx,
		req:           req,
		GoodsInfoList: make([]vo.GetSgzGameZoneItemListRespResultGoodsInfo, 0),
	}
	runCtx.funcs = []func(){
		//1.查询符合条件的账号列表
		runCtx.searchAccountList,
		//2.获取账号详情信息
		runCtx.searchAccountDetail,
		//3.过滤不符合条件的账号
		runCtx.filterAccount,
		//4.整理账号信息
		runCtx.buildResp,
	}
	return runCtx
}

func (runCtx *AccountSearchContext) Process() error {
	for _, f := range runCtx.funcs {
		f()
		if runCtx.err != nil {
			hlog.CtxErrorf(runCtx.ctx, runCtx.err.Error())
			return runCtx.err
		}
	}
	return nil
}

func (runCtx *AccountSearchContext) searchAccountList() {
	//翻页查询
	for i := 0; i < 3; i++ {
		hlog.CtxInfof(runCtx.ctx, "翻页查询第%d页", i+1)
		req := runCtx.buildGetSgzGameZoneItemListReq(cast.ToInt64(i + 1))

		httpResp, err := util.HttpGet(runCtx.ctx, consts.Url_GetSgzGameZoneItemList, nil, util.StructToMap(req))
		if err != nil {
			hlog.CtxErrorf(runCtx.ctx, "url:%s , HttpGet err:%v", consts.Url_GetSgzGameZoneItemList, err)
			runCtx.err = err
			return
		}
		resp := &vo.GetSgzGameZoneItemListResp{}
		err = json.Unmarshal([]byte(httpResp), resp)
		if err != nil {
			hlog.CtxErrorf(runCtx.ctx, "Unmarshal err:%v", err)
			runCtx.err = err
			return
		}
		if !resp.Success {
			hlog.CtxErrorf(runCtx.ctx, "resp result is failed")
			runCtx.err = errors.New(fmt.Sprintf("url:%s , resp result is failed", consts.Url_GetSgzGameZoneItemList))
			return
		}
		runCtx.GoodsInfoList = append(runCtx.GoodsInfoList, resp.Result.GoodsList...)
	}
}

func (runCtx *AccountSearchContext) buildGetSgzGameZoneItemListReq(pageNo int64) *vo.GetSgzGameZoneItemListReq {
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
		Page:             pageNo,
	}
	if len(strings.Trim(runCtx.req.Keyword, " ")) > 0 {
		req.Keyword = runCtx.req.Keyword
	} else {
		req.Keyword = util.ToJsonString(runCtx.ctx, []string{})
	}
	if len(strings.Trim(runCtx.req.ExtConditions, " ")) > 0 {
		req.ExtConditions = runCtx.req.ExtConditions
	} else {
		req.ExtConditions = util.ToJsonString(runCtx.ctx, []string{})
	}
	if len(strings.Trim(runCtx.req.PriceRange, " ")) > 0 {
		req.PriceRange = runCtx.req.PriceRange
	} else {
		req.PriceRange = util.ToJsonString(runCtx.ctx, []string{})
	}
	return req
}

func (runCtx *AccountSearchContext) searchAccountDetail() {
	heroMap := make(map[int64]string, 0)

	for i, goodsItem := range runCtx.GoodsInfoList {
		hlog.CtxInfof(runCtx.ctx, "商品查询进度：%d/%d", i+1, len(runCtx.GoodsInfoList))
		//防止被限流
		time.Sleep(2 * time.Second)

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

			for _, hero := range data.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros {
				heroMap[hero.HeroId] = hero.Name
			}
			return nil
		}, retry.Attempts(3), retry.Delay(1*time.Second))
		if err != nil {
			hlog.CtxErrorf(runCtx.ctx, "url:%s, retry err:%v", goodsItem.DetailUrl, err)
		}
	}

	for i, s := range heroMap {
		fmt.Printf("%d=%s\n", i, s)
	}
}

func (runCtx *AccountSearchContext) filterAccount() {

}

func (runCtx *AccountSearchContext) buildResp() {

}
