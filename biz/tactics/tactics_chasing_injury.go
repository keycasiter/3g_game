package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 追伤
type ChasingInjuryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ChasingInjuryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (c ChasingInjuryTactic) Prepare() {
	panic("implement me")
}

func (c ChasingInjuryTactic) Id() consts.TacticId {
	return consts.ChasingInjury
}

func (c ChasingInjuryTactic) Name() string {
	return "追伤"
}

func (c ChasingInjuryTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (c ChasingInjuryTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (c ChasingInjuryTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (c ChasingInjuryTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (c ChasingInjuryTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (c ChasingInjuryTactic) Execute() {
	panic("implement me")
}

func (c ChasingInjuryTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
