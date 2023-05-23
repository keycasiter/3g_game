package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type AnnihilateInOneFellSwoopTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (a AnnihilateInOneFellSwoopTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.4
	return a
}

func (a AnnihilateInOneFellSwoopTactic) Prepare() {

}

func (a AnnihilateInOneFellSwoopTactic) Id() consts.TacticId {
	return consts.AnnihilateInOneFellSwoop
}

func (a AnnihilateInOneFellSwoopTactic) Name() string {
	return "一举歼灭"
}

func (a AnnihilateInOneFellSwoopTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a AnnihilateInOneFellSwoopTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AnnihilateInOneFellSwoopTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AnnihilateInOneFellSwoopTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a AnnihilateInOneFellSwoopTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AnnihilateInOneFellSwoopTactic) Execute() {
}

func (a AnnihilateInOneFellSwoopTactic) IsTriggerPrepare() bool {
	return a.isTriggerPrepare
}
