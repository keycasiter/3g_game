package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type SittingInAnIsolatedCityTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SittingInAnIsolatedCityTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (s SittingInAnIsolatedCityTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (s SittingInAnIsolatedCityTactic) Id() consts.TacticId {
	return consts.SittingInAnIsolatedCity
}

func (s SittingInAnIsolatedCityTactic) Name() string {
	return "坐守孤城"
}

func (s SittingInAnIsolatedCityTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (s SittingInAnIsolatedCityTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SittingInAnIsolatedCityTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (s SittingInAnIsolatedCityTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (s SittingInAnIsolatedCityTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (s SittingInAnIsolatedCityTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (s SittingInAnIsolatedCityTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
