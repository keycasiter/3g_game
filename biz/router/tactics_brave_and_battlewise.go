package router

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 骁勇善战
// 准备1回合，对敌军单体造成一次兵刃攻击（伤害率300%）并提高自己40点速度，持续2回合
type BraveAndBattlewise struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (b BraveAndBattlewise) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.35
	return b
}

func (b BraveAndBattlewise) Prepare() {

}

func (b BraveAndBattlewise) Id() consts.TacticId {
	return consts.BraveAndBattlewise
}

func (b BraveAndBattlewise) Name() string {
	return "骁勇善战"
}

func (b BraveAndBattlewise) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BraveAndBattlewise) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BraveAndBattlewise) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BraveAndBattlewise) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BraveAndBattlewise) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BraveAndBattlewise) Execute() {
}

func (b BraveAndBattlewise) IsTriggerPrepare() bool {
	return b.isTriggerPrepare
}