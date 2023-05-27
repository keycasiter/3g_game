package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//暗箭难防
type HiddenArrowsAreDifficultToGuardAgainstTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Prepare() {
	panic("implement me")
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Id() consts.TacticId {
	return consts.HiddenArrowsAreDifficultToGuardAgainst
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Name() string {
	return "暗箭难防"
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Execute() {
	panic("implement me")
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
