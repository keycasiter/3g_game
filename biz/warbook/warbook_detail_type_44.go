package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 分而疾战
// 自身受到伤害降低3.5%，受速度影响
type WarBookDetailType_44 struct {
}

func (w *WarBookDetailType_44) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	rate := 0.035 + general.BaseInfo.AbilityAttr.SpeedBase/100/100

	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     rate,
		FromWarbook:    consts.WarBookDetailType_44,
		ProduceGeneral: general,
	})
}
