package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 刚勇无前
// 战斗中，受到兵刃伤害后，下回合行动时，提高20%会心并使下一个攻击的伤害提高65%，持续1回合
type StrongAndBraveWithoutAdvanceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StrongAndBraveWithoutAdvanceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s StrongAndBraveWithoutAdvanceTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	//战斗中，受到兵刃伤害后，下回合行动时，提高20%会心并使下一个攻击的伤害提高65%，持续1回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_WeaponDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			tResp := &vo.TacticsTriggerResult{}
			tGeneral := params.CurrentGeneral
			tRound := params.CurrentRound

			if triggerRound+1 == tRound {
				//攻心效果
				if util.BuffEffectWrapSet(ctx, tGeneral, consts.BuffEffectType_AttackHeart, &vo.EffectHolderParams{
					EffectRate:     0.2,
					EffectRound:    1,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(tGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_AttackHeart,
							TacticId:   s.Id(),
						})

						return revokeResp
					})
				}
				//兵刃伤害提升
				if util.BuffEffectWrapSet(ctx, tGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.65,
					EffectRound:    1,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(tGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
							TacticId:   s.Id(),
						})

						return revokeResp
					})
				}
				//谋略伤害提升
				if util.BuffEffectWrapSet(ctx, tGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.65,
					EffectRound:    1,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(tGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_LaunchStrategyDamageImprove,
							TacticId:   s.Id(),
						})

						return revokeResp
					})
				}
			}

			return tResp
		})

		return triggerResp
	})

}

func (s StrongAndBraveWithoutAdvanceTactic) Id() consts.TacticId {
	return consts.StrongAndBraveWithoutAdvance
}

func (s StrongAndBraveWithoutAdvanceTactic) Name() string {
	return "刚勇无前"
}

func (s StrongAndBraveWithoutAdvanceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (s StrongAndBraveWithoutAdvanceTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s StrongAndBraveWithoutAdvanceTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s StrongAndBraveWithoutAdvanceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (s StrongAndBraveWithoutAdvanceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s StrongAndBraveWithoutAdvanceTactic) Execute() {
}

func (s StrongAndBraveWithoutAdvanceTactic) IsTriggerPrepare() bool {
	return false
}
