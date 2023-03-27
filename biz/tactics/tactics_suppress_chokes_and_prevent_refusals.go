package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：镇扼防拒
// 战法描述：每回合有50%概率（受智力影响）使我军单体（优先选除自己之外的副将）援护所有友军并获得休整状态（每回合恢复一次兵力，治疗率192%，受智力影响）
// 持续1回合，同时使其在1回合内受到普通攻击时，有55%概率（受智力影响）移除攻击者的增益状态
type SuppressChokesAndPreventRefusalsTactic struct {
	tacticsParams model.TacticsParams
}

func (s SuppressChokesAndPreventRefusalsTactic) Init(tacticsParams model.TacticsParams) {
	s.tacticsParams = tacticsParams
}

func (s SuppressChokesAndPreventRefusalsTactic) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DamageRange() consts.GeneralNum {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) EvadeRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusalsTactic) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}
