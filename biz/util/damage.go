package util

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
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

// 普通攻击伤害计算
// ref: https://baijiahao.baidu.com/s?id=1748672336647348844
// @num 我方携带兵力
// @atk 兵刃伤害：我方武将武力；谋略伤害：我方武将智力
// @def 兵刃防御：我方武将智力；谋略防御：我方武将智力
// @inc 伤害增益比例
// @dec 伤害减益比例
func GeneralAttackDamage(num int64,
	atk float64,
	def float64,
	inc float64,
	dec float64) int64 {

	//兵力伤害计算
	numDmg := float64(0)
	//0～3000兵力，伤害为兵力的10%
	if num > 0 && num <= 3000 {
		numDmg = cast.ToFloat64(num) * 0.1
	}
	//3000～6000兵力，1000兵力=50伤害
	if num > 3000 && num <= 6000 {
		numDmg = 3000 * 0.1
		numDmg += cast.ToFloat64((num - 3000) / 1000 * 50)
	}
	//6000～Max兵力，1000兵力=30伤害
	if num > 6000 {
		numDmg = 3000 * 0.1
		numDmg += cast.ToFloat64((num / 1000) * 50)
		numDmg += cast.ToFloat64((num - 3000) / 1000 * 50)
		numDmg += cast.ToFloat64((num - 6000) / 1000 * 30)
	}
	// 总伤害 = (兵力伤害 + 攻击 - 防御) * (1 + 增益比例 - 减益比例)
	return cast.ToInt64((numDmg + (atk - def)) * (1 + inc - dec))
}

// 普通攻击伤害结算
//@attack 攻击武将
//@suffer 被攻击武将
func AttackDamage(ctx context.Context, attackGeneral *vo.BattleGeneral, sufferGeneral *vo.BattleGeneral) {
	soldierNum := attackGeneral.SoldierNum
	defSoldierNum := sufferGeneral.SoldierNum

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
	inc := attackGeneral.BuffEffectHolderMap[consts.BuffEffectType_LaunchWeaponDamageImprove]
	dec := attackGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchWeaponDamageDeduce]

	// 总伤害 = (兵力伤害 + 攻击 - 防御) * (1 + 增益比例 - 减益比例)
	attackDmg := cast.ToInt64((numDmg + (atk - def)) * (1 + inc - dec))
	hlog.CtxInfof(ctx, "numDmg:%f , atk:%f , def:%f , inc:%f , dec:%f", numDmg, atk, def, inc, dec)

	hlog.CtxInfof(ctx, "[%s]对[%s]发动普通攻击",
		attackGeneral.BaseInfo.Name,
		sufferGeneral.BaseInfo.Name,
	)

	//伤害计算
	finalDmg, remainSoldierNum := CalculateDamage(defSoldierNum, attackDmg)
	sufferGeneral.SoldierNum = remainSoldierNum
	hlog.CtxInfof(ctx, "[%s]损失了兵力%d(%d↘%d)", sufferGeneral.BaseInfo.Name, finalDmg, defSoldierNum, remainSoldierNum)
}

// 伤害计算
// @soldierNum 被攻击者当前兵力
// @damage 伤害值
// @return 实际伤害/剩余兵力
func CalculateDamage(soldierNum int64, damage int64) (int64, int64) {
	if soldierNum == 0 {
		return 0, 0
	}

	if damage >= soldierNum {
		return soldierNum, soldierNum
	}

	return damage, soldierNum - damage
}

// 战法伤害计算
// @attack 攻击武将
// @suffer 被攻击武将
// @damage 伤害量
// @return 实际伤害/原兵力/剩余兵力
func TacticDamage(ctx context.Context, attackGeneral *vo.BattleGeneral, sufferGeneral *vo.BattleGeneral, damage int64) (damageNum, soldierNum, remainSoldierNum int64) {
	//是否可以规避
	if rate, ok := sufferGeneral.BuffEffectHolderMap[consts.BuffEffectType_Evade]; ok {
		if GenerateRate(rate) {
			hlog.CtxInfof(ctx, "[%s]处于规避状态，本次伤害无效")
			return 0, sufferGeneral.SoldierNum, sufferGeneral.SoldierNum
		} else {
			hlog.CtxInfof(ctx, "[%s]规避失败")
		}
	}

	//伤害结算
	if sufferGeneral.SoldierNum == 0 {
		return 0, 0, 0
	}
	damageNum = damage
	soldierNum = sufferGeneral.SoldierNum

	if damage >= soldierNum {
		damageNum = soldierNum
	}
	//伤害结算
	sufferGeneral.SoldierNum -= damageNum
	remainSoldierNum = sufferGeneral.SoldierNum
	return
}
