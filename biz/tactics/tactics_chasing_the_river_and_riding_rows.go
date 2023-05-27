package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//溯江摇橹
type ChasingTheRiverAndRidingRowsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ChasingTheRiverAndRidingRowsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (c ChasingTheRiverAndRidingRowsTactic) Prepare() {
	panic("implement me")
}

func (c ChasingTheRiverAndRidingRowsTactic) Id() consts.TacticId {
	return consts.ChasingTheRiverAndRidingRows
}

func (c ChasingTheRiverAndRidingRowsTactic) Name() string {
	return "溯江摇橹"
}

func (c ChasingTheRiverAndRidingRowsTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (c ChasingTheRiverAndRidingRowsTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (c ChasingTheRiverAndRidingRowsTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (c ChasingTheRiverAndRidingRowsTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (c ChasingTheRiverAndRidingRowsTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (c ChasingTheRiverAndRidingRowsTactic) Execute() {
	panic("implement me")
}

func (c ChasingTheRiverAndRidingRowsTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
