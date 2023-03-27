package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/holder"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：士别三日
// 战法描述：战斗前3回合，无法进行普通攻击但获得30%概率规避效果
// 第4回合提高自己68点智力，并对敌军全体造成谋略伤害(伤害率180%，受智力影响)
type ThreeDaysOfSeparationTactic struct {
	tacticsParams model.TacticsParams
}

func (t ThreeDaysOfSeparationTactic) Init(tacticsParams model.TacticsParams) {
	t.tacticsParams = tacticsParams
}

func (t ThreeDaysOfSeparationTactic) Id() int64 {
	return holder.ThreeDaysOfSeparation
}

func (t ThreeDaysOfSeparationTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t ThreeDaysOfSeparationTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t ThreeDaysOfSeparationTactic) TacticsLevel() consts.TacticsLevel {
	return consts.TacticsLevel_S
}

func (t ThreeDaysOfSeparationTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThreeDaysOfSeparationTactic) TriggerRate() float64 {
	return 100.00
}

func (t ThreeDaysOfSeparationTactic) DamageType() consts.DamageType {
	return consts.DamageType_Strategy
}

func (t ThreeDaysOfSeparationTactic) DamageRate() float64 {
	//第四回合，对敌军全体造成谋略伤害，伤害率180%
	if t.tacticsParams.CurrentRound == consts.Battle_Round_Fourth {
		return 1.80
	}
	return 0
}

func (t ThreeDaysOfSeparationTactic) DamageRange() consts.GeneralNum {
	//TODO implement me
	panic("implement me")
}

func (t ThreeDaysOfSeparationTactic) DamageNum() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) IsDamageLockedMaster() bool {
	return false
}

func (t ThreeDaysOfSeparationTactic) IsDamageLockedVice() bool {
	return false
}

func (t ThreeDaysOfSeparationTactic) IncrDamageNum() int64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) IncrDamageRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) DecrDamageNum() int64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) DecrDamageRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) ResumeMilitaryStrengthRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) EnhancedStrategyDamageRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) EnhancedWeaponDamageRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) SuperposeNum() int64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) EvadeRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) IncrForceNum() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) IncrIntelligenceNum() float64 {
	//第四回合开始提高68点智力
	if t.tacticsParams.CurrentRound == consts.Battle_Round_Fourth {
		return 68
	}
	return 0
}

func (t ThreeDaysOfSeparationTactic) IncrCommandNum() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) IncrSpeedNum() float64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) EffectNextRounds() int64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) FrozenNextRounds() int64 {
	return 0
}

func (t ThreeDaysOfSeparationTactic) DebuffEffect() consts.DebuffEffectType {
	return 0
}

func (t ThreeDaysOfSeparationTactic) BuffEffect() consts.BuffEffectType {
	return consts.BuffEffectType_Evade
}

func (t ThreeDaysOfSeparationTactic) IsGeneralAttack() bool {
	//前3回合，无法进行普通攻击
	if t.tacticsParams.CurrentRound <= consts.Battle_Round_Third {
		return false
	}
	return true
}

func (c ThreeDaysOfSeparationTactic) EffectNextRoundDamageRate() float64 {
	return 0
}
