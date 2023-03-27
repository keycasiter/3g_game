package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：鹰视狼顾
// 战法描述：战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)；
// 第5回合起，每回合有80%概率对1-2个敌军单体造成谋略伤害(伤害154%，受智力影响)；
// 自身为主将时，获得16%奇谋几率
type ClearEyedAndMaliciousTactic struct {
	tacticsParams model.TacticsParams
}

func (c ClearEyedAndMaliciousTactic) Init(tacticsParams model.TacticsParams) {
	c.tacticsParams = tacticsParams
}

func (c ClearEyedAndMaliciousTactic) Id() int64 {
	return ClearEyedAndMalicious
}

func (c ClearEyedAndMaliciousTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c ClearEyedAndMaliciousTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (c ClearEyedAndMaliciousTactic) TacticsLevel() consts.TacticsLevel {
	return consts.TacticsLevel_S
}

func (c ClearEyedAndMaliciousTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c ClearEyedAndMaliciousTactic) TriggerRate() float64 {
	return 100.00
}

func (c ClearEyedAndMaliciousTactic) DamageType() consts.DamageType {
	return consts.DamageType_Strategy
}

func (c ClearEyedAndMaliciousTactic) DamageRate() float64 {
	return 1.54
}

func (c ClearEyedAndMaliciousTactic) DamageRange() consts.GeneralNum {
	//第五回合起，80%概率
	if c.tacticsParams.CurrentRound >= consts.Battle_Round_Fifth && util.GenerateRate(0.8) {
		// TODO 描述黑盒的，暂时设置为50%概率，伤害1～2个目标
		if util.GenerateRate(0.5) {
			return consts.GeneralNum_One
		}
		return consts.GeneralNum_Two
	}
	return consts.GeneralNum_Unknow
}

func (c ClearEyedAndMaliciousTactic) IsDamageLockedMaster() bool {
	return false
}

func (c ClearEyedAndMaliciousTactic) IsDamageLockedVice() bool {
	return false
}

func (c ClearEyedAndMaliciousTactic) DamageNum() float64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) IncrDamageNum() int64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) IncrDamageRate() float64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) DecrDamageNum() int64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) DecrDamageRate() float64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) ResumeMilitaryStrengthRate() float64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) EnhancedStrategyDamageRate() float64 {
	rate := float64(0)
	//自身为主将，获得16%奇谋
	if c.tacticsParams.IsMaster {
		rate += 0.16
	}
	//前4回合生效
	if c.tacticsParams.CurrentRound >= consts.Battle_Round_First && c.tacticsParams.CurrentRound <= consts.Battle_Round_Fourth {
		// 触发概率80%
		if util.GenerateRate(0.8) {
			//奇谋率7%
			rate += 0.07
		}
		return rate
	} else {
		return rate
	}
}

func (c ClearEyedAndMaliciousTactic) EnhancedWeaponDamageRate() float64 {
	//前4回合生效
	if c.tacticsParams.CurrentRound >= consts.Battle_Round_First && c.tacticsParams.CurrentRound <= consts.Battle_Round_Fourth {
		// 触发概率80%
		if util.GenerateRate(0.8) {
			//会心率7%
			return 0.07
		}
		return 0
	} else {
		return 0
	}
}

func (c ClearEyedAndMaliciousTactic) SuperposeNum() int64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) EvadeRate() float64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) IncrForceNum() float64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) IncrIntelligenceNum() float64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) IncrCommandNum() float64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) IncrSpeedNum() float64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) EffectNextRounds() int64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) FrozenNextRounds() int64 {
	return 0
}

func (c ClearEyedAndMaliciousTactic) DebuffEffect() consts.DebuffEffectType {
	return 0
}

func (c ClearEyedAndMaliciousTactic) BuffEffect() consts.BuffEffectType {
	return 0
}

func (c ClearEyedAndMaliciousTactic) IsGeneralAttack() bool {
	return true
}

func (c ClearEyedAndMaliciousTactic) EffectNextRoundDamageRate() float64 {
	return 0
}
