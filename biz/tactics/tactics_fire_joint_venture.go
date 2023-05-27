package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//火烧连营
type FireJointVentureTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FireJointVentureTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (f FireJointVentureTactic) Prepare() {
	panic("implement me")
}

func (f FireJointVentureTactic) Id() consts.TacticId {
	return consts.FireJointVenture
}

func (f FireJointVentureTactic) Name() string {
	return "火烧连营"
}

func (f FireJointVentureTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (f FireJointVentureTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (f FireJointVentureTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (f FireJointVentureTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (f FireJointVentureTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (f FireJointVentureTactic) Execute() {
	panic("implement me")
}

func (f FireJointVentureTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
