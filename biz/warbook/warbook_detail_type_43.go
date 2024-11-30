package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 示敌以弱
// 战斗前4回合，使自身造成兵刃伤害和谋略伤害降低15%，第4回合开始，自身主动战法发动率提高7%
type WarBookDetailType_43 struct {
}

func (w *WarBookDetailType_43) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	//兵刃
	if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.15,
		EffectRound:    4,
		FromWarbook:    consts.WarBookDetailType_43,
		ProduceGeneral: general,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			util.DeBuffEffectOfWarbookCostRound(&util.DeBuffEffectOfWarbookCostRoundParams{
				Ctx:               ctx,
				General:           triggerGeneral,
				EffectType:        consts.DebuffEffectType_LaunchWeaponDamageDeduce,
				WarbookDetailType: consts.WarBookDetailType_43,
			})

			return triggerResp
		})
	}
	//谋略
	if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.15,
		EffectRound:    4,
		FromWarbook:    consts.WarBookDetailType_43,
		ProduceGeneral: general,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_EndAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			util.DeBuffEffectOfWarbookCostRound(&util.DeBuffEffectOfWarbookCostRoundParams{
				Ctx:               ctx,
				General:           triggerGeneral,
				EffectType:        consts.DebuffEffectType_LaunchStrategyDamageDeduce,
				WarbookDetailType: consts.WarBookDetailType_43,
			})

			return triggerResp
		})
	}

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if params.CurrentRound == consts.Battle_Round_Fourth {
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_TacticsActiveTriggerWithSelfImprove, &vo.EffectHolderParams{
				EffectRate:     0.07,
				FromWarbook:    consts.WarBookDetailType_43,
				ProduceGeneral: triggerGeneral,
			})
		}

		return triggerResp
	})
}
