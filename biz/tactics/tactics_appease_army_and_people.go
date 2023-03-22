package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：抚辑军民
// 战法描述：战斗前3回合，使我军群体(2人)造成的伤害降低24%，
// 受到的伤害降低24%（受统率影响），
// 战斗第4回合时，恢复其兵力（治疗率126%，受智力影响）
type AppeaseArmyAndPeople struct {
	tacticsParams model.TacticsParams
}

func (a AppeaseArmyAndPeople) Init(tacticsParams model.TacticsParams) {
	a.tacticsParams = tacticsParams
}

func (a AppeaseArmyAndPeople) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) DamageRange() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) EvadeRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeople) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}
