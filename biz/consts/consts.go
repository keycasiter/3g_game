package consts

//环境
const (
	RUN_ENV_DEV  = "dev"
	RUN_ENV_TEST = "util"
	RUN_ENV_PROD = "prod"
)

// 对战阶段
type BattlePhase int

// 对战回合
type BattleRound int

// 武将等级
type GeneralLevel int

// 武将星级
type GeneralStarLevel int

// 武将缘分
type Predestination int

// 兵种能力
type ArmsAbility int

// 兵种类型
type ArmType int

// 阵营
type Group int

/** 对战相关 **/
const (
	// 对战阶段
	Battle_Phase_Prepare  = 1 //准备阶段
	Battle_Phase_Fighting = 2 //对阵阶段

	//对战回合
	Battle_Round_First   = 1 //第一回合
	Battle_Round_Second  = 2 //第二回合
	Battle_Round_Third   = 3 //第三回合
	Battle_Round_Fourth  = 4 //第四回合
	Battle_Round_Fifth   = 5 //第五回合
	Battle_Round_Sixth   = 6 //第六回合
	Battle_Round_Seventh = 7 //第七回合
	Battle_Round_Eighth  = 8 //第八回合
)

/** 武将相关 **/
const (
	// 性别
	Gender_Male   = 1 //男
	Gender_Female = 2 //女

	// 武将阵营
	Group_Unknow   = 0 //未知
	Group_WeiGuo   = 1 //魏国
	Group_ShuGuo   = 2 //蜀国
	Group_WuGuo    = 3 //吴国
	Group_QunXiong = 4 //群雄

	// 武将星级（0 - 5）
	StarLevel_0 = 0
	StarLevel_1 = 1
	StarLevel_2 = 2
	StarLevel_3 = 3
	StarLevel_4 = 4
	StarLevel_5 = 5

	// 统御 （2 - 7）
	ControlLevel_2 = 2
	ControlLevel_3 = 3
	ControlLevel_4 = 4
	ControlLevel_5 = 5
	ControlLevel_6 = 6
	ControlLevel_7 = 7

	// 武将标签
	GeneralTag_Shield        = 1  // 盾
	GeneralTag_Assist        = 2  // 辅
	GeneralTag_YellowTurbans = 3  // 黄
	GeneralTag_Both          = 4  // 兼
	GeneralTag_Control       = 5  // 控
	GeneralTag_Barbarian     = 6  // 蛮
	GeneralTag_Charm         = 7  // 魅
	GeneralTag_Strategy      = 8  // 谋
	GeneralTag_Martial       = 9  // 武
	GeneralTag_Medical       = 10 // 医
	GeneralTag_Fight         = 11 // 战
	GeneralTag_Politics      = 12 // 政
	GeneralTag_Pawn          = 13 // 卒
	GeneralTag_Celestials    = 14 // 仙

	// 兵种适性
	ArmsAbility_S = 1 // S级
	ArmsAbility_A = 2 // A级
	ArmsAbility_B = 3 // B级
	ArmsAbility_C = 4 // C级
)

/** 战法相关 **/
const (
	// 战法类型
	TacticsType_Active        = 1 // 主动
	TacticsType_Passive       = 2 // 被动
	TacticsType_Command       = 3 // 指挥
	TacticsType_Assault       = 4 // 突击
	TacticsType_TroopsTactics = 5 // 阵法
	TacticsType_Arm           = 6 // 兵种

	// 战法品质
	TacticsLevel_S = 1 // S级
	TacticsLevel_A = 2 // A级
	TacticsLevel_B = 3 // B级
	TacticsLevel_C = 4 // C级

	// 战法来源
	TacticsSource_SelfContained = 1 //自带战法
	TacticsSource_Inherit       = 2 //传承战法
	TacticsSource_Event         = 3 //事件战法

	// 战法目标
	TacticsTarget_Enemy_Single = 1 //敌军单体
	TacticsTarget_Enemy_Group  = 2 //敌军群体
	TacticsTarget_Team_Single  = 3 //友军单体
	TacticsTarget_Team_Group   = 4 //友军单体

	// 兵种
	ArmType_Unknow    = 0 //未知
	ArmType_Cavalry   = 1 //骑兵
	ArmType_Mauler    = 2 //盾兵
	ArmType_Archers   = 3 //弓兵
	ArmType_Spearman  = 4 //枪兵
	ArmType_Apparatus = 5 //器械
)

/** 装备相关 **/
const (
	// 装备品质
	EquipLevel_S = 1 //珍品
	EquipLevel_A = 1 //上品
	EquipLevel_B = 1 //精良
	EquipLevel_C = 1 //凡品

	// 装备类型
	EquipType_Weapon   = 1 //武器
	EquipType_Armor    = 2 //防具
	EquipType_Horse    = 3 //坐骑
	EquipType_Treasure = 4 //宝物
)
