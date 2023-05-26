package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 手起刀落
type CutDownTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CutDownTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (c CutDownTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (c CutDownTactic) Id() consts.TacticId {
	return consts.CutDown
}

func (c CutDownTactic) Name() string {
	return "手起刀落"
}

func (c CutDownTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (c CutDownTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CutDownTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (c CutDownTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (c CutDownTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (c CutDownTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (c CutDownTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
