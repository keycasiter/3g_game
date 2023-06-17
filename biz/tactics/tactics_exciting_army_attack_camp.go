package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 振军击营
// 对敌军单体造成兵刃伤害（伤害率235%）及禁疗（无法恢复兵力）状态，持续3回合
// 并使有负面状态的友军单体自带战法发动几率提高10%，持续1回合
type ExcitingArmyAttackCampTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (e ExcitingArmyAttackCampTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
	e.triggerRate = 0.35
	return e
}

func (e ExcitingArmyAttackCampTactic) Prepare() {
}

func (e ExcitingArmyAttackCampTactic) Id() consts.TacticId {
	return consts.ExcitingArmyAttackCamp
}

func (e ExcitingArmyAttackCampTactic) Name() string {
	return "振军击营"
}

func (e ExcitingArmyAttackCampTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (e ExcitingArmyAttackCampTactic) GetTriggerRate() float64 {
	return e.triggerRate
}

func (e ExcitingArmyAttackCampTactic) SetTriggerRate(rate float64) {
	e.triggerRate = rate
}

func (e ExcitingArmyAttackCampTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (e ExcitingArmyAttackCampTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (e ExcitingArmyAttackCampTactic) Execute() {
	ctx := e.tacticsParams.Ctx
	currentGeneral := e.tacticsParams.CurrentGeneral
	// 对敌军单体造成兵刃伤害（伤害率235%）及禁疗（无法恢复兵力）状态，持续3回合
	// 并使有负面状态的友军单体自带战法发动几率提高10%，持续1回合

	//找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, e.tacticsParams)
	//兵刃伤害
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.35)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: e.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemyGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        dmg,
		TacticName:    e.Name(),
	})
	//禁疗
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
		EffectRound:    3,
		FromTactic:     e.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		//注册效果
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := e.tacticsParams.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_ProhibitionTreatment,
				TacticId:   e.Id(),
			})

			return revokeResp
		})
	}

	//并使有负面状态的友军单体自带战法发动几率提高10%，持续1回合
	//找到友军2人
	pairGenerals := util.GetPairGeneralsNotSelf(e.tacticsParams, currentGeneral)
	withDebuffPairGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range pairGenerals {
		if util.DeBuffEffectContainsCheck(general) {
			withDebuffPairGenerals = append(withDebuffPairGenerals, general)
		}
	}
	if len(withDebuffPairGenerals) > 0 {
		hitIdx := util.GenerateHitOneIdx(len(withDebuffPairGenerals))
		pairGeneral := withDebuffPairGenerals[hitIdx]

		//施加效果
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
			EffectRate:     0.1,
			EffectRound:    1,
			FromTactic:     e.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_TacticsActiveTriggerImprove,
					TacticId:   e.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (e ExcitingArmyAttackCampTactic) IsTriggerPrepare() bool {
	return false
}
