package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 驰援
// 战斗首回合，随机援护一名友军，持续1回合
type WarBookDetailType_47 struct {
}

func (w *WarBookDetailType_47) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	pairGeneral := util.GetPairOneGeneral(tacticParams, general)
	if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Intervene, &vo.EffectHolderParams{
		EffectRound:    1,
		FromWarbook:    consts.WarBookDetailType_47,
		ProduceGeneral: general,
	}).IsSuccess {
		pairGeneral.HelpByGeneral = general

		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
				Ctx:               ctx,
				General:           revokeGeneral,
				EffectType:        consts.BuffEffectType_Intervene,
				WarbookDetailType: consts.WarBookDetailType_47,
			})

			return revokeResp
		})
	}
}
