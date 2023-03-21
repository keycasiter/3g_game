package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：士别三日
// 战法描述：战斗前3回合，无法进行普通攻击但获得30%概率规避效果
// 第4回合提高自己68点智力，并对敌军全体造成谋略伤害(伤害率180%，受智力影响)
type ThreeDaysOfSeparation struct {
	tacticsParams model.TacticsParams
}

func (t ThreeDaysOfSeparation) Init(tacticsParams model.TacticsParams) {
	t.tacticsParams = tacticsParams
}

func (t ThreeDaysOfSeparation) Id() int64 {
	return consts.ThreeDaysOfSeparation
}

func (t ThreeDaysOfSeparation) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t ThreeDaysOfSeparation) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t ThreeDaysOfSeparation) TacticsLevel() consts.TacticsLevel {
	return consts.TacticsLevel_S
}

func (t ThreeDaysOfSeparation) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThreeDaysOfSeparation) TriggerRate() float64 {
	return 100.00
}

func (t ThreeDaysOfSeparation) DamageType() consts.DamageType {
	return consts.DamageType_Strategy
}

func (t ThreeDaysOfSeparation) DamageRate() float64 {
	//第四回合，对敌军全体造成谋略伤害，伤害率180%
	if t.tacticsParams.CurrentRound == consts.Battle_Round_Fourth {
		return 1.80
	}
	return 0
}

func (t ThreeDaysOfSeparation) DamageNum() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) DamageRange() int64 {
	//第四回合，对敌军全体造成谋略伤害
	if t.tacticsParams.CurrentRound == consts.Battle_Round_Fourth {
		return 3
	}
	return 0
}

func (t ThreeDaysOfSeparation) IsDamageLockedMaster() bool {
	return false
}

func (t ThreeDaysOfSeparation) IsDamageLockedVice() bool {
	return false
}

func (t ThreeDaysOfSeparation) IncrDamageNum() int64 {
	return 0
}

func (t ThreeDaysOfSeparation) IncrDamageRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) DecrDamageNum() int64 {
	return 0
}

func (t ThreeDaysOfSeparation) DecrDamageRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) ResumeMilitaryStrengthRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) EnhancedStrategyDamageRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) EnhancedWeaponDamageRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) SuperposeNum() int64 {
	return 0
}

func (t ThreeDaysOfSeparation) EvadeRate() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) IncrForceNum() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) IncrIntelligenceNum() float64 {
	//第四回合开始提高68点智力
	if t.tacticsParams.CurrentRound == consts.Battle_Round_Fourth {
		return 68
	}
	return 0
}

func (t ThreeDaysOfSeparation) IncrCommandNum() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) IncrSpeedNum() float64 {
	return 0
}

func (t ThreeDaysOfSeparation) EffectNextRounds() int64 {
	return 0
}

func (t ThreeDaysOfSeparation) FrozenNextRounds() int64 {
	return 0
}

func (t ThreeDaysOfSeparation) DebuffEffect() consts.DebuffEffectType {
	return 0
}

func (t ThreeDaysOfSeparation) BuffEffect() consts.BuffEffectType {
	return 0
}

func (t ThreeDaysOfSeparation) IsGeneralAttack() bool {
	//前3回合，无法进行普通攻击
	if t.tacticsParams.CurrentRound <= consts.Battle_Round_Third {
		return false
	}
	return true
}
