package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

//奇计良谋
//战斗前3回合，使敌军武力最高的武将造成兵刃伤害降低28%（受速度影响），
//使敌军智力最高的武将造成的谋略伤害降低28%（受速度影响）
type CleverPlanAndCleverPlanTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CleverPlanAndCleverPlanTactic) IsTriggerPrepare() bool {
	return false
}

func (c CleverPlanAndCleverPlanTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c CleverPlanAndCleverPlanTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)

	//战斗前3回合，使敌军武力最高的武将造成兵刃伤害降低28%（受速度影响）
	mostForceEnemyGeneral := util.GetMostForceEnemyGeneral(c.tacticsParams)
	forceRate := 0.28 + currentGeneral.BaseInfo.AbilityAttr.SpeedBase/100/100
	util.DebuffEffectWrapSet(ctx, mostForceEnemyGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, forceRate)
	hlog.CtxInfof(ctx, "[%s]造成兵刃伤害降低%.2f%%",
		mostForceEnemyGeneral.BaseInfo.Name,
		forceRate*100,
	)
	util.TacticsTriggerWrapRegister(mostForceEnemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		if triggerRound == consts.Battle_Round_Third {
			if util.DebuffEffectWrapRemove(ctx, triggerGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce) {
				hlog.CtxInfof(ctx, "[%s]造成兵刃伤害提高%.2f%%",
					triggerGeneral.BaseInfo.Name,
					forceRate*100,
				)
			}
		}

		return triggerResp
	})
	//使敌军智力最高的武将造成的谋略伤害降低28%（受速度影响）
	mostIntelligenceEnemyGeneral := util.GetMostIntelligenceEnemyGeneral(c.tacticsParams)
	intelligenceRate := 0.28 + currentGeneral.BaseInfo.AbilityAttr.SpeedBase/100/100
	util.DebuffEffectWrapSet(ctx, mostIntelligenceEnemyGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce, intelligenceRate)
	hlog.CtxInfof(ctx, "[%s]造成谋略伤害降低%.2f%%",
		mostIntelligenceEnemyGeneral.BaseInfo.Name,
		intelligenceRate*100,
	)
	util.TacticsTriggerWrapRegister(mostIntelligenceEnemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		if triggerRound == consts.Battle_Round_Third {
			if util.DebuffEffectWrapRemove(ctx, triggerGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce) {
				hlog.CtxInfof(ctx, "[%s]造成谋略伤害提高%.2f%%",
					triggerGeneral.BaseInfo.Name,
					forceRate*100,
				)
			}
		}

		return triggerResp
	})
}

func (c CleverPlanAndCleverPlanTactic) Id() consts.TacticId {
	return consts.CleverPlanAndCleverPlan
}

func (c CleverPlanAndCleverPlanTactic) Name() string {
	return "奇计良谋"
}

func (c CleverPlanAndCleverPlanTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (c CleverPlanAndCleverPlanTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CleverPlanAndCleverPlanTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CleverPlanAndCleverPlanTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (c CleverPlanAndCleverPlanTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CleverPlanAndCleverPlanTactic) Execute() {

}
