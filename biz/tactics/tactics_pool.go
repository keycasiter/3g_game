package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/interface"
)

// ** 战法池，按类型划分 **
// 主动
var ActiveTacticsMap = make(map[consts.TacticId]bool, 0)

// 被动
var PassiveTacticsMap = make(map[consts.TacticId]bool, 0)

// 指挥
var CommandTacticsMap = make(map[consts.TacticId]bool, 0)

// 突击
var AssaultTacticsMap = make(map[consts.TacticId]bool, 0)

// 阵法
var TroopsTacticsMap = make(map[consts.TacticId]bool, 0)

// 兵种
var ArmTacticsMap = make(map[consts.TacticId]bool, 0)

// ** 战法处理器 **
var TacticsHandlerMap = make(map[consts.TacticId]_interface.Tactics, 0)

func init() {
	initTacticsHandler()
	initTacticsMap()
}

func initTacticsHandler() {
	//鹰视狼顾
	TacticsHandlerMap[consts.ClearEyedAndMalicious] = &ClearEyedAndMaliciousTactic{}
	//士别三日
	TacticsHandlerMap[consts.ThreeDaysOfSeparation] = &ThreeDaysOfSeparationTactic{}
	//熯天炽地
	TacticsHandlerMap[consts.TheSkyIsBlazing] = &TheSkyIsBlazingTactic{}
	//乱世奸雄
	TacticsHandlerMap[consts.TraitorInTroubledTimes] = &TraitorInTroubledTimesTactic{}
	//魅惑
	TacticsHandlerMap[consts.Charming] = &CharmingTactic{}
	//镇扼防拒
	TacticsHandlerMap[consts.SuppressChokesAndPreventRefusals] = &SuppressChokesAndPreventRefusalsTactic{}
	//抚揖军民
	TacticsHandlerMap[consts.AppeaseArmyAndPeople] = &AppeaseArmyAndPeopleTactic{}
	//刮骨疗毒
	TacticsHandlerMap[consts.Curettage] = &CurettageTactic{}
	//锋矢阵
	TacticsHandlerMap[consts.FrontalVectorArray] = &FrontalVectorArrayTactic{}
	//义胆雄心
	TacticsHandlerMap[consts.BraveAmbition] = &BraveAmbitionTactic{}
	//夺魂挟魄
	TacticsHandlerMap[consts.SeizeTheSoul] = &SeizeTheSoulTactic{}
	//火炽原燎
	TacticsHandlerMap[consts.BlazingWildfire] = &BlazingWildfireTactic{}
	//连环计
	TacticsHandlerMap[consts.InterlockedStratagems] = &InterlockedStratagemsTactic{}
	//太平道法
	TacticsHandlerMap[consts.TaipingLaw] = &TaipingLawTactic{}
	//无当飞军
	TacticsHandlerMap[consts.WuDangFlyArmy] = &WuDangFlyArmyTactic{}
	//神机妙算
	TacticsHandlerMap[consts.CleverStrategyAndShrewdTactics] = &CleverStrategyAndShrewdTacticsTactic{}
	//八门金锁阵
	TacticsHandlerMap[consts.EightGateGoldenLockArray] = &EightGateGoldenLockArrayTactic{}
	//草船借箭
	TacticsHandlerMap[consts.BorrowArrowsWithThatchedBoats] = &BorrowArrowsWithThatchedBoatsTactic{}
	//盛气凌敌
	TacticsHandlerMap[consts.OverwhelmingTheEnemyWithVigour] = &OverwhelmingTheEnemyWithVigourTactic{}
	//挫锐
	TacticsHandlerMap[consts.Demoralize] = &DemoralizeTactic{}
	//守而必固
	TacticsHandlerMap[consts.ToKeepAndBeFirm] = &ToKeepAndBeFirmTactic{}
	//横戈跃马
	TacticsHandlerMap[consts.Gallant] = &GallantTactic{}
	//暂避其锋
	TacticsHandlerMap[consts.TakeRefugeFromEnemies] = &TakeRefugeFromEnemiesTactic{}
	//长者之风
	TacticsHandlerMap[consts.TheWindOfTheElderly] = &TheWindOfTheElderlyTactic{}
	//用武通神
	TacticsHandlerMap[consts.UseMartialArtsToConnectWithGods] = &UseMartialArtsToConnectWithGodsTactic{}
	//梦中弑臣
	TacticsHandlerMap[consts.KillingMinisterInDream] = &KillingMinisterInDreamTactic{}
	//奇计良谋
	TacticsHandlerMap[consts.CleverPlanAndCleverPlan] = &CleverPlanAndCleverPlanTactic{}
	//舌战群儒
	TacticsHandlerMap[consts.LectureField] = &LectureFieldTactic{}
	//整装待发
	TacticsHandlerMap[consts.BeFullyEquippedFor] = &BeFullyEquippedForTactic{}
	//御敌屏障
	TacticsHandlerMap[consts.DefensiveBarrier] = &DefensiveBarrierTactic{}
	//骑虎难下
	TacticsHandlerMap[consts.RideTigerHardToGetOff] = &RideTigerHardToGetOffTactic{}
	//才器过人
	TacticsHandlerMap[consts.OutstandingTalent] = &OutstandingTalentTactic{}
}

func initTacticsMap() {
	//被动战法
	PassiveTacticsMap[consts.ThreeDaysOfSeparation] = true
	PassiveTacticsMap[consts.Charming] = true
	PassiveTacticsMap[consts.TaipingLaw] = true
	PassiveTacticsMap[consts.BraveAmbition] = true
	//指挥战法
	CommandTacticsMap[consts.OverwhelmingTheEnemyWithVigour] = true
	CommandTacticsMap[consts.Demoralize] = true
	CommandTacticsMap[consts.ToKeepAndBeFirm] = true
	CommandTacticsMap[consts.Gallant] = true
	CommandTacticsMap[consts.TakeRefugeFromEnemies] = true
	CommandTacticsMap[consts.SuppressChokesAndPreventRefusals] = true
	CommandTacticsMap[consts.AppeaseArmyAndPeople] = true
	CommandTacticsMap[consts.TraitorInTroubledTimes] = true
	PassiveTacticsMap[consts.ClearEyedAndMalicious] = true
	PassiveTacticsMap[consts.CleverStrategyAndShrewdTactics] = true
	PassiveTacticsMap[consts.TheWindOfTheElderly] = true
	PassiveTacticsMap[consts.UseMartialArtsToConnectWithGods] = true
	PassiveTacticsMap[consts.KillingMinisterInDream] = true
	PassiveTacticsMap[consts.CleverPlanAndCleverPlan] = true
	PassiveTacticsMap[consts.LectureField] = true
	PassiveTacticsMap[consts.BeFullyEquippedFor] = true
	PassiveTacticsMap[consts.DefensiveBarrier] = true
	PassiveTacticsMap[consts.RideTigerHardToGetOff] = true
	//阵法
	TroopsTacticsMap[consts.FrontalVectorArray] = true
	TroopsTacticsMap[consts.EightGateGoldenLockArray] = true
	//兵种
	ArmTacticsMap[consts.WuDangFlyArmy] = true
	//主动
	ActiveTacticsMap[consts.Curettage] = true
	ActiveTacticsMap[consts.TheSkyIsBlazing] = true
	ActiveTacticsMap[consts.SeizeTheSoul] = true
	ActiveTacticsMap[consts.BlazingWildfire] = true
	ActiveTacticsMap[consts.InterlockedStratagems] = true
	ActiveTacticsMap[consts.BorrowArrowsWithThatchedBoats] = true
	ActiveTacticsMap[consts.OutstandingTalent] = true
}
