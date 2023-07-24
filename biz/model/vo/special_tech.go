package vo

type QuerySpecialTechCondition struct {
	Id   int64  // 主键ID
	Name string // 名称

	Offset int
	Limit  int
}
