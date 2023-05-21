package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//刚勇无前
//战斗中，受到兵刃伤害后，下回合行动时，提高20%会心并使下一个攻击的伤害提高65%，持续1回合
type StrongAndBraveWithoutAdvance struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StrongAndBraveWithoutAdvance) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s StrongAndBraveWithoutAdvance) Prepare() {
}

func (s StrongAndBraveWithoutAdvance) Id() consts.TacticId {
	return consts.StrongAndBraveWithoutAdvance
}

func (s StrongAndBraveWithoutAdvance) Name() string {
	return "刚勇无前"
}

func (s StrongAndBraveWithoutAdvance) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (s StrongAndBraveWithoutAdvance) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s StrongAndBraveWithoutAdvance) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s StrongAndBraveWithoutAdvance) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (s StrongAndBraveWithoutAdvance) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s StrongAndBraveWithoutAdvance) Execute() {
}

func (s StrongAndBraveWithoutAdvance) IsTriggerPrepare() bool {
	return false
}
