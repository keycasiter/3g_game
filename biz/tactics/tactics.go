package tactics

import "github.com/keycasiter/3g_game/biz/consts"

type Tactics interface {
	/** 基础属性 **/
	//战法ID
	Id() int64
	//战法来源
	TacticsSource() consts.TacticsSource
	//战法类型
	TacticsType() consts.TacticsType
	//战法品质
	TacticsLevel() consts.TacticsLevel
	//战法目标
	TacticsTarget() consts.TacticsTarget
	//支持兵种
	SupportArmTypes() []consts.ArmType

	/** 全局属性 **/
	//发动概率
	TriggerRate() float64

	/** 回合属性 **/
	//当前回合
	CurrentRound(round consts.BattleRound)
	//伤害类型
	DamageType() consts.DamageType
	//增伤数值
	IncrDamageNum() int64
	//增伤率
	IncrDamageRate() float64
	//减伤数值
	DecrDamageNum() int64
	//减伤率
	DecrDamageRate() float64
	//兵力恢复率
	ResumeMilitaryStrengthRate() float64
	//奇谋几率
	EnhancedStrategyDamageRate() float64
	//会心几率
	EnhancedWeaponDamageRate() float64
	//可叠加次数
	SuperposeNum() int64
	//规避率
	EvadeRate() float64
	//提高武力
	IncrForceNum() float64
	//提高智力
	IncrIntelligenceNum() float64
	//提高统率
	IncrCommandNum() float64
	//提高速度
	IncrSpeedNum() float64
	//可持续回合
	EffectNextRounds() int64
	//需冷却回合
	FrozenNextRounds() int64
	//负面效果
	DebuffEffect() consts.DebuffEffectType
	//增益效果
	BuffEffect() consts.BuffEffectType
	//是否可以普通攻击
	IsGeneralAttack() bool
	//是否已经触发
	IsHasTrigger() bool

	/** 武将属性 **/
	//是否为主将
	IsMasterGeneral(isMaster bool)
}