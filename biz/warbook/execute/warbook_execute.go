package execute

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

type WarBookExecutor struct {
	Ctx     context.Context
	General *vo.BattleGeneral
}

func NewWarBookExecutor(ctx context.Context, general *vo.BattleGeneral) *WarBookExecutor {
	return &WarBookExecutor{Ctx: ctx, General: general}
}

func (w *WarBookExecutor) Execute() {
	for _, warbook := range w.General.WarBooks {
		if handler, ok := WarbookHandlerMap[consts.WarBookDetailType(warbook.Id)]; ok {
			handler.Handle(w.Ctx, w.General)
		}
	}
}
