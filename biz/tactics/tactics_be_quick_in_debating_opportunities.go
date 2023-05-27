package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//才辩机捷
type BeQuickInDebatingOpportunitiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BeQuickInDebatingOpportunitiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (b BeQuickInDebatingOpportunitiesTactic) Prepare() {
	panic("implement me")
}

func (b BeQuickInDebatingOpportunitiesTactic) Id() consts.TacticId {
	return consts.BeQuickInDebatingOpportunities
}

func (b BeQuickInDebatingOpportunitiesTactic) Name() string {
	return "才辩机捷"
}

func (b BeQuickInDebatingOpportunitiesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BeQuickInDebatingOpportunitiesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BeQuickInDebatingOpportunitiesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BeQuickInDebatingOpportunitiesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BeQuickInDebatingOpportunitiesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BeQuickInDebatingOpportunitiesTactic) Execute() {
	panic("implement me")
}

func (b BeQuickInDebatingOpportunitiesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
