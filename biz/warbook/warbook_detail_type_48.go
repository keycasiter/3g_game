package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 励军
// 战斗中，使友军单体造成伤害提升3%
type WarBookDetailType_48 struct {
}

func (w *WarBookDetailType_48) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	pairGeneral := util.GetPairOneGeneral(tacticParams, general)

	//兵刃
	util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_48,
		ProduceGeneral: pairGeneral,
	})
	//谋略
	util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_48,
		ProduceGeneral: pairGeneral,
	})
}
