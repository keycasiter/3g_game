package model

import (
	"context"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

type TacticsParams struct {
	Ctx context.Context

	/******** 回合属性变量 *******/
	// 当前对战ID
	CurrentBattleId string
	// 当前回合
	CurrentRound consts.BattleRound
	// 当前武将
	CurrentGeneral *vo.BattleGeneral
	// 当前被攻击武将
	CurrentSufferGeneral *vo.BattleGeneral
	// 当前对战执行动作
	CurrentBattleAction consts.BattleAction

	/******** 对战武将信息变量 *******/
	// 出战武将信息Map
	FightingGeneralMap map[int64]*vo.BattleGeneral
	// 对战武将信息Map
	EnemyGeneralMap map[int64]*vo.BattleGeneral
	// 全部武将信息Map
	AllGeneralMap map[int64]*vo.BattleGeneral
	// 全部武将信息Arr，严格按照顺序[我方主将/我方副将1/我方副将2/敌方主将/敌方副将1/敌方副将2]
	AllGeneralArr []*vo.BattleGeneral
}
