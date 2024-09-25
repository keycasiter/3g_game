package battle

import (
	"context"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/util"
)

type RecSpecialTechListLogic struct {
	Ctx  context.Context
	Req  api.RecSpecialTechListRequest
	Resp api.RecSpecialTechListResponse
}

func NewRecSpecialTechListLogic(ctx context.Context, req api.RecSpecialTechListRequest) *RecSpecialTechListLogic {
	return &RecSpecialTechListLogic{
		Ctx: ctx,
		Req: req,
		Resp: api.RecSpecialTechListResponse{
			Meta: util.BuildSuccMeta(),
		},
	}
}

func (g *RecSpecialTechListLogic) Handle() (api.RecSpecialTechListResponse, error) {
	//todo

	return g.Resp, nil
}
