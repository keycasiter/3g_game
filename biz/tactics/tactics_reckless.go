package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 鲁莽
// 使自己获得洞察（免疫所有控制效果）及42%倒戈（造成兵刃伤害时，恢复自身基于伤害量的一定兵力），持续2回合
// 主动，30%
type RecklessTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RecklessTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 0.35
	return r
}

func (r RecklessTactic) Prepare() {
}

func (r RecklessTactic) Id() consts.TacticId {
	return consts.Reckless
}

func (r RecklessTactic) Name() string {
	return "鲁莽"
}

func (r RecklessTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RecklessTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RecklessTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RecklessTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (r RecklessTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RecklessTactic) Execute() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)

	//使自己获得洞察（免疫所有控制效果）及42%倒戈（造成兵刃伤害时，恢复自身基于伤害量的一定兵力），持续2回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Insight, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     r.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_Insight,
				TacticId:   r.Id(),
			})

			return revokeResp
		})
	}
	//倒戈
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Defection, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     r.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_Defection,
				TacticId:   r.Id(),
			})

			return revokeResp
		})
	}
}

func (r RecklessTactic) IsTriggerPrepare() bool {
	return false
}

func (a RecklessTactic) SetTriggerPrepare(triggerPrepare bool) {
}
