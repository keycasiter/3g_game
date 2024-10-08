package util

import (
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/po"
)

//计算武将兵种适性影响后的属性值
func CalGeneralArmAbility(armAbility consts.ArmsAbility, attr *po.AbilityAttr) {
	switch armAbility {
	case consts.ArmsAbility_S:
		attr.IntelligenceBase = attr.IntelligenceBase * 1.2
		attr.ForceBase = attr.ForceBase * 1.2
		attr.CommandBase = attr.CommandBase * 1.2
		attr.SpeedBase = attr.SpeedBase * 1.2
	case consts.ArmsAbility_A:
		attr.IntelligenceBase = attr.IntelligenceBase * 1
		attr.ForceBase = attr.ForceBase * 1
		attr.CommandBase = attr.CommandBase * 1
		attr.SpeedBase = attr.SpeedBase * 1
	case consts.ArmsAbility_B:
		attr.IntelligenceBase = attr.IntelligenceBase * 0.85
		attr.ForceBase = attr.ForceBase * 0.85
		attr.CommandBase = attr.CommandBase * 0.85
		attr.SpeedBase = attr.SpeedBase * 0.85
	case consts.ArmsAbility_C:
		attr.IntelligenceBase = attr.IntelligenceBase * 0.7
		attr.ForceBase = attr.ForceBase * 0.7
		attr.CommandBase = attr.CommandBase * 0.7
		attr.SpeedBase = attr.SpeedBase * 0.7
	}
}
