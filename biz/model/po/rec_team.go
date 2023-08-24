package po

type RecTeam struct { // nolint:maligned
	Id           int64  `gorm:"column:id" json:"id"`                       //主键ID
	Name         string `gorm:"column:name" json:"name"`                   // 队伍名称
	GeneralIds   string `gorm:"column:general_ids" json:"general_ids"`     //武将信息
	TacticIds    string `gorm:"column:tactic_ids" json:"tactic_ids"`       //战法
	WarbookIds   string `gorm:"column:warbook_ids" json:"warbook_ids"`     //兵书
	TechIds      string `gorm:"column:tech_ids" json:"tech_ids"`           //特技
	Desc         string `gorm:"column:desc" json:"desc"`                   //描述
	EvaluateDesc string `gorm:"column:evaluate_desc" json:"evaluate_desc"` //评级描述
}
