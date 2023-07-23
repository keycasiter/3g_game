package vo

type QueryWarbookCondition struct {
	Id    int64   // 主键ID
	Name  string  // 兵书名称
	Level int32   // 兵书层级
	Ids   []int64 //主键ID集合

	Offset int
	Limit  int
}
