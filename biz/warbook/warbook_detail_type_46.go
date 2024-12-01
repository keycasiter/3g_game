package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 速战
// 战斗中，提高自身24点速度
type WarBookDetailType_46 struct {
}

func (w *WarBookDetailType_46) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrSpeed, &vo.EffectHolderParams{
		EffectValue:    24,
		FromWarbook:    consts.WarBookDetailType_46,
		ProduceGeneral: general,
	})
}
