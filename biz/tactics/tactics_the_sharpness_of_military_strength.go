package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 兵锋
// 使自己及友军单体进入连击（每回合可以普通攻击2次）状态，持续1回合
// 35%
type TheSharpnessOfMilitaryStrengthTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheSharpnessOfMilitaryStrengthTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.35
	return t
}

func (t TheSharpnessOfMilitaryStrengthTactic) Prepare() {

}

func (t TheSharpnessOfMilitaryStrengthTactic) Id() consts.TacticId {
	return consts.TheSharpnessOfMilitaryStrength
}

func (t TheSharpnessOfMilitaryStrengthTactic) Name() string {
	return "兵锋"
}

func (t TheSharpnessOfMilitaryStrengthTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheSharpnessOfMilitaryStrengthTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheSharpnessOfMilitaryStrengthTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheSharpnessOfMilitaryStrengthTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TheSharpnessOfMilitaryStrengthTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheSharpnessOfMilitaryStrengthTactic) Execute() {
}

func (t TheSharpnessOfMilitaryStrengthTactic) IsTriggerPrepare() bool {
	return false
}

func (a TheSharpnessOfMilitaryStrengthTactic) SetTriggerPrepare(triggerPrepare bool) {
}
