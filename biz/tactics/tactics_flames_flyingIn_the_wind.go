package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//焰逐风飞
//对敌军单体熬成谋略攻击(伤害率226%，受智力影响)及震慑状态（无法行动）并有40%概率使其受到谋略伤害提高12%（受智力影响），
//持续1回合
type FlamesFlyingInTheWindTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FlamesFlyingInTheWindTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.35
	return f
}

func (f FlamesFlyingInTheWindTactic) Prepare() {

}

func (f FlamesFlyingInTheWindTactic) Id() consts.TacticId {
	return consts.FlamesFlyingInTheWind
}

func (f FlamesFlyingInTheWindTactic) Name() string {
	return "焰逐风飞"
}

func (f FlamesFlyingInTheWindTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FlamesFlyingInTheWindTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FlamesFlyingInTheWindTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FlamesFlyingInTheWindTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FlamesFlyingInTheWindTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FlamesFlyingInTheWindTactic) Execute() {

}

func (f FlamesFlyingInTheWindTactic) IsTriggerPrepare() bool {
	return false
}