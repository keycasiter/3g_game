package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 闭月
// 选择一名敌军为自身分担25%（受智力影响）伤害，且当目标为敌军武力最高时进入混乱状态，
// 若为敌军智力最高时进入计穷状态，否则造成虚弱及禁疗状态，持续1回合
// 主动，75%
type ClosedMoonTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ClosedMoonTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 0.7
	return c
}

func (c ClosedMoonTactic) Prepare() {

}

func (c ClosedMoonTactic) Id() consts.TacticId {
	return consts.ClosedMoon
}

func (c ClosedMoonTactic) Name() string {
	return "闭月"
}

func (c ClosedMoonTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c ClosedMoonTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c ClosedMoonTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c ClosedMoonTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (c ClosedMoonTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c ClosedMoonTactic) Execute() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)

	// 选择一名敌军为自身分担25%（受智力影响）伤害，
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, c.tacticsParams)
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ShareResponsibilityFor, &vo.EffectHolderParams{
		EffectRate:                      0.25 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100,
		EffectRound:                     1,
		FromTactic:                      c.Id(),
		ShareResponsibilityForByGeneral: enemyGeneral,
		ProduceGeneral:                  currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_ShareResponsibilityFor,
				TacticId:   c.Id(),
			})
			return revokeResp
		})
	}
	// 且当目标为敌军武力最高时进入混乱状态，若为敌军智力最高时进入计穷状态，否则造成虚弱及禁疗状态，持续1回合
	debuff := consts.DebuffEffectType_ProhibitionTreatment
	if enemyGeneral.BaseInfo.Id == util.GetEnemyGeneralWhoIsHighestForce(currentGeneral, c.tacticsParams).BaseInfo.Id {
		//武力最高
		debuff = consts.DebuffEffectType_Chaos
	} else if enemyGeneral.BaseInfo.Id == util.GetEnemyGeneralWhoIsHighestIntelligence(currentGeneral, c.tacticsParams).BaseInfo.Id {
		//智力最高
		debuff = consts.DebuffEffectType_NoStrategy
	}

	if util.DebuffEffectWrapSet(ctx, enemyGeneral, debuff, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     c.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: debuff,
				TacticId:   c.Id(),
			})
			return revokeResp
		})
	}
}

func (c ClosedMoonTactic) IsTriggerPrepare() bool {
	return false
}
