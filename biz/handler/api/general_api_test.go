package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal"
	"testing"
)

func TestGeneralQuery(t *testing.T) {
	//初始化配置文件
	conf.InitConfig()
	//初始化mysql
	dal.InitMysql()

	ctx := context.Background()
	req := &app.RequestContext{}
	req.Request.SetFormData(map[string]string{})

	GeneralQuery(ctx, req)
	fmt.Printf("resp:%s", req.Response.Body())
}
