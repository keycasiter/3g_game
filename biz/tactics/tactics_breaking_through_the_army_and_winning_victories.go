package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//破军威胜
//降低敌军单体70点统率（受武力影响），持续2回合，并对其造成兵刃伤害（伤害率228%）
type BreakingThroughTheArmyAndWinningVictoriesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Prepare() {
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Id() consts.TacticId {
	return consts.BreakingThroughTheArmyAndWinningVictories
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Name() string {
	return "破军威胜"
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) Execute() {

}

func (b BreakingThroughTheArmyAndWinningVictoriesTactic) IsTriggerPrepare() bool {
	return false
}
