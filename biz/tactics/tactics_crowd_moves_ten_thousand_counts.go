package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//众动万计
//受到普通攻击时，有45%概率对攻击来源造成兵刃伤害（伤害率140%），并使其下一次造成的伤害伤减少40%
type CrowdMovesTenThousandCountsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CrowdMovesTenThousandCountsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c CrowdMovesTenThousandCountsTactic) Prepare() {
}

func (c CrowdMovesTenThousandCountsTactic) Id() consts.TacticId {
	return consts.CrowdMovesTenThousandCounts
}

func (c CrowdMovesTenThousandCountsTactic) Name() string {
	return "众动万计"
}

func (c CrowdMovesTenThousandCountsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (c CrowdMovesTenThousandCountsTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CrowdMovesTenThousandCountsTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CrowdMovesTenThousandCountsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (c CrowdMovesTenThousandCountsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CrowdMovesTenThousandCountsTactic) Execute() {

}

func (c CrowdMovesTenThousandCountsTactic) IsTriggerPrepare() bool {
	return false
}
