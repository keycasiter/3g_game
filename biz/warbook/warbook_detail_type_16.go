package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 后发先至
// 发动需要准备的战法时，使自身获得先攻状态，持续1回合
type WarBookDetailType_16 struct {
}

func (w *WarBookDetailType_16) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_ActiveTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		currentTactic := (params.CurrentTactic).(_interface.Tactics)
		if currentTactic.IsTriggerPrepare() {
			if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_FirstAttack, &vo.EffectHolderParams{
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_16,
				ProduceGeneral: general,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_FirstAttack,
						WarbookDetailType: consts.WarBookDetailType_16,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
