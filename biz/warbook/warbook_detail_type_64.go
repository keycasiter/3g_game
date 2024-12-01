package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 三军之众
// 战斗第2-4回合，自身获得急救（25%，受智力影响）
type WarBookDetailType_64 struct {
}

func (w *WarBookDetailType_64) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if params.CurrentRound >= consts.Battle_Round_Second && params.CurrentRound <= consts.Battle_Round_Fourth {
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_EmergencyTreatment, &vo.EffectHolderParams{
				EffectRate:     0.25 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100,
				FromWarbook:    consts.WarBookDetailType_64,
				ProduceGeneral: triggerGeneral,
			})
		}

		return triggerResp
	})
}
