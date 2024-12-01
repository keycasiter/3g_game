package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 归心
// 自身准备战法发动率提升3%
type WarBookDetailType_62 struct {
}

func (w *WarBookDetailType_62) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_TacticsActiveTriggerPrepareImprove, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_62,
		ProduceGeneral: general,
	})
}
