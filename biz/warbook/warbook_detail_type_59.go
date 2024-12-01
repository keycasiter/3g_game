package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 应机立新
// 自身每回合首次造成伤害提升12%
type WarBookDetailType_59 struct {
}

func (w *WarBookDetailType_59) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	roundCntMap := make(map[consts.BattleRound]int, 0)

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_Damage, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if cnt, ok := roundCntMap[params.CurrentRound]; ok {
			cnt++
			roundCntMap[params.CurrentRound] = cnt
		} else {
			roundCntMap[params.CurrentRound] = 1

			//兵刃
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.12,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_59,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_LaunchWeaponDamageImprove,
						WarbookDetailType: consts.WarBookDetailType_59,
					})

					return revokeResp
				})
			}

			//谋略
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.12,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_59,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_LaunchStrategyDamageImprove,
						WarbookDetailType: consts.WarBookDetailType_59,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
