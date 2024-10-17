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

// 飞沙走石
// 主动，50%
// 奇数回合发动时，使敌军群体（2人）智力降低36点（受智力影响），持续1回合，
// 并使其陷入2回合水攻状态（伤害率58%，受智力影响）；
// 偶数回合发动时，使敌军随机单体武力与智力数值对换，持续1回合，并使其陷入2回合沙暴状态（伤害率108%，受智力影响）
type FlyingSandAndRollingPebblesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FlyingSandAndRollingPebblesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.5
	return f
}

func (f FlyingSandAndRollingPebblesTactic) Prepare() {
}

func (f FlyingSandAndRollingPebblesTactic) Id() consts.TacticId {
	return consts.FlyingSandAndRollingPebbles
}

func (f FlyingSandAndRollingPebblesTactic) Name() string {
	return "飞沙走石"
}

func (f FlyingSandAndRollingPebblesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (f FlyingSandAndRollingPebblesTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FlyingSandAndRollingPebblesTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FlyingSandAndRollingPebblesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FlyingSandAndRollingPebblesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FlyingSandAndRollingPebblesTactic) Execute() {
	currentGeneral := f.tacticsParams.CurrentGeneral
	ctx := f.tacticsParams.Ctx
	currentRound := f.tacticsParams.CurrentRound

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	//奇数回合发动时，使敌军群体（2人）智力降低36点（受智力影响），持续1回合，
	//并使其陷入2回合水攻状态（伤害率58%，受智力影响）；
	if currentRound%2 != 0 {
		//找到敌军2人
		enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, f.tacticsParams)
		for _, enemyGeneral := range enemyGenerals {
			effectValue := 36 + cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100)

			//施加效果(降低智力)
			if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
				EffectValue: effectValue,
				EffectRound: 1,
				FromTactic:  f.Id(),
			}).IsSuccess {
				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    enemyGeneral,
					EffectType: consts.DebuffEffectType_DecrIntelligence,
					TacticId:   f.Id(),
				})
			}
			//施加效果(水攻)
			if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_WaterAttack, &vo.EffectHolderParams{
				EffectRound: 2,
				FromTactic:  f.Id(),
			}).IsSuccess {
				if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    enemyGeneral,
					EffectType: consts.DebuffEffectType_WaterAttack,
					TacticId:   f.Id(),
				}) {
					//每回合持续造成伤害（伤害率58%，受智力影响）
					dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.58)
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams: f.tacticsParams,
						AttackGeneral: currentGeneral,
						SufferGeneral: enemyGeneral,
						DamageType:    consts.DamageType_Strategy,
						Damage:        dmg,
						TacticName:    f.Name(),
						EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_WaterAttack),
					})
				}
			}
		}
	}

	//偶数回合发动时，使敌军随机单体武力与智力数值对换，持续1回合，并使其陷入2回合沙暴状态（伤害率108%，受智力影响）
	if currentRound%2 == 0 {
		//找到敌军单体
		enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, f.tacticsParams)

		//兑换效果
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_AggregateStoneIntoGold_Exchange, &vo.EffectHolderParams{
			EffectRound: 1,
			FromTactic:  f.Id(),
		}).IsSuccess {
			//对换
			tmpForce := enemyGeneral.BaseInfo.AbilityAttr.ForceBase
			enemyGeneral.BaseInfo.AbilityAttr.ForceBase = enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase
			enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase = tmpForce

			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				//消耗回合
				if !util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_AggregateStoneIntoGold_Exchange,
					TacticId:   f.Id(),
				}) {
					//对换恢复
					revokeForce := enemyGeneral.BaseInfo.AbilityAttr.ForceBase
					enemyGeneral.BaseInfo.AbilityAttr.ForceBase = enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase
					enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase = revokeForce
				}

				return revokeResp
			})
		}
		//沙暴效果
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Sandstorm, &vo.EffectHolderParams{
			EffectRound: 2,
			FromTactic:  f.Id(),
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				//消耗回合
				if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    enemyGeneral,
					EffectType: consts.DebuffEffectType_Sandstorm,
					TacticId:   f.Id(),
				}) {
					dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.08)
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams: f.tacticsParams,
						AttackGeneral: currentGeneral,
						SufferGeneral: revokeGeneral,
						DamageType:    consts.DamageType_Strategy,
						Damage:        dmg,
						TacticId:      f.Id(),
						TacticName:    f.Name(),
						EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Sandstorm),
					})
				}
				return revokeResp
			})
		}
	}
}

func (f FlyingSandAndRollingPebblesTactic) IsTriggerPrepare() bool {
	return false
}
