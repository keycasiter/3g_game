package vo

type QueryGeneralWarbookCondition struct {
	Id        int64  // 主键ID
	Name      string // 兵书名称
	GeneralID int64  //武将ID
	Type      int64  //兵书类型

	OffSet int
	Limit  int
}
