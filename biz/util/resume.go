package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/spf13/cast"
)

// 恢复兵力结算
// @general 当前武将
// @resumeNum 恢复兵力
func ResumeSoldierNum(ctx context.Context, general *vo.BattleGeneral, resumeNum int64) (finalResumeNum, originNum, finalSoldierNum int64) {
	if !IsCanResume(ctx, general) {
		return 0, general.SoldierNum, general.SoldierNum
	}

	if general.SoldierNum == 0 {
		return 0, general.SoldierNum, general.SoldierNum
	}
	if general.LossSoldierNum == 0 {
		return 0, general.SoldierNum, general.SoldierNum
	}

	//恢复触发器
	if funcs, okk := general.TacticsTriggerMap[consts.BattleAction_SufferResume]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				SufferResumeGeneral: general,
				CurrentResume:       resumeNum,
			}
			f(params)
		}
	}

	//治疗效果降低
	if effectParams, ok := DeBuffEffectGet(general, consts.DebuffEffectType_SufferResumeDeduce); ok {
		effectRate := float64(0)
		for _, param := range effectParams {
			effectRate += param.EffectRate
		}
		resumeNum = cast.ToInt64(cast.ToFloat64(resumeNum) * (1 - effectRate))
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

	hlog.CtxInfof(ctx, "[%s]恢复了兵力%d(%d↗%d)",
		general.BaseInfo.Name,
		finalResumeNum,
		originNum,
		finalSoldierNum,
	)

	return
}
