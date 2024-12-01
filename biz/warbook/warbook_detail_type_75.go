package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 持重
// 第4、5回合，自身免疫混乱
type WarBookDetailType_75 struct {
}

func (w *WarBookDetailType_75) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if params.CurrentRound == consts.Battle_Round_Fourth ||
			params.CurrentRound == consts.Battle_Round_Fifth {

			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_ImmunityChaos, &vo.EffectHolderParams{
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_75,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_ImmunityChaos,
						WarbookDetailType: consts.WarBookDetailType_75,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
