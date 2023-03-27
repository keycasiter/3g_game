package tactics

const (
	//鹰视狼顾
	ClearEyedAndMalicious = iota + 1
	//士别三日
	ThreeDaysOfSeparation
	//熯天炽地
	TheSkyIsBlazing
	//乱世奸雄
	TraitorInTroubledTimes
	//盛气凌敌
	OverwhelmingTheEnemyWithVigour
	//挫锐
	Demoralize
	//守而必固
	ToKeepAndBeFirm
	//横戈跃马
	Gallant
	//暂避其锋
	TakeRefugeFromEnemies
)

// ** 战法池，按类型划分 **
// 主动
var ActiveTacticsMap map[int64]bool

// 被动
var PassiveTacticsMap map[int64]bool

// 指挥
var CommandTacticsMap map[int64]bool

// 突击
var AssaultTacticsMap map[int64]bool

// 阵法
var TroopsTacticsMap map[int64]bool

// 兵种
var ArmTacticsMap map[int64]bool

// ** 战法池 **
var TacticsMap = map[int64]Tactics{}

func init() {
	initTacticsPool()
	initTacticsMap()
}

func initTacticsPool() {
	TacticsMap[ClearEyedAndMalicious] = &ClearEyedAndMaliciousTactic{}
	TacticsMap[ThreeDaysOfSeparation] = &ThreeDaysOfSeparationTactic{}
}

func initTacticsMap() {
	//指挥战法
	CommandTacticsMap[OverwhelmingTheEnemyWithVigour] = true
	CommandTacticsMap[Demoralize] = true
	CommandTacticsMap[ToKeepAndBeFirm] = true
	CommandTacticsMap[Gallant] = true
	CommandTacticsMap[TakeRefugeFromEnemies] = true
}
