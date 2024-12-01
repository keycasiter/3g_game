package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 兵行诡道
// 战斗第2回合，我军单体与敌军随机单体造成伤害降低65%，持续1回合
type WarBookDetailType_68 struct {
}

func (w *WarBookDetailType_68) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		generals := make([]*vo.BattleGeneral, 0)
		generals = append(generals, triggerGeneral)
		generals = append(generals, util.GetEnemyOneGeneralByGeneral(triggerGeneral, tacticParams))

		if params.CurrentRound == consts.Battle_Round_Second {
			for _, general_ := range generals {
				//兵刃造成伤害降低
				if util.DebuffEffectWrapSet(ctx, general_, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
					EffectRate:     0.65,
					EffectRound:    1,
					FromWarbook:    consts.WarBookDetailType_68,
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(general_, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral
						util.DeBuffEffectOfWarbookCostRound(&util.DeBuffEffectOfWarbookCostRoundParams{
							Ctx:               ctx,
							General:           revokeGeneral,
							EffectType:        consts.DebuffEffectType_LaunchWeaponDamageDeduce,
							WarbookDetailType: consts.WarBookDetailType_68,
						})
						return revokeResp
					})
				}
				//谋略造成伤害降低
				if util.DebuffEffectWrapSet(ctx, general_, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
					EffectRate:     0.65,
					EffectRound:    1,
					FromWarbook:    consts.WarBookDetailType_68,
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(general_, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral
						util.DeBuffEffectOfWarbookCostRound(&util.DeBuffEffectOfWarbookCostRoundParams{
							Ctx:               ctx,
							General:           revokeGeneral,
							EffectType:        consts.DebuffEffectType_LaunchStrategyDamageDeduce,
							WarbookDetailType: consts.WarBookDetailType_68,
						})
						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}
