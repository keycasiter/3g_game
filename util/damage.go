package util

import (
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

//浮动伤害计算
//@desc: -10% ~ 10% 伤害浮动
func FluctuateDamage(dmg int64) int64 {
	return cast.ToInt64(cast.ToFloat64(dmg) * Random(-0.1, 0.1))
}

//总伤害计算
// ref: https://baijiahao.baidu.com/s?id=1748672336647348844
//@num 我方携带兵力
//@atk 兵刃伤害：我方武将武力；谋略伤害：我方武将智力
//@def 兵刃防御：我方武将智力；谋略防御：我方武将智力
//@inc 伤害增益比例
//@dec 伤害减益比例
func TotalDamage(num int64,
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
