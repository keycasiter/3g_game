package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 仁德载世
// 每回合治疗我军群体（2人，治疗率68%，受智力影响），并使其受到伤害降低4%，受智力影响，持续1回合
// 且有10%概率对敌军单体施加虚弱状态，持续1回合
// 自身为主将时，施加虚弱状态的概率提高至25%
type BenevolentAndVirtuousThroughoutTheWorldTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerGeneral := params.CurrentGeneral
		triggerResp := &vo.TacticsTriggerResult{}

		//找到我军2人
		pairGenerals := util.GetPairGeneralsTwoArrByGeneral(triggerGeneral, b.tacticsParams)
		for _, general := range pairGenerals {
			resume := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.68)
			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  b.tacticsParams,
				ProduceGeneral: triggerGeneral,
				SufferGeneral:  general,
				ResumeNum:      resume,
				TacticId:       b.Id(),
			})
			//并使其受到伤害降低4%，受智力影响，持续1回合
			rate := 0.04 + (general.BaseInfo.AbilityAttr.IntelligenceRate / 100 / 100)

			//兵刃伤害
			if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     rate,
				EffectRound:    1,
				FromTactic:     b.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
						TacticId:   b.Id(),
					})

					return revokeResp
				})
			}

			//谋略伤害
			if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     rate,
				EffectRound:    1,
				FromTactic:     b.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
						TacticId:   b.Id(),
					})

					return revokeResp
				})
			}
		}

		//有10%概率对敌军单体施加虚弱状态，持续1回合
		//自身为主将时，施加虚弱状态的概率提高至25%
		rate := 0.1
		if currentGeneral.IsMaster {
			rate = 0.25
		}
		if util.GenerateRate(rate) {
			//找到敌军单体
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, b.tacticsParams)
			//施加虚弱状态
			util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_PoorHealth, &vo.EffectHolderParams{
				EffectRound: 1,
				FromTactic:  b.Id(),
			})
		}
		return triggerResp
	})
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Id() consts.TacticId {
	return consts.BenevolentAndVirtuousThroughoutTheWorld
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Name() string {
	return "仁德载世"
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) Execute() {
}

func (b BenevolentAndVirtuousThroughoutTheWorldTactic) IsTriggerPrepare() bool {
	return false
}
