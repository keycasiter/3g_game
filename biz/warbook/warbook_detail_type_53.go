package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 疾战突围
// 发动需要准备的战法时，有20%概率跳过
type WarBookDetailType_53 struct {
}

func (w *WarBookDetailType_53) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerTactic := params.CurrentTactic.(_interface.Tactics)
		triggerResp := &vo.TacticsTriggerResult{}

		if triggerTactic.IsTriggerPrepare() {
			if util.GenerateRate(0.2) {
				triggerTactic.SetTriggerPrepare(true)
			}
		}
		return triggerResp
	})
}
