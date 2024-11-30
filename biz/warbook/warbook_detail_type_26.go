package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 合变
// 每回合结束时，若本回合自身未发动主动战法，则使自身受到伤害降低6%，持续1回合
type WarBookDetailType_26 struct {
}

func (w *WarBookDetailType_26) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if tacticStatistics, ok := tacticParams.BattleTacticStatisticsMap[triggerGeneral.BaseInfo.UniqueId]; ok {
			for _, statistics := range tacticStatistics {
				//主动战法
				if consts.ActiveTacticsMap[consts.TacticId(statistics.TacticId)] {
					//本回合是否发动过
					if m, okk := statistics.RoundTriggerTimes[tacticParams.CurrentPhase]; okk {
						if cnt, okkk := m[tacticParams.CurrentRound]; okkk {
							if cnt > 0 {
								return triggerResp
							}
							//未发动过

							//降低遭受兵刃伤害
							if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
								EffectRate:     0.06,
								EffectRound:    1,
								FromWarbook:    consts.WarBookDetailType_26,
								ProduceGeneral: triggerGeneral,
							}).IsSuccess {
								util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
									revokeResp := &vo.TacticsTriggerResult{}
									revokeGeneral := params.CurrentGeneral

									util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
										Ctx:               ctx,
										General:           revokeGeneral,
										EffectType:        consts.BuffEffectType_SufferWeaponDamageDeduce,
										WarbookDetailType: consts.WarBookDetailType_26,
									})

									return revokeResp
								})
							}
							//降低遭受谋略伤害
							if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
								EffectRate:     0.06,
								EffectRound:    1,
								FromWarbook:    consts.WarBookDetailType_26,
								ProduceGeneral: triggerGeneral,
							}).IsSuccess {
								util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
									revokeResp := &vo.TacticsTriggerResult{}
									revokeGeneral := params.CurrentGeneral

									util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
										Ctx:               ctx,
										General:           revokeGeneral,
										EffectType:        consts.BuffEffectType_SufferStrategyDamageDeduce,
										WarbookDetailType: consts.WarBookDetailType_26,
									})

									return revokeResp
								})
							}
						}
					}
				}
			}
		}

		return triggerResp
	})
}
