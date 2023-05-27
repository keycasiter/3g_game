package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//锦囊妙计
type BrocadeBagAndCleverPlanTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BrocadeBagAndCleverPlanTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (b BrocadeBagAndCleverPlanTactic) Prepare() {
	panic("implement me")
}

func (b BrocadeBagAndCleverPlanTactic) Id() consts.TacticId {
	return consts.BrocadeBagAndCleverPlan
}

func (b BrocadeBagAndCleverPlanTactic) Name() string {
	return "锦囊妙计"
}

func (b BrocadeBagAndCleverPlanTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BrocadeBagAndCleverPlanTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BrocadeBagAndCleverPlanTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BrocadeBagAndCleverPlanTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BrocadeBagAndCleverPlanTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BrocadeBagAndCleverPlanTactic) Execute() {
	panic("implement me")
}

func (b BrocadeBagAndCleverPlanTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
