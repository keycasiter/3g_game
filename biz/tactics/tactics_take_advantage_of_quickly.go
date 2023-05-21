package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//速乘其利
//普通攻击之后，对目标发动一次兵刃攻击（伤害率185%）并计穷（无法发动主动战法）1回合
type TakeAdvantageOfQuickly struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TakeAdvantageOfQuickly) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.35
	return t
}

func (t TakeAdvantageOfQuickly) Prepare() {
}

func (t TakeAdvantageOfQuickly) Id() consts.TacticId {
	return consts.TakeAdvantageOfQuickly
}

func (t TakeAdvantageOfQuickly) Name() string {
	return "速乘其利"
}

func (t TakeAdvantageOfQuickly) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TakeAdvantageOfQuickly) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TakeAdvantageOfQuickly) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TakeAdvantageOfQuickly) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (t TakeAdvantageOfQuickly) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TakeAdvantageOfQuickly) Execute() {

}

func (t TakeAdvantageOfQuickly) IsTriggerPrepare() bool {
	return false
}
