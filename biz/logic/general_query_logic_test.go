package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/dal"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/enum"
	"github.com/keycasiter/3g_game/biz/util"
	"testing"
)

func TestGeneralQueryLogic_Handle(t *testing.T) {
	conf.InitConfig()
	dal.InitMysql()

	ctx := context.Background()
	gender := enum.Gender_Female
	resp := NewGeneralQueryLogic(ctx, api.GeneralQueryRequest{
		Gender: &gender,
	}).Handle()
	fmt.Printf("resp:%s", util.ToJsonString(ctx, resp))
}

func TestTags(t *testing.T) {
	tags := "[\"1\"]"
	arr := []string{}
	json.Unmarshal([]byte(tags), &arr)
	fmt.Printf("%v", arr)
}
