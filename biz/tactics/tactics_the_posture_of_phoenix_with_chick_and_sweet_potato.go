package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//鸱苕凤姿
type ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Prepare() {
	panic("implement me")
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Id() consts.TacticId {
	return consts.ThePostureOfAPhoenixWithAChickAndASweetPotato
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Name() string {
	return "鸱苕凤姿"
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Execute() {
	panic("implement me")
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
