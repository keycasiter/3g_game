package util

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type AppendBattleReportReq struct {
	//对战内含参数
	TacticsParams *model.TacticsParams
	//上报的对战内容
	ReportContent string
}

// 追加对战报告
func AppendBattleReport(req *AppendBattleReportReq) {
	//1.参数校验
	if req.TacticsParams == nil {
		panic(any("AppendBattleReport TacticsParams is nil"))
	}
	if req.TacticsParams.BattleProcessStatisticsMap == nil {
		panic(any("AppendBattleReport TacticsParams BattleReports is nil"))
	}

	//2.追加战报处理
	//战报对象
	battleProcessStatisticsMap := req.TacticsParams.BattleProcessStatisticsMap
	//当前对战阶段
	currentPhase := req.TacticsParams.CurrentPhase
	//当前对战回合
	currentRound := req.TacticsParams.CurrentRound

	//判断对战阶段
	if roundReportMap, ok := battleProcessStatisticsMap[currentPhase]; ok {
		if reports, okk := roundReportMap[currentRound]; okk {
			reports = append(reports, req.ReportContent)
		} else {
			//初始化回合map
			roundReportMap[currentRound] = []string{req.ReportContent}
		}
	} else {
		//初始化阶段map & 回合map
		battleProcessStatisticsMap[currentPhase] = map[consts.BattleRound][]string{
			currentRound: {req.ReportContent},
		}
	}
}

func PrintBattleReport() {

}
