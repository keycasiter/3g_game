package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//酒池肉林
type ExtravagantOrgyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (e ExtravagantOrgyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (e ExtravagantOrgyTactic) Prepare() {
	panic("implement me")
}

func (e ExtravagantOrgyTactic) Id() consts.TacticId {
	return consts.ExtravagantOrgy
}

func (e ExtravagantOrgyTactic) Name() string {
	return "酒池肉林"
}

func (e ExtravagantOrgyTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (e ExtravagantOrgyTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (e ExtravagantOrgyTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (e ExtravagantOrgyTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (e ExtravagantOrgyTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (e ExtravagantOrgyTactic) Execute() {
	panic("implement me")
}

func (e ExtravagantOrgyTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
