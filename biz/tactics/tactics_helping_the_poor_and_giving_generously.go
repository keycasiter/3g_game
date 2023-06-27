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

// 济贫好施
// 战斗第2回合时，将自身属性的40%移交给兵力最低友军，第3～5回合时，每回合治疗我军兵力最低单体（治疗率284%，受智力影响）
// 并使其受到的兵刃伤害和谋略伤害降低26%（受智力影响），持续1回合
// 指挥 100%
type HelpingThePoorAndGivingGenerouslyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 1.0
	return h
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Prepare() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)

	// 战斗第2回合时，将自身属性的40%移交给兵力最低友军，第3～5回合时，每回合治疗我军兵力最低单体（治疗率284%，受智力影响）
	// 并使其受到的兵刃伤害和谋略伤害降低26%（受智力影响），持续1回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//战斗第2回合时，将自身属性的40%移交给兵力最低友军
		if triggerRound == consts.Battle_Round_Second {
			lowestSoliderGeneral := util.GetPairLowestSoldierNumGeneral(h.tacticsParams, triggerGeneral)
			if lowestSoliderGeneral != nil {
				//转移属性获取
				force := currentGeneral.BaseInfo.AbilityAttr.ForceBase * 0.4
				intelligence := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.4
				command := currentGeneral.BaseInfo.AbilityAttr.CommandBase * 0.4
				speed := currentGeneral.BaseInfo.AbilityAttr.SpeedBase * 0.4
				//减少属性
				currentGeneral.BaseInfo.AbilityAttr.ForceBase -= force
				currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase -= intelligence
				currentGeneral.BaseInfo.AbilityAttr.CommandBase -= command
				currentGeneral.BaseInfo.AbilityAttr.SpeedBase -= speed
				//转移友军单体
				lowestSoliderGeneral.BaseInfo.AbilityAttr.ForceBase += force
				lowestSoliderGeneral.BaseInfo.AbilityAttr.IntelligenceBase += intelligence
				lowestSoliderGeneral.BaseInfo.AbilityAttr.CommandBase += command
				lowestSoliderGeneral.BaseInfo.AbilityAttr.SpeedBase += speed
			}
		}
		//第3～5回合时，每回合治疗我军兵力最低单体（治疗率284%，受智力影响）
		//并使其受到的兵刃伤害和谋略伤害降低26%（受智力影响），持续1回合
		if triggerRound >= consts.Battle_Round_Third && triggerRound <= consts.Battle_Round_Fifth {
			lowestSoliderGeneral := util.GetPairLowestSoldierNumGeneral(h.tacticsParams, triggerGeneral)

			resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 2.84)
			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  h.tacticsParams,
				ProduceGeneral: currentGeneral,
				SufferGeneral:  lowestSoliderGeneral,
				ResumeNum:      resumeNum,
			})
			//施加效果
			effectRate := 0.26 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
			if util.BuffEffectWrapSet(ctx, lowestSoliderGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     effectRate,
				EffectRound:    1,
				FromTactic:     h.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(lowestSoliderGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
						TacticId:   h.Id(),
					})

					return revokeResp
				})
			}

			if util.BuffEffectWrapSet(ctx, lowestSoliderGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     effectRate,
				EffectRound:    1,
				FromTactic:     h.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(lowestSoliderGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
						TacticId:   h.Id(),
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Id() consts.TacticId {
	return consts.HelpingThePoorAndGivingGenerously
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Name() string {
	return "济贫好施"
}

func (h HelpingThePoorAndGivingGenerouslyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (h HelpingThePoorAndGivingGenerouslyTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HelpingThePoorAndGivingGenerouslyTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HelpingThePoorAndGivingGenerouslyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (h HelpingThePoorAndGivingGenerouslyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Execute() {
}

func (h HelpingThePoorAndGivingGenerouslyTactic) IsTriggerPrepare() bool {
	return false
}
