package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 避其锐气
// 自身为主将时，战斗前2回合，获得10%规避
type WarBookDetailType_29 struct {
}

func (w *WarBookDetailType_29) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	if !general.IsMaster {
		return
	}

	if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_Evade, &vo.EffectHolderParams{
		EffectRate:     0.1,
		EffectRound:    2,
		FromWarbook:    consts.WarBookDetailType_29,
		ProduceGeneral: general,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
				Ctx:               ctx,
				General:           revokeGeneral,
				EffectType:        consts.BuffEffectType_Evade,
				WarbookDetailType: consts.WarBookDetailType_29,
			})

			return revokeResp
		})
	}
}
