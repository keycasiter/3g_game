package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 独行赴斗
// 嘲讽（强迫目标普通攻击自己）敌军全体，同时提高40%统率，持续2回合
type TravelingAloneToFightTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TravelingAloneToFightTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.5
	return t
}

func (t TravelingAloneToFightTactic) Prepare() {
}

func (t TravelingAloneToFightTactic) Id() consts.TacticId {
	return consts.TravelingAloneToFight
}

func (t TravelingAloneToFightTactic) Name() string {
	return "独行赴斗"
}

func (t TravelingAloneToFightTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TravelingAloneToFightTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TravelingAloneToFightTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TravelingAloneToFightTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TravelingAloneToFightTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TravelingAloneToFightTactic) Execute() {

}

func (t TravelingAloneToFightTactic) IsTriggerPrepare() bool {
	return false
}
