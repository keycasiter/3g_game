package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 锤炼
// 自身受到持续性伤害降低10%
type WarBookDetailType_66 struct {
}

func (w *WarBookDetailType_66) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferContinuousDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.1,
		FromWarbook:    consts.WarBookDetailType_66,
		ProduceGeneral: general,
	})
}
