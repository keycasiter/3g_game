package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 克敌制胜
// 普通攻击之后，对攻击目标再次造成一次谋略伤害(伤害率180%，受智力影响)；
// 若目标处于溃逃或中毒状态，则有85%概率使目标进入虚弱（无法造成伤害）状态，持续1回合
type VanquishTheEnemyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (v VanquishTheEnemyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	v.tacticsParams = tacticsParams
	v.triggerRate = 0.4
	return v
}

func (v VanquishTheEnemyTactic) Prepare() {
}

func (v VanquishTheEnemyTactic) Id() consts.TacticId {
	return consts.VanquishTheEnemy
}

func (v VanquishTheEnemyTactic) Name() string {
	return "克敌制胜"
}

func (v VanquishTheEnemyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (v VanquishTheEnemyTactic) GetTriggerRate() float64 {
	return v.triggerRate
}

func (v VanquishTheEnemyTactic) SetTriggerRate(rate float64) {
	v.triggerRate = rate
}

func (v VanquishTheEnemyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (v VanquishTheEnemyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (v VanquishTheEnemyTactic) Execute() {

}

func (v VanquishTheEnemyTactic) IsTriggerPrepare() bool {
	return false
}
