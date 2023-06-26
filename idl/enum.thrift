namespace go enum

//请求返回码
enum ResponseCode {
   Success = 0
   ParamInvalid = 1

   UnknownError = 999
}

// 兵种适性
enum ArmsAbility {
	S = 1 // S级
	A = 2 // A级
	B = 3 // B级
	C = 4 // C级
}

// 性别
enum Gender {
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

// 战法来源
enum TacticsSource {
	SelfContained = 1 //自带战法
	Inherit       = 2 //传承战法
	Event         = 3 //事件战法
}