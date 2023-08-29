package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/util"
	"io/ioutil"
	"net/http"
)

type GetUserOpenIdResponse struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// GetUserWxOpenId .
// @router /v1/user/get_user_wx_open_id [GET]
func GetUserWxOpenId(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetUserWxOpenIdRequest
	var resp api.GetUserWxOpenIdResponse

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(hertzconsts.StatusBadRequest, err.Error())
		return
	}

	hlog.CtxInfof(ctx, "GetUserWxOpenId Req:%s", util.ToJsonString(ctx, req))

	requestUrl := fmt.Sprintf(conf.GetConfig().Wexin.GetUserOpenIdApiTemplateUrl,
		conf.GetConfig().Wexin.AppId,
		conf.GetConfig().Wexin.Secret,
		req.Code,
	)
	urlResp, err := http.Get(requestUrl)
	if err != nil {
		// handle error
	}
	defer urlResp.Body.Close()
	body, err := ioutil.ReadAll(urlResp.Body)
	if err != nil {
		hlog.CtxErrorf(ctx, "Http Get URL:%s err:%v", requestUrl, err)
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}
	respObj := &GetUserOpenIdResponse{}
	err = json.Unmarshal(body, respObj)
	if err != nil {
		hlog.CtxErrorf(ctx, "parse Object err:%v", err)
		c.JSON(hertzconsts.StatusOK, resp)
		return
	}
	resp.OpenId = respObj.OpenId

	hlog.CtxInfof(ctx, "Http Get URL:%s Resp:%s", requestUrl, util.ToJsonString(ctx, resp))

	c.JSON(hertzconsts.StatusOK, resp)
}
