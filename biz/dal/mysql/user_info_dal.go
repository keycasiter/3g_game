package mysql

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/model/po"
)

func (UserInfoDal) TableName() string {
	return "user_info"
}

type UserInfoDal struct {
}

func NewUserInfo() *UserInfoDal {
	return &UserInfoDal{}
}

func (g *UserInfoDal) QueryUserInfo(ctx context.Context, wxOpenId string) (*po.UserInfo, error) {
	userInfo := &po.UserInfo{}
	conn := DataBase.Model(&po.UserInfo{})

	if err := conn.Where("wx_open_id = ?", wxOpenId).Find(&userInfo).Error; err != nil {
		hlog.CtxErrorf(ctx, "QueryUserInfo err:%v", err)
		return userInfo, err
	}
	return userInfo, nil
}

func (g *UserInfoDal) CheckUserInfo(ctx context.Context, wxOpenId string) (bool, error) {
	userInfo := &po.UserInfo{}
	conn := DataBase.Model(&po.UserInfo{})

	if err := conn.Where("wx_open_id = ?", wxOpenId).Find(&userInfo).Error; err != nil {
		hlog.CtxErrorf(ctx, "CheckUserInfo err:%v", err)
		return false, err
	}
	return userInfo.WxOpenId != "", nil
}

func (g *UserInfoDal) CreateUserInfo(ctx context.Context, userInfo *po.UserInfo) error {
	conn := DataBase.Model(&po.UserInfo{})

	if err := conn.Create(userInfo).Error; err != nil {
		hlog.CtxErrorf(ctx, "CreateUserInfo err:%v", err)
		return err
	}
	return nil
}
