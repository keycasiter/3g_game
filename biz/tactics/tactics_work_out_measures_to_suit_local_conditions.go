package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//兵无常势
//战斗中，自己累积进行第3次普通攻击时，对攻击目标造成谋略伤害（伤害率240%，受智力影响），
//并治疗自己（治疗率180%，受智力影响）
type WorkOutMeasuresToSuitLocalConditionsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 1.0
	return w
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Prepare() {
	panic("implement me")
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Id() consts.TacticId {
	return consts.WorkOutMeasuresToSuitLocalConditions
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Name() string {
	return "兵无常势"
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) Execute() {
}

func (w WorkOutMeasuresToSuitLocalConditionsTactic) IsTriggerPrepare() bool {
	return false
}