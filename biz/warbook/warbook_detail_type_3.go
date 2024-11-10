package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
)

// 军形
// 武将每进阶1次，可使自身造成的伤害及治疗提高2％，受到的伤害降低2％
type WarBookDetailType_3 struct {
}

func (w *WarBookDetailType_3) Handle(ctx context.Context, general *vo.BattleGeneral) {
	for i := 0; i < int(general.Addition.GeneralStarLevel); i++ {
		//兵刃伤害
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRound:    8,
			ProduceGeneral: general,
			FromWarbook:    consts.WarBookDetailType_3,
			EffectRate:     0.02,
		})
		//谋略伤害
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
			EffectRound:    8,
			ProduceGeneral: general,
			FromWarbook:    consts.WarBookDetailType_3,
			EffectRate:     0.02,
		})
		//治疗效果
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_ResumeImprove, &vo.EffectHolderParams{
			EffectRound:    8,
			ProduceGeneral: general,
			FromWarbook:    consts.WarBookDetailType_3,
			EffectRate:     0.02,
		})
		//受到兵刃伤害
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRound:    8,
			ProduceGeneral: general,
			FromWarbook:    consts.WarBookDetailType_3,
			EffectRate:     0.02,
		})
		//受到谋略伤害
		util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRound:    8,
			ProduceGeneral: general,
			FromWarbook:    consts.WarBookDetailType_3,
			EffectRate:     0.02,
		})
	}
}
