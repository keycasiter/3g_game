package _interface

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//战法通用
type TacticsCommon interface {
	// Init 初始化
	Init(tacticsParams model.TacticsParams)
	//战法ID
	Id() int64
	//战法名称
	Name() string
	//战法来源
	TacticsSource() consts.TacticsSource
	//战法类型
	TacticsType() consts.TacticsType
	//支持兵种
	SupportArmTypes() []consts.ArmType
	//发动概率
	TriggerRate() float64
	//累计普攻
	AccumulateGeneralAttack() int
	//累计战法
	AccumulateTacticsAttack() int
}
