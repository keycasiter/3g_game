package tactics

import "github.com/keycasiter/3g_game/biz/tactics/interface"

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
	//魅惑
	Charming
	//抚揖军民
	AppeaseArmyAndPeople
	//镇扼防拒
	SuppressChokesAndPreventRefusals
	//锋矢阵
	FrontalVectorArray
	//刮骨疗毒
	Curettage
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
var TacticsHandlerMap = make(map[int64]_interface.Tactics, 0)

func init() {
	initTacticsHandler()
	initTacticsMap()
}

func initTacticsHandler() {
	//鹰视狼顾
	TacticsHandlerMap[ClearEyedAndMalicious] = &ClearEyedAndMaliciousTactic{}
	//士别三日
	TacticsHandlerMap[ThreeDaysOfSeparation] = &ThreeDaysOfSeparationTactic{}
	//熯天炽地
	TacticsHandlerMap[TheSkyIsBlazing] = &TheSkyIsBlazingTactic{}
	//乱世奸雄
	TacticsHandlerMap[TraitorInTroubledTimes] = &TraitorInTroubledTimesTactic{}
	//魅惑
	TacticsHandlerMap[Charming] = &CharmingTactic{}
	//镇扼防拒
	TacticsHandlerMap[SuppressChokesAndPreventRefusals] = &SuppressChokesAndPreventRefusalsTactic{}
	//抚揖军民
	TacticsHandlerMap[AppeaseArmyAndPeople] = &AppeaseArmyAndPeopleTactic{}
	//刮骨疗毒
	TacticsHandlerMap[Curettage] = &CurettageTactic{}
	//锋矢阵
	TacticsHandlerMap[FrontalVectorArray] = &FrontalVectorArrayTactic{}
}

func initTacticsMap() {
	//被动战法
	PassiveTacticsMap[ThreeDaysOfSeparation] = true
	PassiveTacticsMap[Charming] = true
	//指挥战法
	CommandTacticsMap[OverwhelmingTheEnemyWithVigour] = true
	CommandTacticsMap[Demoralize] = true
	CommandTacticsMap[ToKeepAndBeFirm] = true
	CommandTacticsMap[Gallant] = true
	CommandTacticsMap[TakeRefugeFromEnemies] = true
	CommandTacticsMap[SuppressChokesAndPreventRefusals] = true
	CommandTacticsMap[AppeaseArmyAndPeople] = true
	CommandTacticsMap[TraitorInTroubledTimes] = true
	PassiveTacticsMap[ClearEyedAndMalicious] = true
	//阵法
	TroopsTacticsMap[FrontalVectorArray] = true
	//主动
	ActiveTacticsMap[Curettage] = true
	ActiveTacticsMap[TheSkyIsBlazing] = true
}
