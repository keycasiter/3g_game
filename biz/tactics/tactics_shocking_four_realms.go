package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//震骇四境
type ShockingFourRealmsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s ShockingFourRealmsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s ShockingFourRealmsTactic) Prepare() {
	panic("implement me")
}

func (s ShockingFourRealmsTactic) Id() consts.TacticId {
	return consts.ShockingFourRealms
}

func (s ShockingFourRealmsTactic) Name() string {
	return "震骇四境"
}

func (s ShockingFourRealmsTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s ShockingFourRealmsTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s ShockingFourRealmsTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s ShockingFourRealmsTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s ShockingFourRealmsTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s ShockingFourRealmsTactic) Execute() {
	panic("implement me")
}

func (s ShockingFourRealmsTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
