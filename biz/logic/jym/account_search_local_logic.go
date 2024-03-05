package jym

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"strings"
)

// 账号本地查找逻辑
type AccountSearchLocalContext struct {
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

func NewAccountSearchLocalContext(ctx context.Context, req *vo.AccountSearchReq) *AccountSearchLocalContext {
	runCtx := &AccountSearchLocalContext{
		ctx:              ctx,
		req:              req,
		GoodsInfoList:    make([]vo.GetSgzGameZoneItemListRespResultGoodsInfo, 0),
		GoodsInfoItemMap: make(map[string]*vo.AccountItemInfo, 0),
	}
	runCtx.funcs = []func(){
		//1.查询符合条件的账号
		runCtx.searchAccountList,
		//2.整理账号信息
		runCtx.buildResp,
	}
	return runCtx
}

func (runCtx *AccountSearchLocalContext) Process() error {
	for _, f := range runCtx.funcs {
		f()
		if runCtx.err != nil {
			hlog.CtxErrorf(runCtx.ctx, runCtx.err.Error())
			return runCtx.err
		}
	}
	return nil
}

func (runCtx *AccountSearchLocalContext) searchAccountList() {
	mysql.
}

func (runCtx *AccountSearchLocalContext) buildGetSgzGameZoneItemListReq(pageNo int64) *vo.GetSgzGameZoneItemListReq {
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
	//指定阵容
	if len(runCtx.req.LineUp) > 0 {
		lineUp := ""
		for i, team := range runCtx.req.LineUp {
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

func (runCtx *AccountSearchLocalContext) buildHeros() string {
	heroIds := ""
	for i, heroId := range runCtx.req.DefiniteHeros {
		heroIds += heroId
		if i < len(runCtx.req.DefiniteHeros) {
			heroIds += ","
		}
	}
	return heroIds
}

func (runCtx *AccountSearchLocalContext) buildSpecialTech() string {
	skillIds := ""
	for i, skillId := range runCtx.req.DefiniteSkill {
		skillIds += skillId
		if i < len(runCtx.req.DefiniteHeros) {
			skillIds += ","
		}
	}
	return skillIds
}

func (runCtx *AccountSearchLocalContext) buildResp() {

}
