package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 刚柔
// 受到治疗时，提高9%的治疗效果
type WarBookDetailType_36 struct {
}

func (w *WarBookDetailType_36) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferResumeImprove, &vo.EffectHolderParams{
		EffectRate:     0.09,
		FromWarbook:    consts.WarBookDetailType_36,
		ProduceGeneral: general,
	})
}
