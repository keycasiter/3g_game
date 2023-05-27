package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//精练策数
type RefinedStrategiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RefinedStrategiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (r RefinedStrategiesTactic) Prepare() {
	panic("implement me")
}

func (r RefinedStrategiesTactic) Id() consts.TacticId {
	return consts.RefinedStrategies
}

func (r RefinedStrategiesTactic) Name() string {
	return "精练策数"
}

func (r RefinedStrategiesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (r RefinedStrategiesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (r RefinedStrategiesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (r RefinedStrategiesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (r RefinedStrategiesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (r RefinedStrategiesTactic) Execute() {
	panic("implement me")
}

func (r RefinedStrategiesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
