package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 裸衣血战
// 战斗中无法发动主动战法，战斗前3回合，获得90%连击及10%倒戈，并使自身及敌军单体统率降低40%
type NakedBloodBattleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (n NakedBloodBattleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	n.tacticsParams = tacticsParams
	n.triggerRate = 1.0
	return n
}

func (n NakedBloodBattleTactic) Prepare() {
}

func (n NakedBloodBattleTactic) Id() consts.TacticId {
	return consts.NakedBloodBattle
}

func (n NakedBloodBattleTactic) Name() string {
	return "裸衣血战"
}

func (n NakedBloodBattleTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (n NakedBloodBattleTactic) GetTriggerRate() float64 {
	return n.triggerRate
}

func (n NakedBloodBattleTactic) SetTriggerRate(rate float64) {
	n.triggerRate = rate
}

func (n NakedBloodBattleTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (n NakedBloodBattleTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (n NakedBloodBattleTactic) Execute() {

}

func (n NakedBloodBattleTactic) IsTriggerPrepare() bool {
	return false
}
