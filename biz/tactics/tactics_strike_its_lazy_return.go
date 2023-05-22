package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 击其惰归
// 自身在下回合行动前，若受到超过最大兵力20%的伤害，则恢复自身兵力（治疗率296%，受统率影响）并降低25%受到谋略伤害（受统率影响），持续1回合，
// 否则对敌军全体造成兵刃伤害（伤害率154%）
type StrikeItsLazyReturnTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StrikeItsLazyReturnTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.4
	return s
}

func (s StrikeItsLazyReturnTactic) Prepare() {
}

func (s StrikeItsLazyReturnTactic) Id() consts.TacticId {
	return consts.StrikeItsLazyReturn
}

func (s StrikeItsLazyReturnTactic) Name() string {
	return "击其惰归"
}

func (s StrikeItsLazyReturnTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (s StrikeItsLazyReturnTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s StrikeItsLazyReturnTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s StrikeItsLazyReturnTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s StrikeItsLazyReturnTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s StrikeItsLazyReturnTactic) Execute() {

}

func (s StrikeItsLazyReturnTactic) IsTriggerPrepare() bool {
	return false
}
