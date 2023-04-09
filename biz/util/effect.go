package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"reflect"
)

// 增益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
func BuffEffectWrapSet(general *vo.BattleGeneral, effectType consts.BuffEffectType, v float64) {
	general.BuffEffectHolderMap[effectType] = v
}

// 减益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
func DebuffEffectWrapSet(general *vo.BattleGeneral, effectType consts.DebuffEffectType, v float64) {
	general.DeBuffEffectHolderMap[effectType] = v
}

// 增益效果清除
// @general 要处理的武将
func BuffEffectClean(ctx context.Context, general *vo.BattleGeneral) {
	for effectType, _ := range general.BuffEffectHolderMap {
		hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
			general.BaseInfo.Name,
			effectType,
		)
	}
	general.DeBuffEffectCountMap = map[consts.DebuffEffectType]int64{}
	general.DeBuffEffectHolderMap = map[consts.DebuffEffectType]float64{}
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
	general.DeBuffEffectCountMap = map[consts.DebuffEffectType]int64{}
	general.DeBuffEffectHolderMap = map[consts.DebuffEffectType]float64{}
}

// 战法触发器注册
func TacticsTriggerWrapRegister(general *vo.BattleGeneral, action consts.BattleAction, newFunc func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult) {
	if funcs, ok := general.TacticsTriggerMap[action]; ok {
		funcs = append(funcs, newFunc)
		general.TacticsTriggerMap[action] = funcs
	} else {
		fs := make([]func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult, 0)
		fs = append(fs, newFunc)
		general.TacticsTriggerMap[action] = fs
	}
}

// 战法触发器注销
func TacticsTriggerWrapUnregister(general *vo.BattleGeneral, action consts.BattleAction, removeFunc func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult) {
	if funcs, okk := general.TacticsTriggerMap[action]; okk {
		newFuncs := make([]func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult, 0)
		for _, f := range funcs {
			if reflect.DeepEqual(removeFunc, f) {
				hlog.CtxDebugf(context.Background(), "判断是重复 %v , %v", removeFunc, f)
				continue
			} else {
				newFuncs = append(newFuncs, f)
			}
		}
		general.TacticsTriggerMap[action] = newFuncs
	}
}

// 增益效果次数容器处理
// @holder 效果容器
// @effectType 效果类型
// @cnt 当前次数
// @v 效果值
func TacticsBuffEffectCountWrapIncr(general *vo.BattleGeneral, buffEffect consts.BuffEffectType, incrNum int64, maxNum int64) bool {
	holdNum := int64(0)
	if v, ok := general.BuffEffectCountMap[buffEffect]; ok {
		holdNum = v
	}
	//超限
	totalNum := holdNum + incrNum
	if totalNum > maxNum {
		return false
	}
	//更新
	general.BuffEffectCountMap[buffEffect] = totalNum

	return true
}

// 增益效果次数容器处理
// @holder 效果容器
// @effectType 效果类型
// @cnt 当前次数
// @v 效果值
func TacticsBuffEffectCountWrapDecr(general *vo.BattleGeneral, buffEffect consts.BuffEffectType, decrNum int64) bool {
	holdNum := general.BuffEffectCountMap[buffEffect]
	//消费次数不足
	if decrNum > holdNum {
		return false
	}

	general.BuffEffectCountMap[buffEffect]--
	return true
}

// 增益效果次数容器处理
// @holder 效果容器
// @effectType 效果类型
// @cnt 当前次数
// @v 效果值
func TacticsDebuffEffectCountWrapIncr(general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType, incrNum int64, maxNum int64) bool {
	holdNum := int64(0)
	if v, ok := general.DeBuffEffectCountMap[debuffEffect]; ok {
		holdNum = v
	}
	//超限
	totalNum := holdNum + incrNum
	if totalNum > maxNum {
		return false
	}
	//更新
	general.DeBuffEffectCountMap[debuffEffect] = totalNum

	return true
}

// 战法减益次数消耗
// @general 执行武将
// @debuffEffect 减益效果
// @costNum 消耗次数
func TacticsDebuffEffectCountWrapDecr(general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType, decrNum int64) bool {
	holdNum := general.DeBuffEffectCountMap[debuffEffect]
	//消费次数不足
	if decrNum > holdNum {
		return false
	}

	general.DeBuffEffectCountMap[debuffEffect]--
	return true
}

// 战法增益次数查询
func TacticsBuffCountGet(general *vo.BattleGeneral, buffEffect consts.BuffEffectType) int64 {
	return general.BuffEffectCountMap[buffEffect]
}

// 战法减益次数查询
func TacticsDebuffCountGet(general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType) int64 {
	return general.DeBuffEffectCountMap[debuffEffect]
}
