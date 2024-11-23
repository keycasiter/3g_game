package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 临锋决敌
// 使自己进入连击和免疫缴械状态，持续1回合
type FacingTheEnemyHeadOnTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a FacingTheEnemyHeadOnTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.55
	return a
}

func (a FacingTheEnemyHeadOnTactic) Prepare() {

}

func (a FacingTheEnemyHeadOnTactic) Id() consts.TacticId {
	return consts.FacingTheEnemyHeadOn
}

func (a FacingTheEnemyHeadOnTactic) Name() string {
	return "临锋决敌"
}

func (a FacingTheEnemyHeadOnTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (a FacingTheEnemyHeadOnTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a FacingTheEnemyHeadOnTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a FacingTheEnemyHeadOnTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a FacingTheEnemyHeadOnTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a FacingTheEnemyHeadOnTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)
	// 使自己进入免疫缴械状态，持续1回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ImmunityCancelWeapon, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     a.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_ImmunityCancelWeapon,
				TacticId:   a.Id(),
			})

			return revokeResp
		})
	}

	// 使自己进入连击，持续1回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ContinuousAttack, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     a.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_ContinuousAttack,
				TacticId:   a.Id(),
			})

			return revokeResp
		})
	}
}

func (a FacingTheEnemyHeadOnTactic) IsTriggerPrepare() bool {
	return false
}
