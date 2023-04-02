package util

import (
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//找到当前执行战法武将的队友Map
func GetPairGeneralMap(tacticsParams model.TacticsParams) map[int64]*vo.BattleGeneral {
	pairGeneralMap := make(map[int64]*vo.BattleGeneral, 0)
	currentGeneralId := tacticsParams.CurrentGeneral.BaseInfo.Id
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; ok {
		pairGeneralMap = tacticsParams.FightingGeneralMap
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; ok {
		pairGeneralMap = tacticsParams.EnemyGeneralMap
	}
	return pairGeneralMap
}

//找到当前执行战法武将的队友Arr
func GetPairGeneralArr(tacticsParams model.TacticsParams) []*vo.BattleGeneral {
	pairGeneralArr := make([]*vo.BattleGeneral, 0)
	currentGeneralId := tacticsParams.CurrentGeneral.BaseInfo.Id
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; ok {
		for _, general := range tacticsParams.FightingGeneralMap {
			pairGeneralArr = append(pairGeneralArr, general)
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; ok {
		for _, general := range tacticsParams.EnemyGeneralMap {
			pairGeneralArr = append(pairGeneralArr, general)
		}
	}
	return pairGeneralArr
}
