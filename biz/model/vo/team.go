package vo

type QueryTeamCondition struct {
	Id         int64  // 主键ID
	Name       string // 队伍名称
	GeneralIds []int  // 武将ID
	Group      int32  // 阵营
}
