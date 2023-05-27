package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//监统震军
type SuperviseLeadAndSeizureArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SuperviseLeadAndSeizureArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s SuperviseLeadAndSeizureArmyTactic) Prepare() {
	panic("implement me")
}

func (s SuperviseLeadAndSeizureArmyTactic) Id() consts.TacticId {
	return consts.SuperviseLeadAndSeizureArmy
}

func (s SuperviseLeadAndSeizureArmyTactic) Name() string {
	return "监统震军"
}

func (s SuperviseLeadAndSeizureArmyTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s SuperviseLeadAndSeizureArmyTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s SuperviseLeadAndSeizureArmyTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s SuperviseLeadAndSeizureArmyTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s SuperviseLeadAndSeizureArmyTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s SuperviseLeadAndSeizureArmyTactic) Execute() {
	panic("implement me")
}

func (s SuperviseLeadAndSeizureArmyTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
