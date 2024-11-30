package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 胜战
// 战斗第2回合，自身获得65%破阵状态（造成伤害时无视目标的统御和智力）
type WarBookDetailType_15 struct {
}

func (w *WarBookDetailType_15) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerCurrentGeneral := params.CurrentGeneral

		if params.CurrentRound == consts.Battle_Round_Second {
			if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_BreakFormation, &vo.EffectHolderParams{
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_15,
				ProduceGeneral: general,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerCurrentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           general,
						EffectType:        consts.BuffEffectType_BreakFormation,
						WarbookDetailType: consts.WarBookDetailType_15,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
