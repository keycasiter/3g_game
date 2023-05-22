package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 绝其汲道
// 准备1回合，对敌军群体（2-3人）造成一次兵刃攻击（伤害率162%），使其进入禁疗状态（无法恢复兵力），持续1回合
type EliminateItAndDrawFromItTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (e EliminateItAndDrawFromItTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
	e.triggerRate = 0.5
	return e
}

func (e EliminateItAndDrawFromItTactic) Prepare() {

}

func (e EliminateItAndDrawFromItTactic) Id() consts.TacticId {
	return consts.EliminateItAndDrawFromIt
}

func (e EliminateItAndDrawFromItTactic) Name() string {
	return "绝其汲道"
}

func (e EliminateItAndDrawFromItTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (e EliminateItAndDrawFromItTactic) GetTriggerRate() float64 {
	return e.triggerRate
}

func (e EliminateItAndDrawFromItTactic) SetTriggerRate(rate float64) {
	e.triggerRate = rate
}

func (e EliminateItAndDrawFromItTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (e EliminateItAndDrawFromItTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (e EliminateItAndDrawFromItTactic) Execute() {

}

func (e EliminateItAndDrawFromItTactic) IsTriggerPrepare() bool {
	return e.isTriggerPrepare
}
