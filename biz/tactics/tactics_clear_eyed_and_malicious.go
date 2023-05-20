package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 战法名称：鹰视狼顾
// 战法描述：战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)；
// 第5回合起，每回合有80%概率对1-2个敌军单体造成谋略伤害(伤害154%，受智力影响)；
// 自身为主将时，获得16%奇谋几率
type ClearEyedAndMaliciousTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ClearEyedAndMaliciousTactic) IsTriggerPrepare() bool {
	return false
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

	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.BuffEffectType_ClearEyedAndMalicious_Prepare,
	)
	//战斗前4回合，每回合有80%概率使自身获得7%攻心或奇谋几率(每种效果最多叠加2次)
	//注册效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerGeneral := params.CurrentGeneral
		currentRound := params.CurrentRound
		triggerResp := &vo.TacticsTriggerResult{}

		//前四回合
		if currentRound >= consts.Battle_Round_First && currentRound <= consts.Battle_Round_Fourth {
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
			util.BuffEffectWrapSet(ctx, triggerGeneral, buffEffect, &vo.EffectHolderParams{
				EffectTimes: 2,
			})
		}

		return triggerResp
	})

	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.BuffEffectType_ClearEyedAndMalicious_ClearEyed_Prepare,
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		currentRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral
		triggerResp := &vo.TacticsTriggerResult{}

		//第5回合起，每回合有80%概率对1-2个敌军单体造成谋略伤害(伤害154%，受智力影响)；
		if currentRound == consts.Battle_Round_Fifth {
			if !util.GenerateRate(0.8) {
				hlog.CtxInfof(ctx, "[%s]执行来自[%s]【%s】的「%v」效果因几率没有生效",
					triggerGeneral.BaseInfo.Name,
					triggerGeneral.BaseInfo.Name,
					c.Name(),
					consts.BuffEffectType_ClearEyedAndMalicious_ClearEyed_Prepare,
				)
				return triggerResp
			}

			hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
				triggerGeneral.BaseInfo.Name,
				c.Name(),
				consts.BuffEffectType_ClearEyedAndMalicious_ClearEyed_Prepare,
			)

			//找到1～2敌人
			enemyMap := util.GetEnemyGeneralsOneOrTwoMap(c.tacticsParams)
			for _, enemyGeneral := range enemyMap {
				dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.54)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: c.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: enemyGeneral,
					Damage:        dmg,
					TacticName:    c.Name(),
				})
			}
		}
		return triggerResp
	})

	//自身为主将时，获得16%奇谋几率
	if currentGeneral.IsMaster {
		hlog.CtxInfof(ctx, "[%s]的奇谋提高了16%%",
			currentGeneral.BaseInfo.Name,
		)
		util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_EnhanceStrategy, &vo.EffectHolderParams{
			EffectRate: 0.16,
			FromTactic: c.Id(),
		})
	}
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
