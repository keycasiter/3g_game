package util

import (
	"fmt"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//判断对战结果
func JudgeBattleResult(team1 *vo.BattleTeam, team2 *vo.BattleTeam) consts.BattleResult {
	team1SoliderNum := int64(0)
	team2SoliderNum := int64(0)
	//败
	for _, general := range team1.BattleGenerals {
		team1SoliderNum += general.SoldierNum
		if general.IsMaster && general.SoldierNum == 0 {
			return consts.BattleResult_Lose
		}
	}
	//胜
	for _, general := range team2.BattleGenerals {
		team2SoliderNum += general.SoldierNum
		if general.IsMaster && general.SoldierNum == 0 {
			return consts.BattleResult_Win
		}
	}

	if team1SoliderNum > team2SoliderNum { //优势平
		return consts.BattleResult_Advantage_Draw
	} else if team1SoliderNum < team2SoliderNum { //劣势平
		return consts.BattleResult_Inferiority_Draw
	}

	return consts.BattleResult_Draw
}

//战法触发统计
func TacticReport(tacticParams *model.TacticsParams, generalUniqueId string, tacticId int64, triggerTimes int64, killSoliderNum int64, resumeSoliderNum int64) {
	if tacticStatisticsMap, ok := tacticParams.BattleTacticStatisticsMap[generalUniqueId]; ok {
		if tacticStatistics, okk := tacticStatisticsMap[tacticId]; okk {
			tacticStatistics.TriggerTimes += triggerTimes
			tacticStatistics.KillSoliderNum += killSoliderNum
			tacticStatistics.ResumeSoliderNum += resumeSoliderNum
		} else {
			tacticStatisticsMap[tacticId] = &model.TacticStatistics{
				TacticId:         tacticId,
				TacticName:       fmt.Sprintf("%v", consts.TacticId(tacticId)),
				TriggerTimes:     triggerTimes,
				KillSoliderNum:   killSoliderNum,
				ResumeSoliderNum: resumeSoliderNum,
			}
		}
	}
}

//普攻触发统计
func AttackReport(tacticParams *model.TacticsParams, generalUniqueId string, triggerTimes int64, killSoliderNum int64, resumeSoliderNum int64) {
	if attackStatistics, ok := tacticParams.BattleAttackStatisticsMap[generalUniqueId]; ok {
		attackStatistics.TriggerTimes += triggerTimes
		attackStatistics.KillSoliderNum += killSoliderNum
		attackStatistics.ResumeSoliderNum += resumeSoliderNum
	} else {
		tacticParams.BattleAttackStatisticsMap[generalUniqueId] = &model.TacticStatistics{
			TriggerTimes:     triggerTimes,
			KillSoliderNum:   killSoliderNum,
			ResumeSoliderNum: resumeSoliderNum,
		}
	}
}
