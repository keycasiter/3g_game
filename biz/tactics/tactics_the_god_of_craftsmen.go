package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 工神
// 战斗前3回合，使我军全体获得先攻状态，我军主将造成的兵刃伤害和谋略伤害提高30%，我军副将造成伤害提升15%；
// 第4回合开始，我军全体造成伤害降低15%，持续2回合
// 指挥，100%
type TheGodOfCraftsmenTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheGodOfCraftsmenTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TheGodOfCraftsmenTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	// 战斗前3回合，使我军全体获得先攻状态，我军主将造成的兵刃伤害和谋略伤害提高30%，我军副将造成伤害提升15%；
	pairGenerals := util.GetPairGeneralArr(t.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		//先攻状态
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_FirstAttack, &vo.EffectHolderParams{
			EffectRound:    3,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_FirstAttack,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
		// 第4回合开始，我军全体造成伤害降低15%，持续2回合
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			triggerRound := params.CurrentRound

			if triggerRound == consts.Battle_Round_Fourth {
				//兵刃伤害
				if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
					EffectRound:    2,
					EffectRate:     0.15,
					FromTactic:     t.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_LaunchWeaponDamageDeduce,
							TacticId:   t.Id(),
						})

						return revokeResp
					})
				}
				//谋略伤害
				if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
					EffectRound:    2,
					EffectRate:     0.15,
					FromTactic:     t.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_LaunchStrategyDamageDeduce,
							TacticId:   t.Id(),
						})

						return revokeResp
					})
				}
			}

			return triggerResp
		})
	}
	//我军主将造成的兵刃伤害和谋略伤害提高30%
	pairMasterGeneral := util.GetPairMasterGeneral(t.tacticsParams)
	//兵刃伤害提升
	if util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRound:    3,
		EffectRate:     0.3,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
	//谋略伤害提升
	if util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRound:    3,
		EffectRate:     0.3,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_LaunchStrategyDamageImprove,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
	//我军副将造成伤害提升15%；
	viceGenerals := util.GetPairViceGenerals(t.tacticsParams)
	for _, viceGeneral := range viceGenerals {
		//兵刃伤害提升
		if util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRound:    3,
			EffectRate:     0.15,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(viceGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
		//谋略伤害提升
		if util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
			EffectRound:    3,
			EffectRate:     0.15,
			FromTactic:     t.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(viceGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_LaunchStrategyDamageImprove,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (t TheGodOfCraftsmenTactic) Id() consts.TacticId {
	return consts.TheGodOfCraftsmen
}

func (t TheGodOfCraftsmenTactic) Name() string {
	return "工神"
}

func (t TheGodOfCraftsmenTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t TheGodOfCraftsmenTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheGodOfCraftsmenTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheGodOfCraftsmenTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (t TheGodOfCraftsmenTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheGodOfCraftsmenTactic) Execute() {
}

func (t TheGodOfCraftsmenTactic) IsTriggerPrepare() bool {
	return false
}
