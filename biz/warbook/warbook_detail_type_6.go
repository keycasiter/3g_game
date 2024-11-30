package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 不勇则死
// 战斗第5回合开始，每回合有40%概率进入连击、禁疗状态
type WarBookDetailType_6 struct {
}

func (w *WarBookDetailType_6) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if params.CurrentRound < consts.Battle_Round_Fifth {
			return triggerResp
		}

		if !util.GenerateRate(0.4) {
			return triggerResp
		}

		//连击
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_ContinuousAttack, &vo.EffectHolderParams{
			EffectRound: 1,
			FromWarbook: consts.WarBookDetailType_6,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}

				util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           general,
					EffectType:        consts.BuffEffectType_ContinuousAttack,
					WarbookDetailType: consts.WarBookDetailType_6,
				})

				return revokeResp
			})
		}
		//禁疗
		if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
			EffectRound: 1,
			FromWarbook: consts.WarBookDetailType_6,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}

				util.DeBuffEffectOfWarbookCostRound(&util.DeBuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           general,
					EffectType:        consts.DebuffEffectType_ProhibitionTreatment,
					WarbookDetailType: consts.WarBookDetailType_6,
				})

				return revokeResp
			})
		}

		return triggerResp
	})
}
