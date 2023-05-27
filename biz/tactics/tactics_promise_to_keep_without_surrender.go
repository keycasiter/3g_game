package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type PromiseToKeepWithoutSurrenderTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PromiseToKeepWithoutSurrenderTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (p PromiseToKeepWithoutSurrenderTactic) Prepare() {
	panic("implement me")
}

func (p PromiseToKeepWithoutSurrenderTactic) Id() consts.TacticId {
	return consts.PromiseToKeepWithoutSurrender
}

func (p PromiseToKeepWithoutSurrenderTactic) Name() string {
	return "誓守无降"
}

func (p PromiseToKeepWithoutSurrenderTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (p PromiseToKeepWithoutSurrenderTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (p PromiseToKeepWithoutSurrenderTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (p PromiseToKeepWithoutSurrenderTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (p PromiseToKeepWithoutSurrenderTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (p PromiseToKeepWithoutSurrenderTactic) Execute() {
	panic("implement me")
}

func (p PromiseToKeepWithoutSurrenderTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
