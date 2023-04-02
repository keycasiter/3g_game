package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：刮骨疗毒
// 战法描述：为损失兵力最多的我军单体清除负面状态并为其恢复兵力（治疗率256%，受智力影响）
type CurettageTactic struct {
	tacticsParams model.TacticsParams
}

func (c CurettageTactic) Init(tacticsParams model.TacticsParams) {
	c.tacticsParams = tacticsParams
}

func (c CurettageTactic) Id() int64 {
	return Curettage
}

func (c CurettageTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (c CurettageTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (c CurettageTactic) TacticsLevel() consts.TacticsLevel {
	return consts.TacticsLevel_S
}

func (c CurettageTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CurettageTactic) TriggerRate() float64 {
	return 0.4
}

func (c CurettageTactic) DamageType() consts.DamageType {
	return consts.DamageType_None
}

func (c CurettageTactic) DamageRate() float64 {
	return 0
}

func (c CurettageTactic) DamageNum() float64 {
	return 0
}

func (c CurettageTactic) DamageRange() consts.GeneralNum {
	return 0
}

func (c CurettageTactic) IsLockingMaster() bool {
	return false
}

func (c CurettageTactic) IsLockingVice() bool {
	return false
}

func (c CurettageTactic) IncrDamageNum() int64 {
	return 0
}

func (c CurettageTactic) IncrDamageRate() float64 {
	return 0
}

func (c CurettageTactic) DecrDamageNum() int64 {
	return 0
}

func (c CurettageTactic) DecrDamageRate() float64 {
	return 0
}

func (c CurettageTactic) ResumeMilitaryStrengthRate() float64 {
	return 0
}

func (c CurettageTactic) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) DebuffEffectRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) BuffEffectRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) LockingGeneral() int64 {
	return 0
}
