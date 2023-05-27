package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type AdvancingSecretlyByUnknownPathTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AdvancingSecretlyByUnknownPathTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (a AdvancingSecretlyByUnknownPathTactic) Prepare() {
	panic("implement me")
}

func (a AdvancingSecretlyByUnknownPathTactic) Id() consts.TacticId {
	return consts.AdvancingSecretlyByUnknownPath
}

func (a AdvancingSecretlyByUnknownPathTactic) Name() string {
	return "暗渡陈仓"
}

func (a AdvancingSecretlyByUnknownPathTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (a AdvancingSecretlyByUnknownPathTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (a AdvancingSecretlyByUnknownPathTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (a AdvancingSecretlyByUnknownPathTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (a AdvancingSecretlyByUnknownPathTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (a AdvancingSecretlyByUnknownPathTactic) Execute() {
	panic("implement me")
}

func (a AdvancingSecretlyByUnknownPathTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
