package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type ProvokeTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p ProvokeTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (p ProvokeTactic) Prepare() {
	panic("implement me")
}

func (p ProvokeTactic) Id() consts.TacticId {
	return consts.Provoke
}

func (p ProvokeTactic) Name() string {
	return "挑衅"
}

func (p ProvokeTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (p ProvokeTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (p ProvokeTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (p ProvokeTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (p ProvokeTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (p ProvokeTactic) Execute() {
	panic("implement me")
}

func (p ProvokeTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
