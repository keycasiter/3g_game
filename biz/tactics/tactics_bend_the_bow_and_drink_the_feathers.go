package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 弯弓饮羽
// 普通攻击之后，使攻击目标降低150点统率，并造成计穷（无法发动主动战法）状态，持续1回合
type BendTheBowAndDrinkTheFeathersTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BendTheBowAndDrinkTheFeathersTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BendTheBowAndDrinkTheFeathersTactic) Prepare() {

}

func (b BendTheBowAndDrinkTheFeathersTactic) Id() consts.TacticId {
	return consts.BendTheBowAndDrinkTheFeathers
}

func (b BendTheBowAndDrinkTheFeathersTactic) Name() string {
	return "弯弓饮羽"
}

func (b BendTheBowAndDrinkTheFeathersTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BendTheBowAndDrinkTheFeathersTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BendTheBowAndDrinkTheFeathersTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BendTheBowAndDrinkTheFeathersTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (b BendTheBowAndDrinkTheFeathersTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BendTheBowAndDrinkTheFeathersTactic) Execute() {
}

func (b BendTheBowAndDrinkTheFeathersTactic) IsTriggerPrepare() bool {
	return false
}
