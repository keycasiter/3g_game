package util

import (
	"fmt"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"testing"
)

func TestEffectWrapSet(t *testing.T) {
}

func TestBuffEffectOfTacticCostRound(t *testing.T) {
	res := BuffEffectOfTacticCostRound(&vo.BattleGeneral{
		BuffEffectHolderMap: map[consts.BuffEffectType][]*vo.EffectHolderParams{
			consts.BuffEffectType_AttackHeart: {
				{
					EffectTimes: 2,
					FromTactic:  consts.AbilityToRuleTheCountry,
				},
			},
		},
	},
		consts.BuffEffectType_AttackHeart,
		consts.AbilityToRuleTheCountry,
		1,
	)
	fmt.Printf("result:%v", res)
}
