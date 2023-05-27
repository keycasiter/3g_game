package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//计定谋决
type PlanAndDecideTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PlanAndDecideTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (p PlanAndDecideTactic) Prepare() {
	panic("implement me")
}

func (p PlanAndDecideTactic) Id() consts.TacticId {
	return consts.PlanAndDecide
}

func (p PlanAndDecideTactic) Name() string {
	return "计定谋决"
}

func (p PlanAndDecideTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (p PlanAndDecideTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (p PlanAndDecideTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (p PlanAndDecideTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (p PlanAndDecideTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (p PlanAndDecideTactic) Execute() {
	panic("implement me")
}

func (p PlanAndDecideTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
