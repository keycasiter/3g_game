package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//奇兵间道
type SpecialSoldierPassRoadTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SpecialSoldierPassRoadTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s SpecialSoldierPassRoadTactic) Prepare() {
	panic("implement me")
}

func (s SpecialSoldierPassRoadTactic) Id() consts.TacticId {
	return consts.SpecialSoldierPassRoad
}

func (s SpecialSoldierPassRoadTactic) Name() string {
	return "奇兵间道"
}

func (s SpecialSoldierPassRoadTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s SpecialSoldierPassRoadTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s SpecialSoldierPassRoadTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s SpecialSoldierPassRoadTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s SpecialSoldierPassRoadTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s SpecialSoldierPassRoadTactic) Execute() {
	panic("implement me")
}

func (s SpecialSoldierPassRoadTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
