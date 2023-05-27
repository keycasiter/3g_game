package util

import "github.com/keycasiter/3g_game/biz/model/vo"

//获取武将武力/智力最高一项
func GetGeneralHighestBetweenForceOrIntelligence(general *vo.BattleGeneral) float64 {
	if general.BaseInfo.AbilityAttr.ForceBase > general.BaseInfo.AbilityAttr.IntelligenceBase {
		return general.BaseInfo.AbilityAttr.ForceBase
	}
	return general.BaseInfo.AbilityAttr.IntelligenceBase
}
