package util

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/spf13/cast"
)

// 参考 https://zhuanlan.zhihu.com/p/439300738

//============ 兵力伤害公式 =============
// f1 = 0.1877 * Num^0.32 * (Atk_a - Def_b)(幂函数)
//
//Num是己方武将带兵数量。
//Atk-Def是属性差。造成兵刃伤害时=我方武力-敌方统率；造成谋略伤害时=我方智力-敌方智力。
//func MilitaryStrengthDamage(num int64, atk int64, def int64) int64 {
//	dmg := cast.ToInt64(0.1877 * math.Pow(float64(num), 0.32) * cast.ToFloat64(atk-def))
//	fmt.Printf("military_strength dmg: %d\n", dmg)
//	return dmg
//}

//============ 武将伤害公式 =============
// f2 = 0.0005 * x^2 + 0.9 * x + 4.5
//
//其中 x = (Atk_a - Def_b)
//即武将伤害只与双方属性差值有关。
//该公式将在第三部分【武将伤害】中详细说明。
//func GeneralDamage(atk int64, def int64) int64 {
//	dmg := cast.ToInt64(0.0005*math.Pow(cast.ToFloat64(atk-def), 2) + 0.9*cast.ToFloat64(atk-def) + 4.5)
//	fmt.Printf("general dmg: %d\n", dmg)
//	return dmg
//}

//============ 总伤害公式 =============
//dmg(A->B) = （f1 + f2）*(1 + Inc_A -Dec_A) * (1 + Inc_B - Dec_B) * R * (1 + Crt) * ri
//
//其中：
//f1为兵力伤害公式，见 公式二；
//f2为武将伤害公式，见 公式三。
//Inc、Dec为增减伤，共有4类：己方增伤、己方减伤（伤害降低）、敌方增伤（易伤）、敌方减伤。增减伤主要来自于战法、城内建筑，同类增减伤之间加算。
//R为技能系数，即战法中的伤害系数，普攻默认为100%。
//Crt为会心伤害，即会心时触发的会心伤害加成。
//ri为其他的额外增减伤乘区，同类乘区内加算。不同乘区包括兵种克制系数，士气减伤，兵书，州战法等。
//func TotalDamage(num int64,
//	atk int64,
//	def int64,
//	own_inc float64,
//	own_dec float64,
//	opp_inc float64,
//	opp_dec float64,
//	r float64,
//	crt float64,
//	ri float64,
//) int64 {
//	dmg := cast.ToFloat64(MilitaryStrengthDamage(num, atk, def) + GeneralDamage(atk, def))
//	if own_inc > 0 {
//		dmg = dmg * (1 + own_inc)
//	}
//	if own_dec > 0 {
//		dmg = dmg * (1 - own_dec)
//	}
//	if opp_inc > 0 {
//		dmg = dmg * (1 + opp_inc)
//	}
//	if opp_dec > 0 {
//		dmg = dmg * (1 - opp_dec)
//	}
//	if r > 0 {
//		dmg = dmg * r
//	}
//	if crt > 0 {
//		dmg = dmg * crt
//	}
//	if ri > 0 {
//		dmg = dmg * ri
//	}
//	return cast.ToInt64(dmg)
//}

// 浮动伤害计算
// @desc: -10% ~ 10% 伤害浮动
func FluctuateDamage(dmg int64) int64 {
	return cast.ToInt64(cast.ToFloat64(dmg) * Random(-0.1, 0.1))
}

// 普通攻击伤害结算
// @attack 攻击武将
// @suffer 被攻击武将
func AttackDamage(tacticsParams *model.TacticsParams, attackGeneral *vo.BattleGeneral, sufferGeneral *vo.BattleGeneral, attackDmg int64) {
	defer func() {
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
	soldierNum := attackGeneral.SoldierNum
	defSoldierNum := sufferGeneral.SoldierNum

	//虎痴效果
	if effectParams, ok := BuffEffectGet(attackGeneral, consts.BuffEffectType_TigerIdiot_Locked); ok {
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
	if DeBuffEffectContains(attackGeneral, consts.DebuffEffectType_Taunt) && attackGeneral.TauntByGeneral != nil {
		hlog.CtxInfof(ctx, "[%s]执行来自[%s]的「%v」效果",
			sufferGeneral.BaseInfo.Name,
			sufferGeneral.HelpByGeneral.BaseInfo.Name,
			consts.DebuffEffectType_Taunt,
		)
		AttackDamage(tacticsParams, attackGeneral, attackGeneral.TauntByGeneral, 0)
		return
	}
	//援护效果判断
	if BuffEffectContains(attackGeneral, consts.BuffEffectType_Intervene) && sufferGeneral.HelpByGeneral != nil {
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
		if GenerateRate(rate) {
			hlog.CtxInfof(ctx, "[%s]处于规避状态，本次伤害无效", sufferGeneral.BaseInfo.Name)
			return
		} else {
			hlog.CtxInfof(ctx, "[%s]规避失败", sufferGeneral.BaseInfo.Name)
		}
	}

	//伤害计算方式
	if attackDmg == 0 {
		//需要计算
		attackDmg = calculateAttackDmg(soldierNum, attackGeneral, sufferGeneral)
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
	if v, ok := BuffEffectGetAggrEffectRate(sufferGeneral, consts.BuffEffectType_ShareResponsibilityFor); ok {
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
		if effectParams, ok := DeBuffEffectGet(sufferGeneral, consts.DebuffEffectType_ShockingFourRealms_Prepare); ok {
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

	//记录伤兵
	sufferGeneral.LossSoldierNum += attackDmg
	//记录普通次数
	attackGeneral.ExecuteGeneralAttackNum++
	attackGeneral.ExecuteWeaponAttackNum++
	sufferGeneral.SufferExecuteGeneralAttackNum++
	sufferGeneral.SufferExecuteWeaponAttackNum++

	hlog.CtxInfof(ctx, "[%s]损失了兵力%d(%d↘%d)", sufferGeneral.BaseInfo.Name, finalDmg, defSoldierNum, remainSoldierNum)

	if sufferGeneral.SoldierNum == 0 {
		hlog.CtxInfof(ctx, "[%s]武将兵力为0，无法再战", sufferGeneral.BaseInfo.Name)
	}
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
	//伤害
	Damage int64
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

// 战法伤害计算
// @attack 攻击武将
// @suffer 被攻击武将
// @damage 伤害量
// @return 实际伤害/原兵力/剩余兵力
func TacticDamage(param *TacticDamageParam) (damageNum, soldierNum, remainSoldierNum int64, isEffect bool) {
	ctx := param.TacticsParams.Ctx
	tacticsParams := param.TacticsParams
	sufferGeneral := param.SufferGeneral
	attackGeneral := param.AttackGeneral
	damageType := param.DamageType
	damage := param.Damage
	tacticName := param.TacticName
	effectName := param.EffectName
	isIgnoreDefend := param.IsIgnoreDefend
	isEffect = true

	defer func() {
		// 「伤害开始」触发器
		if funcs, okk := attackGeneral.TacticsTriggerMap[consts.BattleAction_DamageEnd]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: attackGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
		// 「遭受伤害结束」触发器
		if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferDamageEnd]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: attackGeneral,
					AttackGeneral:  attackGeneral,
					DamageType:     damageType,
				}
				f(params)
			}
		}
		if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_TacticEnd]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: attackGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
		//「战法攻击后」触发器
		if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferTacticEnd]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: attackGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
		//「发动兵刃/谋略伤害结束」触发器
		battleAction := consts.BattleAction_StrategyDamageEnd
		if damageType == consts.DamageType_Weapon {
			battleAction = consts.BattleAction_WeaponDamageEnd
		}
		if funcs, okk := attackGeneral.TacticsTriggerMap[battleAction]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: attackGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
		sufferBattleAction := consts.BattleAction_SufferStrategyDamageEnd
		if damageType == consts.DamageType_Weapon {
			sufferBattleAction = consts.BattleAction_SufferWeaponDamageEnd
		}
		if funcs, okk := sufferGeneral.TacticsTriggerMap[sufferBattleAction]; okk {
			for _, f := range funcs {
				params := &vo.TacticsTriggerParams{
					CurrentRound:   tacticsParams.CurrentRound,
					CurrentGeneral: attackGeneral,
					AttackGeneral:  attackGeneral,
				}
				f(params)
			}
		}
		//被伤害效果后触发器
		//战法伤害触发器
		sufferEffectTriggerMapping := map[consts.TacticsType]consts.BattleAction{
			consts.TacticsType_Active:        consts.BattleAction_SufferActiveTacticEnd,
			consts.TacticsType_Passive:       consts.BattleAction_SufferPassiveTacticEnd,
			consts.TacticsType_Assault:       consts.BattleAction_SufferAssaultTacticEnd,
			consts.TacticsType_Arm:           consts.BattleAction_SufferArmTacticEnd,
			consts.TacticsType_Command:       consts.BattleAction_SufferCommandTacticEnd,
			consts.TacticsType_TroopsTactics: consts.BattleAction_SufferTroopsTacticEnd,
		}
		action := sufferEffectTriggerMapping[tacticsParams.TacticsType]
		if funcs, ok := sufferGeneral.TacticsTriggerMap[action]; ok {
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

	//必填参数
	if attackGeneral == nil || sufferGeneral == nil || damage <= 0 || damageType == consts.DamageType_None {
		panic("params err")
	}

	//触发器禁用开关
	if tacticName == "连环计" && param.IsBanInterLockedEffect {
		return
	}

	//虎痴效果
	if effectParams, ok := BuffEffectGet(attackGeneral, consts.BuffEffectType_TigerIdiot_Locked); ok {
		if len(effectParams) > 0 {
			effectParam := effectParams[0]
			sufferGeneral = effectParam.LockingTarget
		}
	}

	// 「伤害开始」触发器
	if funcs, okk := attackGeneral.TacticsTriggerMap[consts.BattleAction_Damage]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	// 「遭受伤害开始」触发器
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
	battleAction := consts.BattleAction_StrategyDamage
	if damageType == consts.DamageType_Weapon {
		battleAction = consts.BattleAction_WeaponDamage
	}
	if funcs, okk := attackGeneral.TacticsTriggerMap[battleAction]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	sufferBattleAction := consts.BattleAction_SufferStrategyDamage
	if damageType == consts.DamageType_Weapon {
		sufferBattleAction = consts.BattleAction_SufferWeaponDamage
	}
	if funcs, okk := sufferGeneral.TacticsTriggerMap[sufferBattleAction]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_Tactic]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}
	if funcs, okk := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferTactic]; okk {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: attackGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}

	//被伤害效果触发器
	//映射转换
	sufferEffectTriggerMapping := map[consts.TacticsType]consts.BattleAction{
		consts.TacticsType_Active:        consts.BattleAction_SufferActiveTactic,
		consts.TacticsType_Passive:       consts.BattleAction_SufferPassiveTactic,
		consts.TacticsType_Assault:       consts.BattleAction_SufferAssaultTactic,
		consts.TacticsType_Arm:           consts.BattleAction_SufferArmTactic,
		consts.TacticsType_Command:       consts.BattleAction_SufferCommandTactic,
		consts.TacticsType_TroopsTactics: consts.BattleAction_SufferTroopsTactic,
	}
	action := sufferEffectTriggerMapping[tacticsParams.TacticsType]
	if funcs, ok := sufferGeneral.TacticsTriggerMap[action]; ok {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: sufferGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}

	//是否可以造成伤害
	if !IsCanDamage(ctx, attackGeneral) {
		return 0, 0, 0, false
	}

	//是否可以规避
	if IsCanEvade(ctx, sufferGeneral) {
		return 0, 0, 0, false
	}

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

	//伤害提升效果
	//自带主动战法伤害提升
	if attackGeneral.EquipTactics[0].Id == param.TacticId && consts.ActiveTacticsMap[attackGeneral.EquipTactics[0].Id] {
		if effectParams, ok := BuffEffectGet(attackGeneral, consts.BuffEffectType_TacticsActiveWithSelfDamageImprove); ok {
			effectRate := float64(0)
			for _, effectParam := range effectParams {
				effectRate += effectParam.EffectRate
			}
			damage = cast.ToInt64(cast.ToFloat64(damage) * (1 + effectRate))
		}
	}
	//主动战法伤害提升
	if consts.ActiveTacticsMap[attackGeneral.EquipTactics[0].Id] {
		if effectParams, ok := BuffEffectGet(attackGeneral, consts.BuffEffectType_TacticsActiveDamageImprove); ok {
			effectRate := float64(0)
			for _, effectParam := range effectParams {
				effectRate += effectParam.EffectRate
			}
			damage = cast.ToInt64(cast.ToFloat64(damage) * (1 + effectRate))
		}
	}
	//伤害提升：兵刃/谋略
	switch param.DamageType {
	case consts.DamageType_Weapon:
		if effectParams, ok := BuffEffectGet(attackGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove); ok {
			effectRate := float64(0)
			for _, effectParam := range effectParams {
				effectRate += effectParam.EffectRate
			}
			damage = cast.ToInt64(cast.ToFloat64(damage) * (1 + effectRate))
		}
		//计数
		attackGeneral.ExecuteWeaponAttackNum++
		sufferGeneral.SufferExecuteWeaponAttackNum++
	case consts.DamageType_Strategy:
		if effectParams, ok := BuffEffectGet(attackGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove); ok {
			effectRate := float64(0)
			for _, effectParam := range effectParams {
				effectRate += effectParam.EffectRate
			}
			damage = cast.ToInt64(cast.ToFloat64(damage) * (1 + effectRate))
		}
		//计数
		attackGeneral.ExecuteStrategyAttackNum++
		sufferGeneral.SufferExecuteStrategyAttackNum++
	}

	//伤害分担
	if v, ok := BuffEffectGetAggrEffectRate(sufferGeneral, consts.BuffEffectType_ShareResponsibilityFor); ok {
		hlog.CtxInfof(ctx, "[%s]由于「%v」效果，本次攻击受到的伤害减少了%.2f%%",
			sufferGeneral.BaseInfo.Name,
			consts.BuffEffectType_ShareResponsibilityFor,
			v*100,
		)
		//分担伤害
		shareDmg := cast.ToInt64(cast.ToFloat64(damage) * v)
		//被分担后的伤害
		damage = cast.ToInt64(cast.ToFloat64(damage) * (1 - v))

		hlog.CtxInfof(ctx, "[%s]执行「%v」效果",
			sufferGeneral.ShareResponsibilityForByGeneral.BaseInfo.Name,
			consts.BuffEffectType_ShareResponsibilityFor,
		)
		TacticDamage(&TacticDamageParam{
			TacticsParams:  tacticsParams,
			AttackGeneral:  attackGeneral,
			SufferGeneral:  sufferGeneral.ShareResponsibilityForByGeneral,
			Damage:         shareDmg,
			TacticName:     tacticName,
			IsIgnoreDefend: isIgnoreDefend,
		})
	}

	//伤害计算
	if sufferGeneral.SoldierNum == 0 {
		return 0, 0, 0, false
	}
	//伤害计算公式：攻击者伤害值 - 防御者统率值 = 实际伤害
	//是否无视防御
	if isIgnoreDefend {
		damageNum = damage
	} else {
		damageNum = damage - cast.ToInt64(sufferGeneral.BaseInfo.AbilityAttr.CommandBase)
	}
	//兜底伤害为负的情况
	if damageNum < 0 {
		damageNum = 0
	}
	//兜底伤害大于剩余兵力情况
	soldierNum = sufferGeneral.SoldierNum
	if damageNum >= soldierNum {
		damageNum = soldierNum
	}
	//记录伤兵
	sufferGeneral.LossSoldierNum += damageNum
	//伤害结算
	sufferGeneral.SoldierNum -= damageNum
	remainSoldierNum = sufferGeneral.SoldierNum

	if effectName == "" {
		hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】的伤害，损失了兵力%d(%d↘%d)",
			sufferGeneral.BaseInfo.Name,
			attackGeneral.BaseInfo.Name,
			tacticName,
			damageNum,
			soldierNum,
			remainSoldierNum,
		)
	} else {
		hlog.CtxInfof(ctx, "[%s]由于[%s]【%s】「%v」的伤害，损失了兵力%d(%d↘%d)",
			sufferGeneral.BaseInfo.Name,
			attackGeneral.BaseInfo.Name,
			tacticName,
			effectName,
			damageNum,
			soldierNum,
			remainSoldierNum,
		)
	}

	return
}
