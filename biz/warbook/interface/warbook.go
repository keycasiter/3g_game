package _interface

import (
	"context"

	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type WarbookHandler interface {
	Handle(ctx context.Context, general *vo.BattleGeneral, TacticsParams *model.TacticsParams)
}
