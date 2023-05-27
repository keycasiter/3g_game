package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//一身是胆
type BebraveAllThroughTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BebraveAllThroughTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (b BebraveAllThroughTactic) Prepare() {
	panic("implement me")
}

func (b BebraveAllThroughTactic) Id() consts.TacticId {
	return consts.BebraveAllThrough
}

func (b BebraveAllThroughTactic) Name() string {
	return "一身是胆"
}

func (b BebraveAllThroughTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BebraveAllThroughTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BebraveAllThroughTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BebraveAllThroughTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BebraveAllThroughTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BebraveAllThroughTactic) Execute() {
	panic("implement me")
}

func (b BebraveAllThroughTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
