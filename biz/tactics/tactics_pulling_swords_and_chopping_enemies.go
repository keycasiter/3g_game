package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 掣刀斫敌
// 使敌军单体受到兵刃伤害提高15%，然后对其造成兵刃伤害（伤害率208%）及震慑状态，持续1回合
type PullingSwordsAndChoppingEnemiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PullingSwordsAndChoppingEnemiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 0.35
	return p
}

func (p PullingSwordsAndChoppingEnemiesTactic) Prepare() {

}

func (p PullingSwordsAndChoppingEnemiesTactic) Id() consts.TacticId {
	return consts.PullingSwordsAndChoppingEnemies
}

func (p PullingSwordsAndChoppingEnemiesTactic) Name() string {
	return "掣刀斫敌"
}

func (p PullingSwordsAndChoppingEnemiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (p PullingSwordsAndChoppingEnemiesTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PullingSwordsAndChoppingEnemiesTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PullingSwordsAndChoppingEnemiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (p PullingSwordsAndChoppingEnemiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PullingSwordsAndChoppingEnemiesTactic) Execute() {

}

func (p PullingSwordsAndChoppingEnemiesTactic) IsTriggerPrepare() bool {
	return false
}
