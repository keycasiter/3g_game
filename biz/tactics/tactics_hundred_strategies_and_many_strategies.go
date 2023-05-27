package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//百计多谋
type HundredStrategiesAndManyStrategiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HundredStrategiesAndManyStrategiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (h HundredStrategiesAndManyStrategiesTactic) Prepare() {
	panic("implement me")
}

func (h HundredStrategiesAndManyStrategiesTactic) Id() consts.TacticId {
	return consts.HundredStrategiesAndManyStrategies
}

func (h HundredStrategiesAndManyStrategiesTactic) Name() string {
	return "百计多谋"
}

func (h HundredStrategiesAndManyStrategiesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (h HundredStrategiesAndManyStrategiesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (h HundredStrategiesAndManyStrategiesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (h HundredStrategiesAndManyStrategiesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (h HundredStrategiesAndManyStrategiesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (h HundredStrategiesAndManyStrategiesTactic) Execute() {
	panic("implement me")
}

func (h HundredStrategiesAndManyStrategiesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
