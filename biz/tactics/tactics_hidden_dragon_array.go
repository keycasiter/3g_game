package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 潜龙阵
// 我军三名武将阵营均不相同时，主将提升15%武力、智力、速度、统率，造成伤害降低15%
// 副将造成伤害提高15%,受到伤害降低15%，且可触发战法的主将效果
// 若我军主将的任意战法拥有主将效果，使其失去该效果且属性提升值降低为5%
type HiddenDragonArrayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HiddenDragonArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 1.0
	return h
}

func (h HiddenDragonArrayTactic) Prepare() {
}

func (h HiddenDragonArrayTactic) Id() consts.TacticId {
	return consts.HiddenDragonArray
}

func (h HiddenDragonArrayTactic) Name() string {
	return "潜龙阵"
}

func (h HiddenDragonArrayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (h HiddenDragonArrayTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HiddenDragonArrayTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HiddenDragonArrayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (h HiddenDragonArrayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HiddenDragonArrayTactic) Execute() {
}

func (h HiddenDragonArrayTactic) IsTriggerPrepare() bool {
	return false
}