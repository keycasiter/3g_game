package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

/*
	属性状态：增加或降低武将各种属性；来自不同战法的同种属性状态可以叠加；来自同一战法的同种属性状态将会刷新持续回合；
			武力/智力/统率/速度/政治/美丽/会心几率/奇谋几率/发动几率/战法造成伤害/受到战法伤害

	持续性状态：每回合武将开始行动时，对武将造成伤害或治疗；同种状态不可叠加，但会刷新

	功能性状态：通常不可叠加，不同来源时不可刷新

	控制状态：不可叠加，不可刷新，负面效果
*/

var (
	//属性状态
	attrBuffEffectMap = map[consts.BuffEffectType]bool{
		consts.BuffEffectType_LaunchStrategyDamageImprove: true,
		consts.BuffEffectType_LaunchWeaponDamageImprove:   true,
	}
	//持续性状态
	continuousDebuffEffectMap = map[consts.DebuffEffectType]bool{
		//灼烧
		consts.DebuffEffectType_Firing: true,
		//水攻
		consts.DebuffEffectType_WaterAttack: true,
		//中毒
		consts.DebuffEffectType_Methysis: true,
		//溃逃
		consts.DebuffEffectType_Escape: true,
		//沙暴
		consts.DebuffEffectType_Sandstorm: true,
		//叛逃
		consts.DebuffEffectType_Defect: true,
	}
	continuousBuffEffectMap = map[consts.BuffEffectType]bool{
		//急救
		consts.BuffEffectType_EmergencyTreatment: true,
		//休整
		consts.BuffEffectType_Rest: true,
	}
	//功能性状态
	functionBuffEffectMap = map[consts.BuffEffectType]bool{
		//急救
		consts.BuffEffectType_EmergencyTreatment: true,
		//休整
		consts.BuffEffectType_Rest: true,
	}
	//控制状态
	controlDebuffEffectMap = map[consts.DebuffEffectType]bool{
		//震慑
		consts.DebuffEffectType_Awe: true,
		//计穷
		consts.DebuffEffectType_NoStrategy: true,
		//缴械
		consts.DebuffEffectType_CancelWeapon: true,
		//混乱
		consts.DebuffEffectType_Chaos: true,
		//虚弱
		consts.DebuffEffectType_PoorHealth: true,
		//禁疗
		consts.DebuffEffectType_ProhibitionTreatment: true,
		//嘲讽
		consts.DebuffEffectType_Taunt: true,
		//伪报
		consts.DebuffEffectType_FalseReport: true,
		//挑拨
		consts.DebuffEffectType_Provoking: true,
		//破坏
		consts.DebuffEffectType_Break: true,
		//捕获
		consts.DebuffEffectType_Capture: true,
	}
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

	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		general.BaseInfo.Name,
		effectType,
	)
	general.BuffEffectHolderMap[effectType] = v
	return true
}

func BuffEffectWrapRemove(ctx context.Context, general *vo.BattleGeneral, effectType consts.BuffEffectType) bool {
	if _, ok := general.BuffEffectHolderMap[effectType]; ok {
		delete(general.BuffEffectHolderMap, effectType)
		delete(general.BuffEffectCountMap, effectType)

		hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
			general.BaseInfo.Name,
			effectType,
		)
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

func BuffEffectGet(general *vo.BattleGeneral, effectType consts.BuffEffectType) (float64, bool) {
	if v, ok := general.BuffEffectHolderMap[effectType]; ok {
		return v, true
	}
	return 0, false
}

func BuffEffectContainsCheck(general *vo.BattleGeneral) bool {
	return len(general.BuffEffectHolderMap) > 0
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

	//嘲讽效果处理
	if effectType == consts.DebuffEffectType_Taunt {
		//是否有洞察
		if BuffEffectContains(general, consts.BuffEffectType_Insight) {
			hlog.CtxInfof(ctx, "[%s]由于「%v」的效果，「%v」对其无效",
				general.BaseInfo.Name,
				consts.BuffEffectType_Insight,
				consts.DebuffEffectType_Taunt,
			)
			return false
		}
	}

	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
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

func DeBuffEffectGet(general *vo.BattleGeneral, effectType consts.DebuffEffectType) (float64, bool) {
	if v, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		return v, true
	}
	return 0, false
}

func DeBuffEffectContainsCheck(general *vo.BattleGeneral) bool {
	return len(general.DeBuffEffectHolderMap) > 0
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
