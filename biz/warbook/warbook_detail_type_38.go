package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 勇毅
// 受到克制兵种伤害减少15%
type WarBookDetailType_38 struct {
}

func (w *WarBookDetailType_38) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferGramArmDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.15,
		FromWarbook:    consts.WarBookDetailType_38,
		ProduceGeneral: general,
	})
}
