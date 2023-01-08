package logic

import (
	"context"
	"github.com/keycasiter/3g_game/biz/dal/mongodb"
	"github.com/keycasiter/3g_game/biz/model/api"
	"github.com/keycasiter/3g_game/biz/model/common"
	"github.com/keycasiter/3g_game/biz/model/po"
	"go.mongodb.org/mongo-driver/bson"
)

type MetadataGeneralLogic struct {
	Ctx context.Context
	Req api.MetadataGeneralPageableListRequest
}

func NewMetadataGeneralLogic(ctx context.Context, req api.MetadataGeneralPageableListRequest) *MetadataGeneralLogic {
	return &MetadataGeneralLogic{Ctx: ctx, Req: req}
}

func (l *MetadataGeneralLogic) MetadataGeneralPageableList() ([]*common.MetadataGeneral, error) {
	&mongodb.NewMetadataGeneralContext().Find(l.Ctx, bson.M{})
	return nil, nil
}
