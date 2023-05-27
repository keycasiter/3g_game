package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//毒泉拒蜀
type PoisonousSpringRefusesShuTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PoisonousSpringRefusesShuTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) Prepare() {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) Id() consts.TacticId {
	return consts.PoisonousSpringRefusesShu
}

func (p PoisonousSpringRefusesShuTactic) Name() string {
	return "毒泉拒蜀"
}

func (p PoisonousSpringRefusesShuTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) Execute() {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
