package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 三里而还
// 战斗第三回合，自身获得反击状态（伤害率50%），持续3回合
type WarBookDetailType_31 struct {
}

func (w *WarBookDetailType_31) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if params.CurrentRound == consts.Battle_Round_Third {
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_StrikeBack, &vo.EffectHolderParams{
				EffectRate:     0.5,
				EffectRound:    3,
				FromWarbook:    consts.WarBookDetailType_31,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_StrikeBack,
						WarbookDetailType: consts.WarBookDetailType_31,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
