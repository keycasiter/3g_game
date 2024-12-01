package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 枕戈坐甲
// 战斗第2-6回合，自身每回合有25%概率获得1次抵御
type WarBookDetailType_63 struct {
}

func (w *WarBookDetailType_63) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if params.CurrentRound >= consts.Battle_Round_Second && params.CurrentRound <= consts.Battle_Round_Sixth {
			if util.GenerateRate(0.25) {
				util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_Defend, &vo.EffectHolderParams{
					EffectTimes:    1,
					FromWarbook:    consts.WarBookDetailType_63,
					ProduceGeneral: triggerGeneral,
				})
			}
		}

		return triggerResp
	})
}
