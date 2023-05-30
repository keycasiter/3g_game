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
	//偃旗息鼓
	TacticsHandlerMap[consts.LowerBannersAndMuffleDrums] = &LowerBannersAndMuffleDrumsTactic{}
	//出其不意
	TacticsHandlerMap[consts.TakeBySurprise] = &TakeBySurpriseTactic{}
	//庐江上甲
	TacticsHandlerMap[consts.LuJiangRiverOverArmoured] = &LuJiangRiverOverArmouredTactic{}
	//谦让
	TacticsHandlerMap[consts.Humility] = &HumilityTactic{}
	//矢志不移
	TacticsHandlerMap[consts.OnesResolveIsUnshaken] = &OnesResolveIsUnshakenTactic{}
	//所向披靡
	TacticsHandlerMap[consts.EverTriumphant] = &EverTriumphantTactic{}
	//破阵摧坚
	TacticsHandlerMap[consts.BreakingThroughTheFormationAndDestroyingTheFirm] = &BreakingThroughTheFormationAndDestroyingTheFirmTactic{}
	//杯蛇鬼车
	TacticsHandlerMap[consts.CupSnakeGhostCar] = &CupSnakeGhostCarTactic{}
	//瞋目横矛
	TacticsHandlerMap[consts.AngryEyeHorizontalSpear] = &AngryEyeHorizontalSpearTactic{}
	//暴戾无仁
	TacticsHandlerMap[consts.ViolentAndHeartless] = &ViolentAndHeartlessTactic{}
	//经天纬地
	TacticsHandlerMap[consts.AbilityToRuleTheCountry] = &AbilityToRuleTheCountryTactic{}
	//暗渡陈仓
	TacticsHandlerMap[consts.AdvancingSecretlyByUnknownPath] = &AdvancingSecretlyByUnknownPathTactic{}
	//四面楚歌
	TacticsHandlerMap[consts.BeBesiegedOnAllSides] = &BeBesiegedOnAllSidesTactic{}
	//十面埋伏
	TacticsHandlerMap[consts.AmbushOnAllSides] = &AmbushOnAllSidesTactic{}
	//古之恶来
	TacticsHandlerMap[consts.AncientEvilComes] = &AncientEvilComesTactic{}
	//文武双全
	TacticsHandlerMap[consts.BeAdeptWithBothPenAndSword] = &BeAdeptWithBothPenAndSwordTactic{}
	//弯弓饮羽
	TacticsHandlerMap[consts.BendTheBowAndDrinkTheFeathers] = &BendTheBowAndDrinkTheFeathersTactic{}
	//妖术
	TacticsHandlerMap[consts.BlackArt] = &BlackArtTactic{}
	//破军威胜
	TacticsHandlerMap[consts.BreakingThroughTheArmyAndWinningVictories] = &BreakingThroughTheArmyAndWinningVictoriesTactic{}
}

func initTacticsMap() {
	//被动战法
	PassiveTacticsMap[consts.ThreeDaysOfSeparation] = true
	PassiveTacticsMap[consts.Charming] = true
	PassiveTacticsMap[consts.TaipingLaw] = true
	PassiveTacticsMap[consts.BraveAmbition] = true
	PassiveTacticsMap[consts.OnesResolveIsUnshaken] = true
	PassiveTacticsMap[consts.BeAdeptWithBothPenAndSword] = true
	//指挥战法
	CommandTacticsMap[consts.OverwhelmingTheEnemyWithVigour] = true
	CommandTacticsMap[consts.Demoralize] = true
	CommandTacticsMap[consts.ToKeepAndBeFirm] = true
	CommandTacticsMap[consts.Gallant] = true
	CommandTacticsMap[consts.TakeRefugeFromEnemies] = true
	CommandTacticsMap[consts.SuppressChokesAndPreventRefusals] = true
	CommandTacticsMap[consts.AppeaseArmyAndPeople] = true
	CommandTacticsMap[consts.TraitorInTroubledTimes] = true
	CommandTacticsMap[consts.ClearEyedAndMalicious] = true
	CommandTacticsMap[consts.CleverStrategyAndShrewdTactics] = true
	CommandTacticsMap[consts.TheWindOfTheElderly] = true
	CommandTacticsMap[consts.UseMartialArtsToConnectWithGods] = true
	CommandTacticsMap[consts.KillingMinisterInDream] = true
	CommandTacticsMap[consts.CleverPlanAndCleverPlan] = true
	CommandTacticsMap[consts.LectureField] = true
	CommandTacticsMap[consts.BeFullyEquippedFor] = true
	CommandTacticsMap[consts.DefensiveBarrier] = true
	CommandTacticsMap[consts.RideTigerHardToGetOff] = true
	CommandTacticsMap[consts.AbilityToRuleTheCountry] = true
	CommandTacticsMap[consts.AncientEvilComes] = true
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
	ActiveTacticsMap[consts.LowerBannersAndMuffleDrums] = true
	ActiveTacticsMap[consts.TakeBySurprise] = true
	ActiveTacticsMap[consts.LuJiangRiverOverArmoured] = true
	ActiveTacticsMap[consts.Humility] = true
	ActiveTacticsMap[consts.EverTriumphant] = true
	ActiveTacticsMap[consts.BreakingThroughTheFormationAndDestroyingTheFirm] = true
	ActiveTacticsMap[consts.CupSnakeGhostCar] = true
	ActiveTacticsMap[consts.AngryEyeHorizontalSpear] = true
	ActiveTacticsMap[consts.AdvancingSecretlyByUnknownPath] = true
	ActiveTacticsMap[consts.BeBesiegedOnAllSides] = true
	ActiveTacticsMap[consts.AmbushOnAllSides] = true
	ActiveTacticsMap[consts.BlackArt] = true
	ActiveTacticsMap[consts.BreakingThroughTheArmyAndWinningVictories] = true
	//突击
	AssaultTacticsMap[consts.ViolentAndHeartless] = true
	AssaultTacticsMap[consts.BendTheBowAndDrinkTheFeathers] = true
}
