package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 象兵
// 将骑兵进阶为横行无忌的象兵：
// 部队基础攻城值提高25%但行军速度降低50%，将受到伤害的25%延后于3回合内逐步结算
// 且结算伤害降低10%，自身灼烧时，获得50%群攻及混乱；
// 若蛮族统领
type ElephantSoldierTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (e ElephantSoldierTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
	e.triggerRate = 1.0
	return e
}

func (e ElephantSoldierTactic) Prepare() {
}

func (e ElephantSoldierTactic) Id() consts.TacticId {
	return consts.ElephantSoldier
}

func (e ElephantSoldierTactic) Name() string {
	return "象兵"
}

func (e ElephantSoldierTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (e ElephantSoldierTactic) GetTriggerRate() float64 {
	return e.triggerRate
}

func (e ElephantSoldierTactic) SetTriggerRate(rate float64) {
	e.triggerRate = rate
}

func (e ElephantSoldierTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (e ElephantSoldierTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
	}
}

func (e ElephantSoldierTactic) Execute() {
}

func (e ElephantSoldierTactic) IsTriggerPrepare() bool {
	return false
}

func (a ElephantSoldierTactic) SetTriggerPrepare(triggerPrepare bool) {
}
