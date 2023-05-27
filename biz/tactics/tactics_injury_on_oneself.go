package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//苦肉计
type InjuryOnOneselfTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i InjuryOnOneselfTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (i InjuryOnOneselfTactic) Prepare() {
	panic("implement me")
}

func (i InjuryOnOneselfTactic) Id() consts.TacticId {
	return consts.InjuryOnOneself
}

func (i InjuryOnOneselfTactic) Name() string {
	return "苦肉计"
}

func (i InjuryOnOneselfTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (i InjuryOnOneselfTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (i InjuryOnOneselfTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (i InjuryOnOneselfTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (i InjuryOnOneselfTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (i InjuryOnOneselfTactic) Execute() {
	panic("implement me")
}

func (i InjuryOnOneselfTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
