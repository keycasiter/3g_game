package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/holder"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：熯天炽地
// 战法描述：准备1回合，对敌军全体施放火攻（伤害率102%，受智力影响），并施加灼烧状态，
// 每回合持续造成伤害（伤害率72%，受智力影响），持续2回合。
// 主动战法 发动率35%
type TheSkyIsBlazingTactic struct {
	tacticsParams model.TacticsParams
}

func (t TheSkyIsBlazingTactic) Init(tacticsParams model.TacticsParams) {
	t.tacticsParams = tacticsParams
}

func (t TheSkyIsBlazingTactic) Id() int64 {
	return holder.TheSkyIsBlazing
}

func (t TheSkyIsBlazingTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheSkyIsBlazingTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TheSkyIsBlazingTactic) TacticsLevel() consts.TacticsLevel {
	return consts.TacticsLevel_S
}

func (t TheSkyIsBlazingTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheSkyIsBlazingTactic) TriggerRate() float64 {
	// 35%
	return 0.35
}

func (t TheSkyIsBlazingTactic) DamageType() consts.DamageType {
	return consts.DamageType_Strategy
}

func (t TheSkyIsBlazingTactic) DamageRate() float64 {
	//伤害率102%
	return 1.02
}

func (t TheSkyIsBlazingTactic) DamageNum() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) DamageRange() consts.GeneralNum {
	return consts.GeneralNum_Three
}

func (t TheSkyIsBlazingTactic) IsDamageLockedMaster() bool {
	return false
}

func (t TheSkyIsBlazingTactic) IsDamageLockedVice() bool {
	return false
}

func (t TheSkyIsBlazingTactic) IncrDamageNum() int64 {
	return 0
}

func (t TheSkyIsBlazingTactic) IncrDamageRate() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) DecrDamageNum() int64 {
	return 0
}

func (t TheSkyIsBlazingTactic) DecrDamageRate() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) ResumeMilitaryStrengthRate() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) EnhancedStrategyDamageRate() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) EnhancedWeaponDamageRate() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) SuperposeNum() int64 {
	return 0
}

func (t TheSkyIsBlazingTactic) EvadeRate() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) IncrForceNum() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) IncrIntelligenceNum() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) IncrCommandNum() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) IncrSpeedNum() float64 {
	return 0
}

func (t TheSkyIsBlazingTactic) EffectNextRounds() int64 {
	return 2
}

func (t TheSkyIsBlazingTactic) FrozenNextRounds() int64 {
	return 0
}

func (t TheSkyIsBlazingTactic) DebuffEffect() consts.DebuffEffectType {
	return consts.DebuffEffectType_Firing
}

func (t TheSkyIsBlazingTactic) BuffEffect() consts.BuffEffectType {
	return 0
}

func (t TheSkyIsBlazingTactic) IsGeneralAttack() bool {
	return true
}

func (t TheSkyIsBlazingTactic) EffectNextRoundDamageRate() float64 {
	//持续伤害率72%
	return 0.72
}
