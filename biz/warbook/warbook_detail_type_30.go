package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 无战而胜
// 造成控制效果时，自身受到兵刃伤害和谋略伤害降低10%，持续1回合
type WarBookDetailType_30 struct {
}

func (w *WarBookDetailType_30) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_DebuffEffectEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if !util.IsControlDeBuffEffect(params.DebuffEffect) {
			return triggerResp
		}

		//受到兵刃伤害
		if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.1,
			EffectRound:    1,
			FromWarbook:    consts.WarBookDetailType_30,
			ProduceGeneral: triggerGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           revokeGeneral,
					EffectType:        consts.BuffEffectType_SufferWeaponDamageDeduce,
					WarbookDetailType: consts.WarBookDetailType_30,
				})

				return revokeResp
			})
		}

		//受到谋略伤害
		if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.1,
			EffectRound:    1,
			FromWarbook:    consts.WarBookDetailType_30,
			ProduceGeneral: triggerGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           revokeGeneral,
					EffectType:        consts.BuffEffectType_SufferStrategyDamageDeduce,
					WarbookDetailType: consts.WarBookDetailType_30,
				})

				return revokeResp
			})
		}

		return triggerResp
	})
}
