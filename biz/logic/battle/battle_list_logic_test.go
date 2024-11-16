package battle

import (
	"context"
	"testing"

	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
)

func TestNewBattleListLogic(t *testing.T) {
	//初始化配置文件
	conf.InitConfig()
	//初始化mysql
	mysql.InitMysql()

	NewBattleListLogic(context.Background(), api.BattleListRequest{
		Uid:      "1462451334",
		PageNo:   1,
		PageSize: 10,
	}).Handle()
}
