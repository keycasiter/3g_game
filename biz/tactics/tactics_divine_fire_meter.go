package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//神火计
type DivineFireMeterTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DivineFireMeterTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (d DivineFireMeterTactic) Prepare() {
	panic("implement me")
}

func (d DivineFireMeterTactic) Id() consts.TacticId {
	return consts.DivineFireMeter
}

func (d DivineFireMeterTactic) Name() string {
	return "神火计"
}

func (d DivineFireMeterTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (d DivineFireMeterTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (d DivineFireMeterTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (d DivineFireMeterTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (d DivineFireMeterTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (d DivineFireMeterTactic) Execute() {
	panic("implement me")
}

func (d DivineFireMeterTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
