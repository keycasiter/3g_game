package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 久战
// 获得3%攻心和倒戈
type WarBookDetailType_60 struct {
}

func (w *WarBookDetailType_60) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_AttackHeart, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_60,
		ProduceGeneral: general,
	})
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_Defection, &vo.EffectHolderParams{
		EffectRate:     0.03,
		FromWarbook:    consts.WarBookDetailType_60,
		ProduceGeneral: general,
	})
}
