package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 弓腰姬
type BowWaistConcubineTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BowWaistConcubineTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (b BowWaistConcubineTactic) Prepare() {
	panic("implement me")
}

func (b BowWaistConcubineTactic) Id() consts.TacticId {
	return consts.BowWaistConcubine
}

func (b BowWaistConcubineTactic) Name() string {
	return "弓腰姬"
}

func (b BowWaistConcubineTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BowWaistConcubineTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BowWaistConcubineTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BowWaistConcubineTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BowWaistConcubineTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BowWaistConcubineTactic) Execute() {
	panic("implement me")
}

func (b BowWaistConcubineTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
