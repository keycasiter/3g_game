package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 分利
// 自身受到兵刃伤害后，统率提升1%；
// 受到谋略伤害后，智力提升1%；
// 分别可叠加5次
type WarBookDetailType_78 struct {
}

func (w *WarBookDetailType_78) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	//受到兵刃伤害
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferWeaponDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.SufferAttackGeneral

		//统率提升
		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
			EffectValue:    cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.CommandBase * 0.1),
			EffectTimes:    1,
			MaxEffectTimes: 5,
			FromWarbook:    consts.WarBookDetailType_78,
			ProduceGeneral: triggerGeneral,
		})

		return triggerResp
	})

	//受到谋略伤害
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferStrategyDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.SufferAttackGeneral

		//智力提升
		util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
			EffectValue:    cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.1),
			EffectTimes:    1,
			MaxEffectTimes: 5,
			FromWarbook:    consts.WarBookDetailType_78,
			ProduceGeneral: triggerGeneral,
		})

		return triggerResp
	})
}
