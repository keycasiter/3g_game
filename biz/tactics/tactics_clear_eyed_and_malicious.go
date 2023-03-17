package tactics

import "github.com/keycasiter/3g_game/biz/consts"

// 鹰视狼顾
type ClearEyedAndMalicious struct {
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
	panic("implement me")
}

func (c ClearEyedAndMalicious) TriggerRate() float64 {
	panic("implement me")
}

func (c ClearEyedAndMalicious) CurrentRound(round consts.BattleRound) {
	panic("implement me")
}

func (c ClearEyedAndMalicious) DamageType() consts.DamageType {
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
	panic("implement me")
}

func (c ClearEyedAndMalicious) EnhancedWeaponDamageRate() float64 {
	panic("implement me")
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

func (c ClearEyedAndMalicious) IsMasterGeneral(isMaster bool) {
	panic("implement me")
}
