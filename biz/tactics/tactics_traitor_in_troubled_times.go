package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：乱世奸雄
// 战法描述：战斗中，使友军群体(2人)造成伤害提高16%（受智力影响），
// 自己受到伤害降低18%（受智力影响），如果自己为主将，副将造成伤害时，会为主将恢复其伤害量10%的兵力
type TraitorInTroubledTimesTactic struct {
	tacticsParams model.TacticsParams
}

func (t TraitorInTroubledTimesTactic) Init(tacticsParams model.TacticsParams) {
	t.tacticsParams = tacticsParams
}

func (t TraitorInTroubledTimesTactic) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) DamageRange() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) EvadeRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (t TraitorInTroubledTimesTactic) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}
