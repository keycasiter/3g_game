package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

// 增益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
func BuffEffectWrapSet(holder map[consts.BuffEffectType]map[consts.BattleRound]float64,
	effectType consts.BuffEffectType,
	round consts.BattleRound,
	v float64) {
	mm := make(map[consts.BattleRound]float64, 0)
	mm[round] = v
	holder[effectType] = mm
}

// 增益效果次数容器处理
// @holder 效果容器
// @effectType 效果类型
// @cnt 当前次数
// @v 效果值
func BuffEffectCountWrapAdd(holder map[consts.BuffEffectType]map[int64]float64,
	effectType consts.BuffEffectType,
	cnt int64,
	v float64) {
	mm := make(map[int64]float64, 0)
	mm[cnt] += v
	holder[effectType] = mm
}

// 减益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
func DebuffEffectWrapSet(holder map[consts.DebuffEffectType]map[consts.BattleRound]float64,
	effectType consts.DebuffEffectType,
	round consts.BattleRound,
	v float64) {
	mm := make(map[consts.BattleRound]float64, 0)
	mm[round] = v
	holder[effectType] = mm
}

// 负面效果清除
// @general 要处理的武将
func DebuffEffectClean(ctx context.Context, general *vo.BattleGeneral) {
	for effectType, _ := range general.DeBuffEffectHolderMap {
		hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
			general.BaseInfo.Name,
			effectType,
		)
	}
	general.DeBuffEffectCountMap = map[consts.DebuffEffectType]map[int64]float64{}
	general.DeBuffEffectTriggerMap = map[consts.DebuffEffectType]map[consts.BattleRound]float64{}
	general.DeBuffEffectHolderMap = map[consts.DebuffEffectType]float64{}
}

// 战法触发器设置
func TacticsTriggerWrapSet(general *vo.BattleGeneral, action consts.BattleAction, f func(params vo.TacticsTriggerParams)) {
	if funcs, ok := general.TacticsTriggerMap[action]; ok {
		funcs = append(funcs, f)
	} else {
		fs := make([]func(params vo.TacticsTriggerParams), 0)
		fs = append(fs, f)
		general.TacticsTriggerMap[action] = fs
	}
}

// 战法增益次数设置
func TacticsBuffCountWrapSet(general *vo.BattleGeneral, buffEffect consts.BuffEffectType, cnt int64, rate float64) {
	if mm, ok := general.BuffEffectCountMap[buffEffect]; ok {
		mm[cnt] = rate
		general.BuffEffectCountMap[buffEffect] = mm
	} else {
		newMm := make(map[int64]float64, 0)
		newMm[cnt] = rate
		general.BuffEffectCountMap[buffEffect] = newMm
	}
}

// 战法减益次数设置
func TacticsDebuffCountWrapSet(general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType, cnt int64, rate float64) {
	if mm, ok := general.DeBuffEffectCountMap[debuffEffect]; ok {
		mm[cnt] = rate
		general.DeBuffEffectCountMap[debuffEffect] = mm
	} else {
		newMm := make(map[int64]float64, 0)
		newMm[cnt] = rate
		general.DeBuffEffectCountMap[debuffEffect] = newMm
	}
}
