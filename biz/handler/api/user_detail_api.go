package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/util"
	"io/ioutil"
	"net/http"
)

// UserInfoDetail @Summary 查询用户信息详情
// @Description 查询用户信息详情
// @Tags 用户
// @Accept json
// @Produce json
// @Router /v1/user/detail [GET]
func UserInfoDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoDetailRequest
	var resp api.UserInfoDetailResponse
	resp.Meta = util.BuildSuccMeta()

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(hertzconsts.StatusBadRequest, err.Error())
		return
	}

	//1.获取登录用户的openId
	hlog.CtxInfof(ctx, "GetUserWxOpenId Req:%s", util.ToJsonString(ctx, req))
	requestUrl := fmt.Sprintf(conf.GetConfig().Wexin.GetUserOpenIdApiTemplateUrl,
		conf.GetConfig().Wexin.AppId,
		conf.GetConfig().Wexin.Secret,
		req.Code,
	)
	urlResp, err := http.Get(requestUrl)
	if err != nil {
		hlog.CtxErrorf(ctx, "Http Get URL:%s err:%v", requestUrl, err)
		resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("err:%v", err))
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}
	defer urlResp.Body.Close()
	body, err := ioutil.ReadAll(urlResp.Body)
	if err != nil {
		hlog.CtxErrorf(ctx, "Http Get URL:%s err:%v", requestUrl, err)
		resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("err:%v", err))
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}
	respObj := &GetUserOpenIdResponse{}
	err = json.Unmarshal(body, respObj)
	if err != nil {
		hlog.CtxErrorf(ctx, "parse Object err:%v", err)
		resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("parse Object err:%v", err))
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}
	hlog.CtxInfof(ctx, "GetUserWxOpenId Resp:%s", util.ToJsonString(ctx, respObj))
	if respObj.OpenId == "" {
		hlog.CtxErrorf(ctx, "openId is empty")
		resp.Meta = util.BuildFailMetaWithMsg("微信openId为空")
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}

	//2.用户信息查询
	userInfo, err := mysql.NewUserInfo().QueryUserInfo(ctx, respObj.OpenId)
	if err != nil {
		hlog.CtxErrorf(ctx, "QueryUserInfo err:%v", err)
		resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("查询用户信息失败 err:%v", err))
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}

	//组合resp
	resp.UserInfo.Uid = userInfo.Uid
	resp.UserInfo.WxOpenId = userInfo.WxOpenId
	resp.UserInfo.NickName = userInfo.NickName
	resp.UserInfo.AvatarUrl = userInfo.AvatarUrl
	resp.UserInfo.Level = int64(userInfo.Level)

	c.JSON(hertzconsts.StatusOK, resp)
}
