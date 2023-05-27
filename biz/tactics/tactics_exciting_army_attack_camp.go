package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//振军击营
type ExcitingArmyAttackCampTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (e ExcitingArmyAttackCampTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (e ExcitingArmyAttackCampTactic) Prepare() {
	panic("implement me")
}

func (e ExcitingArmyAttackCampTactic) Id() consts.TacticId {
	return consts.ExcitingArmyAttackCamp
}

func (e ExcitingArmyAttackCampTactic) Name() string {
	return "振军击营"
}

func (e ExcitingArmyAttackCampTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (e ExcitingArmyAttackCampTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (e ExcitingArmyAttackCampTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (e ExcitingArmyAttackCampTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (e ExcitingArmyAttackCampTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (e ExcitingArmyAttackCampTactic) Execute() {
	panic("implement me")
}

func (e ExcitingArmyAttackCampTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
