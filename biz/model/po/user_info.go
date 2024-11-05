package po

import (
	"time"
)

type UserInfo struct { // nolint:maligned
	Id        int64     `gorm:"column:id" json:"id"`                 // 主键ID
	Uid       int64     `gorm:"column:uid" json:"uid"`               // 用户ID
	NickName  string    `gorm:"column:nick_name" json:"nick_name"`   // 昵称
	AvatarUrl string    `gorm:"column:avatar_url" json:"avatar_url"` // 头像
	WxOpenId  string    `gorm:"column:wx_open_id" json:"wx_open_id"` // 用户微信open_id
	Level     int       `gorm:"column:level" json:"level"`           // 用户等级
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}
