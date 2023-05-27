package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//沉断机谋
type DecisiveStrategyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DecisiveStrategyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (d DecisiveStrategyTactic) Prepare() {
	panic("implement me")
}

func (d DecisiveStrategyTactic) Id() consts.TacticId {
	return consts.DecisiveStrategy
}

func (d DecisiveStrategyTactic) Name() string {
	return "沉断机谋"
}

func (d DecisiveStrategyTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (d DecisiveStrategyTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (d DecisiveStrategyTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (d DecisiveStrategyTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (d DecisiveStrategyTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (d DecisiveStrategyTactic) Execute() {
	panic("implement me")
}

func (d DecisiveStrategyTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
