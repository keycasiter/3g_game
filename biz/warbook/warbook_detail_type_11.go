package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 分险
// 成功发动突击战法后，使自身受到伤害降低6%，持续1回合
type WarBookDetailType_11 struct {
}

func (w *WarBookDetailType_11) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_AssaultTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		//兵刃伤害
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.06,
			EffectRound:    1,
			FromWarbook:    consts.WarBookDetailType_11,
			ProduceGeneral: general,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}

				util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           general,
					EffectType:        consts.BuffEffectType_SufferWeaponDamageDeduce,
					WarbookDetailType: consts.WarBookDetailType_11,
				})

				return revokeResp
			})
		}
		//谋略伤害
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.06,
			EffectRound:    1,
			FromWarbook:    consts.WarBookDetailType_11,
			ProduceGeneral: general,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}

				util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           general,
					EffectType:        consts.BuffEffectType_SufferStrategyDamageDeduce,
					WarbookDetailType: consts.WarBookDetailType_11,
				})

				return revokeResp
			})
		}

		return triggerResp
	})
}
