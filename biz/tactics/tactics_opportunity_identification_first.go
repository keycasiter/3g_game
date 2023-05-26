package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//机鉴先识
type OpportunityIdentificationFirstTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (o OpportunityIdentificationFirstTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (o OpportunityIdentificationFirstTactic) Prepare() {
	panic("implement me")
}

func (o OpportunityIdentificationFirstTactic) Id() consts.TacticId {
	return consts.OpportunityIdentificationFirst
}

func (o OpportunityIdentificationFirstTactic) Name() string {
	return "机鉴先识"
}

func (o OpportunityIdentificationFirstTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (o OpportunityIdentificationFirstTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (o OpportunityIdentificationFirstTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (o OpportunityIdentificationFirstTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (o OpportunityIdentificationFirstTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (o OpportunityIdentificationFirstTactic) Execute() {
	panic("implement me")
}

func (o OpportunityIdentificationFirstTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
