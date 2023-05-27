package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//威风凛凛
type AweInspiringTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AweInspiringTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (a AweInspiringTactic) Prepare() {
	panic("implement me")
}

func (a AweInspiringTactic) Id() consts.TacticId {
	return consts.AweInspiring
}

func (a AweInspiringTactic) Name() string {
	return "威风凛凛"
}

func (a AweInspiringTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (a AweInspiringTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (a AweInspiringTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (a AweInspiringTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (a AweInspiringTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (a AweInspiringTactic) Execute() {
	panic("implement me")
}

func (a AweInspiringTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
