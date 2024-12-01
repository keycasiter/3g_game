package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 乐善好施
// 自身造成治疗时，有概率（取会心和奇谋几率最高一项）使此次治疗提升100%（取会心伤害和奇谋伤害更高一项）
type WarBookDetailType_71 struct {
}

func (w *WarBookDetailType_71) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_Resume, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		enhanceWeaponRate, _ := util.BuffEffectGetAggrEffectRate(triggerGeneral, consts.BuffEffectType_EnhanceWeapon)
		enhanceStrategyRate, _ := util.BuffEffectGetAggrEffectRate(triggerGeneral, consts.BuffEffectType_EnhanceStrategy)

		triggerRate := enhanceWeaponRate
		if enhanceStrategyRate > triggerRate {
			triggerRate = enhanceStrategyRate
		}

		if util.GenerateRate(triggerRate) {
			if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_ResumeImprove, &vo.EffectHolderParams{
				EffectRate:     1,
				EffectRound:    1,
				FromWarbook:    consts.WarBookDetailType_71,
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_ResumeEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfWarbookCostRound(&util.BuffEffectOfWarbookCostRoundParams{
						Ctx:               ctx,
						General:           revokeGeneral,
						EffectType:        consts.BuffEffectType_ResumeImprove,
						WarbookDetailType: consts.WarBookDetailType_71,
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}
