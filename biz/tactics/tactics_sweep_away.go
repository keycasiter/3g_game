package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 横扫
// 使自己获得群攻状态，持续1回合
// 主动，30%
type SweepAwayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SweepAwayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.3
	return s
}

func (s SweepAwayTactic) Prepare() {
}

func (s SweepAwayTactic) Id() consts.TacticId {
	return consts.SweepAway
}

func (s SweepAwayTactic) Name() string {
	return "横扫"
}

func (s SweepAwayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SweepAwayTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SweepAwayTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SweepAwayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SweepAwayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SweepAwayTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	// 使自己获得群攻状态，持续1回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_GroupAttack, &vo.EffectHolderParams{
		EffectRound:    1,
		TriggerRate:    1.0,
		FromTactic:     s.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_GroupAttack,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}
}

func (s SweepAwayTactic) IsTriggerPrepare() bool {
	return false
}

func (a SweepAwayTactic) SetTriggerPrepare(triggerPrepare bool) {
}
