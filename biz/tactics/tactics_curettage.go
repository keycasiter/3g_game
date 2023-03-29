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
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) DamageRange() consts.GeneralNum {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CurettageTactic) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
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
