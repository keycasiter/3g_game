package model

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

type TacticsParams struct {
	// 当前回合
	CurrentRound consts.BattleRound
	// 当前武将
	CurrentGeneral *vo.BattleGeneral
	// 出战队伍信息
	FightingGeneralMap map[int64]*vo.BattleGeneral
	// 对战队伍信息
	EnemyGeneralMap map[int64]*vo.BattleGeneral
}
