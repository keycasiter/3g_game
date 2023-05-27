package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//窃幸乘宠
type StealingLuckAndRidingPetsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StealingLuckAndRidingPetsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s StealingLuckAndRidingPetsTactic) Prepare() {
	panic("implement me")
}

func (s StealingLuckAndRidingPetsTactic) Id() consts.TacticId {
	return consts.StealingLuckAndRidingPets
}

func (s StealingLuckAndRidingPetsTactic) Name() string {
	return "窃幸乘宠"
}

func (s StealingLuckAndRidingPetsTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s StealingLuckAndRidingPetsTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s StealingLuckAndRidingPetsTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s StealingLuckAndRidingPetsTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s StealingLuckAndRidingPetsTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s StealingLuckAndRidingPetsTactic) Execute() {
	panic("implement me")
}

func (s StealingLuckAndRidingPetsTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
