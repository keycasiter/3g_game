package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 机略纵横
type MachineStrategyVerticalAndHorizontalTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MachineStrategyVerticalAndHorizontalTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (m MachineStrategyVerticalAndHorizontalTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (m MachineStrategyVerticalAndHorizontalTactic) Id() consts.TacticId {
	return consts.MachineStrategyVerticalAndHorizontal
}

func (m MachineStrategyVerticalAndHorizontalTactic) Name() string {
	return "机略纵横"
}

func (m MachineStrategyVerticalAndHorizontalTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (m MachineStrategyVerticalAndHorizontalTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (m MachineStrategyVerticalAndHorizontalTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (m MachineStrategyVerticalAndHorizontalTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (m MachineStrategyVerticalAndHorizontalTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (m MachineStrategyVerticalAndHorizontalTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (m MachineStrategyVerticalAndHorizontalTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
