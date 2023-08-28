package mysql

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"testing"
)

func TestTeamDal_QueryTeamList(t *testing.T) {
	conf.InitConfig()
	InitMysql()
	ctx := context.Background()

	list, err := NewTeam().QueryTeamList(ctx, &vo.QueryTeamCondition{
		Name: "",
	})
	if err != nil {
		fmt.Errorf("%v", err)
		t.Fail()
	}
	fmt.Printf("resp: %s", util.ToJsonString(ctx, list))
}
