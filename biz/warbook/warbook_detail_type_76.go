package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 仙姿
// 提升自身武力、智力，提升值相当于魅力值的6%
type WarBookDetailType_76 struct {
}

func (w *WarBookDetailType_76) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
		EffectValue:    cast.ToInt64(general.BaseInfo.AbilityAttr.CharmBase * 0.06),
		FromWarbook:    consts.WarBookDetailType_76,
		ProduceGeneral: general,
	})

	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
		EffectValue:    cast.ToInt64(general.BaseInfo.AbilityAttr.CharmBase * 0.06),
		FromWarbook:    consts.WarBookDetailType_76,
		ProduceGeneral: general,
	})
}
