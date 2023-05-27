package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//镇压黄巾
type SuppressYellowScarvesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SuppressYellowScarvesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s SuppressYellowScarvesTactic) Prepare() {
	panic("implement me")
}

func (s SuppressYellowScarvesTactic) Id() consts.TacticId {
	return consts.SuppressYellowScarves
}

func (s SuppressYellowScarvesTactic) Name() string {
	return "镇压黄巾"
}

func (s SuppressYellowScarvesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s SuppressYellowScarvesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s SuppressYellowScarvesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s SuppressYellowScarvesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s SuppressYellowScarvesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s SuppressYellowScarvesTactic) Execute() {
	panic("implement me")
}

func (s SuppressYellowScarvesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
