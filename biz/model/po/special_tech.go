package po

type SpecialTech struct { // nolint:maligned
	Id    int64  `gorm:"column:id" json:"id"`       // 主键ID
	Name  string `gorm:"column:name" json:"name"`   // 名称
	Type  int64  `gorm:"column:type" json:"type"`   // 类型
	Level int64  `gorm:"column:level" json:"level"` // 等级
	Desc  string `gorm:"column:desc" json:"desc"`   // 描述
}
