package vo

import (
	"github.com/keycasiter/3g_game/biz/consts"
)

// 战报对象
type BattleReport struct {
	BattleRoundReportMap map[consts.BattlePhase]map[consts.BattleRound]string
}
