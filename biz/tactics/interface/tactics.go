package _interface

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法通用
type Tactics interface {
	// 初始化
	Init(tacticsParams *model.TacticsParams) Tactics
	//战法准备
	Prepare()
	//战法ID
	Id() consts.TacticId
	//战法名称
	Name() string
	//战法来源
	TacticsSource() consts.TacticsSource
	//获取发动概率
	GetTriggerRate() float64
	//设置发动概率
	SetTriggerRate(rate float64)
	//战法类型
	TacticsType() consts.TacticsType
	//支持兵种
	SupportArmTypes() []consts.ArmType
	//战法执行
	Execute()
	//是否需要发动准备
	IsTriggerPrepare() bool
}
