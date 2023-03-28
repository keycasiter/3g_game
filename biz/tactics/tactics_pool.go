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
var ActiveTacticsMap = make(map[int64]bool, 0)

// 被动
var PassiveTacticsMap = make(map[int64]bool, 0)

// 指挥
var CommandTacticsMap = make(map[int64]bool, 0)

// 突击
var AssaultTacticsMap = make(map[int64]bool, 0)

// 阵法
var TroopsTacticsMap = make(map[int64]bool, 0)

// 兵种
var ArmTacticsMap = make(map[int64]bool, 0)

// ** 战法处理器 **
var TacticsHandlerMap = make(map[int64]Tactics, 0)

func init() {
	initTacticsHandler()
	initTacticsMap()
}

func initTacticsHandler() {
	TacticsHandlerMap[ClearEyedAndMalicious] = &ClearEyedAndMaliciousTactic{}
	TacticsHandlerMap[ThreeDaysOfSeparation] = &ThreeDaysOfSeparationTactic{}
	TacticsHandlerMap[TheSkyIsBlazing] = &TheSkyIsBlazingTactic{}
	TacticsHandlerMap[TraitorInTroubledTimes] = &TraitorInTroubledTimesTactic{}
}

func initTacticsMap() {
	//被动战法
	PassiveTacticsMap[ClearEyedAndMalicious] = true
	PassiveTacticsMap[ThreeDaysOfSeparation] = true
	//指挥战法
	CommandTacticsMap[OverwhelmingTheEnemyWithVigour] = true
	CommandTacticsMap[Demoralize] = true
	CommandTacticsMap[ToKeepAndBeFirm] = true
	CommandTacticsMap[Gallant] = true
	CommandTacticsMap[TakeRefugeFromEnemies] = true

}
