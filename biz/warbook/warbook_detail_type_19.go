package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 大谋不谋
// 战斗中，每次成功发动主动战法时，有50%概率增加4.5%奇谋和会心，持续2回合；可叠加2次
type WarBookDetailType_19 struct {
}

func (w *WarBookDetailType_19) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if !util.GenerateRate(0.5) {
			return triggerResp
		}

		//会心
		if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_EnhanceWeapon, &vo.EffectHolderParams{
			EffectRate:     0.045,
			EffectRound:    2,
			EffectTimes:    1,
			MaxEffectTimes: 2,
			FromWarbook:    consts.WarBookDetailType_19,
			ProduceGeneral: triggerGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           revokeGeneral,
					EffectType:        consts.BuffEffectType_EnhanceWeapon,
					WarbookDetailType: consts.WarBookDetailType_19,
				})

				return revokeResp
			})
		}

		//奇谋
		if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_EnhanceStrategy, &vo.EffectHolderParams{
			EffectRate:     0.045,
			EffectRound:    2,
			EffectTimes:    1,
			MaxEffectTimes: 2,
			FromWarbook:    consts.WarBookDetailType_19,
			ProduceGeneral: triggerGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           revokeGeneral,
					EffectType:        consts.BuffEffectType_EnhanceStrategy,
					WarbookDetailType: consts.WarBookDetailType_19,
				})

				return revokeResp
			})
		}

		return triggerResp
	})
}
