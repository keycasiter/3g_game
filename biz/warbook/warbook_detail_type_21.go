package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 鬼谋
// 提高4.5%主动战法伤害
type WarBookDetailType_21 struct {
}

func (w *WarBookDetailType_21) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_TacticsActiveDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.045,
		FromWarbook:    consts.WarBookDetailType_21,
		ProduceGeneral: general,
	})
}
