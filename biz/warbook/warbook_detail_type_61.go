package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 远谋
// 自身造成持续性伤害提升6%
type WarBookDetailType_61 struct {
}

func (w *WarBookDetailType_61) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_ContinuousDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.06,
		FromWarbook:    consts.WarBookDetailType_61,
		ProduceGeneral: general,
	})
}
