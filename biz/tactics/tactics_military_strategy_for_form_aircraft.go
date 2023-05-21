package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//形机军略
//对敌军单体造成一次兵刃攻击（伤害率210%）及谋略攻击（伤害率180%，受智力影响）
type MilitaryStrategyForFormAircraftTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MilitaryStrategyForFormAircraftTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	m.tacticsParams = tacticsParams
	m.triggerRate = 0.4
	return m
}

func (m MilitaryStrategyForFormAircraftTactic) Prepare() {
}

func (m MilitaryStrategyForFormAircraftTactic) Id() consts.TacticId {
	return consts.MilitaryStrategyForFormAircraft
}

func (m MilitaryStrategyForFormAircraftTactic) Name() string {
	return "形机军略"
}

func (m MilitaryStrategyForFormAircraftTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (m MilitaryStrategyForFormAircraftTactic) GetTriggerRate() float64 {
	return m.triggerRate
}

func (m MilitaryStrategyForFormAircraftTactic) SetTriggerRate(rate float64) {
	m.triggerRate = rate
}

func (m MilitaryStrategyForFormAircraftTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (m MilitaryStrategyForFormAircraftTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (m MilitaryStrategyForFormAircraftTactic) Execute() {
}

func (m MilitaryStrategyForFormAircraftTactic) IsTriggerPrepare() bool {
	return false
}
