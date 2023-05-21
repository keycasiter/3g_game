package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//百骑劫营
//普通攻击之后，对随机敌军单体发动一次兵刃攻击（伤害率162%）同时有50%几率对敌军主将额外发动一次兵刃攻击（伤害率120%）
type HundredCavalryRobberyBattalionsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HundredCavalryRobberyBattalionsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.4
	return h
}

func (h HundredCavalryRobberyBattalionsTactic) Prepare() {

}

func (h HundredCavalryRobberyBattalionsTactic) Id() consts.TacticId {
	return consts.HundredCavalryRobberyBattalions
}

func (h HundredCavalryRobberyBattalionsTactic) Name() string {
	return "百骑劫营"
}

func (h HundredCavalryRobberyBattalionsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (h HundredCavalryRobberyBattalionsTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HundredCavalryRobberyBattalionsTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HundredCavalryRobberyBattalionsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (h HundredCavalryRobberyBattalionsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
	}
}

func (h HundredCavalryRobberyBattalionsTactic) Execute() {

}

func (h HundredCavalryRobberyBattalionsTactic) IsTriggerPrepare() bool {
	return false
}
