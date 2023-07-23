package po

type GeneralWarbook struct { // nolint:maligned
	Id        int64  `gorm:"column:id" json:"id"`                 // 主键ID
	GeneralId int64  `gorm:"column:general_id" json:"general_id"` // 主键ID
	Name      string `gorm:"column:name" json:"name"`             // 战法名称
	WarBook   string `gorm:"column:warbook" json:"warbook"`       // 兵书
}
