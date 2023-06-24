package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 以逸待劳
// 智力我军群体（2人，治疗率154%，受智力影响），
// 并使其瑕疵3次受到伤害分别降低（50%、37.5%、25%，受智力影响）、且下次受到控制状态时有40%（受智力影响）几率免疫
// 主动，35%
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
