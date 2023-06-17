package tactics

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 火烧连营
// 对敌军单体施加灼烧状态（伤害率82%，受智力影响，自身为主将时，伤害率提高至98%），持续3回合，随机释放2次
// 若目标已有灼烧状态则进行焚营，对敌军全体造成谋略攻击（伤害率62%，受智力影响），之后对另一名敌军施加灼烧状态，若处于灼烧状态的目标同一回合内受到两次焚营伤害，则有30%概率进入震慑状态，持续1回合
// 主动 50%
type FireJointVentureTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FireJointVentureTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.5
	return f
}

func (f FireJointVentureTactic) Prepare() {
}

func (f FireJointVentureTactic) Id() consts.TacticId {
	return consts.FireJointVenture
}

func (f FireJointVentureTactic) Name() string {
	return "火烧连营"
}

func (f FireJointVentureTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (f FireJointVentureTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FireJointVentureTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FireJointVentureTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FireJointVentureTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FireJointVentureTactic) Execute() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral
	//被焚营的目标计数器
	fireCampGenerals := make(map[int64]int, 0)

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)
	// 对敌军单体施加灼烧状态
	// 客服解释：两个目标随机释放2次，结果可能是AA\BB\AB三种情况
	//找到敌军单体
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, f.tacticsParams)

	for i := 0; i < 2; i++ {
		hitIdx := util.GenerateHitOneIdx(len(enemyGenerals))
		enemyGeneral := enemyGenerals[hitIdx]

		// 若目标已有灼烧状态则进行焚营，对敌军全体造成谋略攻击（伤害率62%，受智力影响）
		if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_Firing) {
			//找到敌军全体进行焚营
			for _, general := range enemyGenerals {
				dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.62)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: f.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: general,
					DamageType:    consts.DamageType_Strategy,
					Damage:        dmg,
					TacticName:    f.Name(),
					EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_FireJointVenture_BurningCamp),
				})
				//焚营计数
				if _, ok := fireCampGenerals[general.BaseInfo.UniqueId]; ok {
					fireCampGenerals[general.BaseInfo.UniqueId]++
				} else {
					fireCampGenerals[general.BaseInfo.UniqueId] = 1
				}
				//焚营结算
				if util.DeBuffEffectContains(general, consts.DebuffEffectType_Firing) && fireCampGenerals[general.BaseInfo.UniqueId] == 2 {
					if util.GenerateRate(0.3) {
						//施加震慑效果
						if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
							EffectRound:    1,
							FromTactic:     f.Id(),
							ProduceGeneral: currentGeneral,
						}).IsSuccess {
							//注册消失效果
							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    general,
								EffectType: consts.DebuffEffectType_Awe,
								TacticId:   f.Id(),
							})
						}
					}
				}
			}
			//之后对另一名敌军施加灼烧状态，若处于灼烧状态的目标同一回合内受到两次焚营伤害，则有30%概率进入震慑状态，持续1回合
			//找到另外一名敌军
			anotherEnemyGeneral := util.GetEnemyOneGeneralNotSelfByGeneral(enemyGeneral, f.tacticsParams)

			//施加状态
			if util.DebuffEffectWrapSet(ctx, anotherEnemyGeneral, consts.DebuffEffectType_Firing, &vo.EffectHolderParams{
				EffectRound:    3,
				FromTactic:     f.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(anotherEnemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_Firing,
						TacticId:   f.Id(),
					}) {
						//（伤害率82%，受智力影响，自身为主将时，伤害率提高至98%），持续3回合，随机释放2次
						rate := 0.82
						if currentGeneral.IsMaster {
							rate = 0.98
						}
						dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * rate)
						util.TacticDamage(&util.TacticDamageParam{
							TacticsParams: f.tacticsParams,
							AttackGeneral: currentGeneral,
							SufferGeneral: revokeGeneral,
							DamageType:    consts.DamageType_Strategy,
							Damage:        dmg,
							TacticName:    f.Name(),
							EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Firing),
						})
					}

					return revokeResp
				})
			}
		} else {
			//施加状态
			if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Firing, &vo.EffectHolderParams{
				EffectRound:    3,
				FromTactic:     f.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_Firing,
						TacticId:   f.Id(),
					}) {
						//（伤害率82%，受智力影响，自身为主将时，伤害率提高至98%），持续3回合，随机释放2次
						rate := 0.82
						if currentGeneral.IsMaster {
							rate = 0.98
						}
						dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * rate)
						util.TacticDamage(&util.TacticDamageParam{
							TacticsParams: f.tacticsParams,
							AttackGeneral: currentGeneral,
							SufferGeneral: revokeGeneral,
							DamageType:    consts.DamageType_Strategy,
							Damage:        dmg,
							TacticName:    f.Name(),
							EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Firing),
						})
					}

					return revokeResp
				})
			}
		}
	}
}

func (f FireJointVentureTactic) IsTriggerPrepare() bool {
	return false
}
