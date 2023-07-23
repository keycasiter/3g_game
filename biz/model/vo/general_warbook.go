package vo

type QueryGeneralWarbookCondition struct {
	Id        int64  // 主键ID
	Name      string // 兵书名称
	GeneralID int64  //武将ID

	OffSet int
	Limit  int
}
