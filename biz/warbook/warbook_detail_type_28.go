package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 惜兵爱民
// 每回合有一定几率（50%）恢复一定兵力65%
type WarBookDetailType_28 struct {
}

func (w *WarBookDetailType_28) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if !util.GenerateRate(0.5) {
			return triggerResp
		}

		util.ResumeSoldierNum(&util.ResumeParams{
			Ctx:            ctx,
			TacticsParams:  tacticParams,
			ProduceGeneral: triggerGeneral,
			SufferGeneral:  triggerGeneral,
			ResumeNum:      cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.65),
			WarBookType:    consts.WarBookDetailType_28,
		})

		return triggerResp
	})
}
