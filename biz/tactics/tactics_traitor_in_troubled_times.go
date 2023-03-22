package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：乱世奸雄
// 战法描述：战斗中，使友军群体(2人)造成伤害提高16%（受智力影响），
// 自己受到伤害降低18%（受智力影响），如果自己为主将，副将造成伤害时，会为主将恢复其伤害量10%的兵力
type TraitorInTroubledTimes struct {
	tacticsParams model.TacticsParams
}

func (t TraitorInTroubledTimes) Init(tacticsParams model.TacticsParams) {
	t.tacticsParams = tacticsParams
}

func (t TraitorInTroubledTimes) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) DamageRange() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) EvadeRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimes) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}
