package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//勇冠三军
type PeerlessOrMatchlessBraveryOrValourTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Prepare() {
	panic("implement me")
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Id() consts.TacticId {
	return consts.PeerlessOrMatchlessBraveryOrValour
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Name() string {
	return "勇冠三军"
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) Execute() {
	panic("implement me")
}

func (p PeerlessOrMatchlessBraveryOrValourTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
