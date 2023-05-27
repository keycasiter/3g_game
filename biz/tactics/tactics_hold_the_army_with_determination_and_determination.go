package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//持军毅重
type HoldTheArmyWithDeterminationAndDeterminationTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Prepare() {
	panic("implement me")
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Id() consts.TacticId {
	return consts.HoldTheArmyWithDeterminationAndDetermination
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Name() string {
	return "持军毅重"
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) Execute() {
	panic("implement me")
}

func (h HoldTheArmyWithDeterminationAndDeterminationTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
