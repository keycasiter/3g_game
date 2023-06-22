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

// 准备回合的战法
var ActivePrepareTacticsMap = make(map[consts.TacticId]bool, 0)

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
	//决水溃城
	TacticsHandlerMap[consts.BreakingThroughTheWaterAndCrushingTheCity] = &BreakingThroughTheWaterAndCrushingTheCityTactic{}
	//据水断桥
	TacticsHandlerMap[consts.BrokenBridgeByWater] = &BrokenBridgeByWaterTactic{}
	//众动万计
	TacticsHandlerMap[consts.CrowdMovesTenThousandCounts] = &CrowdMovesTenThousandCountsTactic{}
	//运筹决算
	TacticsHandlerMap[consts.DecisionMakingThroughOperationsResearch] = &DecisionMakingThroughOperationsResearchTactic{}
	//驱散
	TacticsHandlerMap[consts.Disperse] = &DisperseTactic{}
	//箕形阵
	TacticsHandlerMap[consts.DustpanFormation] = &DustpanFormationTactic{}
	//绝其汲道
	TacticsHandlerMap[consts.EliminateItAndDrawFromIt] = &EliminateItAndDrawFromItTactic{}
	//累世立名
	TacticsHandlerMap[consts.EstablishingNameThroughGenerations] = &EstablishingNameThroughGenerationsTactic{}
	//振军击营
	TacticsHandlerMap[consts.ExcitingArmyAttackCamp] = &ExcitingArmyAttackCampTactic{}
	//酒池肉林
	TacticsHandlerMap[consts.ExtravagantOrgy] = &ExtravagantOrgyTactic{}
	//伪书相间
	TacticsHandlerMap[consts.FakeBooksAlternateWithEachOther] = &FakeBooksAlternateWithEachOtherTactic{}
	//落凤
	TacticsHandlerMap[consts.FallingPhoenix] = &FallingPhoenixTactic{}
	//陷阵营
	TacticsHandlerMap[consts.FallIntoCamp] = &FallIntoCampTactic{}
	//火神英风
	TacticsHandlerMap[consts.FireGodHeroStyle] = &FireGodHeroStyleTactic{}
	//火烧连营
	TacticsHandlerMap[consts.FireJointVenture] = &FireJointVentureTactic{}
	//焰逐风飞
	TacticsHandlerMap[consts.FlamesFlyingInTheWind] = &FlamesFlyingInTheWindTactic{}
	//水淹七军
	TacticsHandlerMap[consts.FloodedSeventhArmy] = &FloodedSeventhArmyTactic{}
	//垂心万物
	TacticsHandlerMap[consts.FocusingOnAllThings] = &FocusingOnAllThingsTactic{}
	//料事如神
	TacticsHandlerMap[consts.ForetellLikeProphet] = &ForetellLikeProphetTactic{}
	//挫志怒袭
	TacticsHandlerMap[consts.FrustrationAndAngerAttack] = &FrustrationAndAngerAttackTactic{}
	//符命自立
	TacticsHandlerMap[consts.FumingSelfReliance] = &FumingSelfRelianceTactic{}
	//后发制人
	TacticsHandlerMap[consts.GainMasteryByStrikingOnlyAfterTheEnemyHasStruck] = &GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic{}
	//敛众而击
	TacticsHandlerMap[consts.GatherTheCrowdAndStrike] = &GatherTheCrowdAndStrikeTactic{}
	//合军聚众
	TacticsHandlerMap[consts.GatheringOfTroops] = &GatheringOfTroopsTactic{}
	//将门虎女
	TacticsHandlerMap[consts.GeneralBraveGirl] = &GeneralBraveGirlTactic{}
	//鬼神霆威
	TacticsHandlerMap[consts.GhostGodThunderForce] = &GhostGodThunderForceTactic{}
	//金丹秘术
	TacticsHandlerMap[consts.GoldenPillSecretTechnique] = &GoldenPillSecretTechniqueTactic{}
	//大戟士
	TacticsHandlerMap[consts.GreatHalberdWarrior] = &GreatHalberdWarriorTactic{}
	//枪舞如风
	TacticsHandlerMap[consts.GunDanceLikeTheWind] = &GunDanceLikeTheWindTactic{}
	//唇枪舌战
	TacticsHandlerMap[consts.HaveVerbalBattleWithSomebody] = &HaveVerbalBattleWithSomebodyTactic{}
	//济贫好施
	TacticsHandlerMap[consts.HelpingThePoorAndGivingGenerously] = &HelpingThePoorAndGivingGenerouslyTactic{}
	//暗箭难防
	TacticsHandlerMap[consts.HiddenArrowsAreDifficultToGuardAgainst] = &HiddenArrowsAreDifficultToGuardAgainstTactic{}
	//潜龙阵
	TacticsHandlerMap[consts.HiddenDragonArray] = &HiddenDragonArrayTactic{}
	//暗藏玄机
	TacticsHandlerMap[consts.HiddenMystery] = &HiddenMysteryTactic{}
	//诈降
	TacticsHandlerMap[consts.PretendToSurrender] = &PretendToSurrenderTactic{}
	//高橹连营
	TacticsHandlerMap[consts.HighWoodenPaddlesConnectedToTheCamp] = &HighWoodenPaddlesConnectedToTheCampTactic{}
	//百步穿杨
	TacticsHandlerMap[consts.HitTheTargetAtEveryShot] = &HitTheTargetAtEveryShotTactic{}
	//持军毅重
	TacticsHandlerMap[consts.HoldTheArmyWithDeterminationAndDetermination] = &HoldTheArmyWithDeterminationAndDeterminationTactic{}
	//胡笳余音
	TacticsHandlerMap[consts.HuJiaLingeringSound] = &HuJiaLingeringSoundTactic{}
	//百骑劫营
	TacticsHandlerMap[consts.HundredCavalryRobberyBattalions] = &HundredCavalryRobberyBattalionsTactic{}
	//百计多谋
	TacticsHandlerMap[consts.HundredStrategiesAndManyStrategies] = &HundredStrategiesAndManyStrategiesTactic{}
	//一骑当千
	TacticsHandlerMap[consts.IkkiTousen] = &IkkiTousenTactic{}
	//固若金汤
	TacticsHandlerMap[consts.Impregnable] = &ImpregnableTactic{}
	//婴城自守
	TacticsHandlerMap[consts.InfantryCitySelfDefense] = &InfantryCitySelfDefenseTactic{}
	//智计
	TacticsHandlerMap[consts.IntelligentStrategy] = &IntelligentStrategyTactic{}
	//威谋靡亢
	TacticsHandlerMap[consts.IntenseAndPowerful] = &IntenseAndPowerfulTactic{}
	//铁骑驱驰
	TacticsHandlerMap[consts.IronHorseDrive] = &IronHorseDriveTactic{}
	//绝地反击
	TacticsHandlerMap[consts.JediCounterattack] = &JediCounterattackTactic{}
	//锦帆军
	TacticsHandlerMap[consts.JinFanArmy] = &JinFanArmyTactic{}
	//引弦力战
	TacticsHandlerMap[consts.LeadStringBattle] = &LeadStringBattleTactic{}
	//纵兵劫掠
	TacticsHandlerMap[consts.LeavingSoldiersToPlunder] = &LeavingSoldiersToPlunderTactic{}
	//顾盼生姿
	TacticsHandlerMap[consts.LookAroundCharmingly] = &LookAroundCharminglyTactic{}
	//忠勇义烈
	TacticsHandlerMap[consts.LoyalAndBraveMartyrs] = &LoyalAndBraveMartyrsTactic{}
	//诱敌深入
	TacticsHandlerMap[consts.LureTheEnemyInDeep] = &LureTheEnemyInDeepTactic{}
	//竭力佐谋
	TacticsHandlerMap[consts.MakeEveryEffortToAssistInPlanning] = &MakeEveryEffortToAssistInPlanningTactic{}
	//声东击西
	TacticsHandlerMap[consts.MakeFeintToTheEastButAttackInTheWest] = &MakeFeintToTheEastButAttackInTheWestTactic{}
	//形机军略
	TacticsHandlerMap[consts.MilitaryStrategyForFormAircraft] = &MilitaryStrategyForFormAircraftTactic{}
	//裸衣血战
	TacticsHandlerMap[consts.NakedBloodBattle] = &NakedBloodBattleTactic{}
	//机略纵横
	TacticsHandlerMap[consts.MachineStrategyVerticalAndHorizontal] = &MachineStrategyVerticalAndHorizontalTactic{}
	//长驱直入
	TacticsHandlerMap[consts.MarchInto] = &MarchIntoTactic{}
	//南蛮渠魁
	TacticsHandlerMap[consts.NanManQuKui] = &NanManQuKuiTactic{}
	//兴云布雨
	TacticsHandlerMap[consts.MakeCloudAndRain] = &MakeCloudAndRainTactic{}
	//青囊
	TacticsHandlerMap[consts.MedicalPractice] = &MedicalPracticeTactic{}
	//死战不退
	TacticsHandlerMap[consts.NeverRetreatFromDeadBattle] = &NeverRetreatFromDeadBattleTactic{}
	//机鉴先识
	TacticsHandlerMap[consts.OpportunityIdentificationFirst] = &OpportunityIdentificationFirstTactic{}
	//暴敛四方
	TacticsHandlerMap[consts.OverwhelmingAllDirections] = &OverwhelmingAllDirectionsTactic{}
	//勇冠三军
	TacticsHandlerMap[consts.PeerlessOrMatchlessBraveryOrValour] = &PeerlessOrMatchlessBraveryOrValourTactic{}
	//计定谋决
	TacticsHandlerMap[consts.PlanAndDecide] = &PlanAndDecideTactic{}
	//鸩毒
	TacticsHandlerMap[consts.PoisonedWine] = &PoisonedWineTactic{}
}

func initTacticsMap() {
	//需要准备回合的战法
	ActivePrepareTacticsMap[consts.AdvancingSecretlyByUnknownPath] = true

	//被动战法
	PassiveTacticsMap[consts.LeadStringBattle] = true
	PassiveTacticsMap[consts.PlanAndDecide] = true
	PassiveTacticsMap[consts.NeverRetreatFromDeadBattle] = true
	PassiveTacticsMap[consts.ThreeDaysOfSeparation] = true
	PassiveTacticsMap[consts.Charming] = true
	PassiveTacticsMap[consts.TaipingLaw] = true
	PassiveTacticsMap[consts.BraveAmbition] = true
	PassiveTacticsMap[consts.OnesResolveIsUnshaken] = true
	PassiveTacticsMap[consts.BeAdeptWithBothPenAndSword] = true
	PassiveTacticsMap[consts.CrowdMovesTenThousandCounts] = true
	PassiveTacticsMap[consts.ExtravagantOrgy] = true
	PassiveTacticsMap[consts.FumingSelfReliance] = true
	PassiveTacticsMap[consts.GainMasteryByStrikingOnlyAfterTheEnemyHasStruck] = true
	PassiveTacticsMap[consts.GatheringOfTroops] = true
	PassiveTacticsMap[consts.HelpingThePoorAndGivingGenerously] = true
	PassiveTacticsMap[consts.HighWoodenPaddlesConnectedToTheCamp] = true
	PassiveTacticsMap[consts.JediCounterattack] = true
	PassiveTacticsMap[consts.LoyalAndBraveMartyrs] = true
	PassiveTacticsMap[consts.NakedBloodBattle] = true
	PassiveTacticsMap[consts.MarchInto] = true
	//指挥战法
	CommandTacticsMap[consts.IronHorseDrive] = true
	CommandTacticsMap[consts.OpportunityIdentificationFirst] = true
	CommandTacticsMap[consts.GoldenPillSecretTechnique] = true
	CommandTacticsMap[consts.OverwhelmingTheEnemyWithVigour] = true
	CommandTacticsMap[consts.FocusingOnAllThings] = true
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
	CommandTacticsMap[consts.FireGodHeroStyle] = true
	CommandTacticsMap[consts.PretendToSurrender] = true
	CommandTacticsMap[consts.HundredStrategiesAndManyStrategies] = true
	CommandTacticsMap[consts.NanManQuKui] = true
	CommandTacticsMap[consts.MakeCloudAndRain] = true
	CommandTacticsMap[consts.MedicalPractice] = true
	//阵法
	TroopsTacticsMap[consts.FrontalVectorArray] = true
	TroopsTacticsMap[consts.EightGateGoldenLockArray] = true
	TroopsTacticsMap[consts.DustpanFormation] = true
	TroopsTacticsMap[consts.FallIntoCamp] = true
	TroopsTacticsMap[consts.HiddenDragonArray] = true
	//兵种
	ArmTacticsMap[consts.WuDangFlyArmy] = true
	ArmTacticsMap[consts.GreatHalberdWarrior] = true
	ArmTacticsMap[consts.JinFanArmy] = true
	//主动
	ActiveTacticsMap[consts.PoisonedWine] = true
	ActiveTacticsMap[consts.OverwhelmingAllDirections] = true
	ActiveTacticsMap[consts.MachineStrategyVerticalAndHorizontal] = true
	ActiveTacticsMap[consts.MilitaryStrategyForFormAircraft] = true
	ActiveTacticsMap[consts.MakeEveryEffortToAssistInPlanning] = true
	ActiveTacticsMap[consts.MakeFeintToTheEastButAttackInTheWest] = true
	ActiveTacticsMap[consts.LureTheEnemyInDeep] = true
	ActiveTacticsMap[consts.LookAroundCharmingly] = true
	ActiveTacticsMap[consts.IntenseAndPowerful] = true
	ActiveTacticsMap[consts.LeavingSoldiersToPlunder] = true
	ActiveTacticsMap[consts.IntelligentStrategy] = true
	ActiveTacticsMap[consts.InfantryCitySelfDefense] = true
	ActiveTacticsMap[consts.Impregnable] = true
	ActiveTacticsMap[consts.HitTheTargetAtEveryShot] = true
	ActiveTacticsMap[consts.HuJiaLingeringSound] = true
	ActiveTacticsMap[consts.HoldTheArmyWithDeterminationAndDetermination] = true
	ActiveTacticsMap[consts.HiddenArrowsAreDifficultToGuardAgainst] = true
	ActiveTacticsMap[consts.GeneralBraveGirl] = true
	ActiveTacticsMap[consts.GunDanceLikeTheWind] = true
	ActiveTacticsMap[consts.HaveVerbalBattleWithSomebody] = true
	ActiveTacticsMap[consts.Curettage] = true
	ActiveTacticsMap[consts.GatherTheCrowdAndStrike] = true
	ActiveTacticsMap[consts.FlamesFlyingInTheWind] = true
	ActiveTacticsMap[consts.FloodedSeventhArmy] = true
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
	ActiveTacticsMap[consts.BreakingThroughTheWaterAndCrushingTheCity] = true
	ActiveTacticsMap[consts.BrokenBridgeByWater] = true
	ActiveTacticsMap[consts.DecisionMakingThroughOperationsResearch] = true
	ActiveTacticsMap[consts.Disperse] = true
	ActiveTacticsMap[consts.EliminateItAndDrawFromIt] = true
	ActiveTacticsMap[consts.EstablishingNameThroughGenerations] = true
	ActiveTacticsMap[consts.ExcitingArmyAttackCamp] = true
	ActiveTacticsMap[consts.FakeBooksAlternateWithEachOther] = true
	ActiveTacticsMap[consts.FallingPhoenix] = true
	ActiveTacticsMap[consts.FireJointVenture] = true
	ActiveTacticsMap[consts.ForetellLikeProphet] = true
	ActiveTacticsMap[consts.FrustrationAndAngerAttack] = true

	//突击
	AssaultTacticsMap[consts.PeerlessOrMatchlessBraveryOrValour] = true
	AssaultTacticsMap[consts.ViolentAndHeartless] = true
	AssaultTacticsMap[consts.BendTheBowAndDrinkTheFeathers] = true
	AssaultTacticsMap[consts.GhostGodThunderForce] = true
	AssaultTacticsMap[consts.HiddenMystery] = true
	AssaultTacticsMap[consts.HundredCavalryRobberyBattalions] = true
}
