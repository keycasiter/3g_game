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

// 以逸待劳
// 治疗我军群体（2人，治疗率154%，受智力影响），
// 并使其下次3次受到伤害分别降低（50%、37.5%、25%，受智力影响）、且下次受到控制状态时有40%（受智力影响）几率免疫
// 主动，35%
type WaitAtOnesEaseForTheFatiguedTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WaitAtOnesEaseForTheFatiguedTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 0.35
	return w
}

func (w WaitAtOnesEaseForTheFatiguedTactic) Prepare() {

}

func (w WaitAtOnesEaseForTheFatiguedTactic) Id() consts.TacticId {
	return consts.WaitAtOnesEaseForTheFatigued
}

func (w WaitAtOnesEaseForTheFatiguedTactic) Name() string {
	return "以逸待劳"
}

func (w WaitAtOnesEaseForTheFatiguedTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (w WaitAtOnesEaseForTheFatiguedTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WaitAtOnesEaseForTheFatiguedTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WaitAtOnesEaseForTheFatiguedTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (w WaitAtOnesEaseForTheFatiguedTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (w WaitAtOnesEaseForTheFatiguedTactic) Execute() {
	ctx := w.tacticsParams.Ctx
	currentGeneral := w.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		w.Name(),
	)

	// 治疗我军群体（2人，治疗率154%，受智力影响），
	pairGenerals := util.GetPairGeneralsTwoArrByGeneral(currentGeneral, w.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.54)
		util.ResumeSoldierNum(ctx, pairGeneral, resumeNum)
		// 并使其下次3次受到伤害分别降低（50%、37.5%、25%，受智力影响
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_WaitAtOnesEaseForTheFatigued_Prepare, &vo.EffectHolderParams{
			EffectTimes:    3,
			FromTactic:     w.Id(),
			ProduceGeneral: pairGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_SufferDamage, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				times := util.BuffEffectGetCount(revokeGeneral, consts.BuffEffectType_WaitAtOnesEaseForTheFatigued_Prepare)
				switch times {
				//第一次
				case 3:
					//50%
					if util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_WaitAtOnesEaseForTheFatigued_Prepare,
						TacticId:   w.Id(),
						CostTimes:  1,
					}) {
						//受到兵刃伤害降低
						if util.BuffEffectWrapSet(ctx, revokeGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
							EffectRate:     0.5,
							EffectTimes:    1,
							FromTactic:     w.Id(),
							ProduceGeneral: revokeGeneral,
						}).IsSuccess {
							util.TacticsTriggerWrapRegister(revokeGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								triggerResp := &vo.TacticsTriggerResult{}
								triggerGeneral := params.CurrentGeneral

								util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    triggerGeneral,
									EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
									TacticId:   w.Id(),
								})

								return triggerResp
							})
						}
						//受到谋略伤害降低
						if util.BuffEffectWrapSet(ctx, revokeGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
							EffectRate:     0.5,
							EffectTimes:    1,
							FromTactic:     w.Id(),
							ProduceGeneral: revokeGeneral,
						}).IsSuccess {
							util.TacticsTriggerWrapRegister(revokeGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								triggerResp := &vo.TacticsTriggerResult{}
								triggerGeneral := params.CurrentGeneral

								util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    triggerGeneral,
									EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
									TacticId:   w.Id(),
								})

								return triggerResp
							})
						}
					}
				//第二次
				case 2:
					//37.5%
					if util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_WaitAtOnesEaseForTheFatigued_Prepare,
						TacticId:   w.Id(),
						CostTimes:  1,
					}) {
						//受到兵刃伤害降低
						if util.BuffEffectWrapSet(ctx, revokeGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
							EffectRate:     0.375,
							EffectTimes:    1,
							FromTactic:     w.Id(),
							ProduceGeneral: revokeGeneral,
						}).IsSuccess {
							util.TacticsTriggerWrapRegister(revokeGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								triggerResp := &vo.TacticsTriggerResult{}
								triggerGeneral := params.CurrentGeneral

								util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    triggerGeneral,
									EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
									TacticId:   w.Id(),
								})

								return triggerResp
							})
						}
						//受到谋略伤害降低
						if util.BuffEffectWrapSet(ctx, revokeGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
							EffectRate:     0.375,
							EffectTimes:    1,
							FromTactic:     w.Id(),
							ProduceGeneral: revokeGeneral,
						}).IsSuccess {
							util.TacticsTriggerWrapRegister(revokeGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								triggerResp := &vo.TacticsTriggerResult{}
								triggerGeneral := params.CurrentGeneral

								util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    triggerGeneral,
									EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
									TacticId:   w.Id(),
								})

								return triggerResp
							})
						}
					}
				//第三次
				case 1:
					//25%
					if util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_WaitAtOnesEaseForTheFatigued_Prepare,
						TacticId:   w.Id(),
						CostTimes:  1,
					}) {
						//受到兵刃伤害降低
						if util.BuffEffectWrapSet(ctx, revokeGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
							EffectRate:     0.25,
							EffectTimes:    1,
							FromTactic:     w.Id(),
							ProduceGeneral: revokeGeneral,
						}).IsSuccess {
							util.TacticsTriggerWrapRegister(revokeGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								triggerResp := &vo.TacticsTriggerResult{}
								triggerGeneral := params.CurrentGeneral

								util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    triggerGeneral,
									EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
									TacticId:   w.Id(),
								})

								return triggerResp
							})
						}
						//受到谋略伤害降低
						if util.BuffEffectWrapSet(ctx, revokeGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
							EffectRate:     0.25,
							EffectTimes:    1,
							FromTactic:     w.Id(),
							ProduceGeneral: revokeGeneral,
						}).IsSuccess {
							util.TacticsTriggerWrapRegister(revokeGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								triggerResp := &vo.TacticsTriggerResult{}
								triggerGeneral := params.CurrentGeneral

								util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
									Ctx:        ctx,
									General:    triggerGeneral,
									EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
									TacticId:   w.Id(),
								})

								return triggerResp
							})
						}
					}
				}

				return revokeResp
			})
		}
		//且下次受到控制状态时有40%（受智力影响）几率免疫
		if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_WaitAtOnesEaseForTheFatigued_ImmunityControl, &vo.EffectHolderParams{
			EffectTimes:    1,
			FromTactic:     w.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_SufferDebuffEffect, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				triggerRate := 0.4 + revokeGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
				if util.IsControlDeBuffEffect(params.DebuffEffect) && util.GenerateRate(triggerRate) {
					revokeResp.IsTerminate = true
				}

				return revokeResp
			})
		}
	}
}

func (w WaitAtOnesEaseForTheFatiguedTactic) IsTriggerPrepare() bool {
	return false
}
