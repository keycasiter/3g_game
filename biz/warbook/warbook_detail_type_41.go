package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 临敌不乱
// 战斗第4回合，清除友军单体负面状态
type WarBookDetailType_41 struct {
}

func (w *WarBookDetailType_41) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if params.CurrentRound == consts.Battle_Round_Fourth {
			debuffGeneral := util.GetOnePairGeneralWhoHasDebuff(triggerGeneral, tacticParams)
			if debuffGeneral == nil {
				return triggerResp
			}

			util.DebuffEffectClean(ctx, debuffGeneral)
		}

		return triggerResp
	})
}
