package mongodb

import (
	"context"
	"github.com/keycasiter/3g_game/biz/conf"
	"github.com/keycasiter/3g_game/biz/util"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestMetadataGeneralContext_Find(t *testing.T) {
	conf.InitConfig()
	InitMongodb()

	ctx := &MetadataGeneralContext{}
	res, err := ctx.Find(context.Background(), bson.M{})
	if err != nil {
		t.Fail()
	}
	t.Logf("res:%v", util.ToJsonString(context.Background(), res))
}
