package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：魅惑
// 战法描述：自己受到普通攻击时，有45%几率使攻击者进入混乱（攻击和战法无差别选择目标）、计穷（无法发动主动战法）、虚弱（无法造成伤害）状态的一种，
// 持续1回合，
// 自身为女性时，触发几率额外受智力影响
type Charming struct {
	tacticsParams model.TacticsParams
}

func (c Charming) Init(tacticsParams model.TacticsParams) {
	c.tacticsParams = tacticsParams
}

func (c Charming) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (c Charming) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (c Charming) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (c Charming) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (c Charming) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (c Charming) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) DamageRange() int64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (c Charming) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (c Charming) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) EvadeRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (c Charming) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (c Charming) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (c Charming) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (c Charming) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}
