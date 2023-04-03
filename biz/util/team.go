package util

import (
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 找到当前执行战法武将的队友Map
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

// 找到当前执行战法武将的队友Arr
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

// 找到当前执行战法武将的敌军Map
func GetEnemyGeneralMap(tacticsParams model.TacticsParams) map[int64]*vo.BattleGeneral {
	enemyGeneralMap := make(map[int64]*vo.BattleGeneral, 0)
	currentGeneralId := tacticsParams.CurrentGeneral.BaseInfo.Id
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; !ok {
		enemyGeneralMap = tacticsParams.FightingGeneralMap
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; !ok {
		enemyGeneralMap = tacticsParams.EnemyGeneralMap
	}
	return enemyGeneralMap
}

// 找到当前执行战法武将的敌军Arr
func GetEnemyGeneralArr(tacticsParams model.TacticsParams) []*vo.BattleGeneral {
	enemyGeneralArr := make([]*vo.BattleGeneral, 0)
	currentGeneralId := tacticsParams.CurrentGeneral.BaseInfo.Id
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; !ok {
		for _, general := range tacticsParams.FightingGeneralMap {
			enemyGeneralArr = append(enemyGeneralArr, general)
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; !ok {
		for _, general := range tacticsParams.EnemyGeneralMap {
			enemyGeneralArr = append(enemyGeneralArr, general)
		}
	}
	return enemyGeneralArr
}

// 找到当前执行战法武将的队伍主将
func GetPairMasterGeneral(tacticsParams model.TacticsParams) *vo.BattleGeneral {
	pairGeneralArr := GetPairGeneralArr(tacticsParams)
	for _, general := range pairGeneralArr {
		if general.IsMaster {
			return general
		}
	}
	panic("can't find master general , data error")
	return nil
}

// 找到当前执行战法武将的队伍副将
func GetPairViceGenerals(tacticsParams model.TacticsParams) []*vo.BattleGeneral {
	pairGeneralArr := GetPairGeneralArr(tacticsParams)
	viceGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range pairGeneralArr {
		if !general.IsMaster {
			viceGenerals = append(viceGenerals, general)
		}
	}
	if len(viceGenerals) != 2 {
		panic("vice general num is err , data error")
	}
	return viceGenerals
}

// 找到当前执行战法武将的队伍除自己之外的副将
func GetPairViceGeneralNotSelf(tacticsParams model.TacticsParams) *vo.BattleGeneral {
	pairGeneralArr := GetPairGeneralArr(tacticsParams)
	for _, general := range pairGeneralArr {
		//不是主将，也不是当前执行战法武将自己的副将
		if !general.IsMaster &&
			(tacticsParams.CurrentGeneral.BaseInfo.Id != general.BaseInfo.Id) {
			return general
		}
	}
	panic("can't find vice general not self , data error")
	return nil
}

// 找到当前传入武将之外的两个队友
func GetPairGeneralsNotSelf(tacticsParams model.TacticsParams, targetGeneral *vo.BattleGeneral) []*vo.BattleGeneral {
	//找到队友
	pairGeneralArr := GetPairGeneralArr(tacticsParams)
	pairGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range pairGeneralArr {
		//除传入武将之外的队友
		if general.BaseInfo.Id != targetGeneral.BaseInfo.Id {
			pairGenerals = append(pairGenerals, general)
		}
	}
	if len(pairGenerals) != 2 {
		panic("general not self num is err , data error")
	}
	return pairGenerals
}
