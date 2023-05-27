package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//云聚影从
type CloudGatheringShadowFromTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CloudGatheringShadowFromTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (c CloudGatheringShadowFromTactic) Prepare() {
	panic("implement me")
}

func (c CloudGatheringShadowFromTactic) Id() consts.TacticId {
	return consts.CloudGatheringShadowFrom
}

func (c CloudGatheringShadowFromTactic) Name() string {
	return "云聚影从"
}

func (c CloudGatheringShadowFromTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (c CloudGatheringShadowFromTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (c CloudGatheringShadowFromTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (c CloudGatheringShadowFromTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (c CloudGatheringShadowFromTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (c CloudGatheringShadowFromTactic) Execute() {
	panic("implement me")
}

func (c CloudGatheringShadowFromTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
