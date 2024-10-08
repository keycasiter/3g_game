// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	api "github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/kr/pretty"
)

// BattleExecute @Summary 发起模拟对战
// @Description 发起模拟对战
// @Tags 模拟对战
// @Accept json
// @Produce json
// @Router /v1/battle/do [POST]
func BattleList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.BattleListRequest
	resp := new(api.BattleListResponse)
	resp.Meta = util.BuildSuccMeta()

	err = c.BindAndValidate(&req)
	//日志打印
	hlog.CtxInfof(ctx, "BattleList Req:%s", util.ToJsonString(ctx, req))

	if err != nil {
		hlog.CtxErrorf(ctx, "BattleList Run Err:%v", err)
		resp.Meta = util.BuildFailMetaWithMsg("未知错误")
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}

	////组合resp
	//copier.Copy(resp, serviceResp)
	//buildResponse(resp, serviceResp)

	c.JSON(hertzconsts.StatusOK, resp)
	//日志打印
	pretty.Logf("resp:%s", util.ToJsonString(ctx, resp))
}
