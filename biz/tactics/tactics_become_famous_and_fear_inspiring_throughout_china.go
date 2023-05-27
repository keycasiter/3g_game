package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//威震华夏
type BecomeFamousAndFearInspiringThroughoutChinaTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Prepare() {
	panic("implement me")
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Id() consts.TacticId {
	return consts.BecomeFamousAndFearInspiringThroughoutChina
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Name() string {
	return "威震华夏"
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) Execute() {
	panic("implement me")
}

func (b BecomeFamousAndFearInspiringThroughoutChinaTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
