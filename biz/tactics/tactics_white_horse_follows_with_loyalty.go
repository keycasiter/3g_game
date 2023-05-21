package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type WhiteHorseFollowsWithLoyaltyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 1.0
	return w
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Prepare() {
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Id() consts.TacticId {
	return consts.WhiteHorseFollowsWithLoyalty
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Name() string {
	return "白马义从"
}

func (w WhiteHorseFollowsWithLoyaltyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (w WhiteHorseFollowsWithLoyaltyTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WhiteHorseFollowsWithLoyaltyTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WhiteHorseFollowsWithLoyaltyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (w WhiteHorseFollowsWithLoyaltyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (w WhiteHorseFollowsWithLoyaltyTactic) Execute() {
}

func (w WhiteHorseFollowsWithLoyaltyTactic) IsTriggerPrepare() bool {
	return false
}
