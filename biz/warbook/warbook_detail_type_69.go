package warbook

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 以直报怨
// 自身受到控制状态时，有50%概率对敌军单体造成1次兵刃伤害（伤害率50%）
type WarBookDetailType_69 struct {
}

func (w *WarBookDetailType_69) Handle(ctx context.Context, general *vo.BattleGeneral, tacticParams *model.TacticsParams) {

	util.TacticsTriggerWrapRegister(general, consts.BattleAction_SufferDebuffEffect, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		attackGeneral := params.AttackGeneral

		if util.IsControlDeBuffEffect(params.DebuffEffect) {
			if util.GenerateRate(0.5) {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     tacticParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     attackGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 0.5,
				})
			}
		}

		return triggerResp
	})
}
