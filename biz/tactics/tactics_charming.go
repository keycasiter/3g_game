package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：魅惑
// 战法描述：自己受到普通攻击时，有45%几率使攻击者进入混乱（攻击和战法无差别选择目标）、计穷（无法发动主动战法）、虚弱（无法造成伤害）状态的一种，
// 持续1回合，自身为女性时，触发几率额外受智力影响
type CharmingTactic struct {
	tacticsParams model.TacticsParams
}

func (c CharmingTactic) Init(tacticsParams model.TacticsParams) {
	c.tacticsParams = tacticsParams
}

func (c CharmingTactic) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) DamageRange() consts.GeneralNum {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) DebuffEffectRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) BuffEffectRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (c CharmingTactic) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}
