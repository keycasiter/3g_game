package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//高橹连营
type HighWoodenPaddlesConnectedToTheCampTacitc struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) Prepare() {
	panic("implement me")
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) Id() consts.TacticId {
	return consts.HighWoodenPaddlesConnectedToTheCamp
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) Name() string {
	return "高橹连营"
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) GetTriggerRate() float64 {
	panic("implement me")
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) Execute() {
	panic("implement me")
}

func (h HighWoodenPaddlesConnectedToTheCampTacitc) IsTriggerPrepare() bool {
	panic("implement me")
}
