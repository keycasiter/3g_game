package battle

import (
	"context"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/util"
)

type RecWarBookListLogic struct {
	Ctx  context.Context
	Req  api.RecWarBookListRequest
	Resp api.RecWarBookListResponse
}

func NewRecWarBookListLogic(ctx context.Context, req api.RecWarBookListRequest) *RecWarBookListLogic {
	return &RecWarBookListLogic{
		Ctx: ctx,
		Req: req,
		Resp: api.RecWarBookListResponse{
			Meta: util.BuildSuccMeta(),
		},
	}
}

func (g *RecWarBookListLogic) Handle() (api.RecWarBookListResponse, error) {
	//todo

	return g.Resp, nil
}
