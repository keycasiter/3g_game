package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//金城汤池
type RampartsOfMetalsAndAMoatOfHotWaterTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Prepare() {
	panic("implement me")
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Id() consts.TacticId {
	return consts.RampartsOfMetalsAndAMoatOfHotWater
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Name() string {
	return "金城汤池"
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) Execute() {
	panic("implement me")
}

func (r RampartsOfMetalsAndAMoatOfHotWaterTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
