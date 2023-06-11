package consts

// 环境
const (
	RUN_ENV_DEV  = "dev"
	RUN_ENV_TEST = "test"
	RUN_ENV_PROD = "prod"
)

// 最小值/最大值
const (
	INT_MIN = ^INT_MAX
	INT_MAX = int(^uint(0) >> 1)
)

// 兵力范围
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
	BattleAction_Unknow            BattleAction = iota //未知动作
	BattleAction_BeginAction                           //开始行动
	BattleAction_EndAction                             //结束行动
	BattleAction_Attack                                //普通攻击开始
	BattleAction_AttackEnd                             //普通攻击结束
	BattleAction_ActiveTactic                          //发动主动战法开始
	BattleAction_ActiveTacticEnd                       //发动主动战法结束
	BattleAction_AssaultTactic                         //发动突击战法开始
	BattleAction_AssaultTacticEnd                      //发动突击战法结束
	BattleAction_CommandTactic                         //指挥战法攻击开始
	BattleAction_CommandTacticEnd                      //指挥战法攻击结束
	BattleAction_PassiveTactic                         //被动战法攻击开始
	BattleAction_PassiveTacticEnd                      //被动战法攻击结束
	BattleAction_TroopsTactic                          //阵法战法攻击开始
	BattleAction_TroopsTacticEnd                       //阵法战法攻击结束
	BattleAction_ArmTactic                             //兵种战法攻击开始
	BattleAction_ArmTacticEnd                          //兵种战法攻击借宿
	BattleAction_WeaponDamage                          //发动兵刃伤害开始
	BattleAction_WeaponDamageEnd                       //发动兵刃伤害结束
	BattleAction_StrategyDamage                        //发动谋略伤害开始
	BattleAction_StrategyDamageEnd                     //发动谋略伤害结束
	BattleAction_DebuffEffect                          //施加负面效果开始
	BattleAction_DebuffEffectEnd                       //施加负面效果结束
	BattleAction_BuffEffect                            //施加正面效果开始
	BattleAction_BuffEffectEnd                         //施加正面效果结束

	//遭受伤害
	BattleAction_SufferDamage           //遭受伤害开始
	BattleAction_SufferDamageEnd        //遭受伤害结束
	BattleAction_SufferGeneralAttack    //被普通攻击开始
	BattleAction_SufferGeneralAttackEnd //被普通攻击结束
	BattleAction_SufferActiveTactic     //被主动战法攻击开始
	BattleAction_SufferActiveTacticEnd  //被主动战法攻击结束
	BattleAction_SufferAssaultTactic    //被突击战法攻击开始
	BattleAction_SufferAssaultTacticEnd //被突击战法攻击结束
	BattleAction_SufferCommandTactic    //被指挥战法攻击开始
	BattleAction_SufferCommandTacticEnd //被指挥战法攻击结束
	BattleAction_SufferArmTactic        //被兵种战法攻击开始
	BattleAction_SufferArmTacticEnd     //被兵种战法攻击结束
	BattleAction_SufferTroopsTactic     //被阵法战法攻击开始
	BattleAction_SufferTroopsTacticEnd  //被阵法战法攻击结束
	BattleAction_SufferPassiveTactic    //被被动战法攻击开始
	BattleAction_SufferPassiveTacticEnd //被被动战法攻击结束
	BattleAction_SufferDebuffEffect     //被施加负面效果开始
	BattleAction_SufferDebuffEffectEnd  //被施加负面效果结束
	BattleAction_SufferBuffEffect       //被施加正面效果开始
	BattleAction_SufferBuffEffectEnd    //被施加正面效果结束
)

func (action BattleAction) String() string {
	switch action {
	case BattleAction_BeginAction: //开始行动
		return "开始行动"
	case BattleAction_EndAction: //结束行动
		return "结束行动"
	case BattleAction_Attack: //普通攻击开始
		return "普通攻击开始"
	case BattleAction_AttackEnd: //普通攻击结束
		return "普通攻击结束"
	case BattleAction_SufferGeneralAttack: //被普通攻击开始
		return "被普通攻击开始"
	case BattleAction_SufferGeneralAttackEnd: //被普通攻击结束
		return "被普通攻击结束"
	case BattleAction_ActiveTactic: //发动主动战法开始
		return "发动主动战法开始"
	case BattleAction_ActiveTacticEnd: //发动主动战法结束
		return "发动主动战法结束"
	case BattleAction_SufferActiveTactic: //被主动战法攻击开始
		return "被主动战法攻击开始"
	case BattleAction_SufferActiveTacticEnd: //被主动战法攻击结束
		return "被主动战法攻击结束"
	case BattleAction_AssaultTactic: //发动突击战法攻击开始
		return "发动突击战法攻击开始"
	case BattleAction_AssaultTacticEnd: //发动突击战法攻击结束
		return "发动突击战法攻击结束"
	case BattleAction_SufferAssaultTactic: //被突击战法攻击开始
		return "被突击战法攻击开始"
	case BattleAction_SufferAssaultTacticEnd: //被突击战法攻击结束
		return "被突击战法攻击结束"
	case BattleAction_CommandTactic: //指挥战法攻击开始
		return "指挥战法攻击开始"
	case BattleAction_CommandTacticEnd: //指挥战法攻击结束
		return "指挥战法攻击结束"
	case BattleAction_SufferCommandTactic: //被指挥战法攻击开始
		return "被指挥战法攻击开始"
	case BattleAction_SufferCommandTacticEnd: //被指挥战法攻击结束
		return "被指挥战法攻击结束"
	case BattleAction_PassiveTactic: //被动战法攻击开始
		return "被动战法攻击开始"
	case BattleAction_PassiveTacticEnd: //被动战法攻击结束
		return "被动战法攻击结束"
	case BattleAction_SufferPassiveTactic: //被被动战法攻击开始
		return "被被动战法攻击开始"
	case BattleAction_SufferPassiveTacticEnd: //被被动战法攻击结束
		return "被被动战法攻击结束"
	case BattleAction_TroopsTactic: //阵法战法攻击开始
		return "阵法战法攻击开始"
	case BattleAction_TroopsTacticEnd: //阵法战法攻击结束
		return "阵法战法攻击结束"
	case BattleAction_SufferTroopsTactic: //被阵法战法攻击开始
		return "被阵法战法攻击开始"
	case BattleAction_SufferTroopsTacticEnd: //被阵法战法攻击结束
		return "被阵法战法攻击结束"
	case BattleAction_ArmTactic: //兵种战法攻击开始
		return "兵种战法攻击开始"
	case BattleAction_ArmTacticEnd: //兵种战法攻击结束
		return "兵种战法攻击结束"
	case BattleAction_SufferArmTactic: //被兵种战法攻击开始
		return "被兵种战法攻击开始"
	case BattleAction_SufferArmTacticEnd: //被兵种战法攻击结束
		return "被兵种战法攻击结束"
	case BattleAction_WeaponDamage: //发动兵刃伤害开始
		return "发动兵刃伤害开始"
	case BattleAction_WeaponDamageEnd: //发动兵刃伤害结束
		return "发动兵刃伤害结束"
	case BattleAction_StrategyDamage: //发动谋略伤害开始
		return "发动谋略伤害开始"
	case BattleAction_StrategyDamageEnd: //发动谋略伤害结束
		return "发动谋略伤害结束"
	}
	return ""
}
