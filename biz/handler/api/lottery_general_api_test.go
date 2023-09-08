package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal/cache"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"testing"
)

func TestGeneralLotteryQuery(t *testing.T) {
	//初始化配置文件
	conf.InitConfig()
	//初始化mysql
	mysql.InitMysql()
	cache.InitCache()

	ctx := context.Background()
	req := &app.RequestContext{}
	req.Request.SetFormData(map[string]string{})

	GeneralLotteryInfoQuery(ctx, req)
	fmt.Printf("resp:%s", req.Response.Body())
}
