package consts

// 环境
const (
	RUN_ENV_DEV  = "dev"
	RUN_ENV_TEST = "test"
	RUN_ENV_PROD = "prod"
)

//兵力范围
const (
	//每个武将最低带兵数量
	Min_Soldiers_Num_Per_General = 0
	//每个武将最大带兵数量
	Max_Soldiers_Num_Per_General = 10000
)

/** 武将属性 **/
type AbilityAttr int

const (
	AbilityAttr_Force        AbilityAttr = 1 //武力
	AbilityAttr_Intelligence AbilityAttr = 2 //智力
	AbilityAttr_Command      AbilityAttr = 3 //统率
	AbilityAttr_Speed        AbilityAttr = 4 //速度
)

/** 对战相关 **/

type BattlePhase int

const (
	// 对战阶段
	Battle_Phase_Prepare  BattlePhase = 1 //准备阶段
	Battle_Phase_Fighting BattlePhase = 2 //对阵阶段
)

func (b BattlePhase) String() string {
	switch b {
	case Battle_Phase_Prepare:
		return "准备阶段"
	case Battle_Phase_Fighting:
		return "对阵阶段"
	}
	return ""
}

type BattleRound int

const (
	//对战回合
	Battle_Round_Unknow  BattleRound = 0  //未知回合
	Battle_Round_First   BattleRound = 1  //第一回合
	Battle_Round_Second  BattleRound = 2  //第二回合
	Battle_Round_Third   BattleRound = 3  //第三回合
	Battle_Round_Fourth  BattleRound = 4  //第四回合
	Battle_Round_Fifth   BattleRound = 5  //第五回合
	Battle_Round_Sixth   BattleRound = 6  //第六回合
	Battle_Round_Seventh BattleRound = 7  //第七回合
	Battle_Round_Eighth  BattleRound = 8  //第八回合
	Battle_Round_Prepare BattleRound = -1 //准备回合
)

func (b BattleRound) String() string {
	switch b {
	case Battle_Round_Unknow:
		return "未知回合"
	case Battle_Round_First:
		return "第1回合"
	case Battle_Round_Second:
		return "第2回合"
	case Battle_Round_Third:
		return "第3回合"
	case Battle_Round_Fourth:
		return "第4回合"
	case Battle_Round_Fifth:
		return "第5回合"
	case Battle_Round_Sixth:
		return "第6回合"
	case Battle_Round_Seventh:
		return "第7回合"
	case Battle_Round_Eighth:
		return "第8回合"
	}
	return ""
}

/** 武将相关 **/

type Gender int

const (
	// 性别
	Gender_Male   Gender = 1 //男
	Gender_Female Gender = 2 //女
)

type Group int

const (
	// 武将阵营
	Group_Unknow   Group = 0 //未知
	Group_WeiGuo   Group = 1 //魏国
	Group_ShuGuo   Group = 2 //蜀国
	Group_WuGuo    Group = 3 //吴国
	Group_QunXiong Group = 4 //群雄
)

type StarLevel int

const (
	// 武将星级（0 - 5）
	StarLevel_0 StarLevel = 0
	StarLevel_1 StarLevel = 1
	StarLevel_2 StarLevel = 2
	StarLevel_3 StarLevel = 3
	StarLevel_4 StarLevel = 4
	StarLevel_5 StarLevel = 5
)

type ControlLevel int

const (
	// 统御 （2 - 7）
	ControlLevel_2 ControlLevel = 2
	ControlLevel_3 ControlLevel = 3
	ControlLevel_4 ControlLevel = 4
	ControlLevel_5 ControlLevel = 5
	ControlLevel_6 ControlLevel = 6
	ControlLevel_7 ControlLevel = 7
)

type GeneralTag int

const (
	// 武将标签
	GeneralTag_Shield        GeneralTag = 1  // 盾
	GeneralTag_Assist        GeneralTag = 2  // 辅
	GeneralTag_YellowTurbans GeneralTag = 3  // 黄
	GeneralTag_Both          GeneralTag = 4  // 兼
	GeneralTag_Control       GeneralTag = 5  // 控
	GeneralTag_Barbarian     GeneralTag = 6  // 蛮
	GeneralTag_Charm         GeneralTag = 7  // 魅
	GeneralTag_Strategy      GeneralTag = 8  // 谋
	GeneralTag_Martial       GeneralTag = 9  // 武
	GeneralTag_Medical       GeneralTag = 10 // 医
	GeneralTag_Fight         GeneralTag = 11 // 战
	GeneralTag_Politics      GeneralTag = 12 // 政
	GeneralTag_Pawn          GeneralTag = 13 // 卒
	GeneralTag_Celestials    GeneralTag = 14 // 仙
)

type ArmsAbility int

const (
	// 兵种适性
	ArmsAbility_S ArmsAbility = 1 // S级
	ArmsAbility_A ArmsAbility = 2 // A级
	ArmsAbility_B ArmsAbility = 3 // B级
	ArmsAbility_C ArmsAbility = 4 // C级
)

type TacticsType int

/** 战法相关 **/
const (
	// 战法类型
	TacticsType_Active        TacticsType = 1 // 主动
	TacticsType_Passive       TacticsType = 2 // 被动
	TacticsType_Command       TacticsType = 3 // 指挥
	TacticsType_Assault       TacticsType = 4 // 突击
	TacticsType_TroopsTactics TacticsType = 5 // 阵法
	TacticsType_Arm           TacticsType = 6 // 兵种
)

type TacticsLevel int

const (
	// 战法品质
	TacticsLevel_S TacticsLevel = 1 // S级
	TacticsLevel_A TacticsLevel = 2 // A级
	TacticsLevel_B TacticsLevel = 3 // B级
	TacticsLevel_C TacticsLevel = 4 // C级
)

type TacticsSource int

const (
	// 战法来源
	TacticsSource_SelfContained TacticsSource = 1 //自带战法
	TacticsSource_Inherit       TacticsSource = 2 //传承战法
	TacticsSource_Event         TacticsSource = 3 //事件战法
)

type TacticsTarget int

const (
	// 战法目标
	TacticsTarget_Enemy_Single TacticsTarget = 1 //敌军单体
	TacticsTarget_Enemy_Group  TacticsTarget = 2 //敌军群体
	TacticsTarget_Team_Single  TacticsTarget = 3 //友军单体
	TacticsTarget_Team_Group   TacticsTarget = 4 //友军单体
)

// 对战参战类型
type GeneralBattleType int

const (
	GeneralBattleType_Fighting GeneralBattleType = 1
	GeneralBattleType_Enemy    GeneralBattleType = 2
)

func (g GeneralBattleType) String() string {
	switch g {
	case GeneralBattleType_Fighting:
		return "我方"
	case GeneralBattleType_Enemy:
		return "敌方"
	}
	return ""
}

// 队伍类型
type TeamType int

const (
	TeamType_Fighting TeamType = 1
	TeamType_Enemy    TeamType = 2
)

// 兵种类型
type ArmType int

const (
	// 兵种
	ArmType_Unknow    ArmType = 0 //未知
	ArmType_Cavalry   ArmType = 1 //骑兵
	ArmType_Mauler    ArmType = 2 //盾兵
	ArmType_Archers   ArmType = 3 //弓兵
	ArmType_Spearman  ArmType = 4 //枪兵
	ArmType_Apparatus ArmType = 5 //器械
)

/** 装备相关 **/
type EquipLevel int

const (
	// 装备品质
	EquipLevel_S EquipLevel = 1 //珍品
	EquipLevel_A EquipLevel = 2 //上品
	EquipLevel_B EquipLevel = 3 //精良
	EquipLevel_C EquipLevel = 4 //凡品
)

type EquipType int

const (
	// 装备类型
	EquipType_Weapon   EquipType = 1 //武器
	EquipType_Armor    EquipType = 2 //防具
	EquipType_Horse    EquipType = 3 //坐骑
	EquipType_Treasure EquipType = 4 //宝物
)

/** 伤害类型 **/
type DamageType int

const (
	DamageType_None     DamageType = 0 //无伤害
	DamageType_Weapon   DamageType = 1 //兵刃伤害
	DamageType_Strategy DamageType = 2 //谋略伤害
)

// 负面效果
type DebuffEffectType int

const (
	DebuffEffectType_Unknow                      DebuffEffectType = 0  //未知
	DebuffEffectType_Methysis                    DebuffEffectType = 1  //中毒
	DebuffEffectType_Firing                      DebuffEffectType = 2  //灼烧
	DebuffEffectType_Defect                      DebuffEffectType = 3  //叛逃（受武力或智力最高一项影响，无视防御）
	DebuffEffectType_Sandstorm                   DebuffEffectType = 4  //沙暴（每回合持续造成伤害）
	DebuffEffectType_Chaos                       DebuffEffectType = 5  //混乱（攻击和战法无差别选择目标）
	DebuffEffectType_NoStrategy                  DebuffEffectType = 6  //计穷（无法发动主动战法）
	DebuffEffectType_PoorHealth                  DebuffEffectType = 7  //虚弱（无法造成伤害）
	DebuffEffectType_WaterAttack                 DebuffEffectType = 8  //水攻（每回合持续造成伤害）
	DebuffEffectType_SufferWeaponDamageImprove   DebuffEffectType = 9  //受到兵刃伤害增加
	DebuffEffectType_SufferStrategyDamageImprove DebuffEffectType = 10 //受到谋略伤害增加
	DebuffEffectType_LaunchWeaponDamageDeduce    DebuffEffectType = 11 //造成兵刃伤害减少
	DebuffEffectType_LaunchStrategyDamageDeduce  DebuffEffectType = 12 //造成谋略伤害减少
	DebuffEffectType_CanNotGeneralAttack         DebuffEffectType = 13 //无法普通攻击
	DebuffEffectType_CancelWeapon                DebuffEffectType = 14 //缴械（无法普通攻击）
	DebuffEffectType_DecrForce                   DebuffEffectType = 15 //降低武力
	DebuffEffectType_DecrIntelligence            DebuffEffectType = 16 //降低智力
	DebuffEffectType_DecrCommand                 DebuffEffectType = 17 //降低统率
	DebuffEffectType_DecrSpeed                   DebuffEffectType = 18 //降低速度
)

// 增益效果
type BuffEffectType int

const (
	BuffEffectType_Unknow                      BuffEffectType = 0
	BuffEffectType_Evade                       BuffEffectType = 1  //规避
	BuffEffectType_EnhanceWeapon               BuffEffectType = 2  //会心
	BuffEffectType_EnhanceStrategy             BuffEffectType = 3  //奇谋
	BuffEffectType_GroupAttack                 BuffEffectType = 4  //群攻
	BuffEffectType_FirstAttack                 BuffEffectType = 5  //先攻
	BuffEffectType_Rest                        BuffEffectType = 6  //休整
	BuffEffectType_Defend                      BuffEffectType = 7  //抵御
	BuffEffectType_ContinuousAttack            BuffEffectType = 8  //连击
	BuffEffectType_StrikeBack                  BuffEffectType = 9  //反击
	BuffEffectType_TacticsTriggerImprove       BuffEffectType = 10 //战法发动率提升
	BuffEffectType_LaunchWeaponDamageImprove   BuffEffectType = 11 //造成兵刃伤害增加
	BuffEffectType_LaunchStrategyDamageImprove BuffEffectType = 12 //造成谋略伤害增加
	BuffEffectType_SufferWeaponDamageDeduce    BuffEffectType = 13 //受到兵刃伤害减少
	BuffEffectType_SufferStrategyDamageDeduce  BuffEffectType = 14 //受到谋略伤害减少
	BuffEffectType_IncrForce                   BuffEffectType = 15 //增加武力
	BuffEffectType_IncrIntelligence            BuffEffectType = 16 //增加智力
	BuffEffectType_IncrCommand                 BuffEffectType = 17 //增加统率
	BuffEffectType_IncrSpeed                   BuffEffectType = 18 //增加速度
)

func (b BuffEffectType) String() string {
	switch b {
	case BuffEffectType_Evade:
		return "规避"
	case BuffEffectType_EnhanceWeapon:
		return "会心"
	}
	return ""
}

// 武将等级
type GeneralLevel int

// 武将星级
type GeneralStarLevel int

const (
	GeneralStarLevel_1 GeneralStarLevel = 1
	GeneralStarLevel_2 GeneralStarLevel = 2
	GeneralStarLevel_3 GeneralStarLevel = 3
	GeneralStarLevel_4 GeneralStarLevel = 4
	GeneralStarLevel_5 GeneralStarLevel = 5
)

// 武将缘分
type Predestination int

// 武将数量
type GeneralNum int

const (
	GeneralNum_Unknow GeneralNum = 0
	GeneralNum_One    GeneralNum = 1
	GeneralNum_Two    GeneralNum = 2
	GeneralNum_Three  GeneralNum = 3
)

// 兵书类型
type WarBookType int

const (
	WarBookType_Fighting          WarBookType = 1 //作战
	WarBookType_TruthAndFalsehood WarBookType = 2 //虚实
	WarBookType_MilitaryForm      WarBookType = 3 //军形
	WarBookType_NineChanges       WarBookType = 4 //九变
)

// 兵书枚举
type WarBookDetailType int

const (
	//作战
	WarBookDetailType_TheOddAndTheRightCoexist WarBookDetailType = 1 //奇正相生
	WarBookDetailType_BraveButNotBrave         WarBookDetailType = 2 //蛮勇非勇
	WarBookDetailType_NotBraveWillDie          WarBookDetailType = 3 //不勇则死
	WarBookDetailType_MilitaryAbility          WarBookDetailType = 4 //武略
	WarBookDetailType_VictoriousBattle         WarBookDetailType = 5 //胜战
	WarBookDetailType_PersistentSpirit         WarBookDetailType = 6 //执锐
	WarBookDetailType_MilitaryStrategy         WarBookDetailType = 7 //文韬
	WarBookDetailType_HideKnife                WarBookDetailType = 8 //藏刀
	//虚实
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 9 //大谋不谋
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 10 //以治击乱
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 11 //攻其不备
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //鬼谋
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //妙算
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //将威
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //神机
	//WarBookDetailType_BigPlanWithoutPlan WarBookDetailType = 12 //占卜
	//军形
	//九变
)

// 武将索引
type GeneralIndex int

const (
	GeneralIndex_1 GeneralIndex = 1 //我方武将-主将位
	GeneralIndex_2 GeneralIndex = 2 //我方武将-副将位1
	GeneralIndex_3 GeneralIndex = 3 //我方武将-副将位2
	GeneralIndex_4 GeneralIndex = 4 //敌方武将-主将位
	GeneralIndex_5 GeneralIndex = 5 //敌方武将-副将位1
	GeneralIndex_6 GeneralIndex = 6 //敌方武将-副将位2
)

// 对战行为
type BattleAction int

const (
	BattleAction_Attack       BattleAction = 1 //普通攻击
	BattleAction_SufferAttack BattleAction = 2 //被普通攻击
)
