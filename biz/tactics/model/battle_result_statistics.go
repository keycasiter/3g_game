package model

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

// 对战数据统计
type BattleResultStatistics struct {
	//我军统计
	FightingTeam *TeamBattleStatistics
	//敌军统计
	EnemyTeam *TeamBattleStatistics
}

// 队伍对战数据统计
type TeamBattleStatistics struct {
	//***队伍原始数据***
	//队伍信息
	BattleTeam *vo.BattleTeam

	//***对战数据***
	//对战结果
	BattleResult consts.BattleResult
	//对战统计信息
	GeneralBattleStatisticsList []*GeneralBattleStatistics
}

// 武将对战数据统计
type GeneralBattleStatistics struct {
	//战法统计
	TacticStatisticsList []*TacticStatistics
	//普攻统计
	GeneralAttackStatistics *TacticStatistics

	//回合剩余兵力
	RoundRemainSoliderNum map[consts.BattlePhase]map[consts.BattleRound]int64
}

type TacticStatistics struct {
	//战法ID
	TacticId int64
	//战法名称
	TacticName string
	//战法品质
	TacticQuality int64
	//发动次数
	TriggerTimes int64
	//杀敌
	KillSoliderNum int64
	//救援
	ResumeSoliderNum int64

	//回合发动
	RoundTriggerTimes map[consts.BattlePhase]map[consts.BattleRound]int64
	//回合杀敌
	RoundKillSoliderNum map[consts.BattlePhase]map[consts.BattleRound]int64
	//回合救援
	RoundResumeSoliderNum map[consts.BattlePhase]map[consts.BattleRound]int64
}
