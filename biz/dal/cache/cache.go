package cache

import (
	"context"
	"fmt"

	"github.com/keycasiter/3g_game/biz/dal/mysql"
	"github.com/keycasiter/3g_game/biz/model/po"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

// 武将信息
var CacheGeneralMap = make(map[int64]*po.General, 0)

// 战法信息
var CacheTacticMap = make(map[int64]*po.Tactic, 0)

// 兵书信息
var CacheWarBookMap = make(map[int64]*po.Warbook, 0)

func InitCache() {
	ctx := context.Background()

	//武将信息
	initGeneralCache(ctx)
	//战法信息
	initTacticCache(ctx)
	//兵书信息
	initWarBookCache(ctx)
}

func initGeneralCache(ctx context.Context) {
	generals, err := mysql.NewGeneral().QueryGeneralList(ctx, &vo.QueryGeneralCondition{
		Offset: 0,
		Limit:  10000,
	})
	if err != nil {
		panic(any(fmt.Sprintf("init cache [general] err:%v", err)))
	}
	for _, general := range generals {
		CacheGeneralMap[general.Id] = general
	}
}

func initWarBookCache(ctx context.Context) {
	warbooks, err := mysql.NewWarbook().QueryWarbookList(ctx, &vo.QueryWarbookCondition{
		Offset: 0,
		Limit:  10000,
	})
	if err != nil {
		panic(any(fmt.Sprintf("init cache [tactic] err:%v", err)))
	}
	for _, warbook := range warbooks {
		CacheWarBookMap[warbook.Id] = warbook
	}
}
func initTacticCache(ctx context.Context) {
	tactics, err := mysql.NewTactic().QueryTacticList(ctx, &vo.QueryTacticCondition{
		Offset: 0,
		Limit:  10000,
	})
	if err != nil {
		panic(any(fmt.Sprintf("init cache [tactic] err:%v", err)))
	}
	for _, tactic := range tactics {
		CacheTacticMap[tactic.Id] = tactic
	}
}
