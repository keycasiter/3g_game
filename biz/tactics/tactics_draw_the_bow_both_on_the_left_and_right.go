package tactics

import (
	"fmt"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 左右开弓
// 提高自身13%会心几率（触发时兵刃伤害提高100%），对敌军单体造成一次兵刃攻击（伤害率180%）
// 如果目标为骑兵则额外造成溃散状态，每回合持续造成伤害（伤害率90%，受武力影响），持续2回合
type DrawTheBowBothOnTheLeftAndRightTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.35
	return d
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) Prepare() {
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) Id() consts.TacticId {
	return consts.DrawTheBowBothOnTheLeftAndRight
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) Name() string {
	return "左右开弓"
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) Execute() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral

	// 提高自身13%会心几率（触发时兵刃伤害提高100%）
	//施加效果
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_EnhanceWeapon, &vo.EffectHolderParams{
		EffectRate:  0.13,
		EffectRound: 2,
		FromTactic:  d.Id(),
	}).IsSuccess {

	}
	//对敌军单体造成一次兵刃攻击（伤害率180%）
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, d.tacticsParams)
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.8)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: d.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemyGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        dmg,
		TacticName:    d.Name(),
		TacticId:      d.Id(),
	})

	// 如果目标为骑兵则额外造成溃散状态，每回合持续造成伤害（伤害率90%，受武力影响），持续2回合
	if util.GetEnemyArmType(enemyGeneral, d.tacticsParams) == consts.ArmType_Cavalry {
		//施加效果
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Escape, &vo.EffectHolderParams{
			EffectRound: 2,
			FromTactic:  d.Id(),
		}).IsSuccess {
			//每回合持续伤害/注册消失效果
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				//注册消失效果
				if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    enemyGeneral,
					EffectType: consts.DebuffEffectType_Escape,
					TacticId:   d.Id(),
				}) {
					effectDmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 0.9)
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams: d.tacticsParams,
						AttackGeneral: revokeGeneral,
						SufferGeneral: enemyGeneral,
						DamageType:    consts.DamageType_Weapon,
						Damage:        effectDmg,
						TacticName:    d.Name(),
						TacticId:      d.Id(),
						EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Escape),
					})
				}

				return revokeResp
			})
		}
	}
}

func (d DrawTheBowBothOnTheLeftAndRightTactic) IsTriggerPrepare() bool {
	return false
}
