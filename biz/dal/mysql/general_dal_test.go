package mysql

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
	"testing"
)

func TestGeneral_GetGeneralList(t *testing.T) {
	conf.InitConfig()
	dal.InitMysql()
	ctx := context.Background()

	list, err := NewGeneral().QueryGeneralList(ctx, &vo.QueryGeneralCondition{
		Id:    1,
		Limit: 1,
	})
	if err != nil {
		fmt.Errorf("%v", err)
		t.Fail()
	}
	fmt.Printf("resp: %s", util.ToJsonString(ctx, list))

	fmt.Print(cast.ToBool(list[0].IsSupportCollect))
	fmt.Print(cast.ToBool(list[0].IsSupportDynamics))
}
