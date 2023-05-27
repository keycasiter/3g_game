package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type TheBattleOfCarryingCoffinTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheBattleOfCarryingCoffinTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TheBattleOfCarryingCoffinTactic) Prepare() {
	panic("implement me")
}

func (t TheBattleOfCarryingCoffinTactic) Id() consts.TacticId {
	return consts.TheBattleOfCarryingCoffin
}

func (t TheBattleOfCarryingCoffinTactic) Name() string {
	return "抬棺决战"
}

func (t TheBattleOfCarryingCoffinTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t TheBattleOfCarryingCoffinTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t TheBattleOfCarryingCoffinTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t TheBattleOfCarryingCoffinTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t TheBattleOfCarryingCoffinTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t TheBattleOfCarryingCoffinTactic) Execute() {
	panic("implement me")
}

func (t TheBattleOfCarryingCoffinTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
