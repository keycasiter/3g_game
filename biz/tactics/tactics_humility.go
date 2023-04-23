package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 谦让
// 使我军单体受到伤害降低22%（受智力影响），持续3回合
type HumilityTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HumilityTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.4
	return h
}

func (h HumilityTactic) Prepare() {
}

func (h HumilityTactic) Id() consts.TacticId {
	return consts.Humility
}

func (h HumilityTactic) Name() string {
	return "谦让"
}

func (h HumilityTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (h HumilityTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HumilityTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HumilityTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (h HumilityTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HumilityTactic) Execute() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral
	currentRound := h.tacticsParams.CurrentRound

	//使我军单体受到伤害降低22%（受智力影响），持续3回合
	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)

	//找到我军
	pairGeneral := util.GetPairOneGeneral(h.tacticsParams, currentGeneral)
	//注册效果
	if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Humility_Prepare, 1.0) {
		rate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 0.22
		pairGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce] += rate
		pairGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferStrategyDamageDeduce] += rate

		hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害降低了%.2f%%", pairGeneral.BaseInfo.Name,
			rate*100)
		hlog.CtxInfof(ctx, "[%s]受到的谋略伤害降低了%.2f%%", pairGeneral.BaseInfo.Name,
			rate*100)
		//注册消失效果
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeRound := params.CurrentRound
			revokeGeneral := params.CurrentGeneral

			if currentRound+3 == revokeRound {
				pairGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce] -= rate
				pairGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferStrategyDamageDeduce] -= rate
				util.BuffEffectWrapRemove(ctx, revokeGeneral, consts.BuffEffectType_Humility_Prepare)
				hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害提高了%.2f%%", pairGeneral.BaseInfo.Name,
					rate*100)
				hlog.CtxInfof(ctx, "[%s]受到的谋略伤害提高了%.2f%%", pairGeneral.BaseInfo.Name,
					rate*100)
			}

			return revokeResp
		})
	}
}

func (h HumilityTactic) IsTriggerPrepare() bool {
	return false
}
