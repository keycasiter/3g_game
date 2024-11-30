package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 防备
// 自身受到伤害降低3%
type WarBookDetailType_37 struct {
}

func (w *WarBookDetailType_37) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	//兵刃
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_37,
		ProduceGeneral: general,
	})

	//谋略
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_37,
		ProduceGeneral: general,
	})
}
