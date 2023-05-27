package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//仁德载世
type BenevolentAndVirtuousThroughoutTheWorldTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Prepare() {
	panic("implement me")
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Id() consts.TacticId {
	return consts.BenevolentAndVirtuousThroughoutTheWorld
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Name() string {
	return "仁德载世"
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Execute() {
	panic("implement me")
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
