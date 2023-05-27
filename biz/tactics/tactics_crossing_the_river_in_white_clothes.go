package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//白衣渡江
type CrossingTheRiverInWhiteClothesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CrossingTheRiverInWhiteClothesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (c CrossingTheRiverInWhiteClothesTactic) Prepare() {
	panic("implement me")
}

func (c CrossingTheRiverInWhiteClothesTactic) Id() consts.TacticId {
	return consts.CrossingTheRiverInWhiteClothes
}

func (c CrossingTheRiverInWhiteClothesTactic) Name() string {
	return "白衣渡江"
}

func (c CrossingTheRiverInWhiteClothesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (c CrossingTheRiverInWhiteClothesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (c CrossingTheRiverInWhiteClothesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (c CrossingTheRiverInWhiteClothesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (c CrossingTheRiverInWhiteClothesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (c CrossingTheRiverInWhiteClothesTactic) Execute() {
	panic("implement me")
}

func (c CrossingTheRiverInWhiteClothesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
