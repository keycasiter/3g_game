package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/interface"
)

// ** 战法处理器 **
var TacticsHandlerMap = make(map[consts.TacticId]_interface.Tactics, 0)

func init() {
	initTacticsHandler()
}

func initTacticsHandler() {
	//飞沙走石
	TacticsHandlerMap[consts.FlyingSandAndRollingPebbles] = &FlyingSandAndRollingPebblesTactic{}
	//进言
	TacticsHandlerMap[consts.Introduction] = &IntroductionTactic{}
	//疑城
	TacticsHandlerMap[consts.ShamCastles] = &ShamCastlesTactic{}
	//神上使
	TacticsHandlerMap[consts.DivineEnvoy] = &DivineEnvoyTactic{}
	//扶危定倾
	TacticsHandlerMap[consts.DeliverTheCountryFromDistress] = &DeliverTheCountryFromDistressTactic{}
	//聚石成金
	TacticsHandlerMap[consts.AggregateStoneIntoGold] = &AggregateStoneIntoGoldTactic{}
	//短兵相见
	TacticsHandlerMap[consts.CloseQuarters] = &CloseQuartersTactic{}
	//非攻制胜
	TacticsHandlerMap[consts.NonOffensiveVictory] = &NonOffensiveVictoryTactic{}
	//疾风骤雨
	TacticsHandlerMap[consts.StrongWindAndSwiftRain] = &StrongWindAndSwiftRainTactic{}
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
	//毒泉拒蜀
	TacticsHandlerMap[consts.PoisonousSpringRefusesShu] = &PoisonousSpringRefusesShuTactic{}
	//陷阵突袭
	TacticsHandlerMap[consts.RaidInFormation] = &RaidInFormationTactic{}
	//天降火雨
	TacticsHandlerMap[consts.RainOfFireFromTheSky] = &RainOfFireFromTheSkyTactic{}
	//金城汤池
	TacticsHandlerMap[consts.RampartsOfMetalsAndAMoatOfHotWater] = &RampartsOfMetalsAndAMoatOfHotWaterTactic{}
	//鲁莽
	TacticsHandlerMap[consts.Reckless] = &RecklessTactic{}
	//精练策数
	TacticsHandlerMap[consts.RefinedStrategies] = &RefinedStrategiesTactic{}
	//一力拒守
	TacticsHandlerMap[consts.RefuseToDefendWithOneForce] = &RefuseToDefendWithOneForceTactic{}
	//折冲御侮
	TacticsHandlerMap[consts.RepelForeignAggression] = &RepelForeignAggressionTactic{}
	//乘胜长驱
	TacticsHandlerMap[consts.RidingOnTheVictoryDrive] = &RidingOnTheVictoryDriveTactic{}
	//乘敌不虞
	TacticsHandlerMap[consts.RidingTheEnemyWithoutFear] = &RidingTheEnemyWithoutFearTactic{}
	//奋突
	TacticsHandlerMap[consts.RiseUpBravely] = &RiseUpBravelyTactic{}
	//江天长焰
	TacticsHandlerMap[consts.RiverFireFlame] = &RiverFireFlameTactic{}
	//士争先赴
	TacticsHandlerMap[consts.ScholarsStriveToGoFirst] = &ScholarsStriveToGoFirstTactic{}
	//腹背受敌
	TacticsHandlerMap[consts.ScyllaAndCharybdis] = &ScyllaAndCharybdisTactic{}
	//临机制胜
	TacticsHandlerMap[consts.SeizeTheOpportunityToWin] = &SeizeTheOpportunityToWinTactic{}
	//自愈
	TacticsHandlerMap[consts.SelfHealing] = &SelfHealingTactic{}
	//形一阵
	TacticsHandlerMap[consts.ShapelyArray] = &ShapelyArrayTactic{}
	//沉沙决水
	TacticsHandlerMap[consts.SinkingSandAndBreakingWater] = &SinkingSandAndBreakingWaterTactic{}
	//坐守孤城
	TacticsHandlerMap[consts.SittingInAnIsolatedCity] = &SittingInAnIsolatedCityTactic{}
	//坐断东南
	TacticsHandlerMap[consts.SittingIntheSoutheast] = &SittingIntheSoutheastTactic{}
	//屠几上肉
	TacticsHandlerMap[consts.SlaughterMeatOnTable] = &SlaughterMeatOnTableTactic{}
	//卧薪尝胆
	TacticsHandlerMap[consts.SleepOnTheBrushwoodAndTasteTheGall] = &SleepOnTheBrushwoodAndTasteTheGallTactic{}
	//风声鹤唳
	TacticsHandlerMap[consts.SoundOfTheWindAndTheCryOfTheStork] = &SoundOfTheWindAndTheCryOfTheStorkTactic{}
	//奇兵间道
	TacticsHandlerMap[consts.SpecialSoldierPassRoad] = &SpecialSoldierPassRoadTactic{}
	//窃幸乘宠
	TacticsHandlerMap[consts.StealingLuckAndRidingPets] = &StealingLuckAndRidingPetsTactic{}
	//击其惰归
	TacticsHandlerMap[consts.StrikeItsLazyReturn] = &StrikeItsLazyReturnTactic{}
	//刚勇无前
	TacticsHandlerMap[consts.StrongAndBraveWithoutAdvance] = &StrongAndBraveWithoutAdvanceTactic{}
	//刚烈不屈
	TacticsHandlerMap[consts.StrongAndUnyielding] = &StrongAndUnyieldingTactic{}
	//监统震军
	TacticsHandlerMap[consts.SuperviseLeadAndSeizureArmy] = &SuperviseLeadAndSeizureArmyTactic{}
	//镇压黄巾
	TacticsHandlerMap[consts.SuppressYellowScarves] = &SuppressYellowScarvesTactic{}
	//围师必阙
	TacticsHandlerMap[consts.SurroundingTheTeacherMustBePalace] = &SurroundingTheTeacherMustBePalaceTactic{}
	//群攻
	TacticsHandlerMap[consts.SweepAway] = &SweepAwayTactic{}
	//速乘其利
	TacticsHandlerMap[consts.TakeAdvantageOfQuickly] = &TakeAdvantageOfQuicklyTactic{}
	//强攻
	TacticsHandlerMap[consts.TakeByStorm] = &TakeByStormTactic{}
	//先成其虑
	TacticsHandlerMap[consts.TakeCareOfYourselfFirst] = &TakeCareOfYourselfFirstTactic{}
	//挟势弄权
	TacticsHandlerMap[consts.TakingAdvantageOfTheSituationToGainPower] = &TakingAdvantageOfTheSituationToGainPowerTactic{}
	//气凌三军
	TacticsHandlerMap[consts.TemperamentSurpassesTheThreeArmies] = &TemperamentSurpassesTheThreeArmiesTactic{}
	//万箭齐发
	TacticsHandlerMap[consts.TenThousandArrowsShotAtOnce] = &TenThousandArrowsShotAtOnceTactic{}
	//十胜十败
	TacticsHandlerMap[consts.TenWinsAndTenLosses] = &TenWinsAndTenLossesTactic{}
	//藤甲兵
	TacticsHandlerMap[consts.TengjiaSoldier] = &TengjiaSoldierTactic{}
	//抬棺决战
	TacticsHandlerMap[consts.TheBattleOfCarryingCoffin] = &TheBattleOfCarryingCoffinTactic{}
	//勇者得前
	TacticsHandlerMap[consts.TheBraveLeadTheWay] = &TheBraveLeadTheWayTactic{}
	//千里走单骑
	TacticsHandlerMap[consts.ThousandMileWalkingSingleRider] = &ThousandMileWalkingSingleRiderTactic{}
	//千里驰援
	TacticsHandlerMap[consts.ThousandsOfMilesOfSupport] = &ThousandsOfMilesOfSupportTactic{}
	//三势阵
	TacticsHandlerMap[consts.ThreePotentialArray] = &ThreePotentialArrayTactic{}
	//五雷轰顶
	TacticsHandlerMap[consts.ThunderStruck] = &ThunderStruckTactic{}
	//落雷
	TacticsHandlerMap[consts.Thunderbolt] = &ThunderboltTactic{}
	//虎豹骑
	TacticsHandlerMap[consts.TigerAndLeopardCavalry] = &TigerAndLeopardCavalryTactic{}
	//虎踞鹰扬
	TacticsHandlerMap[consts.TigerCrouchingAndEagleSoaring] = &TigerCrouchingAndEagleSoaringTactic{}
	//虎卫军
	TacticsHandlerMap[consts.TigerGuardArmy] = &TigerGuardArmyTactic{}
	//虎痴
	TacticsHandlerMap[consts.TigerIdiot] = &TigerIdiotTactic{}
	//临战先登
	TacticsHandlerMap[consts.ToAscendBeforeBattle] = &ToAscendBeforeBattleTactic{}
	//焚辎营垒
	TacticsHandlerMap[consts.ToBurnBarracks] = &ToBurnBarracksTactic{}
	//将行其疾
	TacticsHandlerMap[consts.ToCureOnesSpeed] = &ToCureOnesSpeedTactic{}
	//搦战群雄
	TacticsHandlerMap[consts.ToSeizeThePowerOfGroupOfHeroes] = &ToSeizeThePowerOfGroupOfHeroesTactic{}
	//传檄宣威
	TacticsHandlerMap[consts.ToSpreadRumorsAndProclaimPower] = &ToSpreadRumorsAndProclaimPowerTactic{}
	//独行赴斗
	TacticsHandlerMap[consts.TravelingAloneToFight] = &TravelingAloneToFightTactic{}
	//十二奇策
	TacticsHandlerMap[consts.TwelveWonderfulStrategies] = &TwelveWonderfulStrategiesTactic{}
	//暗潮涌动
	TacticsHandlerMap[consts.UndercurrentSurge] = &UndercurrentSurgeTactic{}
	//克敌制胜
	TacticsHandlerMap[consts.VanquishTheEnemy] = &VanquishTheEnemyTactic{}
	//骁健神行
	TacticsHandlerMap[consts.VigorousAndWalk] = &VigorousAndWalkTactic{}
	//以逸待劳
	TacticsHandlerMap[consts.WaitAtOnesEaseForTheFatigued] = &WaitAtOnesEaseForTheFatiguedTactic{}
	//当锋摧决
	TacticsHandlerMap[consts.WhenTheFrontIsDestroyed] = &WhenTheFrontIsDestroyedTactic{}
	//白毦兵
	TacticsHandlerMap[consts.WhiteArmy] = &WhiteArmyTactic{}
	//白马义从
	TacticsHandlerMap[consts.WhiteHorseFollowsWithLoyalty] = &WhiteHorseFollowsWithLoyaltyTactic{}
	//风助火势
	TacticsHandlerMap[consts.WindAssistedFire] = &WindAssistedFireTactic{}
	//校胜帷幄
	TacticsHandlerMap[consts.WinsTent] = &WinsTentTactic{}
	//竭忠尽智
	TacticsHandlerMap[consts.WithAllTheHeart] = &WithAllTheHeartTactic{}
	//兵无常势
	TacticsHandlerMap[consts.WorkOutMeasuresToSuitLocalConditions] = &WorkOutMeasuresToSuitLocalConditionsTactic{}
	//武锋阵
	TacticsHandlerMap[consts.WuFengArray] = &WuFengArrayTactic{}
	//西凉铁骑
	TacticsHandlerMap[consts.XiLiangIronCavalry] = &XiLiangIronCavalryTactic{}
	//燕人咆哮
	TacticsHandlerMap[consts.YanPeopleRoar] = &YanPeopleRoarTactic{}
	//云聚影从
	TacticsHandlerMap[consts.CloudGatheringShadowFrom] = &CloudGatheringShadowFromTactic{}
	//处兹不惑
	TacticsHandlerMap[consts.InChaosNotConfused] = &InChaosNotConfusedTactic{}
	//鸱苕凤姿
	TacticsHandlerMap[consts.ThePostureOfAPhoenixWithAChickAndASweetPotato] = &ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic{}
	//狮子奋迅
	TacticsHandlerMap[consts.TheLionFliesFast] = &TheLionFliesFastTactic{}
	//工神
	TacticsHandlerMap[consts.TheGodOfCraftsmen] = &TheGodOfCraftsmenTactic{}
	//锦帆百翎
	TacticsHandlerMap[consts.JinfanArmyHundredFeathers] = &JinfanArmyHundredFeathersTactic{}
	//肉身铁壁
	TacticsHandlerMap[consts.CorporealIronWall] = &CorporealIronWallTactic{}
	//刀出如霆
	TacticsHandlerMap[consts.TheKnifeLikeThunderbolt] = &TheKnifeLikeThunderboltTactic{}
	//一举歼灭
	TacticsHandlerMap[consts.AnnihilateInOneFellSwoop] = &AnnihilateInOneFellSwoopTactic{}
	//避实击虚
	TacticsHandlerMap[consts.AvoidTheSolidAndStrikeTheWeak] = &AvoidTheSolidAndStrikeTheWeakTactic{}
	//威风凛凛
	TacticsHandlerMap[consts.AweInspiring] = &AweInspiringTactic{}
	//斗智
	TacticsHandlerMap[consts.BattleOfWits] = &BattleOfWitsTactic{}
	//才辩机捷
	TacticsHandlerMap[consts.BeQuickInDebatingOpportunities] = &BeQuickInDebatingOpportunitiesTactic{}
	//倾国倾城
	TacticsHandlerMap[consts.BeautyWhichOverthrowsStatesAndCities] = &BeautyWhichOverthrowsStatesAndCitiesTactic{}
	//一身是胆
	TacticsHandlerMap[consts.BebraveAllThrough] = &BebraveAllThroughTactic{}
	//威震华夏
	TacticsHandlerMap[consts.BecomeFamousAndFearInspiringThroughoutChina] = &BecomeFamousAndFearInspiringThroughoutChinaTactic{}
	//仁德载世
	TacticsHandlerMap[consts.BenevolentAndVirtuousThroughoutTheWorld] = &BenevolentAndVirtuousThroughoutTheWorldTactic{}
	//血刃争锋
	TacticsHandlerMap[consts.BloodBladeBattle] = &BloodBladeBattleTactic{}
	//槊血纵横
	TacticsHandlerMap[consts.BloodyAndUnrestrained] = &BloodyAndUnrestrainedTactic{}
	//弓腰姬
	TacticsHandlerMap[consts.BowWaistConcubine] = &BowWaistConcubineTactic{}
	//骁勇善战
	TacticsHandlerMap[consts.BraveAndBattleWise] = &BraveAndBattleWiseTactic{}
	//勇烈持重
	TacticsHandlerMap[consts.BraveAndResolute] = &BraveAndResoluteTactic{}
	//奋矛英姿
	TacticsHandlerMap[consts.BraveSpearHeroicPose] = &BraveSpearHeroicPoseTactic{}
	//破军威胜
	TacticsHandlerMap[consts.BreakingThroughTheArmyAndWinningVictories] = &BreakingThroughTheArmyAndWinningVictoriesTactic{}
	//锦囊妙计
	TacticsHandlerMap[consts.BrocadeBagAndCleverPlan] = &BrocadeBagAndCleverPlanTactic{}
	//登锋陷阵
	TacticsHandlerMap[consts.ChargeIntoTheEnemyRanks] = &ChargeIntoTheEnemyRanksTactic{}
	//手起刀落
	TacticsHandlerMap[consts.CutDown] = &CutDownTactic{}
	//沉断机谋
	TacticsHandlerMap[consts.DecisiveStrategy] = &DecisiveStrategyTactic{}
	//神射
	TacticsHandlerMap[consts.DivineEjaculation] = &DivineEjaculationTactic{}
	//神火计
	TacticsHandlerMap[consts.DivineFireMeter] = &DivineFireMeterTactic{}
	//神机莫测
	TacticsHandlerMap[consts.DivinelyInspiredStratagem] = &DivinelyInspiredStratagemTactic{}
	//左右开弓
	TacticsHandlerMap[consts.DrawTheBowBothOnTheLeftAndRight] = &DrawTheBowBothOnTheLeftAndRightTactic{}
	//包扎
	TacticsHandlerMap[consts.Dress] = &DressTactic{}
	//象兵
	TacticsHandlerMap[consts.ElephantSoldier] = &ElephantSoldierTactic{}
	//白眉
	TacticsHandlerMap[consts.EyebrowedThrush] = &EyebrowedThrushTactic{}
	//轻勇飞燕
	TacticsHandlerMap[consts.FearlessAndBraveFlyingSwallow] = &FearlessAndBraveFlyingSwallowTactic{}
	//不辱使命
	TacticsHandlerMap[consts.HaveSucceededInCarryingOutAnAssignment] = &HaveSucceededInCarryingOutAnAssignmentTactic{}
	//诈降
	TacticsHandlerMap[consts.PretendToSurrender] = &PretendToSurrenderTactic{}
	//誓守无降
	TacticsHandlerMap[consts.PromiseToKeepWithoutSurrender] = &PromiseToKeepWithoutSurrenderTactic{}
	//傲睨王侯
	TacticsHandlerMap[consts.ProudPrince] = &ProudPrinceTactic{}
	//挑衅
	TacticsHandlerMap[consts.Provoke] = &ProvokeTactic{}
	//掣刀斫敌
	TacticsHandlerMap[consts.PullingSwordsAndChoppingEnemies] = &PullingSwordsAndChoppingEnemiesTactic{}
	//净化
	TacticsHandlerMap[consts.Purify] = &PurifyTactic{}
	//青州兵
	TacticsHandlerMap[consts.QingZhouSoldier] = &QingZhouSoldierTactic{}
	//震骇四境
	TacticsHandlerMap[consts.ShockingFourRealms] = &ShockingFourRealmsTactic{}
	//横扫千军
	TacticsHandlerMap[consts.SweepAwayTheMillionsOfEnemyTroops] = &SweepAwayTheMillionsOfEnemyTroopsTactic{}
	//义心昭烈
	TacticsHandlerMap[consts.TheHeartOfRighteousnessShines] = &TheHeartOfRighteousnessShinesTactic{}
	//兵锋
	TacticsHandlerMap[consts.TheSharpnessOfMilitaryStrength] = &TheSharpnessOfMilitaryStrengthTactic{}
	//挥兵谋胜
	TacticsHandlerMap[consts.WieldTroopsToSeekVictory] = &WieldTroopsToSeekVictoryTactic{}
	//白衣渡江
	TacticsHandlerMap[consts.CrossingTheRiverInWhiteClothes] = &CrossingTheRiverInWhiteClothesTactic{}
	//溯江摇橹
	TacticsHandlerMap[consts.ChasingTheRiverAndRidingRows] = &ChasingTheRiverAndRidingRowsTactic{}
	//江东小霸王
	TacticsHandlerMap[consts.JiangdongLittleOverlord] = &JiangdongLittleOverlordTactic{}
	//国士将风
	TacticsHandlerMap[consts.CountryPersonGeneralStyle] = &CountryPersonGeneralStyleTactic{}
	//追伤
	TacticsHandlerMap[consts.ChasingInjury] = &ChasingInjuryTactic{}
	//闭月
	TacticsHandlerMap[consts.ClosedMoon] = &ClosedMoonTactic{}
	//苦肉计
	TacticsHandlerMap[consts.InjuryOnOneself] = &InjuryOnOneselfTactic{}
}
