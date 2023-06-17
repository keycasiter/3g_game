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

// 水淹七军
// 对敌军群体（2人）施加水攻状态（伤害率96%，受武力影响，自身为主将时，伤害率提升至108%），持续2回合；
// 第二次及之后释放前，使敌军群体（2～3人）受到兵刃伤害提升20%（受武力影响），持续1回合；
// 第三次及之后释放时，有40%概率（受武力影响）立即结算敌军全体的水攻状态（不清除状态）；
// 第四次释放后，对有水攻状态的敌军造成兵刃伤害（伤害率208%）
// 主动，50%
type FloodedSeventhArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	//战法触发次数
	triggerCnt int
}

func (f FloodedSeventhArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.5
	return f
}

func (f FloodedSeventhArmyTactic) Prepare() {
}

func (f FloodedSeventhArmyTactic) Id() consts.TacticId {
	return consts.FloodedSeventhArmy
}

func (f FloodedSeventhArmyTactic) Name() string {
	return "水淹七军"
}

func (f FloodedSeventhArmyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (f FloodedSeventhArmyTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FloodedSeventhArmyTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FloodedSeventhArmyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FloodedSeventhArmyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FloodedSeventhArmyTactic) Execute() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	//释放计数
	f.triggerCnt++

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)
	//第二次及之后释放前，使敌军群体（2～3人）受到兵刃伤害提升20%（受武力影响），持续1回合
	if f.triggerCnt >= 2 {
		//找到敌军群体2～3
		enemyGenerals := util.GetEnemyGeneralsTwoOrThreeMap(f.tacticsParams)
		for _, enemyGeneral := range enemyGenerals {
			effectRate := 0.2 + currentGeneral.BaseInfo.AbilityAttr.ForceBase/100/100
			if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
				EffectRate:     effectRate,
				EffectRound:    1,
				FromTactic:     f.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//消失效果
				util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    currentGeneral,
						EffectType: consts.DebuffEffectType_SufferWeaponDamageImprove,
						TacticId:   f.Id(),
					})

					return revokeResp
				})
			}
		}
	}
	// 第三次及之后释放时，有40%概率（受武力影响）立即结算敌军全体的水攻状态（不清除状态）；
	if f.triggerCnt >= 3 {
		triggerRate := 0.4 + currentGeneral.BaseInfo.AbilityAttr.ForceBase/100/100
		if util.GenerateRate(triggerRate) {
			//找到敌军全体
			allEnemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, f.tacticsParams)
			for _, enemyGeneral := range allEnemyGenerals {
				//是否有水攻状态
				if effectParams, ok := util.DeBuffEffectGet(enemyGeneral, consts.DebuffEffectType_WaterAttack); ok {
					for _, effectParam := range effectParams {
						dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * effectParam.EffectRate)
						util.TacticDamage(&util.TacticDamageParam{
							TacticsParams: f.tacticsParams,
							AttackGeneral: currentGeneral,
							SufferGeneral: enemyGeneral,
							DamageType:    consts.DamageType_Weapon,
							Damage:        dmg,
							TacticName:    f.Name(),
							EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_WaterAttack),
						})
					}
				}
			}
		}
	}
	// 第四次释放后，对有水攻状态的敌军造成兵刃伤害（伤害率208%）
	if f.triggerCnt >= 4 {
		//找到敌军全体
		allEnemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, f.tacticsParams)
		for _, general := range allEnemyGenerals {
			//是否有水攻状态
			if util.DeBuffEffectContains(general, consts.DebuffEffectType_WaterAttack) {
				dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.08)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: f.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: general,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticName:    f.Name(),
				})
			}
		}
	}

	// 对敌军群体（2人）施加水攻状态（伤害率96%，受武力影响，自身为主将时，伤害率提升至108%），持续2回合；
	//找到敌军2人
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, f.tacticsParams)
	//施加状态
	for _, general := range enemyGenerals {
		effectRate := 0.96
		if currentGeneral.IsMaster {
			effectRate = 1.08
		}
		if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_WaterAttack, &vo.EffectHolderParams{
			EffectRate:     effectRate,
			EffectRound:    2,
			FromTactic:     f.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//消失效果
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_WaterAttack,
					TacticId:   f.Id(),
				}) {

					dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * effectRate)
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams: f.tacticsParams,
						AttackGeneral: currentGeneral,
						SufferGeneral: revokeGeneral,
						DamageType:    consts.DamageType_Weapon,
						Damage:        dmg,
						TacticName:    f.Name(),
						EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_WaterAttack),
					})
				}

				return revokeResp
			})
		}
	}

}

func (f FloodedSeventhArmyTactic) IsTriggerPrepare() bool {
	return false
}
