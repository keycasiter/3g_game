package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 武略
// 提高8%普通攻击伤害
type WarBookDetailType_10 struct {
}

func (w *WarBookDetailType_10) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_GeneralAttackDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.08,
		FromWarbook:    consts.WarBookDetailType_10,
		ProduceGeneral: general,
	})
}
