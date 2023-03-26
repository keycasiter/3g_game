package consts

const (
	//鹰视狼顾
	ClearEyedAndMalicious = 1
	//士别三日
	ThreeDaysOfSeparation = 2
	//熯天炽地
	TheSkyIsBlazing = 3
	//乱世奸雄
	TraitorInTroubledTimes = 4
)

//战法池，按类型划分
// 主动
var ActiveTacticsMap map[int]TacticsType

// 被动
var PassiveTacticsMap map[int]TacticsType

// 指挥
var CommandTacticsMap map[int]TacticsType

// 突击
var AssaultTacticsMap map[int]TacticsType

// 阵法
var TroopsTacticsMap map[TacticsType]bool

// 兵种
var ArmTacticsMap map[TacticsType]bool

func init() {

}
