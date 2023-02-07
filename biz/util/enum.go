package util

import "github.com/keycasiter/3g_game/biz/consts"

//兵种类型转义
func TranslateArmType(armType consts.ArmType) string {
	switch armType {
	case consts.ArmType_Cavalry:
		return "骑兵"
	case consts.ArmType_Mauler:
		return "盾兵"
	case consts.ArmType_Archers:
		return "弓兵"
	case consts.ArmType_Spearman:
		return "枪兵"
	case consts.ArmType_Apparatus:
		return "器械"
	}
	return "未知兵种"
}

//兵种适性转义
func TranslateArmsAbility(armsAbility consts.ArmsAbility) string {
	switch armsAbility {
	case consts.ArmsAbility_S:
		return "S"
	case consts.ArmsAbility_A:
		return "A"
	case consts.ArmsAbility_B:
		return "B"
	case consts.ArmsAbility_C:
		return "C"
	}
	return "未知兵种适性"
}

//兵种适性属性加成转义
func TranslateArmsAbilityAddition(armsAbility consts.ArmsAbility) string {
	switch armsAbility {
	case consts.ArmsAbility_S:
		return "120%"
	case consts.ArmsAbility_A:
		return "100%"
	case consts.ArmsAbility_B:
		return "85%"
	case consts.ArmsAbility_C:
		return "70%"
	}
	return "未知兵种适性"
}

//阵营转义
func TranslateGroup(group consts.Group) string {
	switch group {
	case consts.Group_WeiGuo:
		return "魏"
	case consts.Group_ShuGuo:
		return "蜀"
	case consts.Group_WuGuo:
		return "吴"
	case consts.Group_QunXiong:
		return "群%"
	}
	return "未知阵营"
}
