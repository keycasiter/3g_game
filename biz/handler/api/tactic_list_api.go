// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/keycasiter/3g_game/biz/logic/battle"
	api "github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/util"
)

// TacticList @Summary 查询战法列表
// @Description 查询战法列表
// @Tags 模拟对战
// @Accept json
// @Produce json
// @Router /v1/tactic/list [GET]
func TacticList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.TacticListRequest

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(hertzconsts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxInfof(ctx, "TacticList Req:%s", util.ToJsonString(ctx, req))

	resp, err := battle.NewTacticListLogic(ctx, req).Handle()
	if err != nil {
		hlog.CtxErrorf(ctx, "TacticListLogic handle err:%v", err)
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}

	hlog.CtxInfof(ctx, "TacticList Resp:%s", util.ToJsonString(ctx, resp))

	c.JSON(hertzconsts.StatusOK, resp)
}