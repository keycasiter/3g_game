package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 形一阵
// 我军三名武将自带战法类型相同时，战斗中，自身最高属性提高60点，友军群体（2人）造成及受到伤害降低30%
// 此效果每回合降低10%，该效果结束后，每回合使其造成伤害提高16%，受到伤害提高4%，可叠加
type ShapelyArrayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s ShapelyArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s ShapelyArrayTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	// 我军三名武将自带战法类型相同时
	pairGenerals := util.GetPairGeneralArr(s.tacticsParams)

	tacticsType := pairGenerals[0].EquipTactics[0].Type
	for _, general := range pairGenerals {
		if tacticsType != general.EquipTactics[0].Type {
			hlog.CtxInfof(ctx, "三名武将自带战法类型不相同，无法发动")
			return
		}
	}
	//战斗中，自身最高属性提高60点，
	attr, _ := util.GetGeneralHighestAttr(currentGeneral)
	util.ImproveGeneralAttr(currentGeneral, attr, 60)
	//友军群体（2人）造成及受到伤害降低30%，此效果每回合降低10%，该效果结束后，每回合使其造成伤害提高16%，受到伤害提高4%，可叠加
	pairTwoGenerals := util.GetPairGeneralsNotSelf(s.tacticsParams, currentGeneral)
	for _, pairGeneral := range pairTwoGenerals {
		//造成伤害降低
		if util.DebuffEffectWrapSet(ctx, pairGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.3,
			FromTactic:     s.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRate(&util.DebuffEffectOfTacticCostRateParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_LaunchWeaponDamageDeduce,
					TacticId:   s.Id(),
					EffectRate: 0.1,
				})

				return revokeResp
			})
		}
		//受到伤害降低
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.3,
			FromTactic:     s.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRate(&util.BuffEffectOfTacticCostRateParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
					TacticId:   s.Id(),
					EffectRate: 0.1,
				})

				return revokeResp
			})
		}

		//此效果每回合降低10%，该效果结束后，每回合使其造成伤害提高16%，受到伤害提高4%，可叠加
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerRound := params.CurrentRound
			triggerGeneral := params.CurrentGeneral
			triggerResp := &vo.TacticsTriggerResult{}

			if triggerRound >= consts.Battle_Round_Fourth {
				//造成兵刃伤害
				util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.16,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				})
				//造成谋略伤害
				util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.16,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				})
				//受到兵刃伤害
				util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.04,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				})
				//受到谋略伤害
				util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.04,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				})
			}

			return triggerResp
		})
	}
}

func (s ShapelyArrayTactic) Id() consts.TacticId {
	return consts.ShapelyArray
}

func (s ShapelyArrayTactic) Name() string {
	return "形一阵"
}

func (s ShapelyArrayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (s ShapelyArrayTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s ShapelyArrayTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s ShapelyArrayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (s ShapelyArrayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s ShapelyArrayTactic) Execute() {

}

func (s ShapelyArrayTactic) IsTriggerPrepare() bool {
	return false
}
