package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：鹰视狼顾
// 战法描述：战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)；
// 第5回合起，每回合有80%概率对1-2个敌军单体造成谋略伤害(伤害154%，受智力影响)；
// 自身为主将时，获得16%奇谋几率
type ClearEyedAndMaliciousTactic struct {
	tacticsParams model.TacticsParams
}

func (c ClearEyedAndMaliciousTactic) Prepare() {
	currentGeneral := c.tacticsParams.CurrentGeneral

	//自身为主将时，获得16%奇谋几率
	if currentGeneral.IsMaster {
		currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_EnhanceStrategy] += 0.16
	}

	//战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)
	for i := int(consts.Battle_Round_First); i <= int(consts.Battle_Round_Fourth); i++ {
		//80%概率
		if !util.GenerateRate(0.8) {
			continue
		}
		//攻心或奇谋
		chosenIdx := util.GenerateHitOneIdx(2)
		buffs := []consts.BuffEffectType{
			consts.BuffEffectType_EnhanceStrategy,
			consts.BuffEffectType_EnhanceWeapon,
		}
		buffEffect := buffs[chosenIdx]
		//最多累计拦截
		cnt := currentGeneral.BuffEffectAccumulateHolderMap[buffEffect]
		//一种效果最多叠加2次
		if cnt < 2 {
			//获得7%攻心或奇谋几率
			util.BuffEffectWrapSet(currentGeneral.BuffEffectTriggerMap, buffEffect, consts.BattleRound(i), 0.07)
			currentGeneral.BuffEffectAccumulateHolderMap[buffEffect]++
		}
	}
}

func (c ClearEyedAndMaliciousTactic) Init(tacticsParams model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	return c
}

func (c ClearEyedAndMaliciousTactic) Id() int64 {
	return ClearEyedAndMalicious
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

func (c ClearEyedAndMaliciousTactic) Trigger() {
	return
}

func (c ClearEyedAndMaliciousTactic) Execute() {
	return
}

func (c ClearEyedAndMaliciousTactic) Name() string {
	return "鹰视狼顾"
}
