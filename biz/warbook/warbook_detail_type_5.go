package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 奇正相生
// 每回合有15%概率（受智力属性加成）获得连击状态
type WarBookDetailType_5 struct {
}

func (w *WarBookDetailType_5) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		rate := 0.15 + general.BaseInfo.AbilityAttr.IntelligenceBase/100/100
		if !util.GenerateRate(rate) {
			return triggerResp
		}

		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_ContinuousAttack, &vo.EffectHolderParams{
			EffectRound: 1,
			FromWarbook: consts.WarBookDetailType_5,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}

				util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           general,
					EffectType:        consts.BuffEffectType_ContinuousAttack,
					WarbookDetailType: consts.WarBookDetailType_5,
				})

				return revokeResp
			})
		}

		return triggerResp
	})
}
