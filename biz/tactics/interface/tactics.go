package _interface

type Tactics interface {
	//通用
	TacticsCommon
	//增益
	TacticsBuff
	//减益
	TacticsDebuff
	//伤害
	TacticsDamage
	//恢复
	TacticsResume
	//武将锁定
	TacticsLocking
	//战法回合
	TacticsRound
	//附加
	TacticsAttach
}
