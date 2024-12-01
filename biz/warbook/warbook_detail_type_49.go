package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 救主
// 使我军主将受到伤害降低7%，自身受到伤害提升7%
type WarBookDetailType_49 struct {
}

func (w *WarBookDetailType_49) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	if general.IsMaster {
		return
	}

	masterGeneral := util.GetPairMasterGeneral(general, tacticParams)

	//受到兵刃提升
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.07,
		FromWarbook:    consts.WarBookDetailType_49,
		ProduceGeneral: general,
	})
	//受到谋略提升
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.07,
		FromWarbook:    consts.WarBookDetailType_49,
		ProduceGeneral: general,
	})

	//受到兵刃降低
	util.BuffEffectWrapSet(ctx, masterGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.07,
		FromWarbook:    consts.WarBookDetailType_49,
		ProduceGeneral: general,
	})
	//受到谋略降低
	util.BuffEffectWrapSet(ctx, masterGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.07,
		FromWarbook:    consts.WarBookDetailType_49,
		ProduceGeneral: general,
	})
}
