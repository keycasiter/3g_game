package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//顾盼生姿
type LookAroundCharminglyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LookAroundCharminglyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (l LookAroundCharminglyTactic) Prepare() {
	panic("implement me")
}

func (l LookAroundCharminglyTactic) Id() consts.TacticId {
	return consts.LookAroundCharmingly
}

func (l LookAroundCharminglyTactic) Name() string {
	return "顾盼生姿"
}

func (l LookAroundCharminglyTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (l LookAroundCharminglyTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (l LookAroundCharminglyTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (l LookAroundCharminglyTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (l LookAroundCharminglyTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (l LookAroundCharminglyTactic) Execute() {
	panic("implement me")
}

func (l LookAroundCharminglyTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
