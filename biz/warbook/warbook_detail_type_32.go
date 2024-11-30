package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 守而有道
// 自身受到谋略伤害降低5%，受统率影响
type WarBookDetailType_32 struct {
}

func (w *WarBookDetailType_32) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	rate := 0.05 + general.BaseInfo.AbilityAttr.CommandBase/100/100
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     rate,
		FromWarbook:    consts.WarBookDetailType_32,
		ProduceGeneral: general,
	})
}
