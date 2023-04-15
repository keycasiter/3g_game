package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 战法名称：鹰视狼顾
// 战法描述：战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)；
// 第5回合起，每回合有80%概率对1-2个敌军单体造成谋略伤害(伤害154%，受智力影响)；
// 自身为主将时，获得16%奇谋几率
type ClearEyedAndMaliciousTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ClearEyedAndMaliciousTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c ClearEyedAndMaliciousTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c ClearEyedAndMaliciousTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c ClearEyedAndMaliciousTactic) Prepare() {
	currentGeneral := c.tacticsParams.CurrentGeneral
	ctx := c.tacticsParams.Ctx

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)

	//自身为主将时，获得16%奇谋几率
	if currentGeneral.IsMaster {
		currentGeneral.BuffEffectHolderMap[consts.BuffEffectType_EnhanceStrategy] += 0.16
	}

	//战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)
	//注册效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerGeneral := params.CurrentGeneral
		triggerResp := &vo.TacticsTriggerResult{}

		//80%概率
		if !util.GenerateRate(0.8) {
			return triggerResp
		}
		//攻心或奇谋
		chosenIdx := util.GenerateHitOneIdx(2)
		buffs := []consts.BuffEffectType{
			consts.BuffEffectType_EnhanceStrategy,
			consts.BuffEffectType_EnhanceWeapon,
		}
		buffEffect := buffs[chosenIdx]

		//一种效果最多叠加2次
		if !util.TacticsBuffEffectCountWrapIncr(ctx, triggerGeneral, buffEffect, 1, 2, false) {
			return triggerResp
		}
		util.BuffEffectWrapSet(triggerGeneral, buffEffect, 0.07)

		return triggerResp
	})
	//第5回合起，每回合有80%概率对1-2个敌军单体造成谋略伤害(伤害154%，受智力影响)
}

func (c ClearEyedAndMaliciousTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c ClearEyedAndMaliciousTactic) Id() consts.TacticId {
	return consts.ClearEyedAndMalicious
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
