package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 据水断桥
// 对敌军群体（2-3人）造成溃逃状态，每回合持续造成伤害（伤害率78%，受武力影响），并使其造成伤害降低8%（受双方武力之差影响），
// 同时使自身获得16%倒戈（造成兵刃伤害时，恢复自身基于伤害量的一定兵力），持续2回合，该战法发动后回进入1回合冷却
type BrokenBridgeByWaterTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BrokenBridgeByWaterTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BrokenBridgeByWaterTactic) Prepare() {

}

func (b BrokenBridgeByWaterTactic) Id() consts.TacticId {
	return consts.BrokenBridgeByWater
}

func (b BrokenBridgeByWaterTactic) Name() string {
	return "据水断桥"
}

func (b BrokenBridgeByWaterTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (b BrokenBridgeByWaterTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BrokenBridgeByWaterTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BrokenBridgeByWaterTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BrokenBridgeByWaterTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BrokenBridgeByWaterTactic) Execute() {

}

func (b BrokenBridgeByWaterTactic) IsTriggerPrepare() bool {
	return false
}
