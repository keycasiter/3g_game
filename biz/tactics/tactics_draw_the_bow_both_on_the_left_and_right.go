package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 左右开弓
// 提高自身13%会心几率（触发时兵刃伤害提高100%），对敌军单体造成一次兵刃攻击（伤害率180%）
// 如果目标为骑兵则额外造成溃散状态，每回合持续造成伤害（伤害率90%，受武力影响），持续2回合
type DrawTheBowBothOnTheLeftAndRight struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DrawTheBowBothOnTheLeftAndRight) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.35
	return d
}

func (d DrawTheBowBothOnTheLeftAndRight) Prepare() {
}

func (d DrawTheBowBothOnTheLeftAndRight) Id() consts.TacticId {
	return consts.DrawTheBowBothOnTheLeftAndRight
}

func (d DrawTheBowBothOnTheLeftAndRight) Name() string {
	return "左右开弓"
}

func (d DrawTheBowBothOnTheLeftAndRight) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DrawTheBowBothOnTheLeftAndRight) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DrawTheBowBothOnTheLeftAndRight) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DrawTheBowBothOnTheLeftAndRight) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DrawTheBowBothOnTheLeftAndRight) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (d DrawTheBowBothOnTheLeftAndRight) Execute() {
}

func (d DrawTheBowBothOnTheLeftAndRight) IsTriggerPrepare() bool {
	return false
}
