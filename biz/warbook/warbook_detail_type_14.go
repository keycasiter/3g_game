package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 藏刀
// 自身突击战法发动率降低%5，突击战法发动率提高3%
type WarBookDetailType_14 struct {
}

func (w *WarBookDetailType_14) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_TacticsAssaultTriggerImprove, &vo.EffectHolderParams{
		TriggerRate:    0.3,
		FromWarbook:    consts.WarBookDetailType_14,
		ProduceGeneral: general,
	})
}
