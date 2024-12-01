package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 洞若观火
// 自身受到属性降低时，有20%概率免疫
type WarBookDetailType_57 struct {
}

func (w *WarBookDetailType_57) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferDebuffEffect, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		if params.DebuffEffect == consts.DebuffEffectType_DecrCommand ||
			params.DebuffEffect == consts.DebuffEffectType_DecrIntelligence ||
			params.DebuffEffect == consts.DebuffEffectType_DecrSpeed ||
			params.DebuffEffect == consts.DebuffEffectType_DecrForce {

			if util.GenerateRate(0.2) {
				triggerResp.IsTerminate = true
			}
		}

		return triggerResp
	})
}
