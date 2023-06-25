package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 暂避其锋
// 战斗开始后前3回合，使我军智力最高的武将受到兵刃伤害降低40%（受智力影响）
// 使我军武力最高的武将受到的谋略伤害降低40%（受智力影响）
type TakeRefugeFromEnemiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TakeRefugeFromEnemiesTactic) IsTriggerPrepare() bool {
	return false
}

func (t TakeRefugeFromEnemiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TakeRefugeFromEnemiesTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	//战斗开始后前3回合，使我军智力最高的武将受到兵刃伤害降低40%（受智力影响）
	mostIntelligenceGeneral := util.GetMostIntelligencePairGeneral(t.tacticsParams)
	weaponDamageDeduceRate := 0.4 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100

	if util.BuffEffectWrapSet(ctx, mostIntelligenceGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRate: weaponDamageDeduceRate,
		FromTactic: t.Id(),
	}).IsSuccess {
		hlog.CtxInfof(ctx, "[%s]受到的兵刃伤害降低了%.2f%%",
			mostIntelligenceGeneral.BaseInfo.Name,
			weaponDamageDeduceRate*100,
		)
		util.TacticsTriggerWrapRegister(mostIntelligenceGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerRound := params.CurrentRound

			if triggerRound == consts.Battle_Round_Fourth {
				util.BuffEffectWrapRemove(ctx, mostIntelligenceGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, t.Id())
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					mostIntelligenceGeneral.BaseInfo.Name,
					consts.BuffEffectType_SufferWeaponDamageDeduce,
				)
			}

			return triggerResp
		})
	}

	//使我军武力最高的武将受到的谋略伤害降低40%（受智力影响）
	mostForceGeneral := util.GetMostForcePairGeneral(t.tacticsParams)
	strategyDamageDeduceRate := 0.4 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
	if util.BuffEffectWrapSet(ctx, mostForceGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRate: strategyDamageDeduceRate,
		FromTactic: t.Id(),
	}).IsSuccess {
		hlog.CtxInfof(ctx, "[%s]受到的谋略伤害降低了%.2f%%",
			mostForceGeneral.BaseInfo.Name,
			strategyDamageDeduceRate*100,
		)

		util.TacticsTriggerWrapRegister(mostForceGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerRound := params.CurrentRound
			triggerGeneral := params.CurrentGeneral

			if triggerRound == consts.Battle_Round_Fourth {
				util.BuffEffectWrapRemove(ctx, triggerGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, t.Id())
				hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
					mostForceGeneral.BaseInfo.Name,
					consts.BuffEffectType_SufferStrategyDamageDeduce,
				)
			}

			return triggerResp
		})
	}
}

func (t TakeRefugeFromEnemiesTactic) Id() consts.TacticId {
	return consts.TakeRefugeFromEnemies
}

func (t TakeRefugeFromEnemiesTactic) Name() string {
	return "暂避其锋"
}

func (t TakeRefugeFromEnemiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TakeRefugeFromEnemiesTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TakeRefugeFromEnemiesTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TakeRefugeFromEnemiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (t TakeRefugeFromEnemiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TakeRefugeFromEnemiesTactic) Execute() {

}
