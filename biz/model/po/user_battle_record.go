package po

import (
	"time"
)

type UserBattleRecord struct { // nolint:maligned
	Id           int64     `gorm:"column:id" json:"id"`                       // 主键ID
	Uid          int64     `gorm:"column:uid" json:"uid"`                     // 用户ID
	BattleParam  string    `gorm:"column:battle_param" json:"battle_param"`   // 对战参数
	BattleRecord string    `gorm:"column:battle_record" json:"battle_record"` // 对战记录
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`       // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`       // 更新时间
}
