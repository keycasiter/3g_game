package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 白眉
type EyebrowedThrushTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (e EyebrowedThrushTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (e EyebrowedThrushTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (e EyebrowedThrushTactic) Id() consts.TacticId {
	return consts.EyebrowedThrush
}

func (e EyebrowedThrushTactic) Name() string {
	return "白眉"
}

func (e EyebrowedThrushTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (e EyebrowedThrushTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (e EyebrowedThrushTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (e EyebrowedThrushTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (e EyebrowedThrushTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (e EyebrowedThrushTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (e EyebrowedThrushTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
