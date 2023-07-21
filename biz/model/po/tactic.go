package po

type Tactic struct { // nolint:maligned
	Id      int64  `gorm:"column:id" json:"id"`           // 主键ID
	Name    string `gorm:"column:name" json:"name"`       // 战法名称
	Quality int32  `gorm:"column:quality" json:"quality"` // 战法品质
	Source  int32  `gorm:"column:source" json:"source"`   // 战法来源
	Type    int32  `gorm:"column:type" json:"type"`       // 战法类型
	Desc    string `gorm:"column:desc" json:"desc"`       // 描述
}
