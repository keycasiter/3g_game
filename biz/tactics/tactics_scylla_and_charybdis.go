package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//腹背受敌
type ScyllaAndCharybdisTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s ScyllaAndCharybdisTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s ScyllaAndCharybdisTactic) Prepare() {
	panic("implement me")
}

func (s ScyllaAndCharybdisTactic) Id() consts.TacticId {
	return consts.ScyllaAndCharybdis
}

func (s ScyllaAndCharybdisTactic) Name() string {
	return "腹背受敌"
}

func (s ScyllaAndCharybdisTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s ScyllaAndCharybdisTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s ScyllaAndCharybdisTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s ScyllaAndCharybdisTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s ScyllaAndCharybdisTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s ScyllaAndCharybdisTactic) Execute() {
	panic("implement me")
}

func (s ScyllaAndCharybdisTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
