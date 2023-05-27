package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//五雷轰顶
type ThunderStruckTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThunderStruckTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t ThunderStruckTactic) Prepare() {
	panic("implement me")
}

func (t ThunderStruckTactic) Id() consts.TacticId {
	return consts.ThunderStruck
}

func (t ThunderStruckTactic) Name() string {
	return "五雷轰顶"
}

func (t ThunderStruckTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t ThunderStruckTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t ThunderStruckTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t ThunderStruckTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t ThunderStruckTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t ThunderStruckTactic) Execute() {
	panic("implement me")
}

func (t ThunderStruckTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
