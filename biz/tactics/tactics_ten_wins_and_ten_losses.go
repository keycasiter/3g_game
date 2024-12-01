package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 十胜十败
// 战斗前2回合，使我军主将获得洞察状态，受到兵刃伤害和谋略伤害降低50%
// 指挥，100%
type TenWinsAndTenLossesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TenWinsAndTenLossesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TenWinsAndTenLossesTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	pairMasterGeneral := util.GetPairMasterGeneral(currentGeneral, t.tacticsParams)
	// 战斗前2回合，使我军主将获得洞察状态
	if util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_Insight, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     t.Id(),
		ProduceGeneral: pairMasterGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_Insight,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
	//受到兵刃伤害和谋略伤害降低50%
	if util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRound:    2,
		EffectRate:     0.5,
		FromTactic:     t.Id(),
		ProduceGeneral: pairMasterGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
	if util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRound:    2,
		EffectRate:     0.5,
		FromTactic:     t.Id(),
		ProduceGeneral: pairMasterGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
}

func (t TenWinsAndTenLossesTactic) Id() consts.TacticId {
	return consts.TenWinsAndTenLosses
}

func (t TenWinsAndTenLossesTactic) Name() string {
	return "十胜十败"
}

func (t TenWinsAndTenLossesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t TenWinsAndTenLossesTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TenWinsAndTenLossesTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TenWinsAndTenLossesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (t TenWinsAndTenLossesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TenWinsAndTenLossesTactic) Execute() {
}

func (t TenWinsAndTenLossesTactic) IsTriggerPrepare() bool {
	return false
}

func (a TenWinsAndTenLossesTactic) SetTriggerPrepare(triggerPrepare bool) {
}
