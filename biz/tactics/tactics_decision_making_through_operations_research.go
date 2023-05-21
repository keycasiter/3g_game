package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//运筹决算
//准备1回合，对敌军全体发动一次谋略攻击（伤害率176%，受智力影响）
type DecisionMakingThroughOperationsResearchTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (d DecisionMakingThroughOperationsResearchTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.45
	return d
}

func (d DecisionMakingThroughOperationsResearchTactic) Prepare() {
}

func (d DecisionMakingThroughOperationsResearchTactic) Id() consts.TacticId {
	return consts.DecisionMakingThroughOperationsResearch
}

func (d DecisionMakingThroughOperationsResearchTactic) Name() string {
	return "运筹决算"
}

func (d DecisionMakingThroughOperationsResearchTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DecisionMakingThroughOperationsResearchTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DecisionMakingThroughOperationsResearchTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DecisionMakingThroughOperationsResearchTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DecisionMakingThroughOperationsResearchTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DecisionMakingThroughOperationsResearchTactic) Execute() {

}

func (d DecisionMakingThroughOperationsResearchTactic) IsTriggerPrepare() bool {
	return d.isTriggerPrepare
}
