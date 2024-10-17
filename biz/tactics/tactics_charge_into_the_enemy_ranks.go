package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 登锋陷阵
// 对目标单体造成兵刃攻击（伤害率208%）并缴械，同时使自己的统率降低40%，并进入禁疗状态，持续2回合
// 主动 55%
type ChargeIntoTheEnemyRanksTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ChargeIntoTheEnemyRanksTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 0.55
	return c
}

func (c ChargeIntoTheEnemyRanksTactic) Prepare() {
}

func (c ChargeIntoTheEnemyRanksTactic) Id() consts.TacticId {
	return consts.ChargeIntoTheEnemyRanks
}

func (c ChargeIntoTheEnemyRanksTactic) Name() string {
	return "登锋陷阵"
}

func (c ChargeIntoTheEnemyRanksTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c ChargeIntoTheEnemyRanksTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c ChargeIntoTheEnemyRanksTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c ChargeIntoTheEnemyRanksTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (c ChargeIntoTheEnemyRanksTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c ChargeIntoTheEnemyRanksTactic) Execute() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	// 对目标单体造成兵刃攻击（伤害率208%）并缴械，同时使自己的统率降低40%，并进入禁疗状态，持续2回合

	//找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, c.tacticsParams)
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.08)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams: c.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemyGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        dmg,
		TacticId:      c.Id(),
		TacticName:    c.Name(),
	})
	//施加缴械效果
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_CancelWeapon, &vo.EffectHolderParams{
		EffectRound: 2,
		FromTactic:  c.Id(),
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_CancelWeapon,
				TacticId:   c.Id(),
			})

			return revokeResp
		})
	}

	//使自己的统率降低40%
	decrVal := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.CommandBase * 0.4)
	if util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
		EffectValue: decrVal,
		EffectRound: 2,
		FromTactic:  c.Id(),
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrCommand,
				TacticId:   c.Id(),
			})

			return revokeResp
		})
	}

	//并进入禁疗状态
	if util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
		EffectRound: 2,
		FromTactic:  c.Id(),
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_ProhibitionTreatment,
				TacticId:   c.Id(),
			})

			return revokeResp
		})
	}
}

func (c ChargeIntoTheEnemyRanksTactic) IsTriggerPrepare() bool {
	return false
}
