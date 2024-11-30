package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 守势
// 战斗前3回合，受到所有伤害降低5%
type WarBookDetailType_33 struct {
}

func (w *WarBookDetailType_33) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	//兵刃
	if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.05,
		EffectRound:    3,
		FromWarbook:    consts.WarBookDetailType_33,
		ProduceGeneral: general,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
				Ctx:               ctx,
				General:           revokeGeneral,
				EffectType:        consts.BuffEffectType_SufferWeaponDamageDeduce,
				WarbookDetailType: consts.WarBookDetailType_33,
			})

			return revokeResp
		})
	}

	//谋略
	if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.05,
		EffectRound:    3,
		FromWarbook:    consts.WarBookDetailType_33,
		ProduceGeneral: general,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
				Ctx:               ctx,
				General:           revokeGeneral,
				EffectType:        consts.BuffEffectType_SufferStrategyDamageDeduce,
				WarbookDetailType: consts.WarBookDetailType_33,
			})

			return revokeResp
		})
	}
}
