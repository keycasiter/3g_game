package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//闭月
type ClosedMoonTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ClosedMoonTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (c ClosedMoonTactic) Prepare() {
	panic("implement me")
}

func (c ClosedMoonTactic) Id() consts.TacticId {
	return consts.ClosedMoon
}

func (c ClosedMoonTactic) Name() string {
	return "闭月"
}

func (c ClosedMoonTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (c ClosedMoonTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (c ClosedMoonTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (c ClosedMoonTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (c ClosedMoonTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (c ClosedMoonTactic) Execute() {
	panic("implement me")
}

func (c ClosedMoonTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
