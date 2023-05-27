package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//以逸待劳
type WaitAtOnesEaseForTheFatiguedTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WaitAtOnesEaseForTheFatiguedTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (w WaitAtOnesEaseForTheFatiguedTactic) Prepare() {
	panic("implement me")
}

func (w WaitAtOnesEaseForTheFatiguedTactic) Id() consts.TacticId {
	return consts.WaitAtOnesEaseForTheFatigued
}

func (w WaitAtOnesEaseForTheFatiguedTactic) Name() string {
	return "以逸待劳"
}

func (w WaitAtOnesEaseForTheFatiguedTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (w WaitAtOnesEaseForTheFatiguedTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (w WaitAtOnesEaseForTheFatiguedTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (w WaitAtOnesEaseForTheFatiguedTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (w WaitAtOnesEaseForTheFatiguedTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (w WaitAtOnesEaseForTheFatiguedTactic) Execute() {
	panic("implement me")
}

func (w WaitAtOnesEaseForTheFatiguedTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
