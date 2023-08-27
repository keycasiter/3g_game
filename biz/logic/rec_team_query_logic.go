package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/model/po"
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
	//推荐队伍查询
	recTeamlist, err := mysql.NewRecTeam().QueryRecTeamList(g.Ctx, &vo.QueryRecTeamCondition{
		Name:   g.Req.Name,
		Group:  g.Req.Group,
		Offset: util.PageNoToOffset(g.Req.PageNo, g.Req.PageSize),
		Limit:  int(g.Req.PageSize),
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryRecTeamList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}

	generalIds := make([]int64, 0)
	tacticIds := make([]int64, 0)
	warbookIds := make([]int64, 0)
	//整理武将id
	for _, recTeam := range recTeamlist {
		generalIds = append(generalIds, util.StringToIntArray(recTeam.GeneralIds)...)
	}
	//武将信息查询
	generalQueryResp, err := NewGeneralQueryLogic(g.Ctx, api.GeneralQueryRequest{
		PageNo:   0,
		PageSize: int64(len(generalIds)),
		Ids:      generalIds,
	}).Handle()
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryGeneralList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}

	//整理战法ID
	for _, recTeam := range recTeamlist {
		tacticIds = append(tacticIds, util.StringToIntArray(recTeam.TacticIds)...)
	}
	//战法信息查询
	tacticList, err := mysql.NewTactic().QueryTacticList(g.Ctx, &vo.QueryTacticCondition{
		Ids:    tacticIds,
		Offset: 0,
		Limit:  len(tacticIds),
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryTacticList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}

	//整理兵书ID
	for _, recTeam := range recTeamlist {
		warbookIds = append(warbookIds, util.StringToIntArray(recTeam.WarbookIds)...)
	}
	//兵书信息查询
	warbookList, err := mysql.NewWarbook().QueryWarbookList(g.Ctx, &vo.QueryWarbookCondition{
		Ids:    warbookIds,
		Offset: 0,
		Limit:  len(warbookIds),
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryWarbookList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}

	//整理map
	generalMap := make(map[int64]*api.BattleGeneral, 0)
	tacticMap := make(map[int64]*po.Tactic, 0)
	warbookMap := make(map[int64]*po.Warbook, 0)
	for _, general := range generalQueryResp.GetGeneralList() {
		generalMap[general.BaseInfo.Id] = general
	}
	for _, tactic := range tacticList {
		tacticMap[tactic.Id] = tactic
	}
	for _, warbook := range warbookList {
		warbookMap[warbook.Id] = warbook
	}

	//组合resp
	resList := make([]*api.RecTeamGeneral, 0)
	for _, recTeam := range recTeamlist {
		resList = append(resList, &api.RecTeamGeneral{
			GeneralList: makeGeneralList(recTeam, generalMap, tacticMap, warbookMap),
			Name:        recTeam.Name,
			Id:          recTeam.Id,
			ArmType:     recTeam.ArmType,
		})
	}

	g.Resp.RecTeamGeneralList = resList

	return g.Resp, nil
}

func makeGeneralList(recTeam *po.RecTeam,
	generalMap map[int64]*api.BattleGeneral,
	tacticMap map[int64]*po.Tactic,
	warbookMap map[int64]*po.Warbook) []*api.BattleGeneral {

	resList := make([]*api.BattleGeneral, 0)
	generalIdsArr := util.StringToIntArray(recTeam.GeneralIds)
	generalTacticIdsArr := [][]int64{
		util.StringToIntArray(recTeam.TacticIds)[0:2],
		util.StringToIntArray(recTeam.TacticIds)[2:4],
		util.StringToIntArray(recTeam.TacticIds)[4:6],
	}
	generalWarbookIdsArr := [][]int64{
		util.StringToIntArray(recTeam.WarbookIds)[0:4],
		util.StringToIntArray(recTeam.WarbookIds)[4:8],
		util.StringToIntArray(recTeam.WarbookIds)[8:12],
	}
	generalArmTypeAbilityIdsArr := [][]int64{
		util.StringToIntArray(recTeam.ArmTypeAbilityIds)[0:1],
		util.StringToIntArray(recTeam.ArmTypeAbilityIds)[1:2],
		util.StringToIntArray(recTeam.ArmTypeAbilityIds)[2:3],
	}
	//阵容武将
	for idx, generalId := range generalIdsArr {
		if general, ok := generalMap[generalId]; ok {
			newGeneral := &api.BattleGeneral{
				BaseInfo:     general.BaseInfo,
				EquipTactics: make([]*api.Tactics, 0),
				WarBooks:     make([]*api.WarBook, 0),
				ArmsAbility:  make([]enum.ArmsAbility, 0),
				SoldierNum:   10000,
			}
			//阵容战法
			for _, tacticId := range generalTacticIdsArr[idx] {
				if tactic, okk := tacticMap[tacticId]; okk {
					newGeneral.EquipTactics = append(newGeneral.EquipTactics, &api.Tactics{
						Id:            tactic.Id,
						Name:          tactic.Name,
						TacticsSource: enum.TacticsSource(tactic.Source),
						Type:          enum.TacticsType(tactic.Type),
						Quality:       enum.TacticQuality(tactic.Quality),
					})
				}
			}
			//阵容兵书
			for _, warbookId := range generalWarbookIdsArr[idx] {
				if warbook, okk := warbookMap[warbookId]; okk {
					newGeneral.WarBooks = append(newGeneral.WarBooks, &api.WarBook{
						Id:   warbook.Id,
						Name: warbook.Name,
					})
				}
			}
			//兵种适性
			for _, armTypeAbilityId := range generalArmTypeAbilityIdsArr[idx] {
				newGeneral.ArmsAbility = append(newGeneral.ArmsAbility, enum.ArmsAbility(armTypeAbilityId))
			}
			resList = append(resList, newGeneral)
		}
	}

	return resList
}
