package po

type General struct { // nolint:maligned
	Id                int64  `gorm:"column:id" json:"id"`                                   // 主键ID
	Name              string `gorm:"column:name" json:"name"`                               // 姓名
	Gender            int8   `gorm:"column:gender" json:"gender"`                           // 性别
	Control           int32  `gorm:"column:control" json:"control"`                         // 统御
	Group             int8   `gorm:"column:group" json:"group"`                             // 阵营
	Quality           int8   `gorm:"column:quality" json:"quality"`                         // 品质
	Tag               string `gorm:"column:tag" json:"tag"`                                 // 标签
	AbilityAttr       string `gorm:"column:ability_attr" json:"ability_attr"`               // 能力属性
	AvatarUrl         string `gorm:"column:avatar_url" json:"avatar_url"`                   // 头像url
	ArmAttr           string `gorm:"column:arm_attr" json:"arm_attr"`                       // 兵种适性
	SelfTacticId      int32  `gorm:"column:self_tactic_id" json:"self_tactic_id"`           // 自带战法ID
	IsSupportDynamics int8   `gorm:"column:is_support_dynamics" json:"is_support_dynamics"` // 是否支持动态
	IsSupportCollect  int8   `gorm:"column:is_support_collect" json:"is_support_collect"`   // 是否支持典藏
}
