package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 文韬
// 使自身造成的谋略伤害提高2.8% 受智力影响
type WarBookDetailType_13 struct {
}

func (w *WarBookDetailType_13) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	rate := 0.028 + general.BaseInfo.AbilityAttr.IntelligenceBase/100/100
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRate:     rate,
		FromWarbook:    consts.WarBookDetailType_13,
		ProduceGeneral: general,
	})
}
