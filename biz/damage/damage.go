package damage

import (
	"math"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// ** 保底、基础伤害 **
// 兵力  保底 基础
// 1000  20  88
// 2000  40  176
// 3000  53  232
// 4000  64  276
// 5000  73  314
// 6000  77  334
// 7000  81  350
// 8000  84  364
// 9000  87  376
// 10000 90  387

// 获取基础伤害
func GetBaseDmg(soldierNum int64) int64 {
	if soldierNum <= 1000 {
		return 88
	} else if soldierNum <= 2000 {
		return 176
	} else if soldierNum <= 3000 {
		return 232
	} else if soldierNum <= 4000 {
		return 276
	} else if soldierNum <= 5000 {
		return 314
	} else if soldierNum <= 6000 {
		return 334
	} else if soldierNum <= 7000 {
		return 350
	} else if soldierNum <= 8000 {
		return 364
	} else if soldierNum <= 9000 {
		return 376
	} else if soldierNum <= 10000 {
		return 387
	}

	panic("Invalid BaseDmg")
}

// 获取保底伤害
func GetMinimumGuaranteeDmg(soldierNum int64) int64 {
	if soldierNum <= 1000 {
		return 20
	} else if soldierNum <= 2000 {
		return 40
	} else if soldierNum <= 3000 {
		return 53
	} else if soldierNum <= 4000 {
		return 64
	} else if soldierNum <= 5000 {
		return 73
	} else if soldierNum <= 6000 {
		return 77
	} else if soldierNum <= 7000 {
		return 81
	} else if soldierNum <= 8000 {
		return 84
	} else if soldierNum <= 9000 {
		return 87
	} else if soldierNum <= 10000 {
		return 90
	}

	panic("Invalid MinimumGuaranteeDmg")
}

// 浮动伤害计算
// @desc: 浮动伤害：正负4% 9种结果
func FluctuateDamage(dmg int64) int64 {
	rates := []float64{
		-0.01, -0.02, -0.03, -0.04,
		0,
		0.01, 0.02, 0.03, 0.04,
	}
	idx := util.Random(0, 8)
	return cast.ToInt64(cast.ToFloat64(dmg) * (1 + rates[int64(math.Round(idx))]))
}

// 普通攻击伤害结算
// @attack 攻击武将
// @suffer 被攻击武将
func AttackDamage(tacticsParams *model.TacticsParams, attackGeneral *vo.BattleGeneral, sufferGeneral *vo.BattleGeneral, attackDmg int64) {
	defer func() {
		//统计上报
		util.AttackReport(tacticsParams,
			attackGeneral.BaseInfo.UniqueId,
			1,
			attackDmg,
			0,
		)
		//「伤害结束」触发器
		if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_DamageEnd]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: attackGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
		//「遭受伤害结束」触发器
		if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferDamageEnd]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: attackGeneral,
					AttackGeneral:  attackGeneral,
					DamageType:     consts.DamageType_Weapon,
				}
				f(params)
			}
		}
		//「普通攻击结束」触发器
		if funcs, ok := attackGeneral.TacticsTriggerMap[consts.BattleAction_AttackEnd]; ok {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: sufferGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
		//「发动兵刃伤害结束」触发器
		if funcs, okk := attackGeneral.TacticsTriggerMap[consts.BattleAction_WeaponDamageEnd]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: attackGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
		//「被普通攻击结束」触发器
		if funcs, ok := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferGeneralAttackEnd]; ok {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: sufferGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
		//遭受兵刃伤害触发器
		if funcs, ok := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferWeaponDamageEnd]; ok {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: sufferGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
	}()

	ctx := tacticsParams.Ctx
	defSoldierNum := sufferGeneral.SoldierNum

	//虎痴效果
	if effectParams, ok := util.BuffEffectGet(attackGeneral, consts.BuffEffectType_TigerIdiot_Locked); ok {
		if len(effectParams) > 0 {
			effectParam := effectParams[0]
			sufferGeneral = effectParam.LockingTarget
		}
	}

	hlog.CtxInfof(ctx, "[%s]对[%s]发动普通攻击",
		attackGeneral.BaseInfo.Name,
		sufferGeneral.BaseInfo.Name,
	)

	//抵御效果判断
	if effectParams, ok := sufferGeneral.BuffEffectHolderMap[consts.BuffEffectType_Defend]; ok {
		effectType := consts.BuffEffectType_Defend
		for idx, effectParam := range effectParams {
			if effectParam.EffectTimes > 0 {
				effectParam.EffectTimes--
				hlog.CtxInfof(ctx, "[%s]来自【%v】的「%v」效果，本次免疫伤害",
					sufferGeneral.BaseInfo.Name,
					effectParam.FromTactic,
					effectType,
				)
				//清除
				if effectParam.EffectTimes == 0 {
					sufferGeneral.BuffEffectHolderMap[effectType] = append(effectParams[:idx], effectParams[idx+1:]...)
					//如果该效果绑定参数结构体为空，则顺便移除该效果
					if len(sufferGeneral.BuffEffectHolderMap[effectType]) == 0 {
						delete(sufferGeneral.BuffEffectHolderMap, effectType)
						hlog.CtxInfof(ctx, "[%s]的「%v」效果已消失",
							sufferGeneral.BaseInfo.Name,
							effectType,
						)
					} else {
						hlog.CtxInfof(ctx, "[%s]的【%v】「%v」效果已消失",
							sufferGeneral.BaseInfo.Name,
							effectParam.FromTactic,
							effectType,
						)
					}
				}
				return
			}
		}
	}

	//嘲讽效果判断
	if util.DeBuffEffectContains(attackGeneral, consts.DebuffEffectType_Taunt) && attackGeneral.TauntByGeneral != nil {
		hlog.CtxInfof(ctx, "[%s]执行来自[%s]的「%v」效果",
			sufferGeneral.BaseInfo.Name,
			sufferGeneral.HelpByGeneral.BaseInfo.Name,
			consts.DebuffEffectType_Taunt,
		)
		AttackDamage(tacticsParams, attackGeneral, attackGeneral.TauntByGeneral, 0)
		return
	}
	//援护效果判断
	if util.BuffEffectContains(attackGeneral, consts.BuffEffectType_Intervene) && sufferGeneral.HelpByGeneral != nil {
		hlog.CtxInfof(ctx, "[%s]执行来自[%s]的「%v」效果",
			sufferGeneral.BaseInfo.Name,
			sufferGeneral.HelpByGeneral.BaseInfo.Name,
			consts.BuffEffectType_Intervene,
		)
		AttackDamage(tacticsParams, attackGeneral, sufferGeneral.HelpByGeneral, 0)
		return
	}

	//是否可以规避
	if effectParams, ok := sufferGeneral.BuffEffectHolderMap[consts.BuffEffectType_Evade]; ok {
		rate := float64(0)
		for _, param := range effectParams {
			rate += param.EffectRate
		}
		if util.GenerateRate(rate) {
			hlog.CtxInfof(ctx, "[%s]处于规避状态，本次伤害无效", sufferGeneral.BaseInfo.Name)
			return
		} else {
			hlog.CtxInfof(ctx, "[%s]规避失败", sufferGeneral.BaseInfo.Name)
		}
	}

	//伤害计算方式
	if attackDmg == 0 {
		//需要计算
		attackDmg = calculateAttackDmgV2(attackGeneral, sufferGeneral)
	} else {
		//不需要计算，用传入值
	}
	tacticsParams.CurrentDamageNum = attackDmg

	//被伤害效果开始触发器
	if funcs, ok := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferGeneralAttack]; ok {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: sufferGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	//兵刃伤害开始触发器
	if funcs, ok := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferWeaponDamage]; ok {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: sufferGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	//「伤害开始」触发器
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_Damage]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	//「兵刃伤害开始」触发器
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_WeaponDamage]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	//「遭受伤害开始」触发器
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferDamage]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}

	//hlog.CtxInfof(ctx, "兵力基础伤害:%d ,武力/防御差:%.2f , 最终伤害:%d , 攻击者武力:%.2f , 防守者统率:%.2f , 造成+受到兵刃伤害增加:%.2f%% , 造成+受到兵刃伤害减少:%.2f%% , 最终增减伤率:%.2f",
	//	int64(numDmg), atkDefDmg, attackDmg, atk, def, inc*100, dec*100, incDecRate)

	//伤害分担
	if v, ok := util.BuffEffectGetAggrEffectRate(sufferGeneral, consts.BuffEffectType_ShareResponsibilityFor); ok {
		hlog.CtxInfof(ctx, "[%s]由于「%v」效果，本次攻击受到的伤害减少了%.2f%%",
			sufferGeneral.BaseInfo.Name,
			consts.BuffEffectType_ShareResponsibilityFor,
			v*100,
		)
		//分担伤害
		shareDmg := cast.ToInt64(cast.ToFloat64(attackDmg) * v)
		//被分担后的伤害
		attackDmg = cast.ToInt64(cast.ToFloat64(attackDmg) * (1 - v))

		hlog.CtxInfof(ctx, "[%s]执行「%v」效果",
			sufferGeneral.ShareResponsibilityForByGeneral.BaseInfo.Name,
			consts.BuffEffectType_ShareResponsibilityFor,
		)
		AttackDamage(tacticsParams, attackGeneral, sufferGeneral.ShareResponsibilityForByGeneral, shareDmg)
	}

	//首次受攻击触发效果
	if sufferGeneral.SufferExecuteWeaponAttackNum == 0 {
		if effectParams, ok := util.DeBuffEffectGet(sufferGeneral, consts.DebuffEffectType_ShockingFourRealms_Prepare); ok {
			effectRate := float64(0)
			for _, effectParam := range effectParams {
				effectRate += effectParam.EffectRate
			}
			attackDmg = cast.ToInt64(cast.ToFloat64(attackDmg) * (1 + effectRate))
		}
	}

	//伤害计算
	finalDmg := int64(0)
	remainSoldierNum := int64(0)
	if attackDmg > sufferGeneral.SoldierNum {
		finalDmg = sufferGeneral.SoldierNum
		sufferGeneral.SoldierNum = 0
		remainSoldierNum = 0
	} else {
		finalDmg = attackDmg
		sufferGeneral.SoldierNum -= attackDmg
		remainSoldierNum = sufferGeneral.SoldierNum
	}

	//统计数据
	attackGeneral.ExecuteGeneralAttackNum++
	attackGeneral.ExecuteWeaponAttackNum++
	attackGeneral.AccumulateTotalDamageNum += finalDmg
	attackGeneral.AccumulateAttackDamageNum += finalDmg

	sufferGeneral.SufferExecuteGeneralAttackNum++
	sufferGeneral.SufferExecuteWeaponAttackNum++
	sufferGeneral.LossSoldierNum += finalDmg

	hlog.CtxInfof(ctx, "[%s]损失了兵力%d(%d↘%d)", sufferGeneral.BaseInfo.Name, finalDmg, defSoldierNum, remainSoldierNum)

	if sufferGeneral.SoldierNum == 0 {
		hlog.CtxInfof(ctx, "[%s]武将兵力为0，无法再战", sufferGeneral.BaseInfo.Name)
	}
}

func calculateAttackDmgV2(attackGeneral *vo.BattleGeneral, sufferGeneral *vo.BattleGeneral) int64 {
	//**伤害计算公式**
	//最终伤害 = 保底伤害 +（兵力基础伤害+属性差×等级差系数）
	//兵刃 保底伤害 +（兵力基础伤害+(武力-统率)×属性差）x 变量
	//谋略  保底伤害 +（兵力基础伤害+(智力-智力)×属性差）x 变量
	//普通攻击  保底伤害 +（兵力基础伤害+(武力-统率)×属性差）x 1
	attackAttr := float64(0)
	defendAttr := float64(0)
	//兵刃伤害
	attackAttr = attackGeneral.BaseInfo.AbilityAttr.ForceBase
	defendAttr = sufferGeneral.BaseInfo.AbilityAttr.CommandBase

	damageNum :=
		//保底伤害
		GetMinimumGuaranteeDmg(attackGeneral.SoldierNum) +
			cast.ToInt64(
				//兵力基础伤害
				(cast.ToFloat64(GetBaseDmg(attackGeneral.SoldierNum))+
					//属性差
					(attackAttr-defendAttr)*1.44)*
					//变量
					1)

	//最终伤害随机值
	damageNum = FluctuateDamage(damageNum)
	return damageNum
}

func calculateAttackDmg(soldierNum int64, attackGeneral *vo.BattleGeneral, sufferGeneral *vo.BattleGeneral) int64 {
	//兵力伤害计算
	numDmg := float64(0)
	//0～3000兵力，伤害为兵力的10%
	if soldierNum > 0 && soldierNum <= 3000 {
		numDmg = cast.ToFloat64(soldierNum) * 0.1
	}
	//3000～6000兵力，1000兵力=50伤害
	if soldierNum > 3000 && soldierNum <= 6000 {
		numDmg = 3000 * 0.1
		numDmg += cast.ToFloat64((soldierNum - 3000) / 1000 * 50)
	}
	//6000～Max兵力，1000兵力=30伤害
	if soldierNum > 6000 {
		numDmg = 3000 * 0.1
		numDmg += cast.ToFloat64((soldierNum / 1000) * 50)
		numDmg += cast.ToFloat64((soldierNum - 3000) / 1000 * 50)
		numDmg += cast.ToFloat64((soldierNum - 6000) / 1000 * 30)
	}
	atk := attackGeneral.BaseInfo.AbilityAttr.ForceBase
	def := sufferGeneral.BaseInfo.AbilityAttr.CommandBase

	lwdiEffectRate := float64(0)
	swdiEffectRate := float64(0)
	lwdiEffectParams := attackGeneral.BuffEffectHolderMap[consts.BuffEffectType_LaunchWeaponDamageImprove]
	swdiEffectParams := sufferGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_SufferWeaponDamageImprove]
	for _, rate := range lwdiEffectParams {
		lwdiEffectRate += rate.EffectRate
	}
	for _, rate := range swdiEffectParams {
		swdiEffectRate += rate.EffectRate
	}

	inc := lwdiEffectRate + swdiEffectRate

	lwddEffectRate := float64(0)
	swddEffectRate := float64(0)
	lwddEffectParams := attackGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchWeaponDamageDeduce]
	swddEffectParams := sufferGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce]
	for _, rate := range lwddEffectParams {
		lwddEffectRate += rate.EffectRate
	}
	for _, rate := range swddEffectParams {
		swddEffectRate += rate.EffectRate
	}
	dec := lwddEffectRate + swddEffectRate

	// 总伤害 = (兵力伤害 + 攻击 - 防御) * (1 + 增益比例 - 减益比例)
	//（攻击 - 防御）伤害结算
	atkDefDmg := atk - def
	if atkDefDmg < 0 {
		atkDefDmg = 0
	}
	//（1 + 增伤 - 减伤）比率结算
	incDecRate := inc - dec
	if incDecRate < 0 {
		//**** 总体减伤逻辑 ****
		if incDecRate < -0.9 {
			//最大减伤为90%
			incDecRate = 1 - 0.9
		} else {
			//正常减伤
			incDecRate = 1 + incDecRate
		}
	} else {
		//**** 总体增伤逻辑 ****
		incDecRate = 1 + (inc - dec)
	}
	attackDmg := cast.ToInt64((numDmg + atkDefDmg) * incDecRate)
	return attackDmg
}

type TacticDamageParam struct {
	//战法参数
	TacticsParams *model.TacticsParams
	//攻击者
	AttackGeneral *vo.BattleGeneral
	//被攻击者
	SufferGeneral *vo.BattleGeneral
	//伤害类型
	DamageType consts.DamageType
	//伤害提升
	DamageImproveRate float64
	//伤害降低
	DamageDeduceRate float64
	//战法ID
	TacticId consts.TacticId
	//战法名
	TacticName string
	//效果名
	EffectName string
	//是否禁止【连环计】被动器生效
	IsBanInterLockedEffect bool
	//是否无视防御
	IsIgnoreDefend bool
}
