package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type SweepAwayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SweepAwayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s SweepAwayTactic) Prepare() {
	panic("implement me")
}

func (s SweepAwayTactic) Id() consts.TacticId {
	return consts.SweepAway
}

func (s SweepAwayTactic) Name() string {
	return "横扫"
}

func (s SweepAwayTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s SweepAwayTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s SweepAwayTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s SweepAwayTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s SweepAwayTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s SweepAwayTactic) Execute() {
	panic("implement me")
}

func (s SweepAwayTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
