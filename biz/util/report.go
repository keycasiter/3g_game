package util

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

func InitReport() *vo.BattleReport {
	return &vo.BattleReport{}
}

func AppendReport(report *vo.BattleReport, battlePhase consts.BattlePhase, round consts.BattleRound, content string) {

}

func PrintReport() {

}
