package po

import "time"

type UserGeneralLotteryInfo struct {
	Id              int64     `gorm:"column:id" json:"id"`                                 //主键
	Uid             string    `gorm:"column:uid" json:"uid"`                               //用户id
	CardPoolId      int64     `gorm:"column:card_pool_id" json:"card_pool_id"`             //卡池id
	NotHitLev5Times int64     `gorm:"column:not_hit_lev5_times" json:"not_hit_lev5_times"` //未中五星武将次数
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`                 //创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`                 //更新时间
}
