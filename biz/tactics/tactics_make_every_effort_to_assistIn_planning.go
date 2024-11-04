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

// 竭力佐谋
// 使敌军智力最高单体智力降低20%，并有70%概率使自身本回合非自带主动战法发动率提高100%，持续1回合
type MakeEveryEffortToAssistInPlanningTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MakeEveryEffortToAssistInPlanningTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	m.tacticsParams = tacticsParams
	m.triggerRate = 0.55
	return m
}

func (m MakeEveryEffortToAssistInPlanningTactic) Prepare() {

}

func (m MakeEveryEffortToAssistInPlanningTactic) Id() consts.TacticId {
	return consts.MakeEveryEffortToAssistInPlanning
}

func (m MakeEveryEffortToAssistInPlanningTactic) Name() string {
	return "竭力佐谋"
}

func (m MakeEveryEffortToAssistInPlanningTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (m MakeEveryEffortToAssistInPlanningTactic) GetTriggerRate() float64 {
	return m.triggerRate
}

func (m MakeEveryEffortToAssistInPlanningTactic) SetTriggerRate(rate float64) {
	m.triggerRate = rate
}

func (m MakeEveryEffortToAssistInPlanningTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (m MakeEveryEffortToAssistInPlanningTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (m MakeEveryEffortToAssistInPlanningTactic) Execute() {
	ctx := m.tacticsParams.Ctx
	currentGeneral := m.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		m.Name(),
	)
	//使敌军智力最高单体智力降低20%，并有70%概率使自身本回合非自带主动战法发动率提高100%，持续1回合
	enemyGeneral := util.GetEnemyGeneralWhoIsHighestIntelligence(currentGeneral, m.tacticsParams)
	//降低智力
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
		EffectValue:    cast.ToInt64(enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.2),
		EffectRound:    1,
		FromTactic:     m.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrIntelligence,
				TacticId:   m.Id(),
			})
			return revokeResp
		})
	}
	//并有70%概率使自身本回合非自带主动战法发动率提高100%，持续1回合
	if util.GenerateRate(0.7) {
		if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_TacticsActiveTriggerNoSelfImprove, &vo.EffectHolderParams{
			TriggerRate:    1.0,
			EffectRound:    1,
			FromTactic:     m.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_TacticsActiveTriggerNoSelfImprove,
					TacticId:   m.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (m MakeEveryEffortToAssistInPlanningTactic) IsTriggerPrepare() bool {
	return false
}
