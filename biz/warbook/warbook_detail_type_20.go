package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 顺应天时
// 施加持续性状态时，有35%的概率增加1回合
type WarBookDetailType_20 struct {
}

func (w *WarBookDetailType_20) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	//正面效果
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BuffEffect, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		effectType := params.BuffEffect

		if util.IsContinuousBuffEffectType(effectType) {
			if util.GenerateRate(0.35) {
				params.EffectHolderParams.EffectRound++
			}
		}

		return triggerResp
	})

	//负面效果
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_DebuffEffect, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		effectType := params.DebuffEffect

		if util.IsContinuousDeffEffectType(effectType) {
			if util.GenerateRate(0.35) {
				params.EffectHolderParams.EffectRound++
			}
		}

		return triggerResp
	})
}
