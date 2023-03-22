package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：镇扼防拒
// 战法描述：每回合有50%概率（受智力影响）使我军单体（优先选除自己之外的副将）援护所有友军并获得休整状态（每回合恢复一次兵力，治疗率192%，受智力影响）
// 持续1回合，同时使其在1回合内受到普通攻击时，有55%概率（受智力影响）移除攻击者的增益状态
type SuppressChokesAndPreventRefusals struct {
	tacticsParams model.TacticsParams
}

func (s SuppressChokesAndPreventRefusals) Init(tacticsParams model.TacticsParams) {
	s.tacticsParams = tacticsParams
}

func (s SuppressChokesAndPreventRefusals) Id() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) TacticsLevel() consts.TacticsLevel {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) TriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) DamageType() consts.DamageType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) DamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) DamageNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) DamageRange() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) IsDamageLockedMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) IsDamageLockedVice() bool {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) IncrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) IncrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) DecrDamageNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) DecrDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) ResumeMilitaryStrengthRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) EnhancedStrategyDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) EnhancedWeaponDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) SuperposeNum() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) EvadeRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) IncrForceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) IncrIntelligenceNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) IncrCommandNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) IncrSpeedNum() float64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) EffectNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) FrozenNextRounds() int64 {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) DebuffEffect() consts.DebuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) BuffEffect() consts.BuffEffectType {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) IsGeneralAttack() bool {
	//TODO implement me
	panic("implement me")
}

func (s SuppressChokesAndPreventRefusals) EffectNextRoundDamageRate() float64 {
	//TODO implement me
	panic("implement me")
}
