package vo

type QuerySpecialTechCondition struct {
	Id   int64  // 主键ID
	Name string // 名称
	Type int    //装备类型

	Offset int
	Limit  int
}
