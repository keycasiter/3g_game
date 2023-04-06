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
func TacticsTriggerWrapSet(general *vo.BattleGeneral, action consts.BattleAction, f func(params *vo.TacticsTriggerParams)) {
	if funcs, ok := general.TacticsTriggerMap[action]; ok {
		funcs = append(funcs, f)
	} else {
		fs := make([]func(params *vo.TacticsTriggerParams), 0)
		fs = append(fs, f)
		general.TacticsTriggerMap[action] = fs
	}
}

// 战法增益次数设置
func TacticsBuffCountWrapSet(general *vo.BattleGeneral, buffEffect consts.BuffEffectType, cnt int64, rate float64) {
	if mm, ok := general.BuffEffectCountMap[buffEffect]; ok {
		//取出原数量
		for k, _ := range mm {
			//叠加
			cnt += k
		}
		//删除原存储mm
		delete(general.BuffEffectCountMap, buffEffect)
		//生成新mm
		newMm := make(map[int64]float64, 0)
		//刷新
		newMm[cnt] = rate
		general.BuffEffectCountMap[buffEffect] = newMm
	} else {
		newMm := make(map[int64]float64, 0)
		newMm[cnt] = rate
		general.BuffEffectCountMap[buffEffect] = newMm
	}
}

// 战法减益次数设置
func TacticsDebuffCountWrapSet(general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType, cnt int64, rate float64) {
	if mm, ok := general.DeBuffEffectCountMap[debuffEffect]; ok {
		//取出原数量
		for k, _ := range mm {
			//叠加
			cnt += k
		}
		//删除原存储mm
		delete(general.DeBuffEffectCountMap, debuffEffect)
		//生成新mm
		newMm := make(map[int64]float64, 0)
		//刷新
		newMm[cnt] = rate
		general.DeBuffEffectCountMap[debuffEffect] = newMm
	} else {
		newMm := make(map[int64]float64, 0)
		newMm[cnt] = rate
		general.DeBuffEffectCountMap[debuffEffect] = newMm
	}
}

// 战法减益次数消耗
func TacticsDebuffCountCost(general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType, costNum int64) {
	oldNum := int64(0)
	rate := float64(0)
	//兼容消耗数量
	if costNum > oldNum {
		costNum = oldNum
	}

	if mm, ok := general.DeBuffEffectCountMap[debuffEffect]; ok {
		//取出原数量
		for k, v := range mm {
			oldNum = k
			rate = v
		}
		//删除原存储mm
		delete(general.DeBuffEffectCountMap, debuffEffect)
		//生成新mm
		newMm := make(map[int64]float64, 0)
		//刷新
		newMm[oldNum-costNum] = rate
		general.DeBuffEffectCountMap[debuffEffect] = newMm
	} else {
		newMm := make(map[int64]float64, 0)
		newMm[oldNum-costNum] = rate
		general.DeBuffEffectCountMap[debuffEffect] = newMm
	}
}

// 战法增益次数查询
func TacticsBuffCountGet(general *vo.BattleGeneral, buffEffect consts.BuffEffectType) bool {
	if mm, ok := general.BuffEffectCountMap[buffEffect]; ok {
		//0次
		if _, okk := mm[0]; okk {
			return false
		}
		//>0次
		return true
	}
	//不存在效果
	return false
}

// 战法减益次数查询
func TacticsDebuffCountGet(general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType) bool {
	if mm, ok := general.DeBuffEffectCountMap[debuffEffect]; ok {
		//0次
		if _, okk := mm[0]; okk {
			return false
		}
		//>0次
		return true
	}
	//不存在效果
	return false
}
