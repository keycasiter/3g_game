package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//鸩毒
type PoisonedWineTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PoisonedWineTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (p PoisonedWineTactic) Prepare() {
	panic("implement me")
}

func (p PoisonedWineTactic) Id() consts.TacticId {
	return consts.PoisonedWine
}

func (p PoisonedWineTactic) Name() string {
	return "鸩毒"
}

func (p PoisonedWineTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (p PoisonedWineTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (p PoisonedWineTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (p PoisonedWineTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (p PoisonedWineTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (p PoisonedWineTactic) Execute() {
	panic("implement me")
}

func (p PoisonedWineTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
