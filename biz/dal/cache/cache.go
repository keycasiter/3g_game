package cache

import (
	"context"
	"fmt"
	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

//武将信息
var CacheGeneralMap = make(map[int64]*po.General, 0)

func InitCache() {
	ctx := context.Background()

	//武将信息
	initGeneralCache(ctx)
}

func initGeneralCache(ctx context.Context) {
	generals, err := mysql.NewGeneral().QueryGeneralList(ctx, &vo.QueryGeneralCondition{
		Offset: 0,
		Limit:  10000,
	})
	if err != nil {
		panic(fmt.Sprintf("init cache [general] err:%v", err))
	}
	for _, general := range generals {
		CacheGeneralMap[general.Id] = general
	}
}
