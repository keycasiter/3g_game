package consts

// 战法枚举
type TacticId int64

const (
	//未知战法
	UnknownTactic TacticId = iota
	//鹰视狼顾
	ClearEyedAndMalicious
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
	//义胆雄心
	BraveAmbition
	//夺魂挟魄
	SeizeTheSoul
	//火炽原燎
	BlazingWildfire
	//连环计
	InterlockedStratagems
	//太平道法
	TaipingLaw
	//无当飞军
	WuDangFlyArmy
	//神机妙算
	CleverStrategyAndShrewdTactics
	//八门金锁阵
	EightGateGoldenLockArray
	//草船借箭
	BorrowArrowsWithThatchedBoats
	//长者之风
	TheWindOfTheElderly
	//偃旗息鼓
	LowerBannersAndMuffleDrums
	//出其不意
	TakeBySurprise
	//庐江上甲
	LuJiangRiverOverArmoured
	//谦让
	Humility
	//矢志不移
	OnesResolveIsUnshaken
	//所向披靡
	EverTriumphant
	//破阵摧坚
	BreakingThroughTheFormationAndDestroyingTheFirm
	//杯蛇鬼车
	CupSnakeGhostCar
	//四面楚歌
	BeBesiegedOnAllSides
	//瞋目横矛
	AngryEyeHorizontalSpear
	//暴戾无仁
	ViolentAndHeartless
	//沉沙决水
	SinkingSandAndBreakingWater
	//诱敌深入
	LureTheEnemyInDeep
	//一骑当千
	IkkiTousen
	//乘敌不虞
	RidingTheEnemyWithoutFear
	//百骑劫营
	HundredCavalryRobberyBattalions
	//折冲御侮
	RepelForeignAggression
	//合军聚众
	GatheringOfTroops
	//克敌制胜
	VanquishTheEnemy
	//兵锋
	TheSharpnessOfMilitaryStrength
	//风助火势
	WindAssistedFire
	//智计
	IntelligentStrategy
	//勇者得前
	TheBraveLeadTheWay
	//白马义从
	WhiteHorseFollowsWithLoyalty
	//文武双全
	BeAdeptWithBothPenAndSword
	//兵无常势
	WorkOutMeasuresToSuitLocalConditions
	//藤甲兵
	TengjiaSoldier
	//陷阵营
	FallIntoCamp
	//西凉铁骑
	XiLiangIronCavalry
	//虎豹骑
	TigerAndLeopardCavalry
	//白毦兵
	WhiteArmy
	//横扫千军
	SweepAwayTheMillionsOfEnemyTroops
	//用武通神
	UseMartialArtsToConnectWithGods
	//绝地反击
	JediCounterattack
	//梦中弑臣
	KillingMinisterInDream
	//伪书相间
	FakeBooksAlternateWithEachOther
	//奇计良谋
	CleverPlanAndCleverPlan
	//万箭齐发
	TenThousandArrowsShotAtOnce
	//舌战群儒
	LectureField
	//大戟士
	GreatHalberdWarrior
	//义心昭烈
	TheHeartOfRighteousnessShines
	//卧薪尝胆
	SleepOnTheBrushwoodAndTasteTheGall
	//倾国倾城
	BeautyWhichOverthrowsStatesAndCities
	//象兵
	ElephantSoldier
	//虎卫军
	TigerGuardArmy
	//锦帆军
	JinFanArmy
	//青州兵
	QingZhouSoldier
	//乘胜长驱
	RidingOnTheVictoryDrive
	//一力拒守
	RefuseToDefendWithOneForce
	//运筹决算
	DecisionMakingThroughOperationsResearch
	//铁骑驱驰
	IronHorseDrive
	//形机军略
	MilitaryStrategyForFormAircraft
	//箕形阵
	DustpanFormation
	//焰逐风飞
	FlamesFlyingInTheWind
	//威谋靡亢
	IntenseAndPowerful
	//破军威胜
	BreakingThroughTheArmyAndWinningVictories
	//速乘其利
	TakeAdvantageOfQuickly
	//士争先赴
	ScholarsStriveToGoFirst
	//众动万计
	CrowdMovesTenThousandCounts
	//先成其虑
	TakeCareOfYourselfFirst
	//竭力佐谋
	MakeEveryEffortToAssistInPlanning
	//临机制胜
	SeizeTheOpportunityToWin
	//忠勇义烈
	LoyalAndBraveMartyrs
	//气凌三军
	TemperamentSurpassesTheThreeArmies
	//鬼神霆威
	GhostGodThunderForce
	//焚辎营垒
	ToBurnBarracks
	//虎踞鹰扬
	TigerCrouchingAndEagleSoaring
	//三势阵
	ThreePotentialArray
	//当锋摧决
	WhenTheFrontIsDestroyed
	//刚勇无前
	StrongAndBraveWithoutAdvance
	//引弦力战
	LeadStringBattle
	//血刃争锋
	BloodBladeBattle
	//决水溃城
	BreakingThroughTheWaterAndCrushingTheCity
	//裸衣血战
	NakedBloodBattle
	//婴城自守
	InfantryCitySelfDefense
	//独行赴斗
	TravelingAloneToFight
	//据水断桥
	BrokenBridgeByWater
	//武锋阵
	WuFengArray
	//绝其汲道
	EliminateItAndDrawFromIt
	//击其惰归
	StrikeItsLazyReturn
	//千里走单骑
	ThousandMileWalkingSingleRider
	//形一阵
	ShapelyArray
	//潜龙阵
	HiddenDragonArray
	//掣刀斫敌
	PullingSwordsAndChoppingEnemies
	//强攻
	TakeByStorm
	//净化
	Purify
	//驱散
	Disperse
	//挟势弄权
	TakingAdvantageOfTheSituationToGainPower
	//传檄宣威
	ToSpreadRumorsAndProclaimPower
	//整装待发
	BeFullyEquippedFor
	//敛众而击
	GatherTheCrowdAndStrike
	//妖术
	BlackArt
	//一举歼灭
	AnnihilateInOneFellSwoop
	//不辱使命
	HaveSucceededInCarryingOutAnAssignment
	//暗藏玄机
	HiddenMystery
	//诈降
	PretendToSurrender
	//屠几上肉
	SlaughterMeatOnTable
	//左右开弓
	DrawTheBowBothOnTheLeftAndRight
	//骁勇善战
	BraveAndBattleWise
	//御敌屏障
	DefensiveBarrier
	//料事如神
	ForetellLikeProphet
	//挫志怒袭
	FrustrationAndAngerAttack
	//避实击虚
	AvoidTheSolidAndStrikeTheWeak
	//奋突
	RiseUpBravely
	//骑虎难下
	RideTigerHardToGetOff
	//才器过人
	OutstandingTalent
	//弯弓饮羽
	BendTheBowAndDrinkTheFeathers
	//落凤
	FallingPhoenix
	//鲁莽
	Reckless
	//落雷
	Thunderbolt
	//声东击西
	MakeFeintToTheEastButAttackInTheWest
	//风声鹤唳
	SoundOfTheWindAndTheCryOfTheStork
	//天降火雨
	RainOfFireFromTheSky
	//手起刀落
	CutDown
	//后发制人
	GainMasteryByStrikingOnlyAfterTheEnemyHasStruck
	//白眉
	EyebrowedThrush
	//千里驰援
	ThousandsOfMilesOfSupport
	//坐守孤城
	SittingInAnIsolatedCity
	//轻勇飞燕
	FearlessAndBraveFlyingSwallow
	//纵兵劫掠
	LeavingSoldiersToPlunder
	//唇枪舌战
	HaveVerbalBattleWithSomebody
	//机略纵横
	MachineStrategyVerticalAndHorizontal
	//自愈
	SelfHealing
	//暴敛四方
	OverwhelmingAllDirections
	//斗智
	BattleOfWits
	//包扎
	Dress
	//挑衅
	Provoke
	//横扫
	SweepAway
	//追伤
	ChasingInjury
	//机鉴先识
	OpportunityIdentificationFirst
	//古之恶来
	AncientEvilComes
	//神机莫测
	DivinelyInspiredStratagem
	//陷阵突袭
	RaidInFormation
	//水淹七军
	FloodedSeventhArmy
	//计定谋决
	PlanAndDecide
	//死战不退
	NeverRetreatFromDeadBattle
	//经天纬地
	AbilityToRuleTheCountry
	//虎痴
	TigerIdiot
	//金城汤池
	RampartsOfMetalsAndAMoatOfHotWater
	//固若金汤
	Impregnable
	//十面埋伏
	AmbushOnAllSides
	//十二奇策
	TwelveWonderfulStrategies
	//长驱直入
	MarchInto
	//刚烈不屈
	StrongAndUnyielding
	//抬棺决战
	TheBattleOfCarryingCoffin
	//精练策数
	RefinedStrategies
	//震骇四境
	ShockingFourRealms
	//垂心万物
	FocusingOnAllThings
	//持军毅重
	HoldTheArmyWithDeterminationAndDetermination
	//临战先登
	ToAscendBeforeBattle
	//暗渡陈仓
	AdvancingSecretlyByUnknownPath
	//将行其疾
	ToCureOnesSpeed
	//十胜十败
	TenWinsAndTenLosses
	//沉断机谋
	DecisiveStrategy
	//仁德载世
	BenevolentAndVirtuousThroughoutTheWorld
	//槊血纵横
	BloodyAndUnrestrained
	//威震华夏
	BecomeFamousAndFearInspiringThroughoutChina
	//锦囊妙计
	BrocadeBagAndCleverPlan
	//才辩机捷
	BeQuickInDebatingOpportunities
	//誓守无降
	PromiseToKeepWithoutSurrender
	//枪舞如风
	GunDanceLikeTheWind
	//刀出如霆
	TheKnifeLikeThunderbolt
	//将门虎女
	GeneralBraveGirl
	//鸱苕凤姿
	ThePostureOfAPhoenixWithAChickAndASweetPotato
	//奇兵间道
	SpecialSoldierPassRoad
	//百步穿杨
	HitTheTargetAtEveryShot
	//一身是胆
	BebraveAllThrough
	//燕人咆哮
	YanPeopleRoar
	//处兹不惑
	InChaosNotConfused
	//奋矛英姿
	BraveSpearHeroicPose
	//工神
	TheGodOfCraftsmen
	//以逸待劳
	WaitAtOnesEaseForTheFatigued
	//溯江摇橹
	ChasingTheRiverAndRidingRows
	//弓腰姬
	BowWaistConcubine
	//火烧连营
	FireJointVenture
	//江天长焰
	RiverFireFlame
	//暗箭难防
	HiddenArrowsAreDifficultToGuardAgainst
	//国士将风
	CountryPersonGeneralStyle
	//济贫好施
	HelpingThePoorAndGivingGenerously
	//坐断东南
	SittingIntheSoutheast
	//锦帆百翎
	JinfanArmyHundredFeathers
	//肉身铁壁
	CorporealIronWall
	//白衣渡江
	CrossingTheRiverInWhiteClothes
	//神射
	DivineEjaculation
	//江东猛虎
	JiangdongTiger
	//校胜帷幄
	WinsTent
	//神火计
	DivineFireMeter
	//苦肉计
	InjuryOnOneself
	//勇烈持重
	BraveAndResolute
	//江东小霸王
	JiangdongLittleOverlord
	//挥兵谋胜
	WieldTroopsToSeekVictory
	//围师必阙
	SurroundingTheTeacherMustBePalace
	//高橹连营
	HighWoodenPaddlesConnectedToTheCamp
	//符命自立
	FumingSelfReliance
	//南蛮渠魁
	NanManQuKui
	//兴云布雨
	MakeCloudAndRain
	//酒池肉林
	ExtravagantOrgy
	//天下无双
	Unique
	//傲睨王侯
	ProudPrince
	//毒泉拒蜀
	PoisonousSpringRefusesShu
	//监统震军
	SuperviseLeadAndSeizureArmy
	//竭忠尽智
	WithAllTheHeart
	//狮子奋迅
	TheLionFliesFast
	//火神英风
	FireGodHeroStyle
	//累世立名
	EstablishingNameThroughGenerations
	//五雷轰顶
	ThunderStruck
	//窃幸乘宠
	StealingLuckAndRidingPets
	//振军击营
	ExcitingArmyAttackCamp
	//鸩毒
	PoisonedWine
	//登锋陷阵
	ChargeIntoTheEnemyRanks
	//搦战群雄
	ToSeizeThePowerOfGroupOfHeroes
	//勇冠三军
	PeerlessOrMatchlessBraveryOrValour
	//青囊
	MedicalPractice
	//金丹秘术
	GoldenPillSecretTechnique
	//闭月
	ClosedMoon
	//百计多谋
	HundredStrategiesAndManyStrategies
	//顾盼生姿
	LookAroundCharmingly
	//云聚影从
	CloudGatheringShadowFrom
	//胡笳余音
	HuJiaLingeringSound
	//腹背受敌
	ScyllaAndCharybdis
	//威风凛凛
	AweInspiring
	//骁健神行
	VigorousAndWalk
	//暗潮涌动
	UndercurrentSurge
	//镇压黄巾
	SuppressYellowScarves
	//飞沙走石
	FlyingSandAndRollingPebbles
	//扶危定倾
	DeliverTheCountryFromDistress
	//聚石成金
	AggregateStoneIntoGold
	//短兵相见
	CloseQuarters
	//疑城
	ShamCastles
	//神上使
	DivineEnvoy
	//进言
	Introduction
	//非攻制胜
	NonOffensiveVictory
	//疾风骤雨
	StrongWindAndSwiftRain
)

func (b TacticId) String() string {
	switch b {
	//飞沙走石
	case FlyingSandAndRollingPebbles:
		return "飞沙走石"
	//扶危定倾
	case DeliverTheCountryFromDistress:
		return "扶危定倾"
	//聚石成金
	case AggregateStoneIntoGold:
		return "聚石成金"
	//短兵相见
	case CloseQuarters:
		return "短兵相见"
	//疑城
	case ShamCastles:
		return "疑城"
	//神上使
	case DivineEnvoy:
		return "神上使"
	//进言
	case Introduction:
		return "进言"
	//非攻制胜
	case NonOffensiveVictory:
		return "非攻制胜"
	//疾风骤雨
	case StrongWindAndSwiftRain:
		return "疾风骤雨"
	//鹰视狼顾
	case ClearEyedAndMalicious:
		return "鹰视狼顾"
	//士别三日
	case ThreeDaysOfSeparation:
		return "士别三日"
	//熯天炽地
	case TheSkyIsBlazing:
		return "熯天炽地"
	//乱世奸雄
	case TraitorInTroubledTimes:
		return "乱世奸雄"
		//盛气凌敌
	case OverwhelmingTheEnemyWithVigour:
		return "盛气凌敌"
	//挫锐
	case Demoralize:
		return "Demoralize"
	//守而必固
	case ToKeepAndBeFirm:
		return "守而必固"
	//横戈跃马
	case Gallant:
		return "横戈跃马"
		//暂避其锋
	case TakeRefugeFromEnemies:
		return "暂避其锋"
		//魅惑
	case Charming:
		return "魅惑"
		//抚揖军民
	case AppeaseArmyAndPeople:
		return "抚揖军民"
		//镇扼防拒
	case SuppressChokesAndPreventRefusals:
		return "镇扼防拒"
		//锋矢阵
	case FrontalVectorArray:
		return "锋矢阵"
		//刮骨疗毒
	case Curettage:
		return "刮骨疗毒"
		//义胆雄心
	case BraveAmbition:
		return "义胆雄心"
		//夺魂挟魄
	case SeizeTheSoul:
		return "夺魂挟魄"
		//火炽原燎
	case BlazingWildfire:
		return "火炽原燎"
		//连环计
	case InterlockedStratagems:
		return "连环计"
		//太平道法
	case TaipingLaw:
		return "太平道法"
		//无当飞军
	case WuDangFlyArmy:
		return "无当飞军"
		//神机妙算
	case CleverStrategyAndShrewdTactics:
		return "神机妙算"
		//八门金锁阵
	case EightGateGoldenLockArray:
		return "八门金锁阵"
		//草船借箭
	case BorrowArrowsWithThatchedBoats:
		return "草船借箭"
		//长者之风
	case TheWindOfTheElderly:
		return "长者之风"
		//偃旗息鼓
	case LowerBannersAndMuffleDrums:
		return "偃旗息鼓"
		//出其不意
	case TakeBySurprise:
		return "出其不意"
		//庐江上甲
	case LuJiangRiverOverArmoured:
		return "庐江上甲"
		//谦让
	case Humility:
		return "谦让"
		//矢志不移
	case OnesResolveIsUnshaken:
		return "矢志不移"
		//所向披靡
	case EverTriumphant:
		return "所向披靡"
		//破阵摧坚
	case BreakingThroughTheFormationAndDestroyingTheFirm:
		return "破阵摧坚"
		//杯蛇鬼车
	case CupSnakeGhostCar:
		return "杯蛇鬼车"
		//四面楚歌
	case BeBesiegedOnAllSides:
		return "四面楚歌"
		//瞋目横矛
	case AngryEyeHorizontalSpear:
		return "瞋目横矛"
		//暴戾无仁
	case ViolentAndHeartless:
		return "暴戾无仁"
		//沉沙决水
	case SinkingSandAndBreakingWater:
		return "沉沙决水"
		//诱敌深入
	case LureTheEnemyInDeep:
		return "诱敌深入"
		//一骑当千
	case IkkiTousen:
		return "一骑当千"
		//乘敌不虞
	case RidingTheEnemyWithoutFear:
		return "乘敌不虞"
		//百骑劫营
	case HundredCavalryRobberyBattalions:
		return "百骑劫营"
		//折冲御侮
	case RepelForeignAggression:
		return "折冲御侮"
		//合军聚众
	case GatheringOfTroops:
		return "合军聚众"
		//克敌制胜
	case VanquishTheEnemy:
		return "克敌制胜"
		//兵锋
	case TheSharpnessOfMilitaryStrength:
		return "兵锋"
		//风助火势
	case WindAssistedFire:
		return "风助火势"
		//智计
	case IntelligentStrategy:
		return "智计"
		//勇者得前
	case TheBraveLeadTheWay:
		return "勇者得前"
		//白马义从
	case WhiteHorseFollowsWithLoyalty:
		return "白马义从"
		//文武双全
	case BeAdeptWithBothPenAndSword:
		return "文武双全"
		//兵无常势
	case WorkOutMeasuresToSuitLocalConditions:
		return "兵无常势"
		//藤甲兵
	case TengjiaSoldier:
		return "藤甲兵"
		//陷阵营
	case FallIntoCamp:
		return "陷阵营"
		//西凉铁骑
	case XiLiangIronCavalry:
		return "西凉铁骑"
		//虎豹骑
	case TigerAndLeopardCavalry:
		return "虎豹骑"
		//白毦兵
	case WhiteArmy:
		return "白毦兵"
		//横扫千军
	case SweepAwayTheMillionsOfEnemyTroops:
		return "横扫千军"
		//用武通神
	case UseMartialArtsToConnectWithGods:
		return "用武通神"
		//绝地反击
	case JediCounterattack:
		return "绝地反击"
		//梦中弑臣
	case KillingMinisterInDream:
		return "梦中弑臣"
		//伪书相间
	case FakeBooksAlternateWithEachOther:
		return "伪书相间"
		//奇计良谋
	case CleverPlanAndCleverPlan:
		return "奇计良谋"
		//万箭齐发
	case TenThousandArrowsShotAtOnce:
		return "万箭齐发"
		//舌战群儒
	case LectureField:
		return "舌战群儒"
	//大戟士
	case GreatHalberdWarrior:
		return "大戟士"
	//义心昭烈
	case TheHeartOfRighteousnessShines:
		return "义心昭烈"
	//卧薪尝胆
	case SleepOnTheBrushwoodAndTasteTheGall:
		return "卧薪尝胆"
		//倾国倾城
	case BeautyWhichOverthrowsStatesAndCities:
		return "倾国倾城"
	//象兵
	case ElephantSoldier:
		return "象兵"
	//虎卫军
	case TigerGuardArmy:
		return "虎卫军"
		//锦帆军
	case JinFanArmy:
		return "锦帆军"
	//青州兵
	case QingZhouSoldier:
		return "青州兵"
	//乘胜长驱
	case RidingOnTheVictoryDrive:
		return "乘胜长驱"
		//一力拒守
	case RefuseToDefendWithOneForce:
		return "一力拒守"
		//运筹决算
	case DecisionMakingThroughOperationsResearch:
		return "运筹决算"
	//铁骑驱驰
	case IronHorseDrive:
		return "铁骑驱驰"
	//形机军略
	case MilitaryStrategyForFormAircraft:
		return "形机军略"
	//箕形阵
	case DustpanFormation:
		return "箕形阵"
	//焰逐风飞
	case FlamesFlyingInTheWind:
		return "焰逐风飞"
		//威谋靡亢
	case IntenseAndPowerful:
		return "威谋靡亢"
	//破军威胜
	case BreakingThroughTheArmyAndWinningVictories:
		return "破军威胜"
	//速乘其利
	case TakeAdvantageOfQuickly:
		return "速乘其利"
	//士争先赴
	case ScholarsStriveToGoFirst:
		return "士争先赴"
		//先成其虑
	case TakeCareOfYourselfFirst:
		return "先成其虑"
		//竭力佐谋
	case MakeEveryEffortToAssistInPlanning:
		return "竭力佐谋"
		//临机制胜
	case SeizeTheOpportunityToWin:
		return "临机制胜"
	//忠勇义烈
	case LoyalAndBraveMartyrs:
		return "忠勇义烈"
		//气凌三军
	case TemperamentSurpassesTheThreeArmies:
		return "气凌三军"
	//鬼神霆威
	case GhostGodThunderForce:
		return "鬼神霆威"
		//焚辎营垒
	case ToBurnBarracks:
		return "焚辎营垒"
		//虎踞鹰扬
	case TigerCrouchingAndEagleSoaring:
		return "虎踞鹰扬"
	//三势阵
	case ThreePotentialArray:
		return "三势阵"
		//当锋摧决
	case WhenTheFrontIsDestroyed:
		return "当锋摧决"
	//刚勇无前
	case StrongAndBraveWithoutAdvance:
		return "刚勇无前"
	//引弦力战
	case LeadStringBattle:
		return "引弦力战"
	//血刃争锋
	case BloodBladeBattle:
		return "血刃争锋"
	//决水溃城
	case BreakingThroughTheWaterAndCrushingTheCity:
		return "决水溃城"
		//裸衣血战
	case NakedBloodBattle:
		return "裸衣血战"
	//婴城自守
	case InfantryCitySelfDefense:
		return "婴城自守"
		//独行赴斗
	case TravelingAloneToFight:
		return "独行赴斗"
		//据水断桥
	case BrokenBridgeByWater:
		return "据水断桥"
	//武锋阵
	case WuFengArray:
		return "武锋阵"
	//绝其汲道
	case EliminateItAndDrawFromIt:
		return "绝其汲道"
		//击其惰归
	case StrikeItsLazyReturn:
		return "击其惰归"
		//千里走单骑
	case ThousandMileWalkingSingleRider:
		return "千里走单骑"
	//形一阵
	case ShapelyArray:
		return "形一阵"
	//潜龙阵
	case HiddenDragonArray:
		return "潜龙阵"
		//掣刀斫敌
	case PullingSwordsAndChoppingEnemies:
		return "掣刀斫敌"
		//强攻
	case TakeByStorm:
		return "强攻"
		//净化
	case Purify:
		return "净化"
		//驱散
	case Disperse:
		return "驱散"
		//挟势弄权
	case TakingAdvantageOfTheSituationToGainPower:
		return "挟势弄权"
		//传檄宣威
	case ToSpreadRumorsAndProclaimPower:
		return "传檄宣威"
	//整装待发
	case BeFullyEquippedFor:
		return "整装待发"
	//敛众而击
	case GatherTheCrowdAndStrike:
		return "敛众而击"
		//妖术
	case BlackArt:
		return "妖术"
	//一举歼灭
	case AnnihilateInOneFellSwoop:
		return "一举歼灭"
		//不辱使命
	case HaveSucceededInCarryingOutAnAssignment:
		return "不辱使命"
		//暗藏玄机
	case HiddenMystery:
		return "暗藏玄机"
		//诈降
	case PretendToSurrender:
		return "诈降"
		//屠几上肉
	case SlaughterMeatOnTable:
		return "屠几上肉"
		//左右开弓
	case DrawTheBowBothOnTheLeftAndRight:
		return "左右开弓"
		//骁勇善战
	case BraveAndBattleWise:
		return "骁勇善战"
		//御敌屏障
	case DefensiveBarrier:
		return "御敌屏障"
		//众动万计
	case CrowdMovesTenThousandCounts:
		return "众动万计"
		//料事如神
	case ForetellLikeProphet:
		return "料事如神"
		//挫志怒袭
	case FrustrationAndAngerAttack:
		return "挫志怒袭"
		//避实击虚
	case AvoidTheSolidAndStrikeTheWeak:
		return "避实击虚"
		//奋突
	case RiseUpBravely:
		return "奋突"
		//骑虎难下
	case RideTigerHardToGetOff:
		return "骑虎难下"
		//才器过人
	case OutstandingTalent:
		return "才器过人"
		//弯弓饮羽
	case BendTheBowAndDrinkTheFeathers:
		return "弯弓饮羽"
	//落凤
	case FallingPhoenix:
		return "落凤"
	//鲁莽
	case Reckless:
		return "鲁莽"
	//落雷
	case Thunderbolt:
		return "落雷"
	//声东击西
	case MakeFeintToTheEastButAttackInTheWest:
		return "声东击西"
		//风声鹤唳
	case SoundOfTheWindAndTheCryOfTheStork:
		return "风声鹤唳"
	//天降火雨
	case RainOfFireFromTheSky:
		return "天降火雨"
	//手起刀落
	case CutDown:
		return "手起刀落"
	//后发制人
	case GainMasteryByStrikingOnlyAfterTheEnemyHasStruck:
		return "后发制人"
	//白眉
	case EyebrowedThrush:
		return "白眉"
	//千里驰援
	case ThousandsOfMilesOfSupport:
		return "千里驰援"
	//坐守孤城
	case SittingInAnIsolatedCity:
		return "坐守孤城"
	//轻勇飞燕
	case FearlessAndBraveFlyingSwallow:
		return "轻勇飞燕"
	//纵兵劫掠
	case LeavingSoldiersToPlunder:
		return "纵兵劫掠"
	//唇枪舌战
	case HaveVerbalBattleWithSomebody:
		return "唇枪舌战"
	//机略纵横
	case MachineStrategyVerticalAndHorizontal:
		return "机略纵横"
	//自愈
	case SelfHealing:
		return "自愈"
	//暴敛四方
	case OverwhelmingAllDirections:
		return "暴敛四方"
	//斗智
	case BattleOfWits:
		return "斗智"
	//包扎
	case Dress:
		return "包扎"
	//挑衅
	case Provoke:
		return "挑衅"
	//横扫
	case SweepAway:
		return "横扫"
	//追伤
	case ChasingInjury:
		return "追伤"
	//机鉴先识
	case OpportunityIdentificationFirst:
		return "机鉴先识"
	//古之恶来
	case AncientEvilComes:
		return "古之恶来"
	//神机莫测
	case DivinelyInspiredStratagem:
		return "神机莫测"
	//陷阵突袭
	case RaidInFormation:
		return "陷阵突袭"
	//水淹七军
	case FloodedSeventhArmy:
		return "水淹七军"
	//计定谋决
	case PlanAndDecide:
		return "计定谋决"
	//死战不退
	case NeverRetreatFromDeadBattle:
		return "死战不退"
	//经天纬地
	case AbilityToRuleTheCountry:
		return "经天纬地"
	//虎痴
	case TigerIdiot:
		return "虎痴"
	//金城汤池
	case RampartsOfMetalsAndAMoatOfHotWater:
		return "金城汤池"
	//固若金汤
	case Impregnable:
		return "固若金汤"
	//十面埋伏
	case AmbushOnAllSides:
		return "十面埋伏"
	//十二奇策
	case TwelveWonderfulStrategies:
		return "十二奇策"
	//长驱直入
	case MarchInto:
		return "长驱直入"
	//刚烈不屈
	case StrongAndUnyielding:
		return "刚烈不屈"
	//抬棺决战
	case TheBattleOfCarryingCoffin:
		return "抬棺决战"
	//精练策数
	case RefinedStrategies:
		return "精练策数"
	//震骇四境
	case ShockingFourRealms:
		return "震骇四境"
	//垂心万物
	case FocusingOnAllThings:
		return "垂心万物"
	//持军毅重
	case HoldTheArmyWithDeterminationAndDetermination:
		return "持军毅重"
	//临战先登
	case ToAscendBeforeBattle:
		return "临战先登"
	//暗渡陈仓
	case AdvancingSecretlyByUnknownPath:
		return "暗渡陈仓"
	//将行其疾
	case ToCureOnesSpeed:
		return "将行其疾"
	//十胜十败
	case TenWinsAndTenLosses:
		return "十胜十败"
	//沉断机谋
	case DecisiveStrategy:
		return "沉断机谋"
	//仁德载世
	case BenevolentAndVirtuousThroughoutTheWorld:
		return "仁德载世"
	//槊血纵横
	case BloodyAndUnrestrained:
		return "槊血纵横"
	//威震华夏
	case BecomeFamousAndFearInspiringThroughoutChina:
		return "威震华夏"
	//锦囊妙计
	case BrocadeBagAndCleverPlan:
		return "锦囊妙计"
	//才辩机捷
	case BeQuickInDebatingOpportunities:
		return "才辩机捷"
	//誓守无降
	case PromiseToKeepWithoutSurrender:
		return "誓守无降"
	//枪舞如风
	case GunDanceLikeTheWind:
		return "枪舞如风"
	//刀出如霆
	case TheKnifeLikeThunderbolt:
		return "刀出如霆"
	//将门虎女
	case GeneralBraveGirl:
		return "将门虎女"
	//鸱苕凤姿
	case ThePostureOfAPhoenixWithAChickAndASweetPotato:
		return "鸱苕凤姿"
	//奇兵间道
	case SpecialSoldierPassRoad:
		return "奇兵间道"
	//百步穿杨
	case HitTheTargetAtEveryShot:
		return "百步穿杨"
	//一身是胆
	case BebraveAllThrough:
		return "一身是胆"
	//燕人咆哮
	case YanPeopleRoar:
		return "燕人咆哮"
	//处兹不惑
	case InChaosNotConfused:
		return "处兹不惑"
	//奋矛英姿
	case BraveSpearHeroicPose:
		return "奋矛英姿"
	//工神
	case TheGodOfCraftsmen:
		return "工神"
	//以逸待劳
	case WaitAtOnesEaseForTheFatigued:
		return "以逸待劳"
	//溯江摇橹
	case ChasingTheRiverAndRidingRows:
		return "溯江摇橹"
	//弓腰姬
	case BowWaistConcubine:
		return "弓腰姬"
	//火烧连营
	case FireJointVenture:
		return "火烧连营"
	//江天长焰
	case RiverFireFlame:
		return "江天长焰"
	//暗箭难防
	case HiddenArrowsAreDifficultToGuardAgainst:
		return "暗箭难防"
	//国士将风
	case CountryPersonGeneralStyle:
		return "国士将风"
	//济贫好施
	case HelpingThePoorAndGivingGenerously:
		return "济贫好施"
	//坐断东南
	case SittingIntheSoutheast:
		return "坐断东南"
	//锦帆百翎
	case JinfanArmyHundredFeathers:
		return "锦帆百翎"
	//肉身铁壁
	case CorporealIronWall:
		return "肉身铁壁"
	//白衣渡江
	case CrossingTheRiverInWhiteClothes:
		return "白衣渡江"
	//神射
	case DivineEjaculation:
		return "神射"
	//江东猛虎
	case JiangdongTiger:
		return "江东猛虎"
	//校胜帷幄
	case WinsTent:
		return "校胜帷幄"
	//神火计
	case DivineFireMeter:
		return "神火计"
	//苦肉计
	case InjuryOnOneself:
		return "苦肉计"
	//勇烈持重
	case BraveAndResolute:
		return "勇烈持重"
	//江东小霸王
	case JiangdongLittleOverlord:
		return "江东小霸王"
	//挥兵谋胜
	case WieldTroopsToSeekVictory:
		return "挥兵谋胜"
	//围师必阙
	case SurroundingTheTeacherMustBePalace:
		return "围师必阙"
	//高橹连营
	case HighWoodenPaddlesConnectedToTheCamp:
		return "高橹连营"
	//符命自立
	case FumingSelfReliance:
		return "符命自立"
	//南蛮渠魁
	case NanManQuKui:
		return "南蛮渠魁"
	//兴云布雨
	case MakeCloudAndRain:
		return "兴云布雨"
	//酒池肉林
	case ExtravagantOrgy:
		return "酒池肉林"
	//天下无双
	case Unique:
		return "天下无双"
	//傲睨王侯
	case ProudPrince:
		return "傲睨王侯"
	//毒泉拒蜀
	case PoisonousSpringRefusesShu:
		return "毒泉拒蜀"
	//监统震军
	case SuperviseLeadAndSeizureArmy:
		return "监统震军"
	//竭忠尽智
	case WithAllTheHeart:
		return "竭忠尽智"
	//狮子奋迅
	case TheLionFliesFast:
		return "狮子奋迅"
	//火神英风
	case FireGodHeroStyle:
		return "FireGodHeroStyle"
	//累世立名
	case EstablishingNameThroughGenerations:
		return "累世立名"
	//五雷轰顶
	case ThunderStruck:
		return "五雷轰顶"
	//窃幸乘宠
	case StealingLuckAndRidingPets:
		return "窃幸乘宠"
	//振军击营
	case ExcitingArmyAttackCamp:
		return "振军击营"
	//鸩毒
	case PoisonedWine:
		return "鸩毒"
	//登锋陷阵
	case ChargeIntoTheEnemyRanks:
		return "登锋陷阵"
	//搦战群雄
	case ToSeizeThePowerOfGroupOfHeroes:
		return "搦战群雄"
	//勇冠三军
	case PeerlessOrMatchlessBraveryOrValour:
		return "勇冠三军"
	//青囊
	case MedicalPractice:
		return "青囊"
	//金丹秘术
	case GoldenPillSecretTechnique:
		return "金丹秘术"
	//闭月
	case ClosedMoon:
		return "闭月"
	//百计多谋
	case HundredStrategiesAndManyStrategies:
		return "百计多谋"
	//顾盼生姿
	case LookAroundCharmingly:
		return "顾盼生姿"
	//云聚影从
	case CloudGatheringShadowFrom:
		return "云聚影从"
	//胡笳余音
	case HuJiaLingeringSound:
		return "胡笳余音"
	//腹背受敌
	case ScyllaAndCharybdis:
		return "腹背受敌"
	//威风凛凛
	case AweInspiring:
		return "威风凛凛"
	//骁健神行
	case VigorousAndWalk:
		return "骁健神行"
	//暗潮涌动
	case UndercurrentSurge:
		return "暗潮涌动"
	//镇压黄巾
	case SuppressYellowScarves:
		return "镇压黄巾"
	}
	return "未知战法"
}

// ** 战法池，按类型划分 **
// 主动
var ActiveTacticsMap = make(map[TacticId]bool, 0)

// 被动
var PassiveTacticsMap = make(map[TacticId]bool, 0)

// 指挥
var CommandTacticsMap = make(map[TacticId]bool, 0)

// 突击
var AssaultTacticsMap = make(map[TacticId]bool, 0)

// 阵法
var TroopsTacticsMap = make(map[TacticId]bool, 0)

// 兵种
var ArmTacticsMap = make(map[TacticId]bool, 0)

// 准备回合的战法
var ActivePrepareTacticsMap = make(map[TacticId]bool, 0)

func init() {
	initTacticsMap()
}

func initTacticsMap() {
	//需要准备回合的战法
	ActivePrepareTacticsMap[AdvancingSecretlyByUnknownPath] = true
	ActivePrepareTacticsMap[ThunderStruck] = true

	//被动战法
	PassiveTacticsMap[PretendToSurrender] = true
	PassiveTacticsMap[EyebrowedThrush] = true
	PassiveTacticsMap[BraveSpearHeroicPose] = true
	PassiveTacticsMap[BraveAndResolute] = true
	PassiveTacticsMap[BloodBladeBattle] = true
	PassiveTacticsMap[BebraveAllThrough] = true
	PassiveTacticsMap[CorporealIronWall] = true
	PassiveTacticsMap[YanPeopleRoar] = true
	PassiveTacticsMap[ThousandMileWalkingSingleRider] = true
	PassiveTacticsMap[StrongAndUnyielding] = true
	PassiveTacticsMap[SelfHealing] = true
	PassiveTacticsMap[ScholarsStriveToGoFirst] = true
	PassiveTacticsMap[RiseUpBravely] = true
	PassiveTacticsMap[RidingOnTheVictoryDrive] = true
	PassiveTacticsMap[LeadStringBattle] = true
	PassiveTacticsMap[PlanAndDecide] = true
	PassiveTacticsMap[NeverRetreatFromDeadBattle] = true
	PassiveTacticsMap[ThreeDaysOfSeparation] = true
	PassiveTacticsMap[Charming] = true
	PassiveTacticsMap[TaipingLaw] = true
	PassiveTacticsMap[BraveAmbition] = true
	PassiveTacticsMap[OnesResolveIsUnshaken] = true
	PassiveTacticsMap[BeAdeptWithBothPenAndSword] = true
	PassiveTacticsMap[CrowdMovesTenThousandCounts] = true
	PassiveTacticsMap[ExtravagantOrgy] = true
	PassiveTacticsMap[FumingSelfReliance] = true
	PassiveTacticsMap[GainMasteryByStrikingOnlyAfterTheEnemyHasStruck] = true
	PassiveTacticsMap[GatheringOfTroops] = true
	PassiveTacticsMap[HelpingThePoorAndGivingGenerously] = true
	PassiveTacticsMap[HighWoodenPaddlesConnectedToTheCamp] = true
	PassiveTacticsMap[JediCounterattack] = true
	PassiveTacticsMap[LoyalAndBraveMartyrs] = true
	PassiveTacticsMap[NakedBloodBattle] = true
	PassiveTacticsMap[MarchInto] = true
	PassiveTacticsMap[RaidInFormation] = true
	PassiveTacticsMap[SpecialSoldierPassRoad] = true
	PassiveTacticsMap[StrongAndBraveWithoutAdvance] = true
	PassiveTacticsMap[TemperamentSurpassesTheThreeArmies] = true
	PassiveTacticsMap[TigerCrouchingAndEagleSoaring] = true
	PassiveTacticsMap[TigerIdiot] = true
	PassiveTacticsMap[WorkOutMeasuresToSuitLocalConditions] = true
	PassiveTacticsMap[ThePostureOfAPhoenixWithAChickAndASweetPotato] = true
	PassiveTacticsMap[JinfanArmyHundredFeathers] = true
	PassiveTacticsMap[BeQuickInDebatingOpportunities] = true
	PassiveTacticsMap[BloodyAndUnrestrained] = true
	PassiveTacticsMap[DivineEjaculation] = true
	PassiveTacticsMap[DivineFireMeter] = true
	PassiveTacticsMap[ChasingTheRiverAndRidingRows] = true
	PassiveTacticsMap[JiangdongLittleOverlord] = true
	//指挥战法
	CommandTacticsMap[NonOffensiveVictory] = true
	CommandTacticsMap[CountryPersonGeneralStyle] = true
	CommandTacticsMap[CrossingTheRiverInWhiteClothes] = true
	CommandTacticsMap[TheHeartOfRighteousnessShines] = true
	CommandTacticsMap[BrocadeBagAndCleverPlan] = true
	CommandTacticsMap[ProudPrince] = true
	CommandTacticsMap[BowWaistConcubine] = true
	CommandTacticsMap[CloudGatheringShadowFrom] = true
	CommandTacticsMap[WinsTent] = true
	CommandTacticsMap[TenWinsAndTenLosses] = true
	CommandTacticsMap[SurroundingTheTeacherMustBePalace] = true
	CommandTacticsMap[StealingLuckAndRidingPets] = true
	CommandTacticsMap[SittingIntheSoutheast] = true
	CommandTacticsMap[IronHorseDrive] = true
	CommandTacticsMap[OpportunityIdentificationFirst] = true
	CommandTacticsMap[GoldenPillSecretTechnique] = true
	CommandTacticsMap[OverwhelmingTheEnemyWithVigour] = true
	CommandTacticsMap[FocusingOnAllThings] = true
	CommandTacticsMap[Demoralize] = true
	CommandTacticsMap[ToKeepAndBeFirm] = true
	CommandTacticsMap[Gallant] = true
	CommandTacticsMap[TakeRefugeFromEnemies] = true
	CommandTacticsMap[SuppressChokesAndPreventRefusals] = true
	CommandTacticsMap[AppeaseArmyAndPeople] = true
	CommandTacticsMap[TraitorInTroubledTimes] = true
	CommandTacticsMap[ClearEyedAndMalicious] = true
	CommandTacticsMap[CleverStrategyAndShrewdTactics] = true
	CommandTacticsMap[TheWindOfTheElderly] = true
	CommandTacticsMap[UseMartialArtsToConnectWithGods] = true
	CommandTacticsMap[KillingMinisterInDream] = true
	CommandTacticsMap[CleverPlanAndCleverPlan] = true
	CommandTacticsMap[LectureField] = true
	CommandTacticsMap[BeFullyEquippedFor] = true
	CommandTacticsMap[DefensiveBarrier] = true
	CommandTacticsMap[RideTigerHardToGetOff] = true
	CommandTacticsMap[AbilityToRuleTheCountry] = true
	CommandTacticsMap[AncientEvilComes] = true
	CommandTacticsMap[FireGodHeroStyle] = true
	CommandTacticsMap[PretendToSurrender] = true
	CommandTacticsMap[HundredStrategiesAndManyStrategies] = true
	CommandTacticsMap[NanManQuKui] = true
	CommandTacticsMap[MakeCloudAndRain] = true
	CommandTacticsMap[MedicalPractice] = true
	CommandTacticsMap[ProudPrince] = true
	CommandTacticsMap[RampartsOfMetalsAndAMoatOfHotWater] = true
	CommandTacticsMap[RiverFireFlame] = true
	CommandTacticsMap[SuperviseLeadAndSeizureArmy] = true
	CommandTacticsMap[SuppressYellowScarves] = true
	CommandTacticsMap[TheGodOfCraftsmen] = true
	CommandTacticsMap[BenevolentAndVirtuousThroughoutTheWorld] = true
	CommandTacticsMap[WieldTroopsToSeekVictory] = true
	CommandTacticsMap[DeliverTheCountryFromDistress] = true
	//阵法
	TroopsTacticsMap[FrontalVectorArray] = true
	TroopsTacticsMap[EightGateGoldenLockArray] = true
	TroopsTacticsMap[DustpanFormation] = true
	TroopsTacticsMap[FallIntoCamp] = true
	TroopsTacticsMap[HiddenDragonArray] = true
	TroopsTacticsMap[ShapelyArray] = true
	TroopsTacticsMap[ThreePotentialArray] = true
	TroopsTacticsMap[WuFengArray] = true
	//兵种
	ArmTacticsMap[XiLiangIronCavalry] = true
	ArmTacticsMap[WuDangFlyArmy] = true
	ArmTacticsMap[GreatHalberdWarrior] = true
	ArmTacticsMap[JinFanArmy] = true
	ArmTacticsMap[QingZhouSoldier] = true
	ArmTacticsMap[TengjiaSoldier] = true
	ArmTacticsMap[TigerGuardArmy] = true
	ArmTacticsMap[WhiteArmy] = true
	ArmTacticsMap[WhiteHorseFollowsWithLoyalty] = true
	ArmTacticsMap[ElephantSoldier] = true
	//主动
	ActiveTacticsMap[Introduction] = true
	ActiveTacticsMap[DivineEnvoy] = true
	ActiveTacticsMap[ShamCastles] = true
	ActiveTacticsMap[CloseQuarters] = true
	ActiveTacticsMap[StrongWindAndSwiftRain] = true
	ActiveTacticsMap[FlyingSandAndRollingPebbles] = true
	ActiveTacticsMap[TheSharpnessOfMilitaryStrength] = true
	ActiveTacticsMap[ClosedMoon] = true
	ActiveTacticsMap[PullingSwordsAndChoppingEnemies] = true
	ActiveTacticsMap[SweepAwayTheMillionsOfEnemyTroops] = true
	ActiveTacticsMap[ShockingFourRealms] = true
	ActiveTacticsMap[Purify] = true
	ActiveTacticsMap[Provoke] = true
	ActiveTacticsMap[PromiseToKeepWithoutSurrender] = true
	ActiveTacticsMap[FearlessAndBraveFlyingSwallow] = true
	ActiveTacticsMap[HaveSucceededInCarryingOutAnAssignment] = true
	ActiveTacticsMap[DecisiveStrategy] = true
	ActiveTacticsMap[Dress] = true
	ActiveTacticsMap[DrawTheBowBothOnTheLeftAndRight] = true
	ActiveTacticsMap[DivinelyInspiredStratagem] = true
	ActiveTacticsMap[ChargeIntoTheEnemyRanks] = true
	ActiveTacticsMap[BraveAndBattleWise] = true
	ActiveTacticsMap[BecomeFamousAndFearInspiringThroughoutChina] = true
	ActiveTacticsMap[BeautyWhichOverthrowsStatesAndCities] = true
	ActiveTacticsMap[BattleOfWits] = true
	ActiveTacticsMap[AvoidTheSolidAndStrikeTheWeak] = true
	ActiveTacticsMap[AweInspiring] = true
	ActiveTacticsMap[AnnihilateInOneFellSwoop] = true
	ActiveTacticsMap[TheKnifeLikeThunderbolt] = true
	ActiveTacticsMap[TheLionFliesFast] = true
	ActiveTacticsMap[InChaosNotConfused] = true
	ActiveTacticsMap[WithAllTheHeart] = true
	ActiveTacticsMap[WindAssistedFire] = true
	ActiveTacticsMap[WaitAtOnesEaseForTheFatigued] = true
	ActiveTacticsMap[VigorousAndWalk] = true
	ActiveTacticsMap[VanquishTheEnemy] = true
	ActiveTacticsMap[UndercurrentSurge] = true
	ActiveTacticsMap[TwelveWonderfulStrategies] = true
	ActiveTacticsMap[TravelingAloneToFight] = true
	ActiveTacticsMap[ToSeizeThePowerOfGroupOfHeroes] = true
	ActiveTacticsMap[ToSpreadRumorsAndProclaimPower] = true
	ActiveTacticsMap[ToBurnBarracks] = true
	ActiveTacticsMap[ToAscendBeforeBattle] = true
	ActiveTacticsMap[Thunderbolt] = true
	ActiveTacticsMap[ThunderStruck] = true
	ActiveTacticsMap[ThousandsOfMilesOfSupport] = true
	ActiveTacticsMap[TheBattleOfCarryingCoffin] = true
	ActiveTacticsMap[TenThousandArrowsShotAtOnce] = true
	ActiveTacticsMap[TakeByStorm] = true
	ActiveTacticsMap[SweepAway] = true
	ActiveTacticsMap[StrikeItsLazyReturn] = true
	ActiveTacticsMap[SoundOfTheWindAndTheCryOfTheStork] = true
	ActiveTacticsMap[SleepOnTheBrushwoodAndTasteTheGall] = true
	ActiveTacticsMap[SlaughterMeatOnTable] = true
	ActiveTacticsMap[SittingInAnIsolatedCity] = true
	ActiveTacticsMap[SinkingSandAndBreakingWater] = true
	ActiveTacticsMap[SeizeTheOpportunityToWin] = true
	ActiveTacticsMap[ScyllaAndCharybdis] = true
	ActiveTacticsMap[RidingTheEnemyWithoutFear] = true
	ActiveTacticsMap[RefuseToDefendWithOneForce] = true
	ActiveTacticsMap[RefinedStrategies] = true
	ActiveTacticsMap[Reckless] = true
	ActiveTacticsMap[RainOfFireFromTheSky] = true
	ActiveTacticsMap[Purify] = true
	ActiveTacticsMap[PullingSwordsAndChoppingEnemies] = true
	ActiveTacticsMap[Provoke] = true
	ActiveTacticsMap[PromiseToKeepWithoutSurrender] = true
	ActiveTacticsMap[PoisonousSpringRefusesShu] = true
	ActiveTacticsMap[PoisonedWine] = true
	ActiveTacticsMap[OverwhelmingAllDirections] = true
	ActiveTacticsMap[MachineStrategyVerticalAndHorizontal] = true
	ActiveTacticsMap[MilitaryStrategyForFormAircraft] = true
	ActiveTacticsMap[MakeEveryEffortToAssistInPlanning] = true
	ActiveTacticsMap[MakeFeintToTheEastButAttackInTheWest] = true
	ActiveTacticsMap[LureTheEnemyInDeep] = true
	ActiveTacticsMap[LookAroundCharmingly] = true
	ActiveTacticsMap[IntenseAndPowerful] = true
	ActiveTacticsMap[LeavingSoldiersToPlunder] = true
	ActiveTacticsMap[IntelligentStrategy] = true
	ActiveTacticsMap[InfantryCitySelfDefense] = true
	ActiveTacticsMap[Impregnable] = true
	ActiveTacticsMap[HitTheTargetAtEveryShot] = true
	ActiveTacticsMap[HuJiaLingeringSound] = true
	ActiveTacticsMap[HoldTheArmyWithDeterminationAndDetermination] = true
	ActiveTacticsMap[HiddenArrowsAreDifficultToGuardAgainst] = true
	ActiveTacticsMap[GeneralBraveGirl] = true
	ActiveTacticsMap[GunDanceLikeTheWind] = true
	ActiveTacticsMap[HaveVerbalBattleWithSomebody] = true
	ActiveTacticsMap[Curettage] = true
	ActiveTacticsMap[GatherTheCrowdAndStrike] = true
	ActiveTacticsMap[FlamesFlyingInTheWind] = true
	ActiveTacticsMap[FloodedSeventhArmy] = true
	ActiveTacticsMap[TheSkyIsBlazing] = true
	ActiveTacticsMap[SeizeTheSoul] = true
	ActiveTacticsMap[BlazingWildfire] = true
	ActiveTacticsMap[InterlockedStratagems] = true
	ActiveTacticsMap[BorrowArrowsWithThatchedBoats] = true
	ActiveTacticsMap[OutstandingTalent] = true
	ActiveTacticsMap[LowerBannersAndMuffleDrums] = true
	ActiveTacticsMap[TakeBySurprise] = true
	ActiveTacticsMap[LuJiangRiverOverArmoured] = true
	ActiveTacticsMap[Humility] = true
	ActiveTacticsMap[EverTriumphant] = true
	ActiveTacticsMap[BreakingThroughTheFormationAndDestroyingTheFirm] = true
	ActiveTacticsMap[CupSnakeGhostCar] = true
	ActiveTacticsMap[AngryEyeHorizontalSpear] = true
	ActiveTacticsMap[AdvancingSecretlyByUnknownPath] = true
	ActiveTacticsMap[BeBesiegedOnAllSides] = true
	ActiveTacticsMap[AmbushOnAllSides] = true
	ActiveTacticsMap[BlackArt] = true
	ActiveTacticsMap[BreakingThroughTheArmyAndWinningVictories] = true
	ActiveTacticsMap[BreakingThroughTheWaterAndCrushingTheCity] = true
	ActiveTacticsMap[BrokenBridgeByWater] = true
	ActiveTacticsMap[DecisionMakingThroughOperationsResearch] = true
	ActiveTacticsMap[Disperse] = true
	ActiveTacticsMap[EliminateItAndDrawFromIt] = true
	ActiveTacticsMap[EstablishingNameThroughGenerations] = true
	ActiveTacticsMap[ExcitingArmyAttackCamp] = true
	ActiveTacticsMap[FakeBooksAlternateWithEachOther] = true
	ActiveTacticsMap[FallingPhoenix] = true
	ActiveTacticsMap[FireJointVenture] = true
	ActiveTacticsMap[ForetellLikeProphet] = true
	ActiveTacticsMap[FrustrationAndAngerAttack] = true
	ActiveTacticsMap[TakingAdvantageOfTheSituationToGainPower] = true
	ActiveTacticsMap[AggregateStoneIntoGold] = true

	//突击
	AssaultTacticsMap[ChasingInjury] = true
	AssaultTacticsMap[CutDown] = true
	AssaultTacticsMap[WhenTheFrontIsDestroyed] = true
	AssaultTacticsMap[ToCureOnesSpeed] = true
	AssaultTacticsMap[TheBraveLeadTheWay] = true
	AssaultTacticsMap[RepelForeignAggression] = true
	AssaultTacticsMap[PeerlessOrMatchlessBraveryOrValour] = true
	AssaultTacticsMap[ViolentAndHeartless] = true
	AssaultTacticsMap[BendTheBowAndDrinkTheFeathers] = true
	AssaultTacticsMap[GhostGodThunderForce] = true
	AssaultTacticsMap[HiddenMystery] = true
	AssaultTacticsMap[HundredCavalryRobberyBattalions] = true
	AssaultTacticsMap[TakeAdvantageOfQuickly] = true
	AssaultTacticsMap[TakeCareOfYourselfFirst] = true
	AssaultTacticsMap[TigerAndLeopardCavalry] = true
}
