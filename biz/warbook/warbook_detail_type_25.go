package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 占卜
// 每回合行动时，使自身随机获得受到兵刃伤害降低6%或受到谋略伤害降低6%，持续1回合
type WarBookDetailType_25 struct {
}

func (w *WarBookDetailType_25) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.GenerateRate(0.5) {
			//兵刃
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     0.06,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_25,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_SufferWeaponDamageDeduce,
						WarbookDetailType: consts.WarBookDetailType_25,
					})

					return revokeResp
				})
			}
		} else {
			//谋略
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     0.06,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_25,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_SufferStrategyDamageDeduce,
						WarbookDetailType: consts.WarBookDetailType_25,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
