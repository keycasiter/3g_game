package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 锐利
// 自身突击战法发动率提升2%
type WarBookDetailType_72 struct {
}

func (w *WarBookDetailType_72) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_TacticsAssaultTriggerImprove, &vo.EffectHolderParams{
		EffectRate:     0.02,
		FromWarbook:    consts.WarBookDetailType_72,
		ProduceGeneral: general,
	})
}
