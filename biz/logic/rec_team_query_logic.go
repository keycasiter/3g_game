package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
)

type RecTeamQueryLogic struct {
	Ctx  context.Context
	Req  api.RecTeamQueryRequest
	Resp api.RecTeamQueryResponse
}

func NewRecTeamQueryLogic(ctx context.Context, req api.RecTeamQueryRequest) *RecTeamQueryLogic {
	return &RecTeamQueryLogic{
		Ctx: ctx,
		Req: req,
		Resp: api.RecTeamQueryResponse{
			Meta: util.BuildSuccMeta(),
		},
	}
}

func (g *RecTeamQueryLogic) Handle() (api.RecTeamQueryResponse, error) {
	list, err := mysql.NewRecTeam().QueryRecTeamList(g.Ctx, &vo.QueryRecTeamCondition{
		Name:   g.Req.Name,
		Offset: util.PageNoToOffset(g.Req.PageNo, g.Req.PageSize),
		Limit:  int(g.Req.PageSize),
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryRecTeamList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}

	//组合resp
	resList := make([]*api.RecTeamGeneral, 0)
	for _, recTeam := range list {
		resList = append(resList, &api.RecTeamGeneral{
			GeneralIds: util.StringToIntArray(recTeam.GeneralIds),
			TacticIds:  util.StringToIntArray(recTeam.TacticIds),
			WarbookIds: util.StringToIntArray(recTeam.WarbookIds),
		})
	}
	g.Resp.RecTeamGeneralList = resList

	return g.Resp, nil
}
