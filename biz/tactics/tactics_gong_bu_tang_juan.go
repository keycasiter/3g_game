package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 功不唐捐
// 战斗中，自身自带主动战法伤害提升30%并获得30%攻心，战斗第2回合起，自身施加的负面状态有70%概率（受智力影响）不可被驱散
type GongBuTangJuanTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a GongBuTangJuanTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a GongBuTangJuanTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	//战斗中，自身自带主动战法伤害提升30%并获得30%攻心，
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_TacticsActiveDamageImprove, &vo.EffectHolderParams{
		TriggerRate:    0.3,
		EffectRate:     0.3,
		FromTactic:     a.Id(),
		ProduceGeneral: currentGeneral,
	})
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_AttackHeart, &vo.EffectHolderParams{
		TriggerRate:    0.3,
		EffectRate:     0.3,
		FromTactic:     a.Id(),
		ProduceGeneral: currentGeneral,
	})

	//战斗第2回合起，自身施加的负面状态有70%概率（受智力影响）不可被驱散
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_DebuffEffect, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if triggerRound >= consts.Battle_Round_Second {
			triggerRate := 0.7 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
			if util.GenerateRate(triggerRate) {
				params.EffectHolderParams.IsAvoidDispel = true
			}
		}

		return triggerResp
	})
}

func (a GongBuTangJuanTactic) Id() consts.TacticId {
	return consts.GongBuTangJuan
}

func (a GongBuTangJuanTactic) Name() string {
	return "功不唐捐"
}

func (a GongBuTangJuanTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a GongBuTangJuanTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a GongBuTangJuanTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a GongBuTangJuanTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (a GongBuTangJuanTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a GongBuTangJuanTactic) Execute() {
}

func (a GongBuTangJuanTactic) IsTriggerPrepare() bool {
	return false
}

func (a GongBuTangJuanTactic) SetTriggerPrepare(triggerPrepare bool) {
}
