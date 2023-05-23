package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 强攻
// 使自己进入连击（每回合可以普通攻击2次）状态，持续1回合
type TakeByStormTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TakeByStormTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.45
	return t
}

func (t TakeByStormTactic) Prepare() {
}

func (t TakeByStormTactic) Id() consts.TacticId {
	return consts.TakeByStorm
}

func (t TakeByStormTactic) Name() string {
	return "强攻"
}

func (t TakeByStormTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TakeByStormTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TakeByStormTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TakeByStormTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TakeByStormTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TakeByStormTactic) Execute() {
}

func (t TakeByStormTactic) IsTriggerPrepare() bool {
	return false
}
