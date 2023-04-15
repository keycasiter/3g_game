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
func BuffEffectWrapSet(ctx context.Context, general *vo.BattleGeneral, effectType consts.BuffEffectType, v float64) bool {
	if _, ok := general.BuffEffectHolderMap[effectType]; ok {
		hlog.CtxInfof(ctx, "[%s]身上已有同等或更强的「%v」效果",
			general.BaseInfo.Name,
			effectType,
		)
		return false
	}

	hlog.CtxInfof(ctx, "[%s]的「%v」已施加",
		general.BaseInfo.Name,
		effectType,
	)
	general.BuffEffectHolderMap[effectType] = v
	return true
}

func BuffEffectWrapRemove(general *vo.BattleGeneral, effectType consts.BuffEffectType) bool {
	if _, ok := general.BuffEffectHolderMap[effectType]; ok {
		delete(general.BuffEffectHolderMap, effectType)
		delete(general.BuffEffectCountMap, effectType)
		return true
	}
	return false
}

func BuffEffectContains(general *vo.BattleGeneral, effectType consts.BuffEffectType) bool {
	if _, ok := general.BuffEffectHolderMap[effectType]; ok {
		return true
	}
	return false
}

// 减益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
func DebuffEffectWrapSet(ctx context.Context, general *vo.BattleGeneral, effectType consts.DebuffEffectType, v float64) bool {
	if _, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		hlog.CtxInfof(ctx, "[%s]身上已有同等或更强的「%v」效果",
			general.BaseInfo.Name,
			effectType,
		)
		return false
	}

	hlog.CtxInfof(ctx, "[%s]的「%v」已施加",
		general.BaseInfo.Name,
		effectType,
	)
	general.DeBuffEffectHolderMap[effectType] = v
	return true
}

func DebuffEffectWrapRemove(ctx context.Context, general *vo.BattleGeneral, effectType consts.DebuffEffectType) bool {
	if _, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		delete(general.DeBuffEffectHolderMap, effectType)
		delete(general.DeBuffEffectCountMap, effectType)

		hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
			general.BaseInfo.Name,
			effectType,
		)
		return true
	}
	return false
}

func DeBuffEffectContains(general *vo.BattleGeneral, effectType consts.DebuffEffectType) bool {
	if _, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		return true
	}
	return false
}

func TacticFrozenWrapSet(general *vo.BattleGeneral, tacticId consts.TacticId, frozenNum int64, maxNum int64, supportRefresh bool) bool {
	holdNum := int64(0)
	if v, ok := general.TacticsFrozenMap[tacticId]; ok {
		//v最少为1才算刷新效果
		if supportRefresh && frozenNum <= maxNum {
			general.TacticsFrozenMap[tacticId] = frozenNum
			return false
		}

		holdNum = v
	}
	//超限
	totalNum := holdNum + frozenNum
	if totalNum > maxNum {
		return false
	}
	//更新
	general.TacticsFrozenMap[tacticId] = totalNum

	return true
}

func TacticFrozenWrapRemove(general *vo.BattleGeneral, tacticId consts.TacticId) bool {
	if _, ok := general.TacticsFrozenMap[tacticId]; ok {
		delete(general.TacticsFrozenMap, tacticId)
		return true
	}
	return false
}

// 增益效果清除
// @general 要处理的武将
func BuffEffectClean(ctx context.Context, general *vo.BattleGeneral) {
	for effectType, _ := range general.BuffEffectHolderMap {
		//只能清除主动、突击战法效果
		if _, ok := consts.SupprtCleanBuffEffectMap[effectType]; !ok {
			continue
		}
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
		//只能清除主动、突击战法效果
		if _, ok := consts.SupprtCleanDebuffEffectMap[effectType]; !ok {
			continue
		}
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

// 增益效果次数容器处理
// @holder 效果容器
// @effectType 效果类型
// @cnt 当前次数
// @v 效果值
func TacticsBuffEffectCountWrapIncr(ctx context.Context, general *vo.BattleGeneral, buffEffect consts.BuffEffectType, incrNum int64, maxNum int64, supportRefresh bool) bool {
	holdNum := int64(0)
	if v, ok := general.BuffEffectCountMap[buffEffect]; ok {
		//v最少为1才算刷新效果
		if supportRefresh && incrNum <= maxNum {
			general.BuffEffectCountMap[buffEffect] = incrNum
			hlog.CtxInfof(ctx, "[%s]的「%v」的效果已刷新",
				general.BaseInfo.Name,
				buffEffect,
			)
			return false
		}

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
func TacticsBuffEffectCountWrapDecr(ctx context.Context, general *vo.BattleGeneral, buffEffect consts.BuffEffectType, decrNum int64) bool {
	if _, ok := general.BuffEffectHolderMap[buffEffect]; !ok {
		return false
	}

	holdNum := general.BuffEffectCountMap[buffEffect]
	if holdNum == 0 {
		delete(general.BuffEffectHolderMap, buffEffect)
		delete(general.BuffEffectCountMap, buffEffect)
		hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
			general.BaseInfo.Name,
			buffEffect,
		)
		return false
	}

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
func TacticsDebuffEffectCountWrapIncr(ctx context.Context, general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType, incrNum int64, maxNum int64, supportRefresh bool) bool {
	holdNum := int64(0)
	if v, ok := general.DeBuffEffectCountMap[debuffEffect]; ok {
		//v最少为1才算刷新效果
		if supportRefresh && incrNum <= maxNum {
			general.DeBuffEffectCountMap[debuffEffect] = incrNum
			hlog.CtxInfof(ctx, "[%s]的「%v」的效果已刷新",
				general.BaseInfo.Name,
				debuffEffect,
			)
			return false
		}

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
func TacticsDebuffEffectCountWrapDecr(ctx context.Context, general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType, decrNum int64) bool {
	if _, ok := general.DeBuffEffectHolderMap[debuffEffect]; !ok {
		return false
	}

	holdNum := general.DeBuffEffectCountMap[debuffEffect]
	if holdNum == 0 {
		delete(general.DeBuffEffectCountMap, debuffEffect)
		delete(general.DeBuffEffectHolderMap, debuffEffect)
		hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
			general.BaseInfo.Name,
			debuffEffect,
		)
		return false
	}

	//消费次数不足
	if decrNum > holdNum {
		hlog.CtxWarnf(ctx, "效果消费次数不足，expect:%d , real:%d", decrNum, holdNum)
		return false
	}

	general.DeBuffEffectCountMap[debuffEffect]--
	return true
}

// 战法增益次数查询
func TacticsBuffCountGet(general *vo.BattleGeneral, buffEffect consts.BuffEffectType) int64 {
	if cnt, ok := general.BuffEffectCountMap[buffEffect]; ok {
		if cnt == 0 {
			delete(general.BuffEffectCountMap, buffEffect)
		}
		return cnt
	}
	return 0
}

// 战法减益次数查询
func TacticsDebuffCountGet(general *vo.BattleGeneral, debuffEffect consts.DebuffEffectType) int64 {
	if cnt, ok := general.DeBuffEffectCountMap[debuffEffect]; ok {
		if cnt == 0 {
			delete(general.DeBuffEffectCountMap, debuffEffect)
		}
		return cnt
	}
	return 0
}
