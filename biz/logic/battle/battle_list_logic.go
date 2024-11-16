package battle

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/util"
)

type BattleListLogic struct {
	Ctx  context.Context
	Req  api.BattleListRequest
	Resp api.BattleListResponse
}

func NewBattleListLogic(ctx context.Context, req api.BattleListRequest) *BattleListLogic {
	return &BattleListLogic{
		Ctx: ctx,
		Req: req,
		Resp: api.BattleListResponse{
			Meta: util.BuildSuccMeta(),
		},
	}
}

func (g *BattleListLogic) Handle() (api.BattleListResponse, error) {

	list, err := mysql.NewUserBattleRecord().QueryUserBattleRecord(g.Ctx, g.Req.Uid, util.PageNoToOffset(g.Req.PageNo, g.Req.PageSize), int(g.Req.PageSize))
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryUserBattleRecord err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}

	//组合resp
	resList := make([]*api.BattleRecordInfo, 0)
	for _, record := range list {
		//反序列化
		battleRecord := &api.BattleDoResponse{}
		util.ParseJsonObj(g.Ctx, battleRecord, record.BattleRecord)

		resList = append(resList, &api.BattleRecordInfo{
			BattleResult: enum.BattleResult(battleRecord.BattleResultStatistics.FightingTeam.BattleResult),
			FightingTeam: &api.TeamBattleStatistics{
				BattleTeam: &api.BattleTeamStatistics{
					BattleGenerals: battleRecord.BattleResultStatistics.FightingTeam.BattleTeam.BattleGenerals,
				},
			},
			EnemyTeam: &api.TeamBattleStatistics{
				BattleTeam: &api.BattleTeamStatistics{
					BattleGenerals: battleRecord.BattleResultStatistics.EnemyTeam.BattleTeam.BattleGenerals,
				},
			},
		})
	}
	g.Resp = api.BattleListResponse{
		Meta:             util.BuildSuccMeta(),
		BattleRecordList: resList,
	}
	return g.Resp, nil
}
