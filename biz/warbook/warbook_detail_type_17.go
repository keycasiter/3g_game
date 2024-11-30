package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 以治击乱
// 对于控制状态武将伤害提升12%
type WarBookDetailType_17 struct {
}

func (w *WarBookDetailType_17) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_Damage, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		sufferGeneral := params.SufferAttackGeneral

		if util.DeBuffEffectContainsControl(sufferGeneral) {

			//谋略伤害
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.12,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_17,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_LaunchStrategyDamageImprove,
						WarbookDetailType: consts.WarBookDetailType_17,
					})

					return revokeResp
				})
			}

			//兵刃伤害
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.12,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_17,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_LaunchWeaponDamageImprove,
						WarbookDetailType: consts.WarBookDetailType_17,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
