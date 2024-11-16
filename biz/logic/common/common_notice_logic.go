package common

import (
	"context"

	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/util"
)

// 公告
type CommonNoticeLogic struct {
	Ctx   context.Context
	Req   *api.CommonNoticeRequest
	Resp  *api.CommonNoticeResponse
	Funcs []func()
	Err   error
}

func NewCommonNoticeLogic(ctx context.Context, req *api.CommonNoticeRequest) *CommonNoticeLogic {
	g := &CommonNoticeLogic{
		Ctx: ctx,
		Req: req,
	}
	//组合方法
	g.Funcs = []func(){
		//查询公告
		g.queryCommonNotice,
	}
	return g
}

func (g *CommonNoticeLogic) Run() (*api.CommonNoticeResponse, error) {
	for _, f := range g.Funcs {
		f()
		if g.Err != nil {
			return nil, g.Err
		}
	}
	return g.Resp, nil
}

func (c *CommonNoticeLogic) queryCommonNotice() {
	c.Resp = &api.CommonNoticeResponse{
		Meta: util.BuildSuccMeta(),
		NoticeList: []*api.NoticeVo{
			{
				Content: "内测中，欢迎多多提出建议和反馈～",
				Uri:     "",
			},
		},
	}
}
