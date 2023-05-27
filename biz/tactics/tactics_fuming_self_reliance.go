package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//符命自立
type FumingSelfRelianceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FumingSelfRelianceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (f FumingSelfRelianceTactic) Prepare() {
	panic("implement me")
}

func (f FumingSelfRelianceTactic) Id() consts.TacticId {
	return consts.FumingSelfReliance
}

func (f FumingSelfRelianceTactic) Name() string {
	return "符命自立"
}

func (f FumingSelfRelianceTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (f FumingSelfRelianceTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (f FumingSelfRelianceTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (f FumingSelfRelianceTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (f FumingSelfRelianceTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (f FumingSelfRelianceTactic) Execute() {
	panic("implement me")
}

func (f FumingSelfRelianceTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
