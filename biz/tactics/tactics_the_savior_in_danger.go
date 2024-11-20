package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 临危救主
// 战斗中，我军全体每普通攻击2次，有50%概率（受速度影响）治疗我军单体（治疗率85%，受武力或智力更高一项影响），
// 并使我军兵力最低单体受到谋略伤害降低12%（受速度影响），可叠加，持续1回合
type TheSaviorInDangerTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a TheSaviorInDangerTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a TheSaviorInDangerTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	// 战斗中，我军全体每普通攻击2次，有50%概率（受速度影响）治疗我军单体（治疗率85%，受武力或智力更高一项影响），
	allPairGenerals := util.GetPairGeneralArr(currentGeneral, a.tacticsParams)
	allPairGeneralAttackTimes := int64(0)
	for _, pairGeneral := range allPairGenerals {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			//我军全体每普通攻击2次
			for _, general := range allPairGenerals {
				if (general.ExecuteGeneralAttackNum-allPairGeneralAttackTimes)%2 > 0 {
					//已执行次数累计
					allPairGeneralAttackTimes += general.ExecuteGeneralAttackNum
					//有50%概率（受速度影响）
					triggerRate := 0.5 + triggerGeneral.BaseInfo.AbilityAttr.SpeedBase/100/100
					if util.GenerateRate(triggerRate) {
						// 治疗我军单体（治疗率85%，受武力或智力更高一项影响）
						_, val := util.GetGeneralHighestBetweenForceOrIntelligence(triggerGeneral)
						effectRate := 0.85 + val/100/100

						//找到我军单体
						resumePairGeneral := util.GetPairOneGeneral(a.tacticsParams, triggerGeneral)
						resumeNum := int64(float64(resumePairGeneral.SoldierNum) * effectRate)

						util.ResumeSoldierNum(&util.ResumeParams{
							Ctx:            ctx,
							TacticsParams:  a.tacticsParams,
							ProduceGeneral: triggerGeneral,
							SufferGeneral:  resumePairGeneral,
							ResumeNum:      resumeNum,
							TacticId:       a.Id(),
						})
						// 并使我军兵力最低单体受到谋略伤害降低12%（受速度影响），可叠加，持续1回合
						lowestSoliderGeneral := util.GetPairLowestSoldierNumGeneral(a.tacticsParams, triggerGeneral)
						if util.BuffEffectWrapSet(ctx, lowestSoliderGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
							EffectRate:     0.12 + triggerGeneral.BaseInfo.AbilityAttr.SpeedBase/100/100,
							EffectTimes:    1,
							FromTactic:     a.Id(),
							ProduceGeneral: triggerGeneral,
						}).IsSuccess {
							util.TacticsTriggerWrapRegister(lowestSoliderGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								revokeResp := &vo.TacticsTriggerResult{}
								revokeGeneral := params.CurrentGeneral

								util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    revokeGeneral,
									EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
									TacticId:   a.Id(),
								})

								return revokeResp
							})
						}
					}
				}
			}

			return triggerResp
		})
	}
}

func (a TheSaviorInDangerTactic) Id() consts.TacticId {
	return consts.TheSaviorInDanger
}

func (a TheSaviorInDangerTactic) Name() string {
	return "临危救主"
}

func (a TheSaviorInDangerTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a TheSaviorInDangerTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a TheSaviorInDangerTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a TheSaviorInDangerTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a TheSaviorInDangerTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a TheSaviorInDangerTactic) Execute() {
}

func (a TheSaviorInDangerTactic) IsTriggerPrepare() bool {
	return false
}
