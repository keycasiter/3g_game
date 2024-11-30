package vo

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/po"
)

// 对战队伍信息
type BattleTeam struct {
	/** 队伍基础信息 **/
	//队伍类型
	TeamType consts.TeamType
	//队伍兵种
	ArmType consts.ArmType
	//队伍武将信息
	BattleGenerals []*BattleGeneral
	//兵战科技-属性加成
	BuildingTechAttrAddition BuildingTechAttrAddition
	//协力科技-阵营加成
	BuildingTechGroupAddition BuildingTechGroupAddition
}

type EffectHolderParams struct {
	//触发率
	TriggerRate float64
	//影响数值
	EffectValue int64
	//影响百分比
	EffectRate float64
	//影响回合
	EffectRound consts.BattleRound
	//影响次数
	EffectTimes int64
	//最大影响次数
	MaxEffectTimes int64
	//来源于哪个战法
	FromTactic consts.TacticId
	//来源于哪个兵书
	FromWarbook consts.WarBookDetailType
	//是否可以刷新
	IsSupportRefresh bool
	//强制普通攻击目标/被谁嘲讽
	TauntByTarget *BattleGeneral
	//锁定攻击目标
	LockingTarget *BattleGeneral
	//被谁分担伤害
	ShareResponsibilityForByGeneral *BattleGeneral
	//效果施加者
	ProduceGeneral *BattleGeneral
	//是否不可驱散
	IsAvoidDispel bool
}

// 对战武将信息
type BattleGeneral struct {
	//基础信息
	BaseInfo *po.MetadataGeneral
	//佩戴战法
	EquipTactics []*po.Tactics
	//兵书
	WarBooks []*po.Warbook
	//武将对战加成
	Addition *BattleGeneralAddition
	//是否主将
	IsMaster bool
	//原始兵力
	InitSoldierNum int64
	//携带兵力
	SoldierNum int64
	//已损失兵力
	LossSoldierNum int64
	//累计伤害
	AccumulateTotalDamageNum int64
	//累计普通攻击伤害
	AccumulateAttackDamageNum int64
	//累计治疗
	AccumulateTotalResumeNum int64
	//战法伤害统计
	TacticAccumulateDamageMap map[consts.TacticId]int64
	//战法治疗统计
	TacticAccumulateResumeMap map[consts.TacticId]int64
	//战法发动统计
	TacticAccumulateTriggerMap map[consts.TacticId]int64

	//兵书伤害统计
	WarbookAccumulateDamageMap map[consts.WarBookDetailType]int64
	//兵书治疗统计
	WarbookAccumulateResumeMap map[consts.WarBookDetailType]int64
	//兵书发动统计
	WarbookAccumulateTriggerMap map[consts.WarBookDetailType]int64

	//被谁援护
	HelpByGeneral *BattleGeneral `json:"-"`
	//被谁嘲讽
	TauntByGeneral *BattleGeneral `json:"-"`
	//被谁分担伤害
	ShareResponsibilityForByGeneral *BattleGeneral `json:"-"`

	//增益效果变量 map<效果,容器属性>
	BuffEffectHolderMap map[consts.BuffEffectType][]*EffectHolderParams `json:"-"`
	//减益效果变量 map<效果,容器属性>
	DeBuffEffectHolderMap map[consts.DebuffEffectType][]*EffectHolderParams `json:"-"`

	//增益效果次数
	BuffEffectCountMap map[consts.BuffEffectType]int64 `json:"-"`
	//减益效果次数
	DeBuffEffectCountMap map[consts.DebuffEffectType]int64 `json:"-"`

	//战法冷却容器
	TacticFrozenMap map[consts.TacticId]bool `json:"-"`

	//普通攻击次数
	ExecuteGeneralAttackNum int64
	//被普通攻击次数
	SufferExecuteGeneralAttackNum int64
	//兵刃攻击次数
	ExecuteWeaponAttackNum int64
	//谋略攻击次数
	ExecuteStrategyAttackNum int64
	//被兵刃攻击次数
	SufferExecuteWeaponAttackNum int64
	//被谋略攻击次数
	SufferExecuteStrategyAttackNum int64

	//回合剩余兵力
	RoundRemainSoliderNum map[consts.BattlePhase]map[consts.BattleRound]int64

	//*****战法触发器都是按条件（非回合）会触发的******
	//战法触发器 map<触发动作,func(触发函数参数)>
	TacticsTriggerMap map[consts.BattleAction][]func(params *TacticsTriggerParams) *TacticsTriggerResult `json:"-"`
}

// 战法触发参数
type TacticsTriggerParams struct {
	//当前回合
	CurrentRound consts.BattleRound
	//当前执行武将
	CurrentGeneral *BattleGeneral
	//当前发起攻击的武将
	AttackGeneral *BattleGeneral
	//当前被攻击的武将
	SufferAttackGeneral *BattleGeneral
	//当前发起恢复的武将
	ResumeGeneral *BattleGeneral
	//当前被恢复的武将
	SufferResumeGeneral *BattleGeneral
	//当前造成伤害
	CurrentDamage int64
	//当前恢复量
	CurrentResume int64
	//当前执行战法
	CurrentTactic interface{}
	//施加的负面效果
	DebuffEffect consts.DebuffEffectType
	//施加负面效果的战法
	DebuffEffectOfTactic consts.TacticId
	//当前被施加负面战法的武将
	SufferDebuffEffectGeneral *BattleGeneral
	//施加的正面效果
	BuffEffect consts.BuffEffectType
	//施加正面效果的战法
	BuffEffectOfTactic consts.TacticId
	//当前被施加正面战法的武将
	SufferBuffEffectGeneral *BattleGeneral
	//影响回合
	EffectRound consts.BattleRound

	//施加参数
	EffectHolderParams *EffectHolderParams
	//伤害类型
	DamageType consts.DamageType
}

// 战法触发结果
type TacticsTriggerResult struct {
	//是否打断后续战法
	IsTerminate bool
}

// 武将对战加成
type BattleGeneralAddition struct {
	//1. 武将加成
	//1.a. 加点加成
	AbilityAttr po.AbilityAttr
	//1.b. 等级加成
	GeneralLevel consts.GeneralLevel
	//1.c. 红度加成
	GeneralStarLevel consts.GeneralStarLevel
	//1.d. 缘分加成
	Predestination consts.Predestination
	//2. 装备加成
	//TODO
	//3. 特技加成
	//TODO
}

// 建筑科技属性加成
type BuildingTechAttrAddition struct {
	ForceAddition        float64
	IntelligenceAddition float64
	CommandAddition      float64
	SpeedAddition        float64
}

// 建筑科技阵营加成
type BuildingTechGroupAddition struct {
	GroupWeiGuoRate   float64
	GroupShuGuoRate   float64
	GroupWuGuoRate    float64
	GroupQunXiongRate float64
}
