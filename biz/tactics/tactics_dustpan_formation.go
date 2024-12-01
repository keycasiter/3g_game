package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 箕形阵
// 战斗前3回合，使敌军主将造成伤害降低40%（受武力影响），并使我军随机副将受到兵刃伤害降低18%，另一名副将受到谋略伤害降低18%
type DustpanFormationTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DustpanFormationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 1.0
	return d
}

func (d DustpanFormationTactic) Prepare() {
	ctx := d.tacticsParams.Ctx
	currentGeneral := d.tacticsParams.CurrentGeneral
	//战斗前3回合，使敌军主将造成伤害降低40%（受武力影响），并使我军随机副将受到兵刃伤害降低18%，另一名副将受到谋略伤害降低18%

	//施加效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}

		//找到敌军主将
		masterEnemyGeneral := util.GetEnemyMasterGeneral(currentGeneral, d.tacticsParams)
		rate := 0.4 + (currentGeneral.BaseInfo.AbilityAttr.ForceBase / 100 / 100)
		if util.DebuffEffectWrapSet(ctx, masterEnemyGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate:  rate,
			EffectRound: 3,
			FromTactic:  d.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(masterEnemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_LaunchWeaponDamageDeduce,
					TacticId:   d.Id(),
				})

				return revokeResp
			})
		}
		if util.DebuffEffectWrapSet(ctx, masterEnemyGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate:  rate,
			EffectRound: 3,
			FromTactic:  d.Id(),
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(masterEnemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_LaunchStrategyDamageDeduce,
					TacticId:   d.Id(),
				})

				return revokeResp
			})
		}

		//并使我军随机副将受到兵刃伤害降低18%，另一名副将受到谋略伤害降低18%
		pairGenerals := util.GetPairViceGenerals(currentGeneral, d.tacticsParams)

		//效果二选一随机
		buffEffects := []consts.BuffEffectType{}
		if util.GenerateRate(0.5) {
			buffEffects = append(buffEffects, consts.BuffEffectType_SufferWeaponDamageDeduce)
			buffEffects = append(buffEffects, consts.BuffEffectType_SufferStrategyDamageDeduce)
		} else {
			buffEffects = append(buffEffects, consts.BuffEffectType_SufferStrategyDamageDeduce)
			buffEffects = append(buffEffects, consts.BuffEffectType_SufferWeaponDamageDeduce)
		}

		idx := 0
		for _, general := range pairGenerals {
			buffEffect := buffEffects[idx]
			//施加效果
			if util.BuffEffectWrapSet(ctx, general, buffEffect, &vo.EffectHolderParams{
				EffectRate:  0.18,
				EffectRound: 3,
				FromTactic:  d.Id(),
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: buffEffect,
						TacticId:   d.Id(),
					})

					return revokeResp
				})
			}
			idx++
		}

		return triggerResp
	})
}

func (d DustpanFormationTactic) Id() consts.TacticId {
	return consts.DustpanFormation
}

func (d DustpanFormationTactic) Name() string {
	return "箕形阵"
}

func (d DustpanFormationTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DustpanFormationTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DustpanFormationTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DustpanFormationTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (d DustpanFormationTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DustpanFormationTactic) Execute() {
}

func (d DustpanFormationTactic) IsTriggerPrepare() bool {
	return false
}

func (a DustpanFormationTactic) SetTriggerPrepare(triggerPrepare bool) {
}
