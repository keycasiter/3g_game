package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 手起刀落
// 普通攻击之后，对攻击目标再次发起一次兵刃攻击（伤害率214%）
// 突击 30%
type CutDownTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CutDownTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 0.3
	return c
}

func (c CutDownTactic) Prepare() {

}

func (c CutDownTactic) Id() consts.TacticId {
	return consts.CutDown
}

func (c CutDownTactic) Name() string {
	return "手起刀落"
}

func (c CutDownTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (c CutDownTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CutDownTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CutDownTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (c CutDownTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CutDownTactic) Execute() {
	//普通攻击之后，对攻击目标再次发起一次兵刃攻击（伤害率214%）
	sufferGeneral := c.tacticsParams.CurrentSufferGeneral
	currentGeneral := c.tacticsParams.CurrentGeneral

	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams:     c.tacticsParams,
		AttackGeneral:     currentGeneral,
		SufferGeneral:     sufferGeneral,
		DamageType:        consts.DamageType_Weapon,
		DamageImproveRate: 2.14,
		TacticId:          c.Id(),
		TacticName:        c.Name(),
	})
}

func (c CutDownTactic) IsTriggerPrepare() bool {
	return false
}

func (a CutDownTactic) SetTriggerPrepare(triggerPrepare bool) {
}
