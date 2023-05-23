package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 暗藏玄机
// 普通攻击之后，对攻击目标在此发起一次兵刃攻击（伤害率144%），如果目标为敌军主将则额外造成一次谋略攻击（伤害率92%，受智力影响）
type HiddenMysteryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HiddenMysteryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.4
	return h
}

func (h HiddenMysteryTactic) Prepare() {
}

func (h HiddenMysteryTactic) Id() consts.TacticId {
	return consts.HiddenMystery
}

func (h HiddenMysteryTactic) Name() string {
	return "暗藏玄机"
}

func (h HiddenMysteryTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (h HiddenMysteryTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HiddenMysteryTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HiddenMysteryTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (h HiddenMysteryTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HiddenMysteryTactic) Execute() {
}

func (h HiddenMysteryTactic) IsTriggerPrepare() bool {
	return false
}
