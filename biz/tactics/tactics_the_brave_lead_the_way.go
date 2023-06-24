package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 勇者得前
// 普通攻击之后使自己获得一次抵御，可免疫伤害，并使下一个主动战法的伤害率提升80%
type TheBraveLeadTheWayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheBraveLeadTheWayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.45
	return t
}

func (t TheBraveLeadTheWayTactic) Prepare() {
}

func (t TheBraveLeadTheWayTactic) Id() consts.TacticId {
	return consts.TheBraveLeadTheWay
}

func (t TheBraveLeadTheWayTactic) Name() string {
	return "勇者得前"
}

func (t TheBraveLeadTheWayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheBraveLeadTheWayTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheBraveLeadTheWayTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheBraveLeadTheWayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (t TheBraveLeadTheWayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheBraveLeadTheWayTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//普通攻击之后使自己获得一次抵御，可免疫伤害，并使下一个主动战法的伤害率提升80%
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Defend, &vo.EffectHolderParams{
		EffectTimes:    1,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	})
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_TacticsActiveDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.8,
			EffectTimes:    1,
			FromTactic:     t.Id(),
			ProduceGeneral: triggerGeneral,
		}).IsSuccess {
			//注销效果
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_ActiveTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_TacticsActiveDamageImprove,
					TacticId:   t.Id(),
					CostTimes:  1,
				})

				return revokeResp
			})
		}

		return triggerResp
	})
}

func (t TheBraveLeadTheWayTactic) IsTriggerPrepare() bool {
	return false
}
