package util

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/spf13/cast"
)

type ResumeParams struct {
	Ctx context.Context
	//战法上下文
	TacticsParams *model.TacticsParams
	//施加者
	ProduceGeneral *vo.BattleGeneral
	//被施加者
	SufferGeneral *vo.BattleGeneral
	//恢复量
	ResumeNum int64
	//战法ID
	TacticId consts.TacticId
}

// 恢复兵力结算
// @general 当前武将
// @resumeNum 恢复兵力
func ResumeSoldierNum(param *ResumeParams) (finalResumeNum, originNum, finalSoldierNum int64) {
	//参数校验
	if param.ResumeNum == 0 {
		return param.ResumeNum, param.SufferGeneral.SoldierNum, param.SufferGeneral.SoldierNum
	}
	if param.ProduceGeneral == nil {
		panic("ProduceGeneral is nil")
	}
	if param.SufferGeneral == nil {
		panic("SufferGeneral is nil")
	}
	if param.TacticsParams == nil {
		panic("TacticsParams is nil")
	}
	if param.TacticId == 0 {
		panic("TacticId is nil")
	}

	//效果判定
	if !IsCanResume(param.Ctx, param.SufferGeneral) {
		return 0, param.SufferGeneral.SoldierNum, param.SufferGeneral.SoldierNum
	}

	if param.SufferGeneral.SoldierNum == 0 {
		return 0, param.SufferGeneral.SoldierNum, param.SufferGeneral.SoldierNum
	}
	if param.SufferGeneral.LossSoldierNum == 0 {
		return 0, param.SufferGeneral.SoldierNum, param.SufferGeneral.SoldierNum
	}

	//恢复触发器
	if funcs, okk := param.SufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferResume]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				SufferResumeGeneral: param.SufferGeneral,
				CurrentResume:       param.ResumeNum,
			}
			f(params)
		}
	}

	//治疗效果提升
	if effectParams, ok := BuffEffectGet(param.ProduceGeneral, consts.BuffEffectType_ResumeImprove); ok {
		effectRate := float64(0)
		for _, param := range effectParams {
			effectRate += param.EffectRate
		}
		param.ResumeNum = cast.ToInt64(cast.ToFloat64(param.ResumeNum) * (1 + effectRate))
	}

	//治��效果降低

	//治疗效果降低
	if effectParams, ok := DeBuffEffectGet(param.SufferGeneral, consts.DebuffEffectType_SufferResumeDeduce); ok {
		effectRate := float64(0)
		for _, param := range effectParams {
			effectRate += param.EffectRate
		}
		param.ResumeNum = cast.ToInt64(cast.ToFloat64(param.ResumeNum) * (1 - effectRate))
	}

	originNum = param.SufferGeneral.SoldierNum
	//恢复结果大于最大带兵量时
	if param.SufferGeneral.SoldierNum+param.ResumeNum > consts.Max_Soldiers_Num_Per_General {
		finalResumeNum = consts.Max_Soldiers_Num_Per_General - param.SufferGeneral.SoldierNum
		finalSoldierNum = consts.Max_Soldiers_Num_Per_General

		param.SufferGeneral.SoldierNum = consts.Max_Soldiers_Num_Per_General
	} else {
		finalResumeNum = param.ResumeNum
		finalSoldierNum = param.SufferGeneral.SoldierNum + param.ResumeNum

		param.SufferGeneral.SoldierNum += param.ResumeNum
	}

	//统计
	param.ProduceGeneral.TacticAccumulateTriggerMap[param.TacticId] += 1
	param.ProduceGeneral.TacticAccumulateResumeMap[param.TacticId] = param.ResumeNum
	param.ProduceGeneral.AccumulateTotalResumeNum += param.ResumeNum

	TacticReport(param.TacticsParams,
		param.ProduceGeneral.BaseInfo.UniqueId,
		int64(param.TacticId),
		1,
		0,
		finalResumeNum,
	)

	hlog.CtxInfof(param.Ctx, "[%s]恢复了兵力%d(%d↗%d)",
		param.SufferGeneral.BaseInfo.Name,
		finalResumeNum,
		originNum,
		finalSoldierNum,
	)

	return
}
