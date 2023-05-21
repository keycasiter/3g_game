package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//决水溃城
//准备1回合，对敌军群体（2～3人）造成破坏（禁用装备）状态及水攻状态，每回合持续造成伤害（伤害率112%，受智力影响）
//持续2回合，若该战法首回合发动则无需准备
type BreakingThroughTheWaterAndCrushingTheCityTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.45
	return b
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Prepare() {
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Id() consts.TacticId {
	return consts.BreakingThroughTheWaterAndCrushingTheCity
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Name() string {
	return "决水溃城"
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) Execute() {
}

func (b BreakingThroughTheWaterAndCrushingTheCityTactic) IsTriggerPrepare() bool {
	return false
}
