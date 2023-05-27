package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//刚烈不屈
type StrongAndUnyieldingTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StrongAndUnyieldingTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s StrongAndUnyieldingTactic) Prepare() {
	panic("implement me")
}

func (s StrongAndUnyieldingTactic) Id() consts.TacticId {
	return consts.StrongAndUnyielding
}

func (s StrongAndUnyieldingTactic) Name() string {
	return "刚烈不屈"
}

func (s StrongAndUnyieldingTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s StrongAndUnyieldingTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s StrongAndUnyieldingTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s StrongAndUnyieldingTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s StrongAndUnyieldingTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s StrongAndUnyieldingTactic) Execute() {
	panic("implement me")
}

func (s StrongAndUnyieldingTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
