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
	res, err := ctx.Find(context.Background(), bson.M{"name": "姜维"})
	if err != nil {
		t.Fail()
	}
	t.Logf("res:%v", util.ToJsonString(context.Background(), res))
}

func TestMetadataGeneralContext_FindAll(t *testing.T) {
	conf.InitConfig()
	InitMongodb()

	ctx := &MetadataGeneralContext{}
	res, err := ctx.FindAll(context.Background(), bson.M{})
	if err != nil {
		t.Fail()
	}
	t.Logf("res:%v", util.ToJsonString(context.Background(), res))
}

func TestMetadataGeneralContext_FindAll_Condition(t *testing.T) {
	conf.InitConfig()
	InitMongodb()

	arr := []string{"姜维", "庞统", "诸葛亮"}
	ctx := &MetadataGeneralContext{}
	res, err := ctx.FindAll(context.Background(), bson.M{"name": bson.M{"$in": arr}})
	if err != nil {
		t.Fail()
	}
	t.Logf("res:%v", util.ToJsonString(context.Background(), res))
}
