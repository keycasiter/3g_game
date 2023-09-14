package util

import (
	"fmt"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//武将战法触发次数容器 map<武将ID,map<战法,mao<回合，触发次数>>>
var generalTacticCountMap = map[string]map[consts.TacticId]map[consts.BattleRound]int64{}

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
	if generalUniqueId == "" {
		panic(any("generalUniqueId is nil"))
	}
	if tacticId == 0 {
		panic(any("tacticId is nil"))
	}
	if triggerTimes == 0 {
		panic(any("triggerTimes is nil"))
	}

	//按回合次数去重
	if tacticRoundTriggerMap, ok := generalTacticCountMap[generalUniqueId]; ok { //武将判断
		if roundTriggerMap, okk := tacticRoundTriggerMap[consts.TacticId(tacticId)]; okk { //战法判断
			if times, okkk := roundTriggerMap[tacticParams.CurrentRound]; okkk { //回合判断
				if times > 0 {
					triggerTimes = 0
				}
			} else {
				roundTriggerMap[tacticParams.CurrentRound] = triggerTimes
			}
		} else {
			roundTriggerM := make(map[consts.BattleRound]int64, 0)
			roundTriggerM[tacticParams.CurrentRound] = triggerTimes
			tacticRoundTriggerMap[consts.TacticId(tacticId)] = roundTriggerM
		}
	} else {
		//初始化
		roundTriggerM := make(map[consts.BattleRound]int64, 0)
		tacticRoundTriggerM := make(map[consts.TacticId]map[consts.BattleRound]int64, 0)
		//赋值
		roundTriggerM[tacticParams.CurrentRound] = triggerTimes
		tacticRoundTriggerM[consts.TacticId(tacticId)] = roundTriggerM
		generalTacticCountMap[generalUniqueId] = tacticRoundTriggerM
	}

	//按类型次数去重 被动、指挥、兵种、阵法，只在准备阶段触发计数1次
	if (consts.PassiveTacticsMap[consts.TacticId(tacticId)] ||
		consts.CommandTacticsMap[consts.TacticId(tacticId)] ||
		consts.TroopsTacticsMap[consts.TacticId(tacticId)] ||
		consts.ArmTacticsMap[consts.TacticId(tacticId)]) && tacticParams.CurrentRound > 0 {
		triggerTimes = 0
	}

	//累计对战数据
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
	} else {
		m := make(map[int64]*model.TacticStatistics, 0)
		m[tacticId] = &model.TacticStatistics{
			TacticId:         tacticId,
			TacticName:       fmt.Sprintf("%v", consts.TacticId(tacticId)),
			TriggerTimes:     triggerTimes,
			KillSoliderNum:   killSoliderNum,
			ResumeSoliderNum: resumeSoliderNum,
		}
		tacticParams.BattleTacticStatisticsMap[generalUniqueId] = m
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
