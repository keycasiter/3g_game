package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//倾国倾城
type BeautyWhichOverthrowsStatesAndCitiesTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Prepare() {
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Id() consts.TacticId {
	return consts.BeautyWhichOverthrowsStatesAndCities
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Name() string {
	return "倾国倾城"
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) Execute() {
}

func (b BeautyWhichOverthrowsStatesAndCitiesTactic) IsTriggerPrepare() bool {
	return b.isTriggerPrepare
}
