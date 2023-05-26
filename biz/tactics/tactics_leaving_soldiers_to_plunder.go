package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 纵兵劫掠
type LeavingSoldiersToPlunderTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LeavingSoldiersToPlunderTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (l LeavingSoldiersToPlunderTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (l LeavingSoldiersToPlunderTactic) Id() consts.TacticId {
	return consts.LeavingSoldiersToPlunder
}

func (l LeavingSoldiersToPlunderTactic) Name() string {
	return "纵兵劫掠"
}

func (l LeavingSoldiersToPlunderTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (l LeavingSoldiersToPlunderTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (l LeavingSoldiersToPlunderTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (l LeavingSoldiersToPlunderTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (l LeavingSoldiersToPlunderTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (l LeavingSoldiersToPlunderTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (l LeavingSoldiersToPlunderTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
