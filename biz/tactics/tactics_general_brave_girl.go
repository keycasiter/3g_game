package tactics

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 将门虎女
// 对敌军群体（2人）造成兵刃伤害（伤害率128%）及虎嗔效果：
// 下1回合受到额外兵刃伤害（伤害率20%，目标每次受到伤害时，伤害率提高30%，最多叠加3次），
// 若目标在虎嗔效果期间受到3次伤害时，立即结算虎嗔效果并额外造成1回合震慑，并使自己造成的兵刃伤害提升8%（兵刃伤害提升效果可叠加）
type GeneralBraveGirlTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GeneralBraveGirlTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 0.6
	return g
}

func (g GeneralBraveGirlTactic) Prepare() {

}

func (g GeneralBraveGirlTactic) Id() consts.TacticId {
	return consts.GeneralBraveGirl
}

func (g GeneralBraveGirlTactic) Name() string {
	return "将门虎女"
}

func (g GeneralBraveGirlTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (g GeneralBraveGirlTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GeneralBraveGirlTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GeneralBraveGirlTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (g GeneralBraveGirlTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GeneralBraveGirlTactic) Execute() {
	// 对敌军群体（2人）造成兵刃伤害（伤害率128%）及虎嗔效果：
	// 下1回合受到额外兵刃伤害（伤害率20%，目标每次受到伤害时，伤害率提高30%，最多叠加3次），
	// 若目标在虎嗔效果期间受到3次伤害时，立即结算虎嗔效果并额外造成1回合震慑，并使自己造成的兵刃伤害提升8%（兵刃伤害提升效果可叠加）

	ctx := g.tacticsParams.Ctx
	currentGeneral := g.tacticsParams.CurrentGeneral
	currentRound := g.tacticsParams.CurrentRound

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		g.Name(),
	)

	//找到敌军2人
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, g.tacticsParams)
	for _, general := range enemyGenerals {
		//造成伤害
		dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.28)
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams: g.tacticsParams,
			AttackGeneral: currentGeneral,
			SufferGeneral: general,
			DamageType:    consts.DamageType_Weapon,
			Damage:        dmg,
			TacticId:      g.Id(),
			TacticName:    g.Name(),
		})

		//下回合结算注册器
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			settleRound := params.CurrentRound
			settleGeneral := params.CurrentGeneral
			settleResp := &vo.TacticsTriggerResult{}

			if settleRound != currentRound+1 {
				//不满足回合要求
				return settleResp
			}

			//受到额外兵刃伤害（伤害率20%）
			dmgRate := 0.2
			settleDmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * dmgRate)
			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams: g.tacticsParams,
				AttackGeneral: currentGeneral,
				SufferGeneral: settleGeneral,
				DamageType:    consts.DamageType_Weapon,
				Damage:        settleDmg,
				TacticId:      g.Id(),
				TacticName:    g.Name(),
				EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_TigerAnger),
			})
			//目标每次受到伤害时，伤害率提高30%，最多叠加3次
			//兵刃伤害
			effectRate := 0.3
			if util.DebuffEffectWrapSet(ctx, settleGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
				EffectRate:     effectRate,
				EffectRound:    1,
				MaxEffectTimes: 3,
				FromTactic:     g.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//消失效果
				util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    settleGeneral,
						EffectType: consts.DebuffEffectType_SufferWeaponDamageImprove,
						TacticId:   g.Id(),
					})

					return revokeResp
				})
			}
			//谋略伤害
			if util.DebuffEffectWrapSet(ctx, settleGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
				EffectRate:     effectRate,
				EffectRound:    1,
				MaxEffectTimes: 3,
				FromTactic:     g.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//消失效果
				util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    settleGeneral,
						EffectType: consts.DebuffEffectType_SufferStrategyDamageImprove,
						TacticId:   g.Id(),
					})

					return revokeResp
				})
			}

			return settleResp
		})

		//虎嗔效果
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerGeneral := g.tacticsParams.CurrentSufferGeneral
			triggerResp := &vo.TacticsTriggerResult{}

			//虎嗔计数
			if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_TigerAnger_Prepare, &vo.EffectHolderParams{
				EffectTimes:    1,
				MaxEffectTimes: 3,
				FromTactic:     g.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//伤害结算
				if effectParams, ok := util.DeBuffEffectOfTacticGet(triggerGeneral, consts.DebuffEffectType_TigerAnger_Prepare, g.Id()); ok {
					//若目标在虎嗔效果期间受到3次伤害时，立即结算虎嗔效果并额外造成1回合震慑，并使自己造成的兵刃伤害提升8%（兵刃伤害提升效果可叠加）
					effectTimes := int64(0)
					for _, param := range effectParams {
						effectTimes += param.EffectTimes
					}
					//立即结算
					if effectTimes == 3 {
						//立即结算
						//伤害提升计算
						dmgRate := 0.2
						if effectTimes > 0 {
							dmgRate = dmgRate * (1 + 0.3*cast.ToFloat64(effectTimes))
						}
						settleDmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * dmgRate)
						damage.TacticDamage(&damage.TacticDamageParam{
							TacticsParams: g.tacticsParams,
							AttackGeneral: currentGeneral,
							SufferGeneral: triggerGeneral,
							DamageType:    consts.DamageType_Weapon,
							Damage:        settleDmg,
							TacticId:      g.Id(),
							TacticName:    g.Name(),
							EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_TigerAnger),
						})
						//施加震慑
						if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
							EffectRound:    1,
							FromTactic:     g.Id(),
							ProduceGeneral: currentGeneral,
						}).IsSuccess {
							//消失效果
							util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								revokeResp := &vo.TacticsTriggerResult{}
								revokeGeneral := params.CurrentGeneral

								util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    revokeGeneral,
									EffectType: consts.DebuffEffectType_Awe,
									TacticId:   g.Id(),
								})

								return revokeResp
							})
						}
						//并使自己造成的兵刃伤害提升8%（兵刃伤害提升效果可叠加）
						util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
							EffectRate:     0.08,
							FromTactic:     g.Id(),
							ProduceGeneral: currentGeneral,
						})
					}
				}
			}
			return triggerResp
		})
		// 若目标在虎嗔效果期间受到3次伤害时，立即结算虎嗔效果并额外造成1回合震慑，并使自己造成的兵刃伤害提升8%（兵刃伤害提升效果可叠加）
	}
}

func (g GeneralBraveGirlTactic) IsTriggerPrepare() bool {
	return false
}
