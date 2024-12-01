package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 百战
// 自带战法发动率提升3%
type WarBookDetailType_52 struct {
}

func (w *WarBookDetailType_52) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_TacticsAssaultTriggerWithSelfImprove, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_52,
		ProduceGeneral: general,
	})
}
