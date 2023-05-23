package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 不辱使命
// 对敌军单体造成一次兵刃攻击（伤害率220%），并有30%概率施加震慑状态（无法行动），持续1回合
type HaveSucceededInCarryingOutAnAssignmentTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.4
	return h
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Prepare() {

}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Id() consts.TacticId {
	return consts.HaveSucceededInCarryingOutAnAssignment
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Name() string {
	return "不辱使命"
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) Execute() {

}

func (h HaveSucceededInCarryingOutAnAssignmentTactic) IsTriggerPrepare() bool {
	return false
}
