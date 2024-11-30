package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 胜而益强
// 使自身在成功发动普通攻击后，有40%概率恢复一定兵力(治疗率65%)
type WarBookDetailType_8 struct {
}

func (w *WarBookDetailType_8) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}

		if !util.GenerateRate(0.4) {
			return triggerResp
		}

		resumeNum := cast.ToInt64(general.BaseInfo.AbilityAttr.IntelligenceBase * 0.65)
		util.ResumeSoldierNum(&util.ResumeParams{
			Ctx:            ctx,
			ProduceGeneral: general,
			SufferGeneral:  general,
			ResumeNum:      resumeNum,
			WarBookType:    consts.WarBookDetailType_8,
		})

		return triggerResp
	})
}
