package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/spf13/cast"
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
		//造成谋略伤害增加
		consts.BuffEffectType_LaunchStrategyDamageImprove: true,
		//造成兵刃伤害增加
		consts.BuffEffectType_LaunchWeaponDamageImprove: true,
		//受到谋略伤害降低
		consts.BuffEffectType_SufferStrategyDamageDeduce: true,
		//受到兵刃伤害降低
		consts.BuffEffectType_SufferWeaponDamageDeduce: true,
		//倒戈
		consts.BuffEffectType_Defection: true,
		//会心
		consts.BuffEffectType_EnhanceWeapon: true,
		//奇谋
		consts.BuffEffectType_EnhanceStrategy: true,
		//攻心
		consts.BuffEffectType_AttackHeart: true,
	}

	attrDebuffEffectMap = map[consts.DebuffEffectType]bool{
		//受到兵刃伤害增加
		consts.DebuffEffectType_SufferWeaponDamageImprove: true,
		//受到谋略伤害增加
		consts.DebuffEffectType_SufferStrategyDamageImprove: true,
		//造成兵刃伤害减少
		consts.DebuffEffectType_LaunchWeaponDamageDeduce: true,
		//造成谋略伤害减少
		consts.DebuffEffectType_LaunchStrategyDamageDeduce: true,
		//降低武力
		consts.DebuffEffectType_DecrForce: true,
		//降低智力
		consts.DebuffEffectType_DecrIntelligence: true,
		//降低统率
		consts.DebuffEffectType_DecrCommand: true,
		//降低速度
		consts.DebuffEffectType_DecrSpeed: true,
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
		//抵御
		consts.BuffEffectType_Defend: true,
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
//
//	 属性状态：增加或降低武将各种属性；来自不同战法的同种属性状态可以叠加；来自同一战法的同种属性状态将会刷新持续回合；
//				武力/智力/统率/速度/政治/美丽/会心几率/奇谋几率/发动几率/战法造成伤害/受到战法伤害
//		持续性状态：每回合武将开始行动时，对武将造成伤害或治疗；同种状态不可叠加，但会刷新
//		功能性状态：通常不可叠加，不同来源时不可刷新
//		控制状态：不可叠加，不可刷新，负面效果
func BuffEffectWrapSet(ctx context.Context, general *vo.BattleGeneral, effectType consts.BuffEffectType, effectParam *vo.EffectHolderParams) *EffectWrapSetResp {
	/** 1.校验逻辑 **/
	//次数逻辑
	currentTimes := general.BuffEffectCountMap[effectType]

	//处理最大次数默认值为无限大
	if effectParam.MaxEffectTimes <= 0 {
		effectParam.MaxEffectTimes = consts.INT64_MAX
	}

	//超过最大次数限制
	if effectParam.EffectTimes > effectParam.MaxEffectTimes {
		hlog.CtxDebugf(ctx, "[%s]的「%v」效果达到最大叠加次数",
			general.BaseInfo.Name,
			effectType,
		)
		return &EffectWrapSetResp{
			IsSuccess: false,
		}
	}
	//叠加逻辑
	if effectParam.EffectTimes+currentTimes > effectParam.MaxEffectTimes {
		hlog.CtxDebugf(ctx, "[%s]的「%v」效果达到最大叠加次数",
			general.BaseInfo.Name,
			effectType,
		)
		return &EffectWrapSetResp{
			IsSuccess: false,
		}
	} else {
		general.BuffEffectCountMap[effectType] = currentTimes + effectParam.EffectTimes
	}

	/** 触发器逻辑 begin**/
	//施加正面效果开始
	if effectParam.ProduceGeneral != nil {
		if funcs, okk := effectParam.ProduceGeneral.TacticsTriggerMap[consts.BattleAction_BuffEffect]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{}
				f(params)
			}
		}
	}
	//被施加正面效果开始
	if funcs, okk := general.TacticsTriggerMap[consts.BattleAction_SufferBuffEffect]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{}
			f(params)
		}
	}
	defer func() {
		//施加正面效果结束
		if effectParam.ProduceGeneral != nil {
			if funcs, okk := effectParam.ProduceGeneral.TacticsTriggerMap[consts.BattleAction_BuffEffectEnd]; okk {
				for _, f := range funcs {
					params := &vo.TacticsTriggerParams{}
					f(params)
				}
			}
		}
		//被施加正面效果结束
		if funcs, okk := general.TacticsTriggerMap[consts.BattleAction_SufferBuffEffectEnd]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{}
				f(params)
			}
		}
	}()
	/** 触发器逻辑 end**/

	/** 2.执行逻辑 **/
	//是否包含正面效果判断
	_, isContainBuffEffect := general.BuffEffectHolderMap[effectType]
	//是否包含属性正面效果判断
	_, isContainAttrBuffEffect := attrBuffEffectMap[effectType]
	//是否包含持续性正面效果判断
	_, isContainContinuousBuffEffect := continuousBuffEffectMap[effectType]
	//是否包含功能性正面效果判断
	_, isContainFunctionBuffEffectEffect := functionBuffEffectMap[effectType]

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
				return &EffectWrapSetResp{
					IsSuccess:       true,
					IsRefreshEffect: true,
				}
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
			return &EffectWrapSetResp{
				IsSuccess: true,
			}
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
			return &EffectWrapSetResp{
				IsSuccess:       true,
				IsRefreshEffect: true,
			}
		}
	}

	//功能性状态：通常不可叠加，不同来源时不可刷新
	if isContainBuffEffect && isContainFunctionBuffEffectEffect {
		hlog.CtxInfof(ctx, "[%s]身上已有同等或更强的「%v」效果",
			general.BaseInfo.Name,
			effectType,
		)
		return &EffectWrapSetResp{
			IsSuccess:       false,
			IsRefreshEffect: false,
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
	hlog.CtxInfof(ctx, "[%s]的【%v】「%v」效果已施加",
		general.BaseInfo.Name,
		effectParam.FromTactic,
		effectType,
	)

	//分担伤害
	if effectType == consts.BuffEffectType_ShareResponsibilityFor {
		if effectParam.ShareResponsibilityForByGeneral == nil {
			panic(any("分担伤害效果，分担者不能为空"))
		}
		general.ShareResponsibilityForByGeneral = effectParam.ShareResponsibilityForByGeneral
	}

	//增益属性处理
	buffEffectIncr(ctx, general, effectType, effectParam.EffectValue)

	return &EffectWrapSetResp{
		IsSuccess: true,
	}
}

// 移除增益效果
func BuffEffectWrapRemove(ctx context.Context, general *vo.BattleGeneral, effectType consts.BuffEffectType, tacticId consts.TacticId) bool {
	if effectParams, ok := general.BuffEffectHolderMap[effectType]; ok {
		//指定战法id删除
		if tacticId > 0 {
			idx := 0
			for _, effectParam := range effectParams {
				if effectParam.FromTactic == tacticId {
					general.BuffEffectHolderMap[effectType] = append(effectParams[:idx], effectParams[idx+1:]...)
					//如果该效果绑定参数结构体为空，则顺便移除该效果
					if len(general.BuffEffectHolderMap[effectType]) == 0 {
						delete(general.BuffEffectHolderMap, effectType)
						hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
							general.BaseInfo.Name,
							effectType,
						)
					} else {
						hlog.CtxInfof(ctx, "[%s]的【%v】「%v」效果已消失",
							general.BaseInfo.Name,
							effectParam.FromTactic,
							effectType,
						)
					}

					//增益属性移除处理
					buffEffectDecr(ctx, general, effectType, effectParam.EffectValue)

					return true
				}
				idx++
			}
		} else {
			//不指定战法id删除，则删除该效果
			delete(general.BuffEffectHolderMap, effectType)
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				general.BaseInfo.Name,
				effectType,
			)

			return true
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

// 正面效果获取
func BuffEffectOfTacticGet(general *vo.BattleGeneral, effectType consts.BuffEffectType, tacticId consts.TacticId) ([]*vo.EffectHolderParams, bool) {
	res := make([]*vo.EffectHolderParams, 0)
	if effectParams, ok := general.BuffEffectHolderMap[effectType]; ok {
		//按战法Id获取效果
		if tacticId > 0 {
			for _, effectParam := range effectParams {
				if effectParam.FromTactic == tacticId {
					res = append(res, effectParam)
				}
			}
			return res, true
		}
	}
	return nil, false
}

// 获取增益效果总次数
func BuffEffectGetCount(general *vo.BattleGeneral, effectType consts.BuffEffectType) int64 {
	times := int64(0)
	if effectParams, ok := general.BuffEffectHolderMap[effectType]; ok {
		for _, effectParam := range effectParams {
			times += effectParam.EffectTimes
		}
	}
	return times
}

// 获取正面效果种类数量
func BuffEffectHolderCount(general *vo.BattleGeneral) int {
	return len(general.BuffEffectHolderMap)
}

// 获取负面效果种类数量
func DeBuffEffectHolderCount(general *vo.BattleGeneral) int {
	return len(general.DeBuffEffectHolderMap)
}

// 获取增益效果(汇总)
func BuffEffectGetAggrEffectRate(general *vo.BattleGeneral, effectType consts.BuffEffectType) (float64, bool) {
	effectRate := float64(0)
	if v, ok := general.BuffEffectHolderMap[effectType]; ok {
		for _, effectParam := range v {
			effectRate += effectParam.EffectRate
		}
		return effectRate, true
	}
	return 0, false
}

// 是否存在增益效果
func BuffEffectContainsCheck(general *vo.BattleGeneral) bool {
	return len(general.BuffEffectHolderMap) > 0
}

// 增益效果数量
func BuffEffectContainsNum(general *vo.BattleGeneral) int {
	return len(general.BuffEffectHolderMap)
}

// 正面效果次数是否已消耗完毕
func BuffEffectOfTacticIsDeplete(general *vo.BattleGeneral, effectType consts.BuffEffectType, tacticId consts.TacticId) bool {
	if tacticId <= 0 {
		return false
	}

	if effectParams, ok := general.BuffEffectHolderMap[effectType]; ok {
		for _, effectParam := range effectParams {
			//找到指定战法
			if effectParam.FromTactic == tacticId {
				//可用次数是否为0
				if effectParam.EffectTimes == 0 {
					return true
				}
			}
		}
	}
	return false
}

type BuffEffectOfTacticCostRoundParams struct {
	//上下文
	Ctx context.Context
	//操作武将
	General *vo.BattleGeneral
	//正面效果
	EffectType consts.BuffEffectType
	//关联战法
	TacticId consts.TacticId
	//效果消耗完成回调函数
	CostOverTriggerFunc func()
}

// 正面效果消耗
func BuffEffectOfTacticCostRound(params *BuffEffectOfTacticCostRoundParams) bool {
	if params.TacticId <= 0 {
		return false
	}

	if effectParams, ok := params.General.BuffEffectHolderMap[params.EffectType]; ok {
		for idx, effectParam := range effectParams {
			//找到指定战法
			if effectParam.FromTactic == params.TacticId {
				//消耗
				if effectParam.EffectRound > 0 {
					effectParam.EffectRound--
					return true
				}
				//清除
				if effectParam.EffectRound == 0 {
					params.General.BuffEffectHolderMap[params.EffectType] = append(effectParams[:idx], effectParams[idx+1:]...)
					hlog.CtxInfof(params.Ctx, "[%s]的【%v】「%v」效果已消失",
						params.General.BaseInfo.Name,
						params.TacticId,
						params.EffectType,
					)

					//属性加点清理
					buffEffectDecr(params.Ctx, params.General, params.EffectType, effectParam.EffectValue)

					//执行回调函数
					if params.CostOverTriggerFunc != nil {
						params.CostOverTriggerFunc()
					}
					return true
				}
			}
		}
	}
	return false
}

type BuffEffectOfTacticCostTimeParams struct {
	//上下文
	Ctx context.Context
	//操作武将
	General *vo.BattleGeneral
	//正面效果
	EffectType consts.BuffEffectType
	//关联战法
	TacticId consts.TacticId
	//消耗次数
	CostTimes int64
	//效果消耗完成回调函数
	CostOverTriggerFunc func()
}

// 正面效果消耗
func BuffEffectOfTacticCostTime(params *BuffEffectOfTacticCostTimeParams) bool {
	if params.TacticId <= 0 {
		return false
	}

	if effectParams, ok := params.General.BuffEffectHolderMap[params.EffectType]; ok {
		for idx, effectParam := range effectParams {
			//找到指定战法
			if effectParam.FromTactic == params.TacticId {
				//消耗
				if effectParam.EffectTimes > params.CostTimes && params.CostTimes > 0 {
					effectParam.EffectTimes -= params.CostTimes
					return true
				}
				//清除
				if effectParam.EffectTimes == 0 {
					params.General.BuffEffectHolderMap[params.EffectType] = append(effectParams[:idx], effectParams[idx+1:]...)
					hlog.CtxInfof(params.Ctx, "[%s]的【%v】「%v」效果已消失",
						params.General.BaseInfo.Name,
						params.TacticId,
						params.EffectType,
					)

					//属性加点清理
					buffEffectDecr(params.Ctx, params.General, params.EffectType, effectParam.EffectValue)

					//执行回调函数
					if params.CostOverTriggerFunc != nil {
						params.CostOverTriggerFunc()
					}
					return true
				}
			}
		}
	}
	return false
}

type BuffEffectOfTacticDeduceParams struct {
	//上下文
	Ctx context.Context
	//操作武将
	General *vo.BattleGeneral
	//正面效果
	EffectType consts.BuffEffectType
	//关联战法
	TacticId consts.TacticId
	//消耗概率
	EffectRate float64
	//消耗值
	EffectValue int64
	//触发率
	TriggerRate float64
}

// 正面效果降低
func BuffEffectOfTacticDeduce(params *BuffEffectOfTacticDeduceParams) bool {
	//必须指定战法
	if params.TacticId <= 0 {
		return false
	}

	if effectParams, ok := params.General.BuffEffectHolderMap[params.EffectType]; ok {
		for idx, effectParam := range effectParams {
			//找到指定战法
			if effectParam.FromTactic == params.TacticId {
				switch params.EffectType {
				case consts.BuffEffectType_TacticsActiveTriggerImprove:
					//触发率处理
					if params.TriggerRate > 0 {
						effectParam.TriggerRate -= params.TriggerRate
						if effectParam.TriggerRate < 0 {
							effectParam.TriggerRate = 0
						}
						hlog.CtxInfof(params.Ctx, "[%s]的【%v】降低了%.2f",
							params.General.BaseInfo.Name,
							params.EffectType,
							params.TriggerRate*100)
					}

					if effectParam.TriggerRate == 0 {
						params.General.BuffEffectHolderMap[params.EffectType] = append(effectParams[:idx], effectParams[idx+1:]...)
						hlog.CtxInfof(params.Ctx, "[%s]的【%v】「%v」效果已消失",
							params.General.BaseInfo.Name,
							params.TacticId,
							params.EffectType,
						)
					}
				case consts.BuffEffectType_TacticsActiveTriggerPrepareImprove:
					//触发率处理
					if params.TriggerRate > 0 {
						effectParam.TriggerRate -= params.TriggerRate
						if effectParam.TriggerRate < 0 {
							effectParam.TriggerRate = 0
						}
						hlog.CtxInfof(params.Ctx, "[%s]的【%v】降低了%.2f",
							params.General.BaseInfo.Name,
							params.EffectType,
							params.TriggerRate*100)
					}
					if effectParam.TriggerRate == 0 {
						params.General.BuffEffectHolderMap[params.EffectType] = append(effectParams[:idx], effectParams[idx+1:]...)
						hlog.CtxInfof(params.Ctx, "[%s]的【%v】「%v」效果已消失",
							params.General.BaseInfo.Name,
							params.TacticId,
							params.EffectType,
						)
					}
				}
			}
		}
	}
	return false
}

type EffectWrapSetResp struct {
	//是否设置成功
	IsSuccess bool
	//是否进行效果刷新操作
	IsRefreshEffect bool
}

// 减益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
//
//	 属性状态：增加或降低武将各种属性；来自不同战法的同种属性状态可以叠加；来自同一战法的同种属性状态将会刷新持续回合；
//				武力/智力/统率/速度/政治/美丽/会心几率/奇谋几率/发动几率/战法造成伤害/受到战法伤害
//		持续性状态：每回合武将开始行动时，对武将造成伤害或治疗；同种状态不可叠加，但会刷新
//		功能性状态：通常不可叠加，不同来源时不可刷新
//		控制状态：不可叠加，不可刷新，负面效果
func DebuffEffectWrapSet(ctx context.Context, general *vo.BattleGeneral, effectType consts.DebuffEffectType, effectParam *vo.EffectHolderParams) *EffectWrapSetResp {
	/** 1.校验逻辑 **/
	//次数逻辑
	currentTimes := general.DeBuffEffectCountMap[effectType]

	//处理最大次数默认值为无限大
	if effectParam.MaxEffectTimes <= 0 {
		effectParam.MaxEffectTimes = consts.INT64_MAX
	}
	//超过最大次数限制
	if effectParam.EffectTimes > effectParam.MaxEffectTimes {
		hlog.CtxDebugf(ctx, "[%s]的「%v」效果达到最大叠加次数",
			general.BaseInfo.Name,
			effectType,
		)
		return &EffectWrapSetResp{
			IsSuccess: false,
		}
	}
	//叠加逻辑
	if effectParam.EffectTimes+currentTimes > effectParam.MaxEffectTimes {
		hlog.CtxDebugf(ctx, "[%s]的「%v」效果达到最大叠加次数",
			general.BaseInfo.Name,
			effectType,
		)
		return &EffectWrapSetResp{
			IsSuccess: false,
		}
	} else {
		general.DeBuffEffectCountMap[effectType] = currentTimes + effectParam.EffectTimes
	}

	/** 触发器逻辑 begin**/
	//施加负面效果开始
	if effectParam.ProduceGeneral != nil {
		if funcs, okk := effectParam.ProduceGeneral.TacticsTriggerMap[consts.BattleAction_DebuffEffect]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					EffectHolderParams: effectParam,
				}
				f(params)
			}
		}
	}
	//被施加负面效果开始
	if funcs, okk := general.TacticsTriggerMap[consts.BattleAction_SufferDebuffEffect]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{}
			resp := f(params)
			if resp.IsTerminate {
				return &EffectWrapSetResp{
					IsSuccess: false,
				}
			}
		}
	}
	defer func() {
		//施加负面效果结束
		if effectParam.ProduceGeneral != nil {
			if funcs, okk := effectParam.ProduceGeneral.TacticsTriggerMap[consts.BattleAction_DebuffEffectEnd]; okk {
				for _, f := range funcs {
					params := &vo.TacticsTriggerParams{
						DebuffEffect:              effectType,
						DebuffEffectOfTactic:      effectParam.FromTactic,
						SufferDebuffEffectGeneral: general,
					}
					f(params)
				}
			}
		}
		//被施加负面效果结束
		if funcs, okk := general.TacticsTriggerMap[consts.BattleAction_SufferDebuffEffectEnd]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					AttackGeneral: effectParam.ProduceGeneral,
					DebuffEffect:  effectType,
					EffectRound:   effectParam.EffectRound,
				}
				f(params)
			}
		}
	}()
	/** 触发器逻辑 end**/

	/** 执行逻辑 **/
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
				return &EffectWrapSetResp{
					IsSuccess:       true,
					IsRefreshEffect: true,
				}
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
			return &EffectWrapSetResp{
				IsSuccess: true,
			}
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
			return &EffectWrapSetResp{
				IsSuccess:       true,
				IsRefreshEffect: true,
			}
		}
	}

	//控制状态：不可叠加，不可刷新，负面效果
	if isContainDebuffEffect && isContainControlDebuffEffect {
		hlog.CtxInfof(ctx, "[%s]身上已有同等或更强的「%v」效果",
			general.BaseInfo.Name,
			effectType,
		)
		return &EffectWrapSetResp{}
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
			return &EffectWrapSetResp{}
		} else {
			general.TauntByGeneral = effectParam.ProduceGeneral
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
	hlog.CtxInfof(ctx, "[%s]来自【%v】「%v」效果已施加",
		general.BaseInfo.Name,
		effectParam.FromTactic,
		effectType,
	)

	//属性加点效果设置
	debuffEffectDecr(ctx, general, effectType, effectParam.EffectValue)

	return &EffectWrapSetResp{
		IsSuccess: true,
	}
}

// 移除负面效果
func DebuffEffectWrapRemove(ctx context.Context, general *vo.BattleGeneral, effectType consts.DebuffEffectType, tacticId consts.TacticId) bool {
	if effectParams, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		//指定战法id删除
		if tacticId > 0 {
			idx := 0
			for _, effectParam := range effectParams {
				if effectParam.FromTactic == tacticId {
					general.DeBuffEffectHolderMap[effectType] = append(effectParams[:idx], effectParams[idx+1:]...)
					//如果该效果绑定参数结构体为空，则顺便移除该效果
					if len(general.DeBuffEffectHolderMap[effectType]) == 0 {
						delete(general.DeBuffEffectHolderMap, effectType)
						hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
							general.BaseInfo.Name,
							effectType,
						)
					} else {
						hlog.CtxInfof(ctx, "[%s]的【%v】「%v」效果已消失",
							general.BaseInfo.Name,
							effectParam.FromTactic,
							effectType,
						)
					}

					//减益效果恢复
					debuffEffectIncr(ctx, general, effectType, effectParam.EffectValue)

					return true
				}
				idx++
			}
		} else {
			//不指定战法id删除，则删除该效果
			delete(general.DeBuffEffectHolderMap, effectType)
			hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
				general.BaseInfo.Name,
				effectType,
			)
			return true
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

// 控制效果判断
func DeBuffEffectContainsControl(general *vo.BattleGeneral) bool {
	for effectType, _ := range general.DeBuffEffectHolderMap {
		if _, ok := controlDebuffEffectMap[effectType]; ok {
			return true
		}
	}
	return false
}

// 控制效果判断
func IsControlDeBuffEffect(effectType consts.DebuffEffectType) bool {
	if _, ok := controlDebuffEffectMap[effectType]; ok {
		return true
	}
	return false
}

// 负面效果判断
func DeBuffEffectOfTacticContains(general *vo.BattleGeneral, effectType consts.DebuffEffectType, tacticId consts.TacticId) bool {
	if effectParams, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		for _, effectParam := range effectParams {
			if effectParam.FromTactic == tacticId {
				return true
			}
		}
	}
	return false
}

// 负面效果次数是否已消耗完毕
func DeBuffEffectOfTacticIsDeplete(general *vo.BattleGeneral, effectType consts.DebuffEffectType, tacticId consts.TacticId) bool {
	if tacticId <= 0 {
		return false
	}

	if effectParams, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		for _, effectParam := range effectParams {
			//找到指定战法
			if effectParam.FromTactic == tacticId {
				//可用次数是否为0
				if effectParam.EffectTimes == 0 {
					return true
				}
			}
		}
	}
	return false
}

type DebuffEffectOfTacticCostRoundParams struct {
	//上下文
	Ctx context.Context
	//操作武将
	General *vo.BattleGeneral
	//负面效果
	EffectType consts.DebuffEffectType
	//关联战法
	TacticId consts.TacticId
	//效果消耗完成回调函数
	CostOverTriggerFunc func()
}

// 负面效果消耗
func DeBuffEffectOfTacticCostRound(params *DebuffEffectOfTacticCostRoundParams) bool {
	if params.TacticId <= 0 {
		return false
	}

	if effectParams, ok := params.General.DeBuffEffectHolderMap[params.EffectType]; ok {
		for idx, effectParam := range effectParams {
			//找到指定战法
			if effectParam.FromTactic == params.TacticId {
				//消耗
				if effectParam.EffectRound > 0 {
					effectParam.EffectRound--
					return true
				}

				//清除
				if effectParam.EffectRound == 0 {
					params.General.DeBuffEffectHolderMap[params.EffectType] = append(effectParams[:idx], effectParams[idx+1:]...)
					hlog.CtxInfof(params.Ctx, "[%s]来自【%v】「%v」效果已消失",
						params.General.BaseInfo.Name,
						params.TacticId,
						params.EffectType,
					)

					//减益效果恢复
					debuffEffectIncr(params.Ctx, params.General, params.EffectType, effectParam.EffectValue)

					//执行回调函数
					if params.CostOverTriggerFunc != nil {
						params.CostOverTriggerFunc()
					}
					return true
				}
			}
		}
	}
	return false
}

type DebuffEffectOfTacticCostRateParams struct {
	//上下文
	Ctx context.Context
	//操作武将
	General *vo.BattleGeneral
	//负面效果
	EffectType consts.DebuffEffectType
	//关联战法
	TacticId consts.TacticId
	//影响率
	EffectRate float64
	//效果消耗完成回调函数
	CostOverTriggerFunc func()
}

// 负面效果消耗
func DeBuffEffectOfTacticCostRate(params *DebuffEffectOfTacticCostRateParams) bool {
	if params.TacticId <= 0 {
		return false
	}

	if effectParams, ok := params.General.DeBuffEffectHolderMap[params.EffectType]; ok {
		for idx, effectParam := range effectParams {
			//找到指定战法
			if effectParam.FromTactic == params.TacticId {
				//消耗
				if effectParam.EffectRate > 0 {
					effectParam.EffectRate -= params.EffectRate
					return true
				}

				//清除
				if effectParam.EffectRate == 0 {
					params.General.DeBuffEffectHolderMap[params.EffectType] = append(effectParams[:idx], effectParams[idx+1:]...)
					hlog.CtxInfof(params.Ctx, "[%s]来自【%v】「%v」效果已消失",
						params.General.BaseInfo.Name,
						params.TacticId,
						params.EffectType,
					)

					//减益效果恢复
					debuffEffectIncr(params.Ctx, params.General, params.EffectType, effectParam.EffectValue)

					//执行回调函数
					if params.CostOverTriggerFunc != nil {
						params.CostOverTriggerFunc()
					}
					return true
				}
			}
		}
	}
	return false
}

type BuffEffectOfTacticCostRateParams struct {
	//上下文
	Ctx context.Context
	//操作武将
	General *vo.BattleGeneral
	//正面效果
	EffectType consts.BuffEffectType
	//关联战法
	TacticId consts.TacticId
	//影响率
	EffectRate float64
	//效果消耗完成回调函数
	CostOverTriggerFunc func()
}

// 正面效果消耗
func BuffEffectOfTacticCostRate(params *BuffEffectOfTacticCostRateParams) bool {
	if params.TacticId <= 0 {
		return false
	}

	if effectParams, ok := params.General.BuffEffectHolderMap[params.EffectType]; ok {
		for idx, effectParam := range effectParams {
			//找到指定战法
			if effectParam.FromTactic == params.TacticId {
				//消耗
				if effectParam.EffectRate > 0 {
					effectParam.EffectRate -= params.EffectRate
					return true
				}

				//清除
				if effectParam.EffectRate == 0 {
					params.General.BuffEffectHolderMap[params.EffectType] = append(effectParams[:idx], effectParams[idx+1:]...)
					hlog.CtxInfof(params.Ctx, "[%s]来自【%v】「%v」效果已消失",
						params.General.BaseInfo.Name,
						params.TacticId,
						params.EffectType,
					)

					//减益效果恢复
					buffEffectIncr(params.Ctx, params.General, params.EffectType, effectParam.EffectValue)

					//执行回调函数
					if params.CostOverTriggerFunc != nil {
						params.CostOverTriggerFunc()
					}
					return true
				}
			}
		}
	}
	return false
}

func buffEffectIncr(ctx context.Context, general *vo.BattleGeneral, effectType consts.BuffEffectType, effectValue int64) {
	//属性加点效果设置
	switch effectType {
	case consts.BuffEffectType_IncrForce:
		general.BaseInfo.AbilityAttr.ForceBase += cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的武力提高了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.BuffEffectType_IncrIntelligence:
		general.BaseInfo.AbilityAttr.IntelligenceBase += cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的智力提高了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.BuffEffectType_IncrCommand:
		general.BaseInfo.AbilityAttr.CommandBase += cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的统率提高了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.BuffEffectType_IncrSpeed:
		general.BaseInfo.AbilityAttr.SpeedBase += cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的速度提高了%.2d",
			general.BaseInfo.Name,
			effectValue)
	}
}

func buffEffectDecr(ctx context.Context, general *vo.BattleGeneral, effectType consts.BuffEffectType, effectValue int64) {
	//属性加点效果设置
	switch effectType {
	case consts.BuffEffectType_IncrForce:
		general.BaseInfo.AbilityAttr.ForceBase -= cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的武力降低了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.BuffEffectType_IncrIntelligence:
		general.BaseInfo.AbilityAttr.IntelligenceBase -= cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的智力降低了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.BuffEffectType_IncrCommand:
		general.BaseInfo.AbilityAttr.CommandBase -= cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的统率降低了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.BuffEffectType_IncrSpeed:
		general.BaseInfo.AbilityAttr.SpeedBase -= cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的速度降低了%.2d",
			general.BaseInfo.Name,
			effectValue)
	}
}

func debuffEffectIncr(ctx context.Context, general *vo.BattleGeneral, effectType consts.DebuffEffectType, effectValue int64) {
	//属性加点效果设置
	switch effectType {
	case consts.DebuffEffectType_DecrForce:
		general.BaseInfo.AbilityAttr.ForceBase += cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的武力提高了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.DebuffEffectType_DecrIntelligence:
		general.BaseInfo.AbilityAttr.IntelligenceBase += cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的智力提高了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.DebuffEffectType_DecrCommand:
		general.BaseInfo.AbilityAttr.CommandBase += cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的统率提高了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.DebuffEffectType_DecrSpeed:
		general.BaseInfo.AbilityAttr.SpeedBase += cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的速度提高了%.2d",
			general.BaseInfo.Name,
			effectValue)
	}
}

func debuffEffectDecr(ctx context.Context, general *vo.BattleGeneral, effectType consts.DebuffEffectType, effectValue int64) {
	//属性加点效果设置
	switch effectType {
	case consts.DebuffEffectType_DecrForce:
		general.BaseInfo.AbilityAttr.ForceBase -= cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的武力降低了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.DebuffEffectType_DecrIntelligence:
		general.BaseInfo.AbilityAttr.IntelligenceBase -= cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的智力降低了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.DebuffEffectType_DecrCommand:
		general.BaseInfo.AbilityAttr.CommandBase -= cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的统率降低了%.2d",
			general.BaseInfo.Name,
			effectValue)
	case consts.DebuffEffectType_DecrSpeed:
		general.BaseInfo.AbilityAttr.SpeedBase -= cast.ToFloat64(effectValue)
		hlog.CtxInfof(ctx, "[%s]的速度降低了%.2d",
			general.BaseInfo.Name,
			effectValue)
	}
}

// 负面效果获取
func DeBuffEffectOfTacticGet(general *vo.BattleGeneral, effectType consts.DebuffEffectType, tacticId consts.TacticId) ([]*vo.EffectHolderParams, bool) {
	res := make([]*vo.EffectHolderParams, 0)
	if effectParams, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		//按战法Id获取效果
		if tacticId > 0 {
			for _, effectParam := range effectParams {
				if effectParam.FromTactic == tacticId {
					res = append(res, effectParam)
				}
			}
			return res, true
		}
	}
	return nil, false
}

// 负面效果获取
func DeBuffEffectGet(general *vo.BattleGeneral, effectType consts.DebuffEffectType) ([]*vo.EffectHolderParams, bool) {
	if effectParams, ok := general.DeBuffEffectHolderMap[effectType]; ok {
		return effectParams, true
	}
	return nil, false
}

// 检查是否有负面效果
func DeBuffEffectContainsCheck(general *vo.BattleGeneral) bool {
	return len(general.DeBuffEffectHolderMap) > 0
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

// 清除准备战法冷却
func TacticFrozenClean(ctx context.Context, general *vo.BattleGeneral) {
	general.TacticFrozenMap = map[consts.TacticId]bool{}
}

// 清除准备战法冷却
func TacticFrozenOfTacticClean(ctx context.Context, general *vo.BattleGeneral, tacticId consts.TacticId) {
	general.TacticFrozenMap[tacticId] = false
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

// 是否可以行动
func IsCanBeginAction(ctx context.Context, general *vo.BattleGeneral) bool {
	//震慑
	if effectParams, ok := general.DeBuffEffectHolderMap[consts.DebuffEffectType_Awe]; ok {
		if len(effectParams) > 0 && GenerateRate(effectParams[0].EffectRate) {
			hlog.CtxInfof(ctx, "武将[%s]处于「%v」状态，无法行动",
				general.BaseInfo.Name,
				consts.DebuffEffectType_Awe,
			)
			return false
		}
	}
	return true
}

// 是否可以发动主动战法
func IsCanActiveTactic(ctx context.Context, general *vo.BattleGeneral) bool {
	//无法发动主动战法
	if effectParams, ok := general.DeBuffEffectHolderMap[consts.DebuffEffectType_CanNotActiveTactic]; ok {
		if len(effectParams) > 0 && GenerateRate(effectParams[0].EffectRate) {
			hlog.CtxInfof(ctx, "武将[%s]处于「%v」状态，无法发动主动战法",
				general.BaseInfo.Name,
				consts.DebuffEffectType_CanNotActiveTactic,
			)
			return false
		}
	}
	//计穷
	if effectParams, ok := general.DeBuffEffectHolderMap[consts.DebuffEffectType_NoStrategy]; ok {
		if len(effectParams) > 0 && GenerateRate(effectParams[0].EffectRate) {
			hlog.CtxInfof(ctx, "武将[%s]处于「%v」状态，无法发动主动战法",
				general.BaseInfo.Name,
				consts.DebuffEffectType_NoStrategy,
			)
			return false
		}
	}
	return true
}

// 是否可以造成伤害
func IsCanDamage(ctx context.Context, general *vo.BattleGeneral) bool {
	//虚弱
	if effectParams, ok := general.DeBuffEffectHolderMap[consts.DebuffEffectType_PoorHealth]; ok {
		if len(effectParams) > 0 && GenerateRate(effectParams[0].EffectRate) {
			hlog.CtxInfof(ctx, "武将[%s]处于「%v」状态，无法造成伤害",
				general.BaseInfo.Name,
				consts.DebuffEffectType_PoorHealth,
			)
			return false
		}
	}
	return true
}

// 是否可以普通攻击
func IsCanGeneralAttack(ctx context.Context, general *vo.BattleGeneral) bool {
	//缴械
	if effectParams, ok := general.DeBuffEffectHolderMap[consts.DebuffEffectType_CancelWeapon]; ok {
		if len(effectParams) > 0 && GenerateRate(effectParams[0].EffectRate) {
			hlog.CtxInfof(ctx, "武将[%s]处于「%v」状态，无法普通攻击",
				general.BaseInfo.Name,
				consts.DebuffEffectType_CancelWeapon,
			)
			return false
		}
	}
	//无法普通攻击
	if effectParams, ok := general.DeBuffEffectHolderMap[consts.DebuffEffectType_CanNotGeneralAttack]; ok {
		if len(effectParams) > 0 && GenerateRate(effectParams[0].EffectRate) {
			hlog.CtxInfof(ctx, "武将[%s]处于「%v」状态，无法普通攻击",
				general.BaseInfo.Name,
				consts.DebuffEffectType_CanNotGeneralAttack,
			)
			return false
		}
	}
	//金城汤池[无法普通攻击]
	if effectParams, ok := general.DeBuffEffectHolderMap[consts.DebuffEffectType_RampartsOfMetalsAndAMoatOfHotWaterTactic_CanNotGeneralAttack]; ok {
		if len(effectParams) > 0 && GenerateRate(effectParams[0].EffectRate) {
			hlog.CtxInfof(ctx, "武将[%s]处于「%v」状态，无法普通攻击",
				general.BaseInfo.Name,
				consts.DebuffEffectType_RampartsOfMetalsAndAMoatOfHotWaterTactic_CanNotGeneralAttack,
			)
			return false
		}
	}
	return true
}

// 是否可以恢复兵力
func IsCanResume(ctx context.Context, general *vo.BattleGeneral) bool {
	//禁疗
	if _, ok := general.DeBuffEffectHolderMap[consts.DebuffEffectType_ProhibitionTreatment]; ok {
		hlog.CtxInfof(ctx, "武将[%s]处于「%v」状态，无法恢复兵力",
			general.BaseInfo.Name,
			consts.DebuffEffectType_ProhibitionTreatment,
		)
		return false
	}
	return true
}

// 是否可以规避
func IsCanEvade(ctx context.Context, general *vo.BattleGeneral) bool {
	if effectParams, ok := general.BuffEffectHolderMap[consts.BuffEffectType_Evade]; ok {
		rate := float64(0)
		for _, param := range effectParams {
			rate += param.EffectRate
		}
		if GenerateRate(rate) {
			hlog.CtxInfof(ctx, "[%s]处于规避状态，本次伤害无效", general.BaseInfo.Name)
			return true
		} else {
			hlog.CtxInfof(ctx, "[%s]规避失败", general.BaseInfo.Name)
			return false
		}
	}
	return false
}
