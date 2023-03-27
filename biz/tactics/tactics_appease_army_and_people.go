package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：抚辑军民
// 战法描述：战斗前3回合，使我军群体(2人)造成的伤害降低24%，
// 受到的伤害降低24%（受统率影响），
// 战斗第4回合时，恢复其兵力（治疗率126%，受智力影响）
type AppeaseArmyAndPeopleTactic struct {
	tacticsParams model.TacticsParams
}

func (a AppeaseArmyAndPeopleTactic) Init(tacticsParams model.TacticsParams) {
	a.tacticsParams = tacticsParams
}

func (a AppeaseArmyAndPeopleTactic) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) DamageRange() consts.GeneralNum {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) EvadeRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (a AppeaseArmyAndPeopleTactic) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}
