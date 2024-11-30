package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 诱敌之策
// 自身受到兵刃伤害和谋略伤害提高16%，友军造成的兵刃伤害和谋略伤害提高5%
type WarBookDetailType_42 struct {
}

func (w *WarBookDetailType_42) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	//兵刃
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.16,
		FromWarbook:    consts.WarBookDetailType_42,
		ProduceGeneral: general,
	})
	//谋略
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.16,
		FromWarbook:    consts.WarBookDetailType_42,
		ProduceGeneral: general,
	})

	pairGenerals := util.GetPairGeneralsNotSelf(tacticParams, general)
	for _, pairGeneral := range pairGenerals {
		//兵刃
		util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.05,
			FromWarbook:    consts.WarBookDetailType_42,
			ProduceGeneral: general,
		})
		//谋略
		util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.05,
			FromWarbook:    consts.WarBookDetailType_42,
			ProduceGeneral: general,
		})
	}
}
