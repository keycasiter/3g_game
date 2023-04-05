package util

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

//恢复兵力结算
//@general 当前武将
//@resumeNum 恢复兵力
func ResumeSoldierNum(general *vo.BattleGeneral, resumeNum int64) (finalResumeNum, originNum, finalSoldierNum int64) {
	if general.SoldierNum == 0 {
		return 0, general.SoldierNum, general.SoldierNum
	}
	if general.LossSoldierNum == 0 {
		return 0, general.SoldierNum, general.SoldierNum
	}

	originNum = general.SoldierNum
	//恢复结果大于最大带兵量时
	if general.SoldierNum+resumeNum > consts.Max_Soldiers_Num_Per_General {
		finalResumeNum = consts.Max_Soldiers_Num_Per_General - general.SoldierNum
		finalSoldierNum = consts.Max_Soldiers_Num_Per_General

		general.SoldierNum = consts.Max_Soldiers_Num_Per_General
	} else {
		finalResumeNum = resumeNum
		finalSoldierNum = general.SoldierNum + resumeNum

		general.SoldierNum += resumeNum
	}

	return
}
