package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 统军
// 自身统率提升20
type WarBookDetailType_67 struct {
}

func (w *WarBookDetailType_67) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
		EffectValue:    20,
		FromWarbook:    consts.WarBookDetailType_67,
		ProduceGeneral: general,
	})
}
