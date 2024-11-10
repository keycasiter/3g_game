package _interface

import (
	"context"

	"github.com/keycasiter/3g_game/biz/model/vo"
)

type WarbookHandler interface {
	Handle(ctx context.Context, general *vo.BattleGeneral)
}
