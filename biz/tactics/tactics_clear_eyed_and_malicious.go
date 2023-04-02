package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 战法名称：鹰视狼顾
// 战法描述：战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)；
// 第5回合起，每回合有80%概率对1-2个敌军单体造成谋略伤害(伤害154%，受智力影响)；
// 自身为主将时，获得16%奇谋几率
type ClearEyedAndMaliciousTactic struct {
	tacticsParams model.TacticsParams
}

func (c ClearEyedAndMaliciousTactic) Name() string {
	return "鹰视狼顾"
}

func (c ClearEyedAndMaliciousTactic) BuffEffect() map[int64]map[consts.BuffEffectType]map[consts.BattleRound]float64 {
	return nil
}

func (c ClearEyedAndMaliciousTactic) DebuffEffect() map[int64]map[consts.DebuffEffectType]map[consts.BattleRound]float64 {
	panic("implement me")
}

func (c ClearEyedAndMaliciousTactic) Damage() map[int64]map[consts.BattleRound]map[consts.DamageType]float64 {
	panic("implement me")
}

func (c ClearEyedAndMaliciousTactic) Resume() map[int64]map[consts.BattleRound]float64 {
	panic("implement me")
}

func (c ClearEyedAndMaliciousTactic) GetCurrentRound() consts.BattleRound {
	panic("implement me")
}

func (c ClearEyedAndMaliciousTactic) LastTriggerRound() consts.BattleRound {
	panic("implement me")
}

func (c ClearEyedAndMaliciousTactic) Init(tacticsParams model.TacticsParams) {
	c.tacticsParams = tacticsParams
}

func (c ClearEyedAndMaliciousTactic) Id() int64 {
	return ClearEyedAndMalicious
}

func (c ClearEyedAndMaliciousTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c ClearEyedAndMaliciousTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (c ClearEyedAndMaliciousTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c ClearEyedAndMaliciousTactic) TriggerRate() float64 {
	return 1.0
}
