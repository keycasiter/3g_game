package redis

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/spf13/cast"
)

var RedisClient *redis.Client

func InitRedis() {
	ctx := context.Background()
	hlog.CtxInfof(ctx, "Redis初始化...")

	config := conf.GetConfig()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       cast.ToInt(config.Redis.DB),
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "Redis初始化失败... err:%v", err)
		panic(err)
	} else {
		hlog.CtxInfof(ctx, "Redis初始化成功...")
	}
}
