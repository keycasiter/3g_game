package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 自愈
type SelfHealingTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SelfHealingTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (s SelfHealingTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (s SelfHealingTactic) Id() consts.TacticId {
	return consts.SelfHealing
}

func (s SelfHealingTactic) Name() string {
	return "自愈"
}

func (s SelfHealingTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (s SelfHealingTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SelfHealingTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (s SelfHealingTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (s SelfHealingTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (s SelfHealingTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (s SelfHealingTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
