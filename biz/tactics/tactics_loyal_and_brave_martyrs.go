package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 忠勇义烈
// 战斗中，自身每回合有60%概率获得以下效果：
// 主动战法发动率提升6%（受武力影响）
// 武力、统率、智力提升45；28%倒戈，持续1回合，每种效果独立判定
type LoyalAndBraveMartyrsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LoyalAndBraveMartyrsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 1.0
	return l
}

func (l LoyalAndBraveMartyrsTactic) Prepare() {
	ctx := l.tacticsParams.Ctx
	currentGeneral := l.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		l.Name(),
	)

	//战斗中，自身每回合有60%概率获得以下效果：
	//主动战法发动率提升6%（受武力影响）
	//武力、统率、智力提升45；28%倒戈，持续1回合，每种效果独立判定
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		//主动战法发动率提升
		if util.GenerateRate(0.6) {
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
				TriggerRate:    0.06 + currentGeneral.BaseInfo.AbilityAttr.ForceBase/100/100,
				EffectRound:    1,
				FromTactic:     l.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_TacticsActiveTriggerImprove,
						TacticId:   l.Id(),
					})

					return revokeResp
				})
			}
		}
		//武力
		if util.GenerateRate(0.6) {
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
				EffectValue:    45,
				EffectRound:    1,
				FromTactic:     l.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_IncrForce,
						TacticId:   l.Id(),
					})

					return revokeResp
				})
			}
		}
		//统率
		if util.GenerateRate(0.6) {
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
				EffectValue:    45,
				EffectRound:    1,
				FromTactic:     l.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_IncrCommand,
						TacticId:   l.Id(),
					})

					return revokeResp
				})
			}
		}
		//智力
		if util.GenerateRate(0.6) {
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
				EffectValue:    45,
				EffectRound:    1,
				FromTactic:     l.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_IncrIntelligence,
						TacticId:   l.Id(),
					})

					return revokeResp
				})
			}
		}
		//倒戈
		if util.GenerateRate(0.6) {
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_Defection, &vo.EffectHolderParams{
				EffectRate:     0.28,
				EffectRound:    1,
				FromTactic:     l.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_Defection,
						TacticId:   l.Id(),
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (l LoyalAndBraveMartyrsTactic) Id() consts.TacticId {
	return consts.LoyalAndBraveMartyrs
}

func (l LoyalAndBraveMartyrsTactic) Name() string {
	return "忠勇义烈"
}

func (l LoyalAndBraveMartyrsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (l LoyalAndBraveMartyrsTactic) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LoyalAndBraveMartyrsTactic) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LoyalAndBraveMartyrsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (l LoyalAndBraveMartyrsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LoyalAndBraveMartyrsTactic) Execute() {
}

func (l LoyalAndBraveMartyrsTactic) IsTriggerPrepare() bool {
	return false
}
