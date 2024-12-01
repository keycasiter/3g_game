package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 善战
// 每次造成伤害后，提升0.5%伤害，最多叠加8次
type WarBookDetailType_77 struct {
}

func (w *WarBookDetailType_77) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_DamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		//兵刃伤害提升
		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.005,
			EffectTimes:    1,
			MaxEffectTimes: 8,
			FromWarbook:    consts.WarBookDetailType_77,
			ProduceGeneral: triggerGeneral,
		})
		//谋略伤害提升
		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.005,
			EffectTimes:    1,
			MaxEffectTimes: 8,
			FromWarbook:    consts.WarBookDetailType_77,
			ProduceGeneral: triggerGeneral,
		})

		return triggerResp
	})
}
