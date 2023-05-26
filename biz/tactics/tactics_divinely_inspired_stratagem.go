package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type DivinelyInspiredStratagemTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DivinelyInspiredStratagemTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (d DivinelyInspiredStratagemTactic) Prepare() {
	panic("implement me")
}

func (d DivinelyInspiredStratagemTactic) Id() consts.TacticId {
	return consts.DivinelyInspiredStratagem
}

func (d DivinelyInspiredStratagemTactic) Name() string {
	return "神机莫测"
}

func (d DivinelyInspiredStratagemTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (d DivinelyInspiredStratagemTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (d DivinelyInspiredStratagemTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (d DivinelyInspiredStratagemTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (d DivinelyInspiredStratagemTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (d DivinelyInspiredStratagemTactic) Execute() {
	panic("implement me")
}

func (d DivinelyInspiredStratagemTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
