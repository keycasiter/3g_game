package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 审时度势
// 自身受到伤害提升5%，造成治疗提升12%
type WarBookDetailType_73 struct {
}

func (w *WarBookDetailType_73) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	//治疗效果
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_ResumeImprove, &vo.EffectHolderParams{
		EffectRate:     0.12,
		FromWarbook:    consts.WarBookDetailType_73,
		ProduceGeneral: general,
	})
	//受到兵刃伤害
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.05,
		FromWarbook:    consts.WarBookDetailType_73,
		ProduceGeneral: general,
	})
	//受到谋略伤害
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.05,
		FromWarbook:    consts.WarBookDetailType_73,
		ProduceGeneral: general,
	})
}
