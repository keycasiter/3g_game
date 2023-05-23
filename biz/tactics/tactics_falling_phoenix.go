package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 落凤
// 对随机敌军单体造成兵刃攻击（伤害率250%），并计穷（无法发动主动战法）1回合
type FallingPhoenixTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FallingPhoenixTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.35
	return f
}

func (f FallingPhoenixTactic) Prepare() {
}

func (f FallingPhoenixTactic) Id() consts.TacticId {
	return consts.FallingPhoenix
}

func (f FallingPhoenixTactic) Name() string {
	return "落凤"
}

func (f FallingPhoenixTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FallingPhoenixTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FallingPhoenixTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FallingPhoenixTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FallingPhoenixTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FallingPhoenixTactic) Execute() {
}

func (f FallingPhoenixTactic) IsTriggerPrepare() bool {
	return false
}
