package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//神射
type DivineEjaculationTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DivineEjaculationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (d DivineEjaculationTactic) Prepare() {
	panic("implement me")
}

func (d DivineEjaculationTactic) Id() consts.TacticId {
	return consts.DivineEjaculation
}

func (d DivineEjaculationTactic) Name() string {
	return "神射"
}

func (d DivineEjaculationTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (d DivineEjaculationTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (d DivineEjaculationTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (d DivineEjaculationTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (d DivineEjaculationTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (d DivineEjaculationTactic) Execute() {
	panic("implement me")
}

func (d DivineEjaculationTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
