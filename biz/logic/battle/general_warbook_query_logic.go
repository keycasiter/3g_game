package battle

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
	"strings"
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
	generalWarbookList, err := mysql.NewGeneralWarbook().QueryGeneralWarbookList(g.Ctx, &vo.QueryGeneralWarbookCondition{
		GeneralID: g.Req.GeneralId,
		OffSet:    0,
		Limit:     1,
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryGeneralWarbookList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}
	if len(generalWarbookList) == 0 {
		g.Resp.Meta = util.BuildFailMetaWithMsg("未查询到武将战法")
		return g.Resp, nil
	}

	//单个武将兵书信息结构
	warbookHolder := make([]*TypeList, 0)
	util.ParseJsonObj(g.Ctx, &warbookHolder, generalWarbookList[0].WarBook)

	//查询全部战法信息
	warbookList, err := mysql.NewWarbook().QueryWarbookList(g.Ctx, &vo.QueryWarbookCondition{
		Offset: 0,
		Limit:  999,
	})
	if err != nil {
		hlog.CtxErrorf(g.Ctx, "QueryWarbookList err:%v", err)
		g.Resp.Meta = util.BuildFailMeta()
		return g.Resp, err
	}
	//整理兵书listToMap
	warbookMap := make(map[int64]*po.Warbook, 0)
	for _, warbook := range warbookList {
		warbookMap[warbook.Id] = warbook
	}

	//整理resp
	resMap := make(map[int64]map[int64][]*api.WarBook, 0)
	//解析整体
	for _, typeItem := range warbookHolder {
		//按层级解析
		levMap := make(map[int64][]*api.WarBook, 0)
		for _, warbookEnum := range typeItem.List {
			warbookHolderList := make([]*api.WarBook, 0)
			//拆分
			enumArr := strings.Split(warbookEnum.Enum, ",")
			for _, enumId := range enumArr {
				//转义
				if warbook, ok := warbookMap[cast.ToInt64(enumId)]; ok {
					warbookHolderList = append(warbookHolderList, &api.WarBook{
						Id:    warbook.Id,
						Name:  warbook.Name,
						Type:  int64(warbook.Type),
						Level: int64(warbook.Level),
					})
				}
			}
			levMap[cast.ToInt64(warbookEnum.Lev)] = warbookHolderList

		}
		//按req条件过滤结果
		if g.Req.WarbookType > 0 && g.Req.WarbookType == enum.WarbookType(cast.ToInt64(typeItem.Type)) {
			resMap[cast.ToInt64(typeItem.Type)] = levMap
			break
		} else {
			continue
		}

		resMap[cast.ToInt64(typeItem.Type)] = levMap
	}

	g.Resp.Meta = util.BuildSuccMeta()
	g.Resp.WarBookMapList = resMap
	return g.Resp, nil
}

type TypeList struct {
	Type string         `json:"type"`
	List []*WarbookEnum `json:"list"`
}

type WarbookEnum struct {
	Lev  int    `json:"lev"`
	Enum string `json:"enum"`
}
