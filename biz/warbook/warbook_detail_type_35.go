package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 铁甲
// 受到普通攻击和突击战法伤害降低6%
type WarBookDetailType_35 struct {
}

func (w *WarBookDetailType_35) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	//普通攻击
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferGeneralAttackDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.06,
		FromWarbook:    consts.WarBookDetailType_35,
		ProduceGeneral: general,
	})
}
