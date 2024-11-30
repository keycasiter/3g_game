package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 将威
// 提高自身2.5%会心和奇谋几率
type WarBookDetailType_23 struct {
}

func (w *WarBookDetailType_23) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	//会心
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_EnhanceWeapon, &vo.EffectHolderParams{
		EffectRate:     0.025,
		FromWarbook:    consts.WarBookDetailType_23,
		ProduceGeneral: general,
	})
	//奇谋
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_EnhanceStrategy, &vo.EffectHolderParams{
		EffectRate:     0.025,
		FromWarbook:    consts.WarBookDetailType_23,
		ProduceGeneral: general,
	})
}
