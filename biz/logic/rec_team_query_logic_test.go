package logic

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"testing"
)

func TestRecTeamQueryLogic_Handle(t *testing.T) {
	//初始化配置文件
	conf.InitConfig()
	//初始化mysql
	mysql.InitMysql()

	resp, err := NewRecTeamQueryLogic(context.Background(), api.RecTeamQueryRequest{
		PageNo:   1,
		PageSize: 10,
	}).Handle()
	if err != nil {
		fmt.Errorf("err:%v", err)
		t.Fail()
	}
	fmt.Printf("resp:%v", resp)
}
