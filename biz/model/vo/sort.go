package vo

type BattleGeneralsOrderBySpeed []*BattleGeneral

func (generals BattleGeneralsOrderBySpeed) Len() int {
	return len(generals)
}

func (generals BattleGeneralsOrderBySpeed) Less(i, j int) bool {
	return generals[i].BaseInfo.AbilityAttr.SpeedBase > generals[j].BaseInfo.AbilityAttr.SpeedBase
}

func (generals BattleGeneralsOrderBySpeed) Swap(i, j int) {
	temp := generals[i]
	generals[i] = generals[j]
	generals[j] = temp
}
