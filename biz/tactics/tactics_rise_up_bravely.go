package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 奋突
// 普通攻击之后，使自己造成兵刃伤害提高12%，最多叠加3次，并且有35%概率使目标缴械（无法进行普通攻击），持续1回合
type RiseUpBravelyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RiseUpBravelyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 1.0
	return r
}

func (r RiseUpBravelyTactic) Prepare() {

}

func (r RiseUpBravelyTactic) Id() consts.TacticId {
	return consts.RiseUpBravely
}

func (r RiseUpBravelyTactic) Name() string {
	return "奋突"
}

func (r RiseUpBravelyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RiseUpBravelyTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RiseUpBravelyTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RiseUpBravelyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (r RiseUpBravelyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RiseUpBravelyTactic) Execute() {

}

func (r RiseUpBravelyTactic) IsTriggerPrepare() bool {
	return false
}
