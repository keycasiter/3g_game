package po

import (
	"database/sql"
	"time"
)

type JymGoods struct { // nolint:maligned
	Id          int64        `gorm:"column:id" json:"id"`                     // 主键
	GoodsUrl    string       `gorm:"column:goods_url" json:"goods_url"`       // 商品链接
	GoodsDetail string       `gorm:"column:goods_detail" json:"goods_detail"` // 商品信息
	CreateTime  time.Time    `gorm:"column:create_time" json:"create_time"`   // 创建时间
	UpdateTime  sql.NullTime `gorm:"column:update_time" json:"update_time"`   // 更新时间
	Status      int32        `gorm:"column:status" json:"status"`             // 状态
}
