package logic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
)

type SpecialTechQueryLogic struct {
	Ctx  context.Context
	Req  api.SpecialTechQueryRequest
	Resp api.SpecialTechQueryResponse
}

func NewSpecialTechQueryLogic(ctx context.Context, req api.SpecialTechQueryRequest) *SpecialTechQueryLogic {
	return &SpecialTechQueryLogic{
		Ctx: ctx,
		Req: req,
		Resp: api.SpecialTechQueryResponse{
			Meta: util.BuildSuccMeta(),
		},
	}
}

func (s *SpecialTechQueryLogic) Handle() (api.SpecialTechQueryResponse, error) {
	//查询特技列表
	list, err := mysql.NewSpecialTech().QuerySpecialTechList(s.Ctx, &vo.QuerySpecialTechCondition{
		Id:     s.Req.Id,
		Name:   s.Req.Name,
		Type:   int(s.Req.Type),
		Offset: util.PageNoToOffset(s.Req.PageNo, s.Req.PageSize),
		Limit:  int(s.Req.PageSize),
	})
	if err != nil {
		hlog.CtxErrorf(s.Ctx, "QuerySpecialTechList err:%v", err)
		s.Resp.Meta = util.BuildFailMeta()
		return s.Resp, err
	}

	//组合resp
	resList := make([]*api.SpecialTech, 0)
	for _, tech := range list {
		resList = append(resList, &api.SpecialTech{
			Id:   tech.Id,
			Name: tech.Name,
			Type: enum.EquipType(tech.Type),
		})
	}

	s.Resp.SpecialTechList = resList
	return s.Resp, nil
}
