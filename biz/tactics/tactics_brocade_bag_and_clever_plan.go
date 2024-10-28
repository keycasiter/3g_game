package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 锦囊妙计
// 战斗中，奇数回合32%概率（受智力影响），偶数回合有75%概率（受智力影响）使友军单体的自带主动战法发动率100%且有25%概率跳过1回合准备，持续1回合；
// 自身为主将时，触发时额外对敌军单体造成谋略伤害（伤害率45%，受智力及双方智力差影响），并使跳过概率提升到80%
// 指挥 100%
type BrocadeBagAndCleverPlanTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BrocadeBagAndCleverPlanTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BrocadeBagAndCleverPlanTactic) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral
	currentRound := b.tacticsParams.CurrentRound

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		roundTriggerResp := &vo.TacticsTriggerResult{}

		//奇数回合32%概率（受智力影响），偶数回合有75%概率（受智力影响）
		effectRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase / 100 / 100
		rate := 0.32 + effectRate
		if currentRound%2 == 0 {
			rate = 0.75 + effectRate
		}

		if util.GenerateRate(rate) {
			//找到友军单体
			pairGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, b.tacticsParams)
			//找到自带主动战法
			for idx, equipTactic := range pairGeneral.EquipTactics {
				//自带 + 主动
				if idx == 0 && consts.ActiveTacticsMap[equipTactic.Id] {
					util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						triggerGeneral := params.CurrentGeneral
						triggerResp := &vo.TacticsTriggerResult{}
						//主动战法触发率提升
						if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
							TriggerRate: 1.0,
							EffectRound: 1,
							FromTactic:  b.Id(),
						}).IsSuccess {
							//注册消失效果
							util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_ActiveTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								revokeResp := &vo.TacticsTriggerResult{}
								revokeGeneral := params.CurrentGeneral

								util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    revokeGeneral,
									EffectType: consts.BuffEffectType_TacticsActiveTriggerImprove,
									TacticId:   b.Id(),
								})

								return revokeResp
							})
						}
						return triggerResp
					})

					//有25%概率跳过1回合准备，持续1回合，主将时跳过概率提升到80%
					skipTriggerRate := 0.25
					if currentGeneral.IsMaster {
						skipTriggerRate = 0.8
					}
					if util.GenerateRate(skipTriggerRate) {
						if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_ActiveTactic_SkipPrepareRound, &vo.EffectHolderParams{
							EffectRound: 1,
							FromTactic:  b.Id(),
						}).IsSuccess {
							//注册消失效果
							util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								revokeResp := &vo.TacticsTriggerResult{}
								revokeGeneral := params.CurrentGeneral

								util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    revokeGeneral,
									EffectType: consts.BuffEffectType_ActiveTactic_SkipPrepareRound,
									TacticId:   b.Id(),
								})

								return revokeResp
							})
						}
					}
				}
			}

			//触发时额外对敌军单体造成谋略伤害（伤害率45%，受智力及双方智力差影响）
			//敌军单体
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, b.tacticsParams)
			diff := util.CalculateAttrDiff(
				enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase,
				currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase,
			)
			dmgRate := diff/100/100 + 0.45
			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams:     b.tacticsParams,
				AttackGeneral:     currentGeneral,
				SufferGeneral:     enemyGeneral,
				DamageType:        consts.DamageType_Strategy,
				DamageImproveRate: dmgRate,
				TacticId:          b.Id(),
				TacticName:        b.Name(),
			})
		}

		return roundTriggerResp
	})
}

func (b BrocadeBagAndCleverPlanTactic) Id() consts.TacticId {
	return consts.BrocadeBagAndCleverPlan
}

func (b BrocadeBagAndCleverPlanTactic) Name() string {
	return "锦囊妙计"
}

func (b BrocadeBagAndCleverPlanTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (b BrocadeBagAndCleverPlanTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BrocadeBagAndCleverPlanTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BrocadeBagAndCleverPlanTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (b BrocadeBagAndCleverPlanTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BrocadeBagAndCleverPlanTactic) Execute() {

}

func (b BrocadeBagAndCleverPlanTactic) IsTriggerPrepare() bool {
	return false
}
