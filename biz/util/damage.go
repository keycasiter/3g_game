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
func AttackDamage(tacticsParams *model.TacticsParams, attackGeneral *vo.BattleGeneral, sufferGeneral *vo.BattleGeneral) {
	ctx := tacticsParams.Ctx
	soldierNum := attackGeneral.SoldierNum
	defSoldierNum := sufferGeneral.SoldierNum

	hlog.CtxInfof(ctx, "[%s]对[%s]发动普通攻击",
		attackGeneral.BaseInfo.Name,
		sufferGeneral.BaseInfo.Name,
	)

	//被伤害效果触发器
	if funcs, ok := sufferGeneral.TacticsTriggerMap[consts.BattleAction_SufferAttack]; ok {
		for _, f := range funcs {
			params := &vo.TacticsTriggerParams{
				CurrentRound:   tacticsParams.CurrentRound,
				CurrentGeneral: sufferGeneral,
				AttackGeneral:  attackGeneral,
			}
			f(params)
		}
	}

	//援护效果判断
	if sufferGeneral.HelpByGeneral != nil {
		hlog.CtxInfof(ctx, "[%s]执行来自[%s]的「%v」效果",
			sufferGeneral.BaseInfo.Name,
			sufferGeneral.HelpByGeneral.BaseInfo.Name,
			consts.BuffEffectType_Intervene,
		)
		AttackDamage(tacticsParams, attackGeneral, sufferGeneral.HelpByGeneral)
		return
	}

	//是否可以规避
	if rate, ok := sufferGeneral.BuffEffectHolderMap[consts.BuffEffectType_Evade]; ok {
		if GenerateRate(rate) {
			hlog.CtxInfof(ctx, "[%s]处于规避状态，本次伤害无效", sufferGeneral.BaseInfo.Name)
			return
		} else {
			hlog.CtxInfof(ctx, "[%s]规避失败", sufferGeneral.BaseInfo.Name)
		}
	}

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
	inc := attackGeneral.BuffEffectHolderMap[consts.BuffEffectType_LaunchWeaponDamageImprove] +
		sufferGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_SufferWeaponDamageImprove]
	dec := attackGeneral.DeBuffEffectHolderMap[consts.DebuffEffectType_LaunchWeaponDamageDeduce] +
		sufferGeneral.BuffEffectHolderMap[consts.BuffEffectType_SufferWeaponDamageDeduce]

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

	//hlog.CtxInfof(ctx, "兵力基础伤害:%d ,武力/防御差:%.2f , 最终伤害:%d , 攻击者武力:%.2f , 防守者统率:%.2f , 造成+受到兵刃伤害增加:%.2f%% , 造成+受到兵刃伤害减少:%.2f%% , 最终增减伤率:%.2f",
	//	int64(numDmg), atkDefDmg, attackDmg, atk, def, inc*100, dec*100, incDecRate)

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

	hlog.CtxInfof(ctx, "[%s]损失了兵力%d(%d↘%d)", sufferGeneral.BaseInfo.Name, finalDmg, defSoldierNum, remainSoldierNum)

	if sufferGeneral.SoldierNum == 0 {
		hlog.CtxInfof(ctx, "[%s]武将兵力为0，无法再战", sufferGeneral.BaseInfo.Name)
	}
}

// 战法伤害计算
// @attack 攻击武将
// @suffer 被攻击武将
// @damage 伤害量
// @return 实际伤害/原兵力/剩余兵力
func TacticDamage(tacticsParams *model.TacticsParams, attackGeneral *vo.BattleGeneral, sufferGeneral *vo.BattleGeneral, damage int64) (damageNum, soldierNum, remainSoldierNum int64, isEffect bool) {
	ctx := tacticsParams.Ctx
	isEffect = true

	//是否可以规避
	if rate, ok := sufferGeneral.BuffEffectHolderMap[consts.BuffEffectType_Evade]; ok {
		if GenerateRate(rate) {
			hlog.CtxInfof(ctx, "[%s]处于规避状态，本次伤害无效", sufferGeneral.BaseInfo.Name)
			return 0, sufferGeneral.SoldierNum, sufferGeneral.SoldierNum, false
		} else {
			hlog.CtxInfof(ctx, "[%s]规避失败", sufferGeneral.BaseInfo.Name)
		}
	}

	//伤害结算
	if sufferGeneral.SoldierNum == 0 {
		return 0, 0, 0, false
	}
	damageNum = damage
	soldierNum = sufferGeneral.SoldierNum

	if damage >= soldierNum {
		damageNum = soldierNum
	}
	//记录伤兵
	sufferGeneral.LossSoldierNum += damageNum
	//伤害结算
	sufferGeneral.SoldierNum -= damageNum
	remainSoldierNum = sufferGeneral.SoldierNum

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

	return
}
