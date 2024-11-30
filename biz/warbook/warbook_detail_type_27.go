package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 严阵以待
// 战斗前2回合，为友军群体2人分担15%伤害
type WarBookDetailType_27 struct {
}

func (w *WarBookDetailType_27) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	pairGenerals := util.GetPairGeneralsNotSelf(tacticParams, general)

	for _, pairGeneral := range pairGenerals {
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_ShareResponsibilityFor, &vo.EffectHolderParams{
			EffectRate:                      0.15,
			EffectRound:                     2,
			FromWarbook:                     consts.WarBookDetailType_27,
			ShareResponsibilityForByGeneral: general,
			ProduceGeneral:                  general,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
					Ctx:               ctx,
					General:           revokeGeneral,
					EffectType:        consts.BuffEffectType_ShareResponsibilityFor,
					WarbookDetailType: consts.WarBookDetailType_27,
				})

				return revokeResp
			})
		}
	}
}
