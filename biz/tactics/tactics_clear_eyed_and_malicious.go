package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：鹰视狼顾
// 战法描述：战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)；
// 第5回合起，每回合有80%概率对1-2个敌军单体造成谋略伤害(伤害154%，受智力影响)；
// 自身为主将时，获得16%奇谋几率
type ClearEyedAndMalicious struct {
	generalId    string
	isMaster     bool
	currentRound consts.BattleRound
}

func (c ClearEyedAndMalicious) Init(generalId string, isMaster bool, currentRound consts.BattleRound) {
	c.generalId = generalId
	c.isMaster = isMaster
	c.currentRound = currentRound
}

func (c ClearEyedAndMalicious) Id() int64 {
	return 1
}

func (c ClearEyedAndMalicious) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (c ClearEyedAndMalicious) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (c ClearEyedAndMalicious) TacticsLevel() consts.TacticsLevel {
	return consts.TacticsLevel_S
}

func (c ClearEyedAndMalicious) TacticsTarget() consts.TacticsTarget {
	return consts.TacticsTarget_Team_Group
}

func (c ClearEyedAndMalicious) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c ClearEyedAndMalicious) TriggerRate() float64 {
	return 100.00
}

func (c ClearEyedAndMalicious) DamageType() consts.DamageType {
	return consts.DamageType_Strategy
}

func (c ClearEyedAndMalicious) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c ClearEyedAndMalicious) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c ClearEyedAndMalicious) IncrDamageNum() int64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) IncrDamageRate() float64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) DecrDamageNum() int64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) DecrDamageRate() float64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) ResumeMilitaryStrengthRate() float64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) EnhancedStrategyDamageRate() float64 {
	rate := float64(0)
	//自身为主将，获得16%奇谋
	if c.isMaster {
		rate += 0.16
	}
	//前4回合生效
	if c.currentRound >= consts.Battle_Round_First && c.currentRound <= consts.Battle_Round_Fourth {
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

func (c ClearEyedAndMalicious) EnhancedWeaponDamageRate() float64 {
	//前4回合生效
	if c.currentRound >= consts.Battle_Round_First && c.currentRound <= consts.Battle_Round_Fourth {
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

func (c ClearEyedAndMalicious) SuperposeNum() int64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) EvadeRate() float64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) IncrForceNum() float64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) IncrIntelligenceNum() float64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) IncrCommandNum() float64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) IncrSpeedNum() float64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) EffectNextRounds() int64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) FrozenNextRounds() int64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) DebuffEffect() consts.DebuffEffectType {
	panic("implement me")
}

func (c ClearEyedAndMalicious) BuffEffect() consts.BuffEffectType {
	panic("implement me")
}

func (c ClearEyedAndMalicious) IsGeneralAttack() bool {
	panic("implement me")
}

func (c ClearEyedAndMalicious) IsHasTrigger() bool {
	panic("implement me")
}
