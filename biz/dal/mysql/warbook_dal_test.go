package mysql

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/util"
	"testing"
)

func TestWarBookDal_QueryWarbookList(t *testing.T) {
	conf.InitConfig()
	dal.InitMysql()
	ctx := context.Background()

	list, err := NewWarbook().QueryWarbookList(ctx, &vo.QueryWarbookCondition{
		Name: "",
	})
	if err != nil {
		fmt.Errorf("%v", err)
		t.Fail()
	}
	fmt.Printf("resp: %s", util.ToJsonString(ctx, list))
}
