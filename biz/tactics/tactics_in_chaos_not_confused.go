package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//处兹不惑
type InChaosNotConfusedTacitc struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i InChaosNotConfusedTacitc) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (i InChaosNotConfusedTacitc) Prepare() {
	panic("implement me")
}

func (i InChaosNotConfusedTacitc) Id() consts.TacticId {
	return consts.InChaosNotConfused
}

func (i InChaosNotConfusedTacitc) Name() string {
	return "处兹不惑"
}

func (i InChaosNotConfusedTacitc) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (i InChaosNotConfusedTacitc) GetTriggerRate() float64 {
	panic("implement me")
}

func (i InChaosNotConfusedTacitc) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (i InChaosNotConfusedTacitc) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (i InChaosNotConfusedTacitc) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (i InChaosNotConfusedTacitc) Execute() {
	panic("implement me")
}

func (i InChaosNotConfusedTacitc) IsTriggerPrepare() bool {
	panic("implement me")
}
