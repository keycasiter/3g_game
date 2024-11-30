package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 妙算
// 造成控制效果时，50%概率降低目标16点智力和统率，持续1回合
type WarBookDetailType_22 struct {
}

func (w *WarBookDetailType_22) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_DebuffEffect, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		sufferGeneral := params.SufferAttackGeneral

		if util.IsControlDeBuffEffect(params.DebuffEffect) {
			if util.GenerateRate(0.5) {
				//降低智力
				if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
					EffectValue:    16,
					EffectRound:    1,
					FromWarbook:    consts.WarBookDetailType_22,
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfWarbookCostRound(&util.DeBuffEffectOfWarbookCostRoundParams{
							Ctx:               ctx,
							General:           revokeGeneral,
							EffectType:        consts.DebuffEffectType_DecrIntelligence,
							WarbookDetailType: consts.WarBookDetailType_22,
						})

						return revokeResp
					})
				}
				//降低统率
				if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
					EffectValue:    16,
					EffectRound:    1,
					FromWarbook:    consts.WarBookDetailType_22,
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfWarbookCostRound(&util.DeBuffEffectOfWarbookCostRoundParams{
							Ctx:               ctx,
							General:           revokeGeneral,
							EffectType:        consts.DebuffEffectType_DecrCommand,
							WarbookDetailType: consts.WarBookDetailType_22,
						})

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}
