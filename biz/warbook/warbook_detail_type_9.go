package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 蛮勇非勇
// 提高自身智力属性，其值等于自身武力属性的16%
type WarBookDetailType_9 struct {
}

func (w *WarBookDetailType_9) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	incrVal := cast.ToInt64(general.BaseInfo.AbilityAttr.IntelligenceBase * 0.16)
	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
		EffectValue:    incrVal,
		FromWarbook:    consts.WarBookDetailType_9,
		ProduceGeneral: general,
	})
}
