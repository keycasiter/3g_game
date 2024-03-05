package mysql

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/conf"
	"testing"
)

func TestJymGoodsDal_QueryJymGoodsBySkill(t *testing.T) {
	//初始化配置文件
	conf.InitConfig()
	//初始化mysql
	InitMysql()
	ctx := context.Background()

	list, err := NewJymGoods().QueryJymGoodsBySkill(ctx, []string{
		//"刮骨疗毒","杯蛇鬼车",
	})
	if err != nil {
		t.Errorf("err:%v", err)
	}
	for _, vo := range list {
		fmt.Printf("%v\n", vo.GoodsUrl)
	}
}
