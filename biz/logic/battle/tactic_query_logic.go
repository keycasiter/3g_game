package battle

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
)

type TacticListLogic struct {
	Ctx  context.Context
	Req  api.TacticListRequest
	Resp api.TacticListResponse
}

func NewTacticListLogic(ctx context.Context, req api.TacticListRequest) *TacticListLogic {
	return &TacticListLogic{
		Ctx: ctx,
		Req: req,
		Resp: api.TacticListResponse{
			Meta: util.BuildSuccMeta(),
		},
	}
}

func (g *TacticListLogic) Handle() (api.TacticListResponse, error) {
	//查询战法列表
	sources := make([]int32, 0)
	for _, source := range g.Req.Sources {
		sources = append(sources, int32(source))
	}
	list, err := mysql.NewTactic().QueryTacticList(g.Ctx, &vo.QueryTacticCondition{
		Id:      g.Req.Id,
		Name:    g.Req.Name,
		Quality: int32(g.Req.Quality),
		Source:  int32(g.Req.Source),
		Sources: sources,
		Type:    int32(g.Req.Type),
		Offset:  util.PageNoToOffset(g.Req.PageNo, g.Req.PageSize),
		Limit:   int(g.Req.PageSize),
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryTacticList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}

	//组合resp
	resList := make([]*api.Tactics, 0)
	for _, tactic := range list {
		resList = append(resList, &api.Tactics{
			Id:            tactic.Id,
			Name:          tactic.Name,
			TacticsSource: enum.TacticsSource(tactic.Source),
			Type:          enum.TacticsType(tactic.Type),
			Quality:       enum.TacticQuality(tactic.Quality),
		})
	}
	g.Resp.TacticList = resList

	return g.Resp, nil
}
