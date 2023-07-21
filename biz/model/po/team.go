package po

type Team struct { // nolint:maligned
	Id         int64  `gorm:"column:id" json:"id"`                   // 主键ID
	Name       string `gorm:"column:name" json:"name"`               // 队伍名称
	GeneralIds string `gorm:"column:general_ids" json:"general_ids"` // 武将ID
	Group      int32  `gorm:"column:group" json:"group"`             // 阵营
	Desc       string `gorm:"column:desc" json:"desc"`               // 描述
}
