package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/util"
)

type GetUserOpenIdResponse struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// UserInfoDetail @Summary 用户登录
// @Description 用户登录
// @Tags 用户
// @Accept json
// @Produce json
// @Router /v1/user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserLoginRequest
	var resp api.UserLoginResponse
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

	//2.注册逻辑
	isExist, userInfo, err := mysql.NewUserInfo().CheckUserInfo(ctx, respObj.OpenId)
	if err != nil {
		hlog.CtxErrorf(ctx, "CheckUserInfo err:%v", err)
		resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("检测用户失败 err:%v", err))
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}
	//不存在则注册
	if !isExist {
		uid := util.GenerateUID()
		nowTime := time.Now()
		err := mysql.NewUserInfo().CreateUserInfo(ctx, &po.UserInfo{
			Uid:       uid,
			NickName:  req.NickName,
			AvatarUrl: req.AvatarUrl,
			WxOpenId:  respObj.OpenId,
			Level:     int(consts.UserLevel_Common),
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
		})
		if err != nil {
			hlog.CtxErrorf(ctx, "CreateUserInfo err:%v", err)
			resp.Meta = util.BuildFailMetaWithMsg(fmt.Sprintf("创建用户失败 err:%v", err))
			c.JSON(hertzconsts.StatusOK, resp)
			return
		}

		resp.UserInfo = &api.UserInfo{}
		resp.UserInfo.Uid = uid
		resp.UserInfo.WxOpenId = respObj.OpenId
		resp.UserInfo.NickName = req.NickName
		resp.UserInfo.AvatarUrl = req.AvatarUrl
		resp.UserInfo.Level = int64(consts.UserLevel_Common)
	} else {
		//存在

		resp.UserInfo = &api.UserInfo{}
		resp.UserInfo.Uid = userInfo.Uid
		resp.UserInfo.WxOpenId = userInfo.WxOpenId
		resp.UserInfo.NickName = userInfo.NickName
		resp.UserInfo.AvatarUrl = userInfo.AvatarUrl
		resp.UserInfo.Level = int64(userInfo.Level)
	}

	c.JSON(hertzconsts.StatusOK, resp)
}
