package po

type Warbook struct { // nolint:maligned
	Id    int64  `gorm:"column:id" json:"id"`       // 主键ID
	Name  string `gorm:"column:name" json:"name"`   // 兵书名称
	Level int32  `gorm:"column:level" json:"level"` // 兵书层级
	Type  int32  `gorm:"column:type" json:"type"`   // 兵书类型
	Desc  string `gorm:"column:desc" json:"desc"`   // 描述
}
