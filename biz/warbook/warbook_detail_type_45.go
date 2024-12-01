package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 掩虚
// 若每回合行动时，自身为损失兵力最多者，自身受到伤害降低6%
type WarBookDetailType_45 struct {
}

func (w *WarBookDetailType_45) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		lowestSoliderNumGeneral := util.GetPairLowestSoldierNumGeneral(tacticParams, triggerGeneral)

		if lowestSoliderNumGeneral.BaseInfo.UniqueId == triggerGeneral.BaseInfo.UniqueId {
			//兵刃
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     0.06,
				FromWarbook:    consts.WarBookDetailType_45,
				ProduceGeneral: triggerGeneral,
			})
			//谋略
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     0.06,
				FromWarbook:    consts.WarBookDetailType_45,
				ProduceGeneral: triggerGeneral,
			})
		}

		return triggerResp
	})
}
