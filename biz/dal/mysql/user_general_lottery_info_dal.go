package mysql

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/po"
	"time"
)

func (UserGeneralLotteryInfoDal) TableName() string {
	return "user_general_lottery_info"
}

type UserGeneralLotteryInfoDal struct {
}

func NewUserGeneralLotteryInfo() *UserGeneralLotteryInfoDal {
	return &UserGeneralLotteryInfoDal{}
}

func (g *UserGeneralLotteryInfoDal) QueryUserGeneralLotteryInfo(ctx context.Context, uid string, cardPoolId int64) (*po.UserGeneralLotteryInfo, error) {
	UserGeneralLotteryInfo := &po.UserGeneralLotteryInfo{}
	conn := DataBase.Model(&po.UserGeneralLotteryInfo{})

	if err := conn.Where("uid = ? and card_pool_id = ?", uid, cardPoolId).Find(&UserGeneralLotteryInfo).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryUserGeneralLotteryInfo err:%v", err)
		return UserGeneralLotteryInfo, err
	}
	return UserGeneralLotteryInfo, nil
}

func (g *UserGeneralLotteryInfoDal) CreateUserGeneralLotteryInfo(ctx context.Context, UserGeneralLotteryInfo *po.UserGeneralLotteryInfo) error {
	conn := DataBase.Model(&po.UserGeneralLotteryInfo{})

	if err := conn.Create(UserGeneralLotteryInfo).Error; err != nil {
		hlog.CtxErrorf(ctx, "CreateUserGeneralLotteryInfo err:%v", err)
		return err
	}
	return nil
}

func (g *UserGeneralLotteryInfoDal) UpdateUserGeneralLotteryInfo(ctx context.Context, userGeneralLotteryInfo *po.UserGeneralLotteryInfo) error {
	conn := DataBase.Model(&po.UserGeneralLotteryInfo{})

	m := map[string]interface{}{
		"updated_at":         time.Now(),
		"not_hit_lev5_times": userGeneralLotteryInfo.NotHitLev5Times,
	}

	if err := conn.Where("uid = ? and card_pool_id = ?", userGeneralLotteryInfo.Uid, userGeneralLotteryInfo.CardPoolId).
		Updates(m).Error; err != nil {
		hlog.CtxErrorf(ctx, "CreateUserGeneralLotteryInfo err:%v", err)
		return err
	}
	return nil
}
