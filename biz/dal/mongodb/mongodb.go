package mongodb

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var MongodbClient *mongo.Client
var Mongodb3gGame *mongo.Database

//pool连接池模式
func InitMongodb() {

	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.Config.Mongodb.ConnectTimeout))
	defer cancel()
	// 通过传进来的uri连接相关的配置
	clientOptions := options.Client().ApplyURI(conf.Config.Mongodb.Uri)
	// 发起连接
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		hlog.CtxErrorf(ctx, "mongodb connect err:%v", err)
		panic(err)
	}
	// 判断服务可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		hlog.CtxErrorf(ctx, "mongodb Ping err:%v", err)
		panic(err)
	}
	MongodbClient = client
	Mongodb3gGame = client.Database(conf.Config.Mongodb.Database)
}
