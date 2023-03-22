package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：熯天炽地
// 战法描述：准备1回合，对敌军全体施放火攻（伤害率102%，受智力影响），并施加灼烧状态，
// 每回合持续造成伤害（伤害率72%，受智力影响），持续2回合。
// 主动战法 发动率35%
type TheSkyIsBlazing struct {
	tacticsParams model.TacticsParams
}

func (t TheSkyIsBlazing) Init(tacticsParams model.TacticsParams) {
	t.tacticsParams = tacticsParams
}

func (t TheSkyIsBlazing) Id() int64 {
	return consts.TheSkyIsBlazing
}

func (t TheSkyIsBlazing) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheSkyIsBlazing) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TheSkyIsBlazing) TacticsLevel() consts.TacticsLevel {
	return consts.TacticsLevel_S
}

func (t TheSkyIsBlazing) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheSkyIsBlazing) TriggerRate() float64 {
	// 35%
	return 0.35
}

func (t TheSkyIsBlazing) DamageType() consts.DamageType {
	return consts.DamageType_Strategy
}

func (t TheSkyIsBlazing) DamageRate() float64 {
	//伤害率102%
	return 1.02
}

func (t TheSkyIsBlazing) DamageNum() float64 {
	return 0
}

func (t TheSkyIsBlazing) DamageRange() int64 {
	return consts.GeneralNum_Three
}

func (t TheSkyIsBlazing) IsDamageLockedMaster() bool {
	return false
}

func (t TheSkyIsBlazing) IsDamageLockedVice() bool {
	return false
}

func (t TheSkyIsBlazing) IncrDamageNum() int64 {
	return 0
}

func (t TheSkyIsBlazing) IncrDamageRate() float64 {
	return 0
}

func (t TheSkyIsBlazing) DecrDamageNum() int64 {
	return 0
}

func (t TheSkyIsBlazing) DecrDamageRate() float64 {
	return 0
}

func (t TheSkyIsBlazing) ResumeMilitaryStrengthRate() float64 {
	return 0
}

func (t TheSkyIsBlazing) EnhancedStrategyDamageRate() float64 {
	return 0
}

func (t TheSkyIsBlazing) EnhancedWeaponDamageRate() float64 {
	return 0
}

func (t TheSkyIsBlazing) SuperposeNum() int64 {
	return 0
}

func (t TheSkyIsBlazing) EvadeRate() float64 {
	return 0
}

func (t TheSkyIsBlazing) IncrForceNum() float64 {
	return 0
}

func (t TheSkyIsBlazing) IncrIntelligenceNum() float64 {
	return 0
}

func (t TheSkyIsBlazing) IncrCommandNum() float64 {
	return 0
}

func (t TheSkyIsBlazing) IncrSpeedNum() float64 {
	return 0
}

func (t TheSkyIsBlazing) EffectNextRounds() int64 {
	return consts.Battle_Round_Second
}

func (t TheSkyIsBlazing) FrozenNextRounds() int64 {
	return 0
}

func (t TheSkyIsBlazing) DebuffEffect() consts.DebuffEffectType {
	return consts.DebuffEffectType_Firing
}

func (t TheSkyIsBlazing) BuffEffect() consts.BuffEffectType {
	return 0
}

func (t TheSkyIsBlazing) IsGeneralAttack() bool {
	return true
}

func (t TheSkyIsBlazing) EffectNextRoundDamageRate() float64 {
	//持续伤害率72%
	return 0.72
}
