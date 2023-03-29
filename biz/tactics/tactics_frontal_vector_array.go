package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：锋矢阵
// 战法描述：战斗中，使我军主将造成的伤害提升30%，受到的伤害提升20%，我军副将造成的伤害降低15%，受到的伤害降低25%
type FrontalVectorArrayTactic struct {
	tacticsParams model.TacticsParams
}

func (f FrontalVectorArrayTactic) Init(tacticsParams model.TacticsParams) {
	f.tacticsParams = tacticsParams
}

func (f FrontalVectorArrayTactic) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) DamageRange() consts.GeneralNum {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) DebuffEffectRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) BuffEffectRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (f FrontalVectorArrayTactic) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}
