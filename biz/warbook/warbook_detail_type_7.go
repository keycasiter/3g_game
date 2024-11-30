package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 一鼓作气
// 增加突击战法18%伤害，随后每次发动突击战法该效果降低1/3
type WarBookDetailType_7 struct {
}

func (w *WarBookDetailType_7) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {
	util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_TacticsAssaultDamageImprove, &vo.EffectHolderParams{
			EffectRate:     0.18,
			FromWarbook:    consts.WarBookDetailType_7,
			ProduceGeneral: general,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_AssaultTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}

				effectParams, ok := util.BuffEffectOfWarBookGet(general, consts.BuffEffectType_TacticsAssaultDamageImprove, consts.WarBookDetailType_7)
				if ok {
					for _, param := range effectParams {
						param.EffectRate = param.EffectRate * (1 - 1/3)
					}
				}

				return revokeResp
			})
		}

		return triggerResp
	})
}
