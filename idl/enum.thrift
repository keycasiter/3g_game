namespace go enum

//请求返回码
enum ResponseCode {
   Success = 0
   ParamInvalid = 1

   UnknownError = 999
}

// 判定
enum Enable {
    UNKNOW = 0
	YES = 1
	NO = 2
}

// 兵种适性
enum ArmsAbility {
    Unknow = 0 //未知
	S = 1 // S级
	A = 2 // A级
	B = 3 // B级
	C = 4 // C级
}

// 兵书类型
enum WarbookType {
    Unknow = 0 //未知
	Battle = 1 // 作战（红色兵书）
	MilitaryForm = 2 // 虚实（紫色兵书）
	FalsehoodVersusReality = 3 // 军形（蓝色兵书）
	NineVariations = 4 // 九变（绿色兵书）
	BeginPlan = 5 // 始计（黄色兵书）
	UseInterval = 6 // 用间（灰色兵书）
}


// 性别
enum Gender {
    Unknow = 0 //未知
	Male   = 1 //男
	Female = 2 //女
}

// 武将阵营
enum Group {
	Unknow   = 0 //未知
	WeiGuo   = 1 //魏国
	ShuGuo   = 2 //蜀国
	WuGuo    = 3 //吴国
	QunXiong = 4 //群雄
}

// 统御 （2 - 7）
enum ControlLevel {
	Level_2 = 2
	Level_3 = 3
	Level_4 = 4
	Level_5 = 5
	Level_6 = 6
	Level_7 = 7
}

// 武将标签
enum GeneralTag {
	Shield        = 1  // 盾
	Assist        = 2  // 辅
	YellowTurbans = 3  // 黄
	Both          = 4  // 兼
	Control       = 5  // 控
	Barbarian     = 6  // 蛮
	Charm         = 7  // 魅
	Strategy      = 8  // 谋
	Martial       = 9  // 武
	Medical       = 10 // 医
	Fight         = 11 // 战
	Politics      = 12 // 政
	Pawn          = 13 // 卒
	Celestials    = 14 // 仙
}

// 战法品质
enum TacticQuality {
	S = 1 // S级
	A = 2 // A级
	B = 3 // B级
	C = 4 // C级
}

// 战法来源
enum TacticsSource {
	SelfContained = 1 //自带战法
	Inherit       = 2 //传承战法
	Event         = 3 //事件战法
}

// 武将品质
enum GeneralQuality {
	S = 1 // S级
	A = 2 // A级
	B = 3 // B级
	C = 4 // C级
}

// 对战参战类型
enum GeneralBattleType {
	Fighting = 1
	Enemy    = 2
}

// 战法类型
enum TacticsType {
	Active        = 1 // 主动
	Passive       = 2 // 被动
	Command       = 3 // 指挥
	Assault       = 4 // 突击
	TroopsTactics = 5 // 阵法
	Arm           = 6 // 兵种
}

// 兵种类型
enum ArmType {
	Unknow    = 0 //未知
	Cavalry   = 1 //骑兵
	Mauler    = 2 //盾兵
	Archers   = 3 //弓兵
	Spearman  = 4 //枪兵
	Apparatus = 5 //器械
}

// 队伍类型
enum TeamType {
	Fighting = 1
	Enemy    = 2
}

// 武将星级
enum GeneralStarLevel {
	One = 1
	Two = 2
	Three = 3
	Four = 4
	Five = 5
}

// 装备类型
enum EquipType {
	Weapon = 1
	Armor = 2
	Horse = 3
	Treasure = 4
}

// 装备等级
enum EquipLevel {
	S = 1
	A = 2
	B = 3
	C = 4
}

enum BattleResult{
	Win              = 1 //胜利
    Lose             = 2 //失败
    Draw             = 3 //平局
    Advantage_Draw   = 4 //优势平局
    Inferiority_Draw = 5 //劣势平局
}

//对战阶段
enum BattlePhase{
    Unknow   = 0 //未知
    Prepare  = 1 // 准备阶段
    Fighting = 2 // 对阵阶段
}

//对战回合
enum BattleRound {
    Unknow  = 0
    First   = 1
    Second  = 2
    Third   = 3
    Fourth  = 4
    Fifth   = 5
    Sixth   = 6
    Seventh = 7
    Eighth  = 8
    Prepare = -1
}