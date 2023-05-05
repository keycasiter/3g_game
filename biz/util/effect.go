package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
)

/*	属性状态：增加或降低武将各种属性；来自不同战法的同种属性状态可以叠加；来自同一战法的同种属性状态将会刷新持续回合；
			武力/智力/统率/速度/政治/美丽/会心几率/奇谋几率/发动几率/战法造成伤害/受到战法伤害
	持续性状态：每回合武将开始行动时，对武将造成伤害或治疗；同种状态不可叠加，但会刷新
	功能性状态：通常不可叠加，不同来源时不可刷新
	控制状态：不可叠加，不可刷新，负面效果
*/

var (
	//**属性状态**
	attrBuffEffectMap = map[consts.BuffEffectType]bool{
		consts.BuffEffectType_LaunchStrategyDamageImprove: true,
		consts.BuffEffectType_LaunchWeaponDamageImprove:   true,
		consts.BuffEffectType_SufferStrategyDamageDeduce:  true,
		consts.BuffEffectType_SufferWeaponDamageDeduce:    true,
	}

	attrDebuffEffectMap = map[consts.DebuffEffectType]bool{
		consts.DebuffEffectType_SufferWeaponDamageImprove:   true,
		consts.DebuffEffectType_SufferStrategyDamageImprove: true,
		consts.DebuffEffectType_LaunchWeaponDamageDeduce:    true,
		consts.DebuffEffectType_LaunchStrategyDamageDeduce:  true,
	}
	//**持续性状态**
	//不可叠加，但会刷新回合
	//灼烧、水攻、中毒、溃逃、沙暴、叛逃
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
		//群攻
		consts.BuffEffectType_GroupAttack: true,
		//反击
		consts.BuffEffectType_StrikeBack: true,
		//倒戈
		consts.BuffEffectType_Defection: true,
		//会心
		consts.BuffEffectType_EnhanceWeapon: true,
		//奇谋
		consts.BuffEffectType_EnhanceStrategy: true,
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
		//无法普通攻击
		consts.DebuffEffectType_CanNotGeneralAttack: true,
	}
)

// 增益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
//  属性状态：增加或降低武将各种属性；来自不同战法的同种属性状态可以叠加；来自同一战法的同种属性状态将会刷新持续回合；
//			武力/智力/统率/速度/政治/美丽/会心几率/奇谋几率/发动几率/战法造成伤害/受到战法伤害
//	持续性状态：每回合武将开始行动时，对武将造成伤害或治疗；同种状态不可叠加，但会刷新
//	功能性状态：通常不可叠加，不同来源时不可刷新
//	控制状态：不可叠加，不可刷新，负面效果
func BuffEffectWrapSet(ctx context.Context, general *vo.BattleGeneral, effectType consts.BuffEffectType, effectParam *vo.EffectHolderParams) bool {
	//是否包含正面效果判断
	_, isContainBuffEffect := general.BuffEffectHolderMap[effectType]
	//是否包含属性正面效果判断
	_, isContainAttrBuffEffect := attrBuffEffectMap[effectType]
	//是否包含持续性正面效果判断
	_, isContainContinuousBuffEffect := continuousBuffEffectMap[effectType]

	//属性正面状态
	// 增加或降低武将各种属性；来自不同战法的同种属性状态可以叠加；来自同一战法的同种属性状态将会刷新持续回合；
	// 武力/智力/统率/速度/政治/美丽/会心几率/奇谋几率/发动几率/战法造成伤害/受到战法伤害
	if isContainBuffEffect && isContainAttrBuffEffect {
		//来自同一战法的属性效果
		for _, effectHolderParam := range general.BuffEffectHolderMap[effectType] {
			if effectHolderParam.FromTactic == effectParam.FromTactic {
				effectHolderParam.EffectRound = effectParam.EffectRound
				hlog.CtxInfof(ctx, "[%s]来自【%v】「%v」效果已刷新",
					general.BaseInfo.Name,
					effectParam.FromTactic,
					effectType,
				)
				return true
			}
		}
		//不来自同一战法的属性效果
		if effectHolderParams, ok := general.BuffEffectHolderMap[effectType]; ok {
			effectHolderParams = append(effectHolderParams, effectParam)
			hlog.CtxInfof(ctx, "[%s]来自【%v】「%v」效果已施加",
				general.BaseInfo.Name,
				effectParam.FromTactic,
				effectType,
			)
			return true
		}
	}

	//持续性负面效果：不可叠加，刷新回合
	if isContainBuffEffect && isContainContinuousBuffEffect {
		for _, effectHolderParam := range general.BuffEffectHolderMap[effectType] {
			effectHolderParam.EffectRound = effectParam.EffectRound
			hlog.CtxInfof(ctx, "[%s]来自【%v】「%v」效果已刷新",
				general.BaseInfo.Name,
				effectParam.FromTactic,
				effectType,
			)
			return true
		}
	}

	//施加效果
	if isContainBuffEffect {
		general.BuffEffectHolderMap[effectType] = append(general.BuffEffectHolderMap[effectType], effectParam)
	} else {
		general.BuffEffectHolderMap[effectType] = []*vo.EffectHolderParams{
			effectParam,
		}
	}
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		general.BaseInfo.Name,
		effectType,
	)
	return true
}

// 移除增益效果
func BuffEffectWrapRemove(ctx context.Context, general *vo.BattleGeneral, effectType consts.BuffEffectType, tacticId consts.TacticId) bool {
	if effectParams, ok := general.BuffEffectHolderMap[effectType]; ok {
		idx := 0
		for _, effectParam := range effectParams {
			if effectParam.FromTactic == tacticId {
				general.BuffEffectHolderMap[effectType] = append(effectParams[:idx], effectParams[idx+1:]...)
				hlog.CtxInfof(ctx, "[%s]的【%v】「%v」效果已消失",
					general.BaseInfo.Name,
					effectParam.FromTactic,
					effectType,
				)
				return true
			}
			idx++
		}
	}
	return false
}

// 是否包含某个增益效果
func BuffEffectContains(general *vo.BattleGeneral, effectType consts.BuffEffectType) bool {
	if _, ok := general.BuffEffectHolderMap[effectType]; ok {
		return true
	}
	return false
}

// 获取增益效果
func BuffEffectGet(general *vo.BattleGeneral, effectType consts.BuffEffectType) ([]*vo.EffectHolderParams, bool) {
	if v, ok := general.BuffEffectHolderMap[effectType]; ok {
		return v, true
	}
	return nil, false
}

// 是否存在增益效果
func BuffEffectContainsCheck(general *vo.BattleGeneral) bool {
	return len(general.BuffEffectHolderMap) > 0
}

// 减益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
//  属性状态：增加或降低武将各种属性；来自不同战法的同种属性状态可以叠加；来自同一战法的同种属性状态将会刷新持续回合；
//			武力/智力/统率/速度/政治/美丽/会心几率/奇谋几率/发动几率/战法造成伤害/受到战法伤害
//	持续性状态：每回合武将开始行动时，对武将造成伤害或治疗；同种状态不可叠加，但会刷新
//	功能性状态：通常不可叠加，不同来源时不可刷新
//	控制状态：不可叠加，不可刷新，负面效果
func DebuffEffectWrapSet(ctx context.Context, general *vo.BattleGeneral, effectType consts.DebuffEffectType, effectParam *vo.EffectHolderParams) bool {
	//是否包含负面效果判断
	_, isContainDebuffEffect := general.DeBuffEffectHolderMap[effectType]
	//是否包含属性负面效果判断
	_, isContainAttrDebuffEffect := attrDebuffEffectMap[effectType]
	//是否包含持续性负面效果判断
	_, isContainContinuousDebuffEffect := continuousDebuffEffectMap[effectType]
	//是否包含控制负面效果判断
	_, isContainControlDebuffEffect := controlDebuffEffectMap[effectType]

	//属性负面状态
	// 增加或降低武将各种属性；来自不同战法的同种属性状态可以叠加；来自同一战法的同种属性状态将会刷新持续回合；
	// 武力/智力/统率/速度/政治/美丽/会心几率/奇谋几率/发动几率/战法造成伤害/受到战法伤害
	if isContainDebuffEffect && isContainAttrDebuffEffect {
		//来自同一战法的属性效果
		for _, effectHolderParam := range general.DeBuffEffectHolderMap[effectType] {
			if effectHolderParam.FromTactic == effectParam.FromTactic {
				effectHolderParam.EffectRound = effectParam.EffectRound
				hlog.CtxInfof(ctx, "[%s]来自【%v】「%v」效果已刷新",
					general.BaseInfo.Name,
					effectParam.FromTactic,
					effectType,
				)
				return true
			}
		}
		//不来自同一战法的属性效果
		if effectHolderParams, ok := general.DeBuffEffectHolderMap[effectType]; ok {
			effectHolderParams = append(effectHolderParams, effectParam)
			hlog.CtxInfof(ctx, "[%s]来自【%v】「%v」效果已施加",
				general.BaseInfo.Name,
				effectParam.FromTactic,
				effectType,
			)
			return true
		}
	}

	//持续性负面效果：不可叠加，刷新回合
	if isContainDebuffEffect && isContainContinuousDebuffEffect {
		for _, effectHolderParam := range general.DeBuffEffectHolderMap[effectType] {
			effectHolderParam.EffectRound = effectParam.EffectRound
			hlog.CtxInfof(ctx, "[%s]来自【%v】「%v」效果已刷新",
				general.BaseInfo.Name,
				effectParam.FromTactic,
				effectType,
			)
			return true
		}
	}

	//控制状态：不可叠加，不可刷新，负面效果
	if isContainDebuffEffect && isContainControlDebuffEffect {
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

	//施加效果
	if isContainDebuffEffect {
		general.DeBuffEffectHolderMap[effectType] = append(general.DeBuffEffectHolderMap[effectType], effectParam)
	} else {
		general.DeBuffEffectHolderMap[effectType] = []*vo.EffectHolderParams{
			effectParam,
		}
	}
	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		general.BaseInfo.Name,
		effectType,
	)
	return true
}

// 移除负面效果
func DebuffEffectWrapRemove(ctx context.Context, general *vo.BattleGeneral, effectType consts.DebuffEffectType, tacticId consts.TacticId) bool {
	if effectParams, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		idx := 0
		for _, effectParam := range effectParams {
			if effectParam.FromTactic == tacticId {
				general.DeBuffEffectHolderMap[effectType] = append(effectParams[:idx], effectParams[idx+1:]...)
				hlog.CtxInfof(ctx, "[%s]的【%v】「%v」效果已消失",
					general.BaseInfo.Name,
					effectParam.FromTactic,
					effectType,
				)
				return true
			}
			idx++
		}
	}
	return false
}

// 负面效果判断
func DeBuffEffectContains(general *vo.BattleGeneral, effectType consts.DebuffEffectType) bool {
	if _, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		return true
	}
	return false
}

// 负面效果获取
func DeBuffEffectGet(general *vo.BattleGeneral, effectType consts.DebuffEffectType) ([]*vo.EffectHolderParams, bool) {
	if v, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		return v, true
	}
	return nil, false
}

// 检查是否有负面效果
func DeBuffEffectContainsCheck(general *vo.BattleGeneral) bool {
	return len(general.DeBuffEffectHolderMap) > 0
}

// 设置战法冻结
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

//战法冻结清零
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
	general.DeBuffEffectHolderMap = map[consts.DebuffEffectType][]*vo.EffectHolderParams{}
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
	general.DeBuffEffectHolderMap = map[consts.DebuffEffectType][]*vo.EffectHolderParams{}
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
