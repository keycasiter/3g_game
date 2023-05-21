package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//竭力佐谋
//使敌军智力最高单体智力降低20%，并有70%概率使自身本回合非自带主动战法发动率提高100%，持续1回合
type MakeEveryEffortToAssistInPlanningTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MakeEveryEffortToAssistInPlanningTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	m.tacticsParams = tacticsParams
	m.triggerRate = 0.55
	return m
}

func (m MakeEveryEffortToAssistInPlanningTactic) Prepare() {

}

func (m MakeEveryEffortToAssistInPlanningTactic) Id() consts.TacticId {
	return consts.MakeEveryEffortToAssistInPlanning
}

func (m MakeEveryEffortToAssistInPlanningTactic) Name() string {
	return "竭力佐谋"
}

func (m MakeEveryEffortToAssistInPlanningTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (m MakeEveryEffortToAssistInPlanningTactic) GetTriggerRate() float64 {
	return m.triggerRate
}

func (m MakeEveryEffortToAssistInPlanningTactic) SetTriggerRate(rate float64) {
	m.triggerRate = rate
}

func (m MakeEveryEffortToAssistInPlanningTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (m MakeEveryEffortToAssistInPlanningTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (m MakeEveryEffortToAssistInPlanningTactic) Execute() {

}

func (m MakeEveryEffortToAssistInPlanningTactic) IsTriggerPrepare() bool {
	return false
}
