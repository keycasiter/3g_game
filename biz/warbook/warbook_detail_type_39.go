package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 援其必攻
// 为损失兵力最多的友军恢复兵力时，使其受到伤害降低12%，持续1回合
type WarBookDetailType_39 struct {
}

func (w *WarBookDetailType_39) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_Resume, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		sufferResumeGeneral := params.SufferResumeGeneral

		maxLossGeneral := util.GetPairMaxLossSoldierNumGeneral(triggerGeneral, tacticParams)
		if maxLossGeneral.BaseInfo.UniqueId == sufferResumeGeneral.BaseInfo.UniqueId {
			//谋略
			if util.BuffEffectWrapSet(ctx, sufferResumeGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     0.12,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_39,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(sufferResumeGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_SufferStrategyDamageDeduce,
						WarbookDetailType: consts.WarBookDetailType_39,
					})

					return revokeResp
				})
			}
			//兵刃
			if util.BuffEffectWrapSet(ctx, sufferResumeGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     0.12,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_39,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(sufferResumeGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_SufferWeaponDamageDeduce,
						WarbookDetailType: consts.WarBookDetailType_39,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
