package logic

import (
	"context"
	"encoding/json"
	"github.com/avast/retry-go"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
	"regexp"
	"strings"
	"time"
)

type AccountCheckContext struct {
	ctx context.Context

	//校验链接
	req *vo.AccountCheckReq
	//商品详情
	goodsDetail *vo.AccountItemInfo

	err   error
	funcs []func()
}

func NewAccountCheckLogic(ctx context.Context, req *vo.AccountCheckReq) *AccountCheckContext {
	runCtx := &AccountCheckContext{ctx: ctx, req: req}

	runCtx.funcs = []func(){
		//获取商品详情
		runCtx.getAccountDetail,
		//检测商品
		runCtx.checkAccountDetail,
	}
	return runCtx
}

func (runCtx *AccountCheckContext) Process() error {
	for _, f := range runCtx.funcs {
		f()
		if runCtx.err != nil {
			hlog.CtxErrorf(runCtx.ctx, "AccountCheckContext Process err:%v", runCtx.err)
			return runCtx.err
		}
	}
	return nil
}

func (runCtx *AccountCheckContext) getAccountDetail() {
	err := retry.Do(func() error {
		httpRes, err := util.HttpGet(runCtx.ctx, runCtx.req.CheckGoodsUrl, nil, nil)
		if err != nil {
			hlog.CtxErrorf(runCtx.ctx, "url:%s , HttpGet err:%v", runCtx.req.CheckGoodsUrl, err)
			return err
		}
		//fmt.Printf("httpRes:\n%s", httpRes)
		reg, err := regexp.Compile("window.__INITIAL_STATE__ =.*\\n")
		if err != nil {
			hlog.CtxErrorf(runCtx.ctx, "url:%s , regexp err:%v", runCtx.req.CheckGoodsUrl, err)
			return err
		}
		jsonStr := reg.FindString(httpRes)
		jsonStr = strings.ReplaceAll(jsonStr, "window.__INITIAL_STATE__ =", "")
		//fmt.Printf("jsonStr:\n%s", jsonStr)

		data := &vo.AccountItemInfo{}
		err = json.Unmarshal([]byte(jsonStr), data)
		if err != nil {
			//hlog.CtxErrorf(runCtx.ctx, "json unmarshal err:%v ,\njsonStr:%s\nurl:%s\n:resp:%s", err, jsonStr, runCtx.checkUrl, httpRes)
			hlog.CtxErrorf(runCtx.ctx, "json unmarshal err:%v ,\njsonStr:%s", err, jsonStr)
			return err
		}

		//整理商品详情结果
		runCtx.goodsDetail = data
		return nil
	}, retry.Attempts(3), retry.Delay(1*time.Second))
	if err != nil {
		hlog.CtxErrorf(runCtx.ctx, "url:%s, retry err:%v", runCtx.req.CheckGoodsUrl, err)
	}
}

func (runCtx *AccountCheckContext) checkAccountDetail() {
	//整理指定武将
	definiteHeroMap := make(map[string]bool, 0)
	for _, heroId := range runCtx.req.DefiniteHeros {
		definiteHeroMap[heroId] = true
	}
	//整理账号全部武将

	//指定武将是否要求觉醒
	if runCtx.req.IsDefiniteHeroMustAwake {
		//遍历账号所有武将
		for _, heroInfo := range runCtx.goodsDetail.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros {
			//指定武将开三兵书
			if _, ok := definiteHeroMap[cast.ToString(heroInfo.HeroId)]; ok {
				if !heroInfo.IsAwake {
					hlog.CtxInfof(runCtx.ctx, "武将未觉醒：%s ,检测失败", heroInfo.Name)
				}
			}
		}
	}
	//指定武将是否开三兵书
	if runCtx.req.IsDefiniteHeroMustTalent3 {
		//遍历账号所有武将
		for _, heroInfo := range runCtx.goodsDetail.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Heros {
			//指定武将开三兵书
			if _, ok := definiteHeroMap[cast.ToString(heroInfo.HeroId)]; ok {
				if !heroInfo.IsUnlockTalent3 {
					hlog.CtxInfof(runCtx.ctx, "武将三兵书未开：%s ,检测失败", heroInfo.Name)
				}
			}
		}
	}

	//指定战法
	if len(runCtx.req.MustTactic) > 0 {
		//整理战法map
		tacticMap := make(map[string]bool, 0)
		for _, tacticName := range runCtx.req.MustTactic {
			tacticMap[tacticName] = true
		}

		//当前账号所有战法
		currentAccountAllTacticMap := make(map[string]bool, 0)
		//遍历账号所有战法
		for _, skillInfo := range runCtx.goodsDetail.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Skills {
			if strings.Trim(skillInfo.Name, " ") == "" {
				continue
			}
			currentAccountAllTacticMap[skillInfo.Name] = true
		}

		//匹配指定战法是否满足条件
		for tacticName, _ := range tacticMap {
			if _, ok := currentAccountAllTacticMap[tacticName]; !ok {
				//不满足要求，直接跳过
				hlog.CtxInfof(runCtx.ctx, "战法不存在：%s ,检测失败", tacticName)
			}
		}
	}

	//指定特技
	if len(runCtx.req.MustSpecialTech) > 0 {
		//整理特技map
		specialTechMap := make(map[string]bool, 0)
		for _, techName := range runCtx.req.MustSpecialTech {
			specialTechMap[techName] = true
		}
		//当前账号所有特技
		currentAccountAllTechMap := make(map[string]bool, 0)
		//遍历账号所有特技
		for _, equipment := range runCtx.goodsDetail.ApiData.ItemLingxiRoleDetail.S3RoleCustomizeInfo.Storage.Equipments {
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
		for techName, _ := range specialTechMap {
			if _, ok := currentAccountAllTechMap[techName]; !ok {
				//不满足特技要求，直接跳过
				hlog.CtxInfof(runCtx.ctx, "特技不存在：%s ,检测失败", techName)
			}
		}
	}
}
