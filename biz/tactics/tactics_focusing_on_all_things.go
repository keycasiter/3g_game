package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type FocusingOnAllThingsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FocusingOnAllThingsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (f FocusingOnAllThingsTactic) Prepare() {
	panic("implement me")
}

func (f FocusingOnAllThingsTactic) Id() consts.TacticId {
	return consts.FocusingOnAllThings
}

func (f FocusingOnAllThingsTactic) Name() string {
	return "垂心万物"
}

func (f FocusingOnAllThingsTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (f FocusingOnAllThingsTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (f FocusingOnAllThingsTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (f FocusingOnAllThingsTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (f FocusingOnAllThingsTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (f FocusingOnAllThingsTactic) Execute() {
	panic("implement me")
}

func (f FocusingOnAllThingsTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
