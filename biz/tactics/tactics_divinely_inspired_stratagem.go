package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 神机莫测
// 使敌军单体混乱2回合，并对自身外的敌我全体依次判定：
// 若未混乱则有35%概率使其混乱2回合，友军已混乱时，解除其负面状态（贾诩为主将时，提高目标友军12%造成伤害，受智力影响，持续2回合）；
// 敌军已混乱时，对其造成谋略攻击（伤害率175%，受智力影响）
// 主动，65%
type DivinelyInspiredStratagemTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DivinelyInspiredStratagemTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.65
	return d
}

func (d DivinelyInspiredStratagemTactic) Prepare() {
}

func (d DivinelyInspiredStratagemTactic) Id() consts.TacticId {
	return consts.DivinelyInspiredStratagem
}

func (d DivinelyInspiredStratagemTactic) Name() string {
	return "神机莫测"
}

func (d DivinelyInspiredStratagemTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (d DivinelyInspiredStratagemTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DivinelyInspiredStratagemTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DivinelyInspiredStratagemTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DivinelyInspiredStratagemTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DivinelyInspiredStratagemTactic) Execute() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral

	// 使敌军单体混乱2回合
	//找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, d.tacticsParams)
	//施加混乱效果
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
		EffectRound: 2,
		FromTactic:  d.Id(),
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_Chaos,
				TacticId:   d.Id(),
			})

			return revokeResp
		})
	}
	//并对自身外的敌我全体依次判定：
	allNotSelfGenerals := util.GetAllGeneralsNotSelfByGeneral(currentGeneral, d.tacticsParams)
	for _, general := range allNotSelfGenerals {
		//若未混乱则有35%概率使其混乱2回合
		if !util.DeBuffEffectContains(general, consts.DebuffEffectType_Chaos) {
			if util.GenerateRate(0.35) {
				//施加效果
				util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
					EffectRound: 2,
					FromTactic:  d.Id(),
				})
			}
		} else {
			//已混乱
			//友军已混乱时，解除其负面状态（贾诩为主将时，提高目标友军12%造成伤害，受智力影响，持续2回合）
			if util.IsPair(currentGeneral, general, d.tacticsParams) {
				util.DebuffEffectWrapRemove(ctx, general, consts.DebuffEffectType_Chaos, d.Id())

				if currentGeneral.IsMaster {
					//施加效果
					rate := 0.12 + (currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase / 100 / 100)
					if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
						EffectRate:  rate,
						EffectRound: 2,
						FromTactic:  d.Id(),
					}).IsSuccess {
						//注册消失效果
						util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
								TacticId:   d.Id(),
							})

							return revokeResp
						})
					}
				}
			} else {
				// 敌军已混乱时，对其造成谋略攻击（伤害率175%，受智力影响）
				dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.75
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     d.tacticsParams,
					AttackGeneral:     currentGeneral,
					SufferGeneral:     general,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: dmgRate,
					TacticId:          d.Id(),
					TacticName:        d.Name(),
				})
			}
		}
	}

}

func (d DivinelyInspiredStratagemTactic) IsTriggerPrepare() bool {
	return false
}
