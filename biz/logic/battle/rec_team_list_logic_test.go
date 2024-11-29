package battle

import (
	"context"
	"testing"

	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
)

func TestNewRecTeamListLogic(t *testing.T) {
	conf.InitConfig()
	mysql.InitMysql()
	//{"Name":"","Group":1,"ArmType":2,"PageNo":1,"PageSize":10}
	handler := NewRecTeamListLogic(context.Background(), api.RecTeamListRequest{
		Name:     "",
		Group:    1,
		ArmType:  2,
		PageNo:   1,
		PageSize: 10,
	})
	handler.Handle()
}
