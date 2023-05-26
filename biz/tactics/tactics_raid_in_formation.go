package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type RaidInFormationTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RaidInFormationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (r RaidInFormationTactic) Prepare() {
	panic("implement me")
}

func (r RaidInFormationTactic) Id() consts.TacticId {
	return consts.RaidInFormation
}

func (r RaidInFormationTactic) Name() string {
	return "陷阵突袭"
}

func (r RaidInFormationTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (r RaidInFormationTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (r RaidInFormationTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (r RaidInFormationTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (r RaidInFormationTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (r RaidInFormationTactic) Execute() {
	panic("implement me")
}

func (r RaidInFormationTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
