package po

type JymGoods struct { // nolint:maligned
	Id          int64  `gorm:"column:id" json:"id"`                     // 主键
	GoodsUrl    string `gorm:"column:goods_url" json:"goods_url"`       // 商品链接
	GoodsDetail string `gorm:"column:goods_detail" json:"goods_detail"` // 商品信息
}
