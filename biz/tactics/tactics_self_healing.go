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

// 自愈
// 战斗中，使自己获得休整状态，每回合回复一定兵力（回复率100%）
// 被动，100%
type SelfHealingTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SelfHealingTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s SelfHealingTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	// 战斗中，使自己获得休整状态，每回合回复一定兵力（回复率100%）
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Rest, &vo.EffectHolderParams{
		FromTactic:     s.Id(),
		EffectRound:    8,
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    triggerGeneral,
				EffectType: consts.BuffEffectType_Rest,
				TacticId:   s.Id(),
			}) {
				resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1)
				util.ResumeSoldierNum(&util.ResumeParams{
					Ctx:            ctx,
					TacticsParams:  s.tacticsParams,
					ProduceGeneral: currentGeneral,
					SufferGeneral:  currentGeneral,
					ResumeNum:      resumeNum,
				})
			}

			return triggerResp
		})
	}
}

func (s SelfHealingTactic) Id() consts.TacticId {
	return consts.SelfHealing
}

func (s SelfHealingTactic) Name() string {
	return "自愈"
}

func (s SelfHealingTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SelfHealingTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SelfHealingTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SelfHealingTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (s SelfHealingTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SelfHealingTactic) Execute() {

}

func (s SelfHealingTactic) IsTriggerPrepare() bool {
	return false
}
