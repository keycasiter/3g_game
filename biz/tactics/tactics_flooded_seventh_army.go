package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//水淹七军
type FloodedSeventhArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FloodedSeventhArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (f FloodedSeventhArmyTactic) Prepare() {
	panic("implement me")
}

func (f FloodedSeventhArmyTactic) Id() consts.TacticId {
	return consts.FloodedSeventhArmy
}

func (f FloodedSeventhArmyTactic) Name() string {
	return "水淹七军"
}

func (f FloodedSeventhArmyTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (f FloodedSeventhArmyTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (f FloodedSeventhArmyTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (f FloodedSeventhArmyTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (f FloodedSeventhArmyTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (f FloodedSeventhArmyTactic) Execute() {
	panic("implement me")
}

func (f FloodedSeventhArmyTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
