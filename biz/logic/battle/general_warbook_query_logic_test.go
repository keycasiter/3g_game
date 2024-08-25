package battle

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/util"
	"testing"
)

func TestGeneralWarBookQueryLogic_Handle(t *testing.T) {
	conf.InitConfig()
	mysql.InitMysql()
	ctx := context.Background()

	resp, err := NewGeneralWarBookListLogic(ctx, api.GeneralWarBookListRequest{
		GeneralId: 1,
	}).Handle()
	if err != nil {
		fmt.Errorf("err:%v", err)
		return
	}
	fmt.Printf("resp:%v", util.ToJsonString(ctx, resp))
}
