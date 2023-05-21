package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//气凌三军
//受到普通攻击时对攻击者进行一次反击（伤害率52%），自身为副将时，伤害率提升至74%
type TemperamentSurpassesTheThreeArmies struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TemperamentSurpassesTheThreeArmies) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TemperamentSurpassesTheThreeArmies) Prepare() {
}

func (t TemperamentSurpassesTheThreeArmies) Id() consts.TacticId {
	return consts.TemperamentSurpassesTheThreeArmies
}

func (t TemperamentSurpassesTheThreeArmies) Name() string {
	return "气凌三军"
}

func (t TemperamentSurpassesTheThreeArmies) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TemperamentSurpassesTheThreeArmies) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TemperamentSurpassesTheThreeArmies) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TemperamentSurpassesTheThreeArmies) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t TemperamentSurpassesTheThreeArmies) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TemperamentSurpassesTheThreeArmies) Execute() {

}

func (t TemperamentSurpassesTheThreeArmies) IsTriggerPrepare() bool {
	return false
}
