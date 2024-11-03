package vo

type QueryRecTeamCondition struct {
	Offset     int
	Limit      int
	Name       string
	GeneralIds []int64
	Group      int64
	ArmType    int64
}
