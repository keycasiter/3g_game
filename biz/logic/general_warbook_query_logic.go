package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
)

type GeneralWarBookQueryLogic struct {
	Ctx  context.Context
	Req  api.GeneralWarBookQueryRequest
	Resp api.GeneralWarBookQueryResponse
}

func NewGeneralWarBookQueryLogic(ctx context.Context, req api.GeneralWarBookQueryRequest) *GeneralWarBookQueryLogic {
	return &GeneralWarBookQueryLogic{
		Ctx: ctx,
		Req: req,
		Resp: api.GeneralWarBookQueryResponse{
			Meta: util.BuildSuccMeta(),
		},
	}
}

func (g *GeneralWarBookQueryLogic) Handle() (api.GeneralWarBookQueryResponse, error) {
	//查询武将战法
	generalWarbooklist, err := mysql.NewGeneralWarbook().QueryGeneralWarbookList(g.Ctx, &vo.QueryGeneralWarbookCondition{
		GeneralID: g.Req.GeneralId,
		OffSet:    0,
		Limit:     1,
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryGeneralWarbookList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}
	if len(generalWarbooklist) == 0 {
		g.Resp.Meta = util.BuildFailMetaWithMsg("未查询到武将战法")
		return g.Resp, nil
	}
	generalWarbook := generalWarbooklist[0]
	warbookIds := make([]int64, 0)
	warbookMap := make(map[int64]map[int64][]*po.Warbook, 0)
	util.ParseJsonObj(g.Ctx, warbookMap, generalWarbook.WarBook)
	for _, warbookTypeMap := range warbookMap {
		for _, warbookList := range warbookTypeMap {
			for _, warbook := range warbookList {
				warbookIds = append(warbookIds, warbook.Id)
			}
		}
	}

	//查询战法列表
	list, err := mysql.NewWarbook().QueryWarbookList(g.Ctx, &vo.QueryWarbookCondition{
		Ids:    warbookIds,
		Offset: 0,
		Limit:  len(warbookIds),
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryWarbookList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}
	//整理战法map
	resMap := make(map[int64]map[int64][]*api.WarBook, 0)
	for _, warbook := range list {
		if warbookMapList, ok := resMap[int64(warbook.Type)]; ok {
			if warbookList, okk := warbookMapList[int64(warbook.Level)]; okk {
				warbookList = append(warbookList, &api.WarBook{
					Id:    warbook.Id,
					Name:  warbook.Name,
					Type:  int64(warbook.Type),
					Level: int64(warbook.Level),
				})
			}
		}
	}

	g.Resp.Meta = util.BuildSuccMeta()
	g.Resp.WarBookMapList = resMap
	return g.Resp, nil
}
