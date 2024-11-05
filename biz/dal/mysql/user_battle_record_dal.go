package mysql

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/po"
)

func (UserBattleRecordDal) TableName() string {
	return "user_battle_record"
}

type UserBattleRecordDal struct {
}

func NewUserBattleRecord() *UserBattleRecordDal {
	return &UserBattleRecordDal{}
}

func (g *UserBattleRecordDal) QueryUserBattleRecord(ctx context.Context, uid string, offset, limit int) ([]*po.UserBattleRecord, error) {
	userBattleRecords := make([]*po.UserBattleRecord, 0)
	conn := DataBase.Model(&po.UserBattleRecord{})

	conn.Where("uid = ?", uid)

	if offset > 0 && limit > 0 {
		conn = conn.Offset(offset).Limit(limit)
	}

	if err := conn.Find(&userBattleRecords).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryUserBattleRecord err:%v", err)
		return userBattleRecords, err
	}
	return userBattleRecords, nil
}

func (g *UserBattleRecordDal) CreateUserBattleRecord(ctx context.Context, userBattleRecord *po.UserBattleRecord) error {
	conn := DataBase.Model(&po.UserBattleRecord{})

	if err := conn.Create(userBattleRecord).Error; err != nil {
		hlog.CtxErrorf(ctx, "CreateUserBattleRecord err:%v", err)
		return err
	}
	return nil
}
