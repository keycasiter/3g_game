package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 谋定后动
// 自身所有不需要准备的主动战法增加1回合准备，该战法伤害提高100%
type WarBookDetailType_54 struct {
}

func (w *WarBookDetailType_54) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		currentTactic := params.CurrentTactic.(_interface.Tactics)

		if currentTactic.TacticsType() == consts.TacticsType_Active &&
			!currentTactic.IsTriggerPrepare() {
			currentTactic.SetTriggerPrepare(true)
		}

		util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerActiveTacticResp := &vo.TacticsTriggerResult{}
			triggerActiveGeneral := params.CurrentGeneral

			//当前战法
			if currentTactic.Id() == params.EffectHolderParams.FromTactic {
				if util.BuffEffectWrapSet(ctx, triggerActiveGeneral, consts.BuffEffectType_TacticsActiveDamageImprove, &vo.EffectHolderParams{
					EffectRate:     1,
					EffectRound:    1,
					FromWarbook:    consts.WarBookDetailType_54,
					ProduceGeneral: general,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(triggerActiveGeneral, consts.BattleAction_ActiveTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral
						util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
							Ctx:               ctx,
							General:           revokeGeneral,
							EffectType:        consts.BuffEffectType_TacticsActiveDamageImprove,
							WarbookDetailType: consts.WarBookDetailType_54,
						})

						return revokeResp
					})
				}
			}

			return triggerActiveTacticResp
		})

		return triggerResp
	})
}
