package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 静心
// 自身受到主动战法伤害降低7%
type WarBookDetailType_34 struct {
}

func (w *WarBookDetailType_34) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferActiveTacticDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.07,
		FromWarbook:    consts.WarBookDetailType_34,
		ProduceGeneral: general,
	})
}
