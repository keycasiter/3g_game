package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 进言
// 主动 35%
// 使友军单体主动战法发动几率提高8%，并提高40点智力，持续2回合
type IntroductionTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i IntroductionTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.35
	return i
}

func (i IntroductionTactic) Prepare() {

}

func (i IntroductionTactic) Id() consts.TacticId {
	return consts.Introduction
}

func (i IntroductionTactic) Name() string {
	return "进言"
}

func (i IntroductionTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (i IntroductionTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i IntroductionTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i IntroductionTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i IntroductionTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i IntroductionTactic) Execute() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		i.Name(),
	)

	//使友军单体主动战法发动几率提高8%，并提高40点智力，持续2回合
	//找到友军单体
	pairGeneral := util.GetPairOneGeneral(i.tacticsParams, currentGeneral)

	//主动战法发动率
	if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
		TriggerRate: 0.08,
		EffectRound: 2,
		FromTactic:  i.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_TacticsActiveTriggerImprove,
				TacticId:   i.Id(),
			})

			return revokeResp
		})
	}
	//智力提升
	if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
		EffectRound: 2,
		EffectValue: 40,
		FromTactic:  i.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrIntelligence,
				TacticId:   i.Id(),
			})

			return revokeResp
		})
	}
}

func (i IntroductionTactic) IsTriggerPrepare() bool {
	return false
}

func (a IntroductionTactic) SetTriggerPrepare(triggerPrepare bool) {
}
