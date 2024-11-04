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

// 计定谋决
// 偷取敌方智力最高武将24点智力（受智力影响）；
// 且每当自身试图发动主动战法前，回复我军单体一定兵力（治疗率64%，受智力影响）
// 被动 100%
type PlanAndDecideTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PlanAndDecideTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 1.0
	return p
}

func (p PlanAndDecideTactic) Prepare() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)
	// 偷取敌方智力最高武将24点智力（受智力影响）；
	// 且每当自身试图发动主动战法前，回复我军单体一定兵力（治疗率64%，受智力影响）
	highestGeneral := util.GetEnemyGeneralWhoIsHighestIntelligence(currentGeneral, p.tacticsParams)
	effectVal := cast.ToInt64(24 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100)
	util.DebuffEffectWrapSet(ctx, highestGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
		EffectValue:    effectVal,
		FromTactic:     p.Id(),
		ProduceGeneral: currentGeneral,
	})
	util.BuffEffectWrapSet(ctx, highestGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
		EffectValue:    effectVal,
		FromTactic:     p.Id(),
		ProduceGeneral: currentGeneral,
	})
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		pairGeneral := util.GetPairOneGeneral(p.tacticsParams, currentGeneral)
		resumeNum := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.64)
		util.ResumeSoldierNum(&util.ResumeParams{
			Ctx:            ctx,
			TacticsParams:  p.tacticsParams,
			ProduceGeneral: triggerGeneral,
			SufferGeneral:  pairGeneral,
			ResumeNum:      resumeNum,
			TacticId:       p.Id(),
		})
		return triggerResp
	})
}

func (p PlanAndDecideTactic) Id() consts.TacticId {
	return consts.PlanAndDecide
}

func (p PlanAndDecideTactic) Name() string {
	return "计定谋决"
}

func (p PlanAndDecideTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (p PlanAndDecideTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PlanAndDecideTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PlanAndDecideTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (p PlanAndDecideTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PlanAndDecideTactic) Execute() {
}

func (p PlanAndDecideTactic) IsTriggerPrepare() bool {
	return false
}
