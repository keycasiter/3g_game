package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 避实击虚
// 对统率最低的敌军单体发动一次兵刃攻击(伤害率185%)
type AvoidTheSolidAndStrikeTheWeakTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.5
	return a
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) Prepare() {
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) Id() consts.TacticId {
	return consts.AvoidTheSolidAndStrikeTheWeak
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) Name() string {
	return "避实击虚"
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) Execute() {
	currentGeneral := a.tacticsParams.CurrentGeneral

	//找到敌军统率最低的武将
	sufferGeneral := util.GetEnemyGeneralWhoLowestCommand(currentGeneral, a.tacticsParams)
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 1.85)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: a.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: sufferGeneral,
		Damage:        dmg,
		DamageType:    consts.DamageType_Weapon,
		TacticName:    a.Name(),
		TacticId:      a.Id(),
	})
}

func (a AvoidTheSolidAndStrikeTheWeakTactic) IsTriggerPrepare() bool {
	return false
}
