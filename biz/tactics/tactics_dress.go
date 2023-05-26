package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type DressTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DressTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (d DressTactic) Prepare() {
	panic("implement me")
}

func (d DressTactic) Id() consts.TacticId {
	return consts.Dress
}

func (d DressTactic) Name() string {
	return "包扎"
}

func (d DressTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (d DressTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (d DressTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (d DressTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (d DressTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (d DressTactic) Execute() {
	panic("implement me")
}

func (d DressTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
