package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 无功而励
// 每回合行动时，若自身处于缴械、计穷、震慑或混乱状态，则恢复一定兵力（恢复率80%）
type WarBookDetailType_40 struct {
}

func (w *WarBookDetailType_40) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.DeBuffEffectContains(triggerGeneral, consts.DebuffEffectType_CancelWeapon) ||
			util.DeBuffEffectContains(triggerGeneral, consts.DebuffEffectType_NoStrategy) ||
			util.DeBuffEffectContains(triggerGeneral, consts.DebuffEffectType_Awe) ||
			util.DeBuffEffectContains(triggerGeneral, consts.DebuffEffectType_Chaos) {
			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  tacticParams,
				ProduceGeneral: triggerGeneral,
				SufferGeneral:  triggerGeneral,
				ResumeNum:      cast.ToInt64(cast.ToFloat64(triggerGeneral.LossSoldierNum) * 0.8),
				WarBookType:    consts.WarBookDetailType_40,
			})
		}

		return triggerResp
	})
}
