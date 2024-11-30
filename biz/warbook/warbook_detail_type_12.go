package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 执锐
// 使自身造成伤害提高3%
type WarBookDetailType_12 struct {
}

func (w *WarBookDetailType_12) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_12,
		ProduceGeneral: general,
	})
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_12,
		ProduceGeneral: general,
	})
}
