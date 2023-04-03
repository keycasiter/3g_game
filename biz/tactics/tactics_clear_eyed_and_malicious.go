package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：鹰视狼顾
// 战法描述：战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)；
// 第5回合起，每回合有80%概率对1-2个敌军单体造成谋略伤害(伤害154%，受智力影响)；
// 自身为主将时，获得16%奇谋几率
type ClearEyedAndMaliciousTactic struct {
	tacticsParams model.TacticsParams
}

func (c ClearEyedAndMaliciousTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (c ClearEyedAndMaliciousTactic) Name() string {
	return "鹰视狼顾"
}

func (c ClearEyedAndMaliciousTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	return c
}

func (c ClearEyedAndMaliciousTactic) Id() int64 {
	return ClearEyedAndMalicious
}

func (c ClearEyedAndMaliciousTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (c ClearEyedAndMaliciousTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c ClearEyedAndMaliciousTactic) TriggerRate() float64 {
	return 1.0
}
