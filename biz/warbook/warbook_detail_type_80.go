package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 出奇制胜
// 自身造成伤害降低50%，会心伤害和奇谋伤害降低50%，第2回合起自带主动战法发动率提升30%
type WarBookDetailType_80 struct {
}

func (w *WarBookDetailType_80) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	//造成伤害降低
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.5,
		FromWarbook:    consts.WarBookDetailType_80,
		ProduceGeneral: general,
	})
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
		EffectRate:     0.5,
		FromWarbook:    consts.WarBookDetailType_80,
		ProduceGeneral: general,
	})
	//会心伤害和奇谋伤害降低
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_EnhanceWeaponDeduce, &vo.EffectHolderParams{
		EffectRate:     0.5,
		FromWarbook:    consts.WarBookDetailType_80,
		ProduceGeneral: general,
	})
	util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_EnhanceStrategyDeduce, &vo.EffectHolderParams{
		EffectRate:     0.5,
		FromWarbook:    consts.WarBookDetailType_80,
		ProduceGeneral: general,
	})

	//第2回合起自带主动战法发动率提升30%
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if params.CurrentRound == consts.Battle_Round_Second {
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_TacticsActiveTriggerWithSelfImprove, &vo.EffectHolderParams{
				TriggerRate:    0.3,
				FromWarbook:    consts.WarBookDetailType_80,
				ProduceGeneral: triggerGeneral,
			})
		}

		return triggerResp
	})
}
