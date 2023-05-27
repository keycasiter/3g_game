package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//槊血纵横
type BloodyAndUnrestrainedTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BloodyAndUnrestrainedTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) Prepare() {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) Id() consts.TacticId {
	return consts.BloodyAndUnrestrained
}

func (b BloodyAndUnrestrainedTactic) Name() string {
	return "槊血纵横"
}

func (b BloodyAndUnrestrainedTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) Execute() {
	panic("implement me")
}

func (b BloodyAndUnrestrainedTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
