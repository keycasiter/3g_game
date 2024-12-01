package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 疑城
// 主动 75%
// 使自己免疫混乱，并提高68点武力，持续2回合
type ShamCastlesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s ShamCastlesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.75
	return s
}

func (s ShamCastlesTactic) Prepare() {

}

func (s ShamCastlesTactic) Id() consts.TacticId {
	return consts.ShamCastles
}

func (s ShamCastlesTactic) Name() string {
	return "疑城"
}

func (s ShamCastlesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s ShamCastlesTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s ShamCastlesTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s ShamCastlesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s ShamCastlesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s ShamCastlesTactic) Execute() {
	currentGeneral := s.tacticsParams.CurrentGeneral
	ctx := s.tacticsParams.Ctx

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	//使自己免疫混乱，并提高68点武力，持续2回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ImmunityChaos, &vo.EffectHolderParams{
		EffectRound: 2,
		FromTactic:  s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_ImmunityChaos,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}
	//提高武力
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
		EffectRound: 2,
		EffectValue: 68,
		FromTactic:  s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrForce,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}
}

func (s ShamCastlesTactic) IsTriggerPrepare() bool {
	return false
}

func (a ShamCastlesTactic) SetTriggerPrepare(triggerPrepare bool) {
}
