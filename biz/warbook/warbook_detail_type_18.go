package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 攻其不备
// 对敌军兵力最高者造成的伤害提高10%
type WarBookDetailType_18 struct {
}

func (w *WarBookDetailType_18) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_Damage, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		sufferGeneral := params.SufferAttackGeneral

		enemyGeneralWhoHighestSoliderNum := util.GetEnemyOneGeneralByHighestSolider(sufferGeneral, tacticParams)

		if enemyGeneralWhoHighestSoliderNum.BaseInfo.UniqueId == sufferGeneral.BaseInfo.UniqueId {
			//谋略
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.1,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_18,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_LaunchStrategyDamageImprove,
						WarbookDetailType: consts.WarBookDetailType_18,
					})

					return revokeResp
				})
			}

			//兵刃
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.1,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_18,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_LaunchStrategyDamageImprove,
						WarbookDetailType: consts.WarBookDetailType_18,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
