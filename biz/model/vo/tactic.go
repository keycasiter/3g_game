package vo

type QueryTacticCondition struct {
	Id      int64  // 主键ID
	Name    string // 战法名称
	Quality int32  // 战法品质
	Source  int32  // 战法来源
	Type    int32  // 战法类型
}