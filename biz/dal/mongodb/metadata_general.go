package mongodb

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/po"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	METADATA_GENERAL_COLLECTION = "metadata_general"
)

type MetadataGeneralContext struct {
}

func NewMetadataGeneralContext() *MetadataGeneralContext {
	return &MetadataGeneralContext{}
}

func (*MetadataGeneralContext) Insert(ctx context.Context, m *po.MetadataGeneral) error {
	objId, err := Mongodb3gGame.Collection(METADATA_GENERAL_COLLECTION).InsertOne(ctx, &m)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s insert err:%v", METADATA_GENERAL_COLLECTION, err)
		return err
	}
	hlog.CtxInfof(ctx, "%s insert succ，objId:%v", METADATA_GENERAL_COLLECTION, objId)
	return nil
}

func (*MetadataGeneralContext) Find(ctx context.Context, m bson.M) (*po.MetadataGeneral, error) {
	var result *po.MetadataGeneral
	err := Mongodb3gGame.Collection(METADATA_GENERAL_COLLECTION).FindOne(ctx, m).Decode(&result)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s find err:%v", METADATA_GENERAL_COLLECTION, err)
		return result, err
	}
	hlog.CtxInfof(ctx, "%s find succ，objId:%v", METADATA_GENERAL_COLLECTION, result.Id)
	return result, nil
}

func (*MetadataGeneralContext) Find(ctx context.Context, m bson.M) (*po.MetadataGeneral, error) {
	var result []*po.MetadataGeneral
	err := Mongodb3gGame.Collection(METADATA_GENERAL_COLLECTION).Find(ctx, m).Decode(&result)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s find err:%v", METADATA_GENERAL_COLLECTION, err)
		return result, err
	}
	hlog.CtxInfof(ctx, "%s find succ，objId:%v", METADATA_GENERAL_COLLECTION, result.Id)
	return result, nil
}
