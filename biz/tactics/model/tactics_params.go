package model

import (
	"context"

	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

// 战法参数
type TacticsParams struct {
	Ctx context.Context

	/******** 战法信息相关 *******/
	// 战法类型
	TacticsType consts.TacticsType

	/******** 回合属性变量 *******/
	// 当前对战ID
	CurrentBattleId string
	//当前对战阶段
	CurrentPhase consts.BattlePhase
	// 当前回合
	CurrentRound consts.BattleRound
	// 当前武将
	CurrentGeneral *vo.BattleGeneral
	// 当前被攻击武将
	CurrentSufferGeneral *vo.BattleGeneral
	// 当前伤害量
	CurrentDamageNum int64
	// 当前治疗量
	CurrentResumeNum int64
	// 当前对战执行动作
	CurrentBattleAction consts.BattleAction

	/******** 对战武将信息变量 *******/
	// 出战队伍信息
	FightingTeam *vo.BattleTeam
	// 对战队伍信息
	EnemyTeam *vo.BattleTeam
	// 出战武将信息Map
	FightingGeneralMap map[string]*vo.BattleGeneral
	// 对战武将信息Map
	EnemyGeneralMap map[string]*vo.BattleGeneral
	// 全部武将信息Map
	AllGeneralMap map[string]*vo.BattleGeneral
	// 全部武将信息Arr
	AllGeneralArr []*vo.BattleGeneral

	/******** 对战战报信息 *******/
	//对战回合描述 map<对战阶段,<回合,[]战报内容>>
	BattleProcessStatisticsMap map[consts.BattlePhase]map[consts.BattleRound][]string
	//对战武将战法数据统计 map<武将uniqueId,map<战法ID,统计>>
	BattleTacticStatisticsMap map[string]map[int64]*TacticStatistics
	//对战武将普攻数据统计 map<武将uniqueId,统计>
	BattleAttackStatisticsMap map[string]*TacticStatistics
}
