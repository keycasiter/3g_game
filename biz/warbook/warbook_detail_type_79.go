package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 精准
// 第2、3回合，获得必中状态
type WarBookDetailType_79 struct {
}

func (w *WarBookDetailType_79) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if params.CurrentRound == consts.Battle_Round_Second ||
			params.CurrentRound == consts.Battle_Round_Third {

			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_MustHit, &vo.EffectHolderParams{
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_79,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_MustHit,
						WarbookDetailType: consts.WarBookDetailType_79,
					})

					return revokeResp
				})
			}

			return triggerResp
		}

		return triggerResp
	})
}
