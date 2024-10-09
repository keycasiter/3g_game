package jym

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/jym"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 账号查找逻辑
type AccountSearchContext struct {
	ctx context.Context

	req   *jym.AccountSearchRequest
	resp  *jym.AccountSearchResponse
	err   error
	funcs []func()

	//****中间变量****
	//查询到的所有商品列表信息
	GoodsInfoList []vo.GetSgzGameZoneItemListRespResultGoodsInfo
	//查询到的所有商品详情信息
	GoodsInfoItemMap map[string]*vo.AccountItemInfo
}

func NewAccountSearchContext(ctx context.Context, req *jym.AccountSearchRequest, resp *jym.AccountSearchResponse) *AccountSearchContext {
	runCtx := &AccountSearchContext{
		ctx:              ctx,
		req:              req,
		resp:             resp,
		GoodsInfoList:    make([]vo.GetSgzGameZoneItemListRespResultGoodsInfo, 0),
		GoodsInfoItemMap: make(map[string]*vo.AccountItemInfo, 0),
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
	for i := 0; i < int(runCtx.req.PageNum); i++ {
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

		//当前页面不足15个，不需要再翻页了，交易猫默认一页15条
		if len(resp.Result.GoodsList) < 15 {
			hlog.CtxInfof(runCtx.ctx, "总数量：%d,翻页查询第%d页，本页商品数量：%d，不需要继续翻页", resp.Result.TotalCnt, i+1, len(resp.Result.GoodsList))
			break
		}
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
		Keyword:          runCtx.req.Keyword,
		Page:             pageNo,
	}
	//指定特技等
	extConditions := &vo.ExtConditions{}
	if len(runCtx.req.MustSpecialTechList) > 0 {
		extConditions.EquipSkill = runCtx.buildSpecialTech()
	}
	//指定英雄
	if len(runCtx.req.DefiniteHeroList) > 0 {
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
	//指定阵容
	if len(runCtx.req.LineUpList) > 0 {
		lineUp := ""
		for i, team := range runCtx.req.LineUpList {
			lineUp += team
			if i < len(lineUp) {
				lineUp += ","
			}
		}
		extConditions.LineUp = lineUp
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

func (runCtx *AccountSearchContext) buildHeros() string {
	heroIds := ""
	for i, heroId := range runCtx.req.DefiniteHeroList {
		heroIds += heroId
		if i < len(runCtx.req.DefiniteHeroList) {
			heroIds += ","
		}
	}
	return heroIds
}

func (runCtx *AccountSearchContext) buildSpecialTech() string {
	skillIds := ""
	for i, skillId := range runCtx.req.DefiniteSkillList {
		skillIds += skillId
		if i < len(runCtx.req.DefiniteHeroList) {
			skillIds += ","
		}
	}
	return skillIds
}

func (runCtx *AccountSearchContext) searchAccountDetail() {
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

// 过滤商品结果
func (runCtx *AccountSearchContext) filterAccount() {
	//整理指定武将
	definiteHeroMap := make(map[string]bool, 0)
	for _, heroId := range runCtx.req.DefiniteHeroList {
		definiteHeroMap[heroId] = true
	}
	//整理账号全部武将

	//指定武将是否要求觉醒
	if runCtx.req.IsDefiniteHeroMustAwake {
		newHolder := make(map[string]*vo.AccountItemInfo, 0)
		//遍历账号
		for goodsItemUrl, accountItemInfo := range runCtx.GoodsInfoItemMap {
			//遍历账号所有武将
			isMatch := true
			for _, heroInfo := range accountItemInfo.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros {
				//指定武将开三兵书
				if _, ok := definiteHeroMap[cast.ToString(heroInfo.HeroId)]; ok {
					if !heroInfo.IsAwake {
						hlog.CtxInfof(runCtx.ctx, "商品：%s，武将未觉醒：%s ,跳过", goodsItemUrl, heroInfo.Name)
						isMatch = false
						break
					}
				}
			}
			//符合条件存储
			if isMatch {
				newHolder[goodsItemUrl] = accountItemInfo
			}
		}
		runCtx.GoodsInfoItemMap = newHolder
	}
	//指定武将是否开三兵书
	if runCtx.req.IsDefiniteHeroMustTalent3 {
		newHolder := make(map[string]*vo.AccountItemInfo, 0)
		//遍历账号
		for goodsItemUrl, accountItemInfo := range runCtx.GoodsInfoItemMap {
			isMatch := true
			//遍历账号所有武将
			for _, heroInfo := range accountItemInfo.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros {
				//指定武将开三兵书
				if _, ok := definiteHeroMap[cast.ToString(heroInfo.HeroId)]; ok {
					if !heroInfo.IsUnlockTalent3 {
						hlog.CtxInfof(runCtx.ctx, "商品：%s，武将三兵书未开：%s ,跳过", goodsItemUrl, heroInfo.Name)
						isMatch = false
						break
					}
				}
			}
			//符合条件存储
			if isMatch {
				newHolder[goodsItemUrl] = accountItemInfo
			}
		}
		runCtx.GoodsInfoItemMap = newHolder
	}

	//指定战法
	if len(runCtx.req.MustTacticList) > 0 {
		//整理战法map
		tacticMap := make(map[string]bool, 0)
		for _, tacticName := range runCtx.req.MustTacticList {
			tacticMap[tacticName] = true
		}
		//符合条件的账号
		newHolder := make(map[string]*vo.AccountItemInfo, 0)
		//遍历账号
		for goodsItemUrl, accountItemInfo := range runCtx.GoodsInfoItemMap {
			//当前账号所有战法
			currentAccountAllTacticMap := make(map[string]bool, 0)
			//遍历账号所有战法
			for _, skillInfo := range accountItemInfo.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Skills {
				if strings.Trim(skillInfo.Name, " ") == "" {
					continue
				}
				currentAccountAllTacticMap[skillInfo.Name] = true
			}

			//匹配指定战法是否满足条件
			isMatch := true
			for tacticName, _ := range tacticMap {
				if _, ok := currentAccountAllTacticMap[tacticName]; !ok {
					//不满足要求，直接跳过
					hlog.CtxInfof(runCtx.ctx, "商品：%s, 区服:%s , 标价:%.2f , 战法不存在：%s ,跳过", goodsItemUrl,
						accountItemInfo.ApiData.ItemBaseInfo.ServerName,
						accountItemInfo.ApiData.ItemBaseInfo.SellPrice, tacticName)
					isMatch = false
					break
				}
			}

			if isMatch {
				//符合条件存储
				newHolder[goodsItemUrl] = accountItemInfo
			}
		}
		runCtx.GoodsInfoItemMap = newHolder
	}

	//指定特技
	if len(runCtx.req.MustSpecialTechList) > 0 {
		//整理特技map
		specialTechMap := make(map[string]bool, 0)
		for _, techName := range runCtx.req.MustSpecialTechList {
			specialTechMap[techName] = true
		}
		//符合条件的账号
		newHolder := make(map[string]*vo.AccountItemInfo, 0)
		//遍历账号
		for goodsItemUrl, accountItemInfo := range runCtx.GoodsInfoItemMap {
			//当前账号所有特技
			currentAccountAllTechMap := make(map[string]bool, 0)
			//遍历账号所有特技
			for _, equipment := range accountItemInfo.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Storage.Equipments {
				//只看5星装备即可,没特技跳过
				if equipment.Star != 5 || strings.Trim(equipment.SkillDesc, " ") == "" {
					continue
				}
				//指定特技
				for _, skillName := range equipment.SkillDescList {
					currentAccountAllTechMap[skillName] = true
				}
			}

			//匹配指定特技是否满足条件
			isMatch := true
			for techName, _ := range specialTechMap {
				if _, ok := currentAccountAllTechMap[techName]; !ok {
					//不满足特技要求，直接跳过
					hlog.CtxInfof(runCtx.ctx, "商品：%s，特级不存在：%s ,跳过", goodsItemUrl, techName)
					isMatch = false
					break
				}
			}

			if isMatch {
				//符合条件存储
				newHolder[goodsItemUrl] = accountItemInfo
			}
		}
		runCtx.GoodsInfoItemMap = newHolder
	}
}

func (runCtx *AccountSearchContext) buildResp() {
	list := make([]*jym.ApiData, 0)
	for _, info := range runCtx.GoodsInfoItemMap {
		vo := &jym.ApiData{}
		copier.Copy(vo, info.ApiData)
		list = append(list, vo)
	}
	runCtx.resp.ApiDatas = list
}
func (runCtx *AccountSearchContext) buildRespPrint() {
	//总结
	fmt.Println("######################## 总结 #########################")
	fmt.Printf("符合条件数量:%d\n", len(runCtx.GoodsInfoItemMap))
	fmt.Println("######################## 商品链接 #########################")
	for goodsDetailUrl, goodsItemInfo := range runCtx.GoodsInfoItemMap {
		fmt.Printf("%s , 标价:%.2f\n", goodsDetailUrl, goodsItemInfo.ApiData.ItemBaseInfo.SellPrice)
	}
	//打印结果
	for detailUrl, detailItem := range runCtx.GoodsInfoItemMap {
		fmt.Println("#################################################")
		//fmt.Printf("商品ID:%s\n", detailItem.ApiData.ItemBaseInfo.ItemId)
		fmt.Printf("【商品链接】%s\n", detailUrl)
		fmt.Printf("【标题】%s\n", detailItem.ApiData.ItemBaseInfo.Title)
		fmt.Printf("【卖家标价】%.2f\n", detailItem.ApiData.ItemBaseInfo.SellPrice)
		fmt.Printf("【区服】%s\n", detailItem.ApiData.ItemBaseInfo.ServerName)
		fmt.Printf("【收藏人数】%d\n", detailItem.ApiData.ItemQuality.FavoriteNum)
		fmt.Printf("【卖点】%s\n", util.ToJsonString(runCtx.ctx, detailItem.ApiData.SellPointTags)+" "+util.ToJsonString(runCtx.ctx, detailItem.ApiData.SecondSellPointTags))

		fmt.Printf("\n【检测如下战法均存在】\n")
		tacticNames := ""
		for _, tacticName := range runCtx.req.MustTacticList {
			tacticNames += tacticName + ","
		}
		fmt.Printf(tacticNames)

		fmt.Printf("\n【检测如下指定武将均存在】\n")
		heroNames := ""
		for _, heroId := range runCtx.req.DefiniteHeroList {
			heroNames += consts.HeroMap[heroId] + ","
		}
		fmt.Printf(heroNames)

		fmt.Printf("\n【特技情况】\n")
		techNames := ""
		for _, equipment := range detailItem.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Storage.Equipments {
			if equipment.Star == 5 && strings.Trim(equipment.SkillDesc, " ") != "" {
				for _, techName := range equipment.SkillDescList {
					techNames += techName + " "
				}
			}
		}
		fmt.Println(techNames)

		fmt.Printf("【账号武将情况】\n")
		//按红度排序
		sort.Slice(detailItem.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros, func(i, j int) bool {
			if detailItem.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros[i].Stage >
				detailItem.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros[j].Stage {
				return true
			}
			return false
		})
		for _, hero := range detailItem.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros {
			fmt.Printf("%s ,红度:%d,觉醒：%v , 三兵书:%v\n", hero.Name, hero.Stage, util.BoolToString(hero.IsAwake), util.BoolToString(hero.IsUnlockTalent3))
		}
	}
}
