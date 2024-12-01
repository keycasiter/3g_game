package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 神清气净
// 自身累计受到2次控制状态时候，清除我军单体负面状态，仅可触发1次
type WarBookDetailType_58 struct {
}

func (w *WarBookDetailType_58) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	cnt := 0

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferDebuffEffectEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.IsControlDeBuffEffect(params.DebuffEffect) {
			cnt++

			if cnt == 2 {
				pairGeneral := util.GetOnePairGeneralWhoHasDebuff(triggerGeneral, tacticParams)
				if pairGeneral == nil {
					return triggerResp
				}
				util.DebuffEffectClean(ctx, pairGeneral)
			}
		}

		return triggerResp
	})
}
