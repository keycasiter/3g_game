package consts

type TacticId int64

const (
	//鹰视狼顾
	ClearEyedAndMalicious TacticId = iota + 1
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
	BeAdeptWithBothThePenAndTheSword
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
	BraveAndBattlewise
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
	//落雷
	Thunderbolt
	//声东击西
	MakeFeintToTheEastButAttackInTheWest
	//风声鹤唳
	TheSoundOfTheWindAndTheCryOfTheStork
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
)

func (b TacticId) String() string {
	switch b {
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
	case BeAdeptWithBothThePenAndTheSword:
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
	case BraveAndBattlewise:
		return "骁勇善战"
		//御敌屏障
	case DefensiveBarrier:
		return "御敌屏障"
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
	//落雷
	case Thunderbolt:
		return "落雷"
	//声东击西
	case MakeFeintToTheEastButAttackInTheWest:
		return "声东击西"
		//风声鹤唳
	case TheSoundOfTheWindAndTheCryOfTheStork:
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
	}
	return "未知战法"
}
