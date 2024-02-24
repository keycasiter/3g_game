package mysql

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/po"
)

func (JymGoodsDal) TableName() string {
	return "jym_goods"
}

type JymGoodsDal struct {
}

func NewJymGoods() *JymGoodsDal {
	return &JymGoodsDal{}
}

func (g *JymGoodsDal) BatchSaveJymGoods(ctx context.Context, pos []*po.JymGoods) error {
	err := DataBase.Model(&po.JymGoods{}).CreateInBatches(pos, 10).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "batchSaveJymGoods CreateInBatches err:%v", err)
		return err
	}
	return nil
}
