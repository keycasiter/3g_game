package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 散仙
// 获得6%规避效果
type WarBookDetailType_50 struct {
}

func (w *WarBookDetailType_50) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_Evade, &vo.EffectHolderParams{
		EffectRate:     0.06,
		FromWarbook:    consts.WarBookDetailType_50,
		ProduceGeneral: general,
	})
}
