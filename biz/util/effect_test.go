package util

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"testing"
)

func TestEffectWrapSet(t *testing.T) {
	mm := make(map[consts.BuffEffectType]map[consts.BattleRound]float64, 0)
	BuffEffectWrapSet(mm, consts.BuffEffectType_Evade, consts.Battle_Round_Fourth, 1)
}
