package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 以退为进
// 自身受到伤害降低15%，但每回合受到伤害提升4%
type WarBookDetailType_70 struct {
}

func (w *WarBookDetailType_70) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	//兵刃伤害
	if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.15,
		FromWarbook:    consts.WarBookDetailType_70,
		ProduceGeneral: general,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.04,
				FromWarbook:    consts.WarBookDetailType_70,
				ProduceGeneral: triggerGeneral,
			})

			return triggerResp
		})
	}
	//谋略伤害
	if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.15,
		FromWarbook:    consts.WarBookDetailType_70,
		ProduceGeneral: general,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.04,
				FromWarbook:    consts.WarBookDetailType_70,
				ProduceGeneral: triggerGeneral,
			})

			return triggerResp
		})
	}
}
