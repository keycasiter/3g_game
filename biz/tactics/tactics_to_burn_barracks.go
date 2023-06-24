package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 焚辎营垒
// 对敌军群体（2人）造成谋略伤害（伤害率146%，受智力影响）并使其进入禁疗状态，持续1回合
type ToBurnBarracks struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ToBurnBarracks) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.4
	return t
}

func (t ToBurnBarracks) Prepare() {
}

func (t ToBurnBarracks) Id() consts.TacticId {
	return consts.ToBurnBarracks
}

func (t ToBurnBarracks) Name() string {
	return "焚辎营垒"
}

func (t ToBurnBarracks) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t ToBurnBarracks) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ToBurnBarracks) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ToBurnBarracks) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t ToBurnBarracks) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ToBurnBarracks) Execute() {
}

func (t ToBurnBarracks) IsTriggerPrepare() bool {
	return false
}
