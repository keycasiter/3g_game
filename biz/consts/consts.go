package consts

// 环境
const (
	RUN_ENV_DEV  = "dev"
	RUN_ENV_TEST = "test"
	RUN_ENV_PROD = "prod"
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
	Battle_Round_Unknow  BattleRound = 0 //未知回合
	Battle_Round_First   BattleRound = 1 //第一回合
	Battle_Round_Second  BattleRound = 2 //第二回合
	Battle_Round_Third   BattleRound = 3 //第三回合
	Battle_Round_Fourth  BattleRound = 4 //第四回合
	Battle_Round_Fifth   BattleRound = 5 //第五回合
	Battle_Round_Sixth   BattleRound = 6 //第六回合
	Battle_Round_Seventh BattleRound = 7 //第七回合
	Battle_Round_Eighth  BattleRound = 8 //第八回合
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
	DamageType_Weapon   DamageType = 1 //兵刃伤害
	DamageType_Strategy DamageType = 2 //谋略伤害
)

// 负面效果
type DebuffEffectType int

const (
	DebuffEffectType_Methysis   DebuffEffectType = 1 //中毒
	DebuffEffectType_Firing     DebuffEffectType = 2 //灼烧
	DebuffEffectType_Defect     DebuffEffectType = 3 //叛逃（受武力或智力最高一项影响，无视防御）
	DebuffEffectType_Sandstorm  DebuffEffectType = 4 //沙暴
	DebuffEffectType_Chaos      DebuffEffectType = 5 //混乱（攻击和战法无差别选择目标）
	DebuffEffectType_NoScheme   DebuffEffectType = 6 //计穷（无法发动主动战法）
	DebuffEffectType_PoorHealth DebuffEffectType = 7 //虚弱（无法造成伤害）
)

// 增益效果
type BuffEffectType int

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
	GeneralNum_One   GeneralNum = 1
	GeneralNum_Two   GeneralNum = 2
	GeneralNum_Three GeneralNum = 3
)
