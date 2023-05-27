package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//刀出如霆
type TheKnifeLikeThunderboltTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheKnifeLikeThunderboltTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) Prepare() {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) Id() consts.TacticId {
	return consts.TheKnifeLikeThunderbolt
}

func (t TheKnifeLikeThunderboltTactic) Name() string {
	return "刀出如霆"
}

func (t TheKnifeLikeThunderboltTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) Execute() {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
