package util

import (
	"github.com/keycasiter/3g_game/biz/model/vo"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"sort"
)

// 找到当前执行战法武将的队友Map
func GetPairGeneralMap(tacticsParams *model.TacticsParams) map[int64]*vo.BattleGeneral {
	pairGeneralMap := make(map[int64]*vo.BattleGeneral, 0)
	currentGeneralId := tacticsParams.CurrentGeneral.BaseInfo.UniqueId
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; ok {
		pairGeneralMap = tacticsParams.FightingGeneralMap
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; ok {
		pairGeneralMap = tacticsParams.EnemyGeneralMap
	}
	return pairGeneralMap
}

// 找到当前执行战法武将的队友Arr
func GetPairGeneralArr(tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	pairGeneralArr := make([]*vo.BattleGeneral, 0)
	currentGeneralId := tacticsParams.CurrentGeneral.BaseInfo.UniqueId
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
func GetEnemyGeneralMap(tacticsParams *model.TacticsParams) map[int64]*vo.BattleGeneral {
	enemyGeneralMap := make(map[int64]*vo.BattleGeneral, 0)
	currentGeneralId := tacticsParams.CurrentGeneral.BaseInfo.UniqueId
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; !ok {
		enemyGeneralMap = tacticsParams.FightingGeneralMap
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; !ok {
		enemyGeneralMap = tacticsParams.EnemyGeneralMap
	}
	return enemyGeneralMap
}

// 找到当前执行战法武将的敌军Arr
func GetEnemyGeneralArr(tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	enemyGeneralArr := make([]*vo.BattleGeneral, 0)
	currentGeneralId := tacticsParams.CurrentGeneral.BaseInfo.UniqueId
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
func GetPairMasterGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	pairGeneralArr := GetPairGeneralArr(tacticsParams)
	for _, general := range pairGeneralArr {
		if general.IsMaster {
			return general
		}
	}
	panic("can't find master general , data error")
	return nil
}

// 找到当前执行战法武将的敌军主将
func GetEnemyMasterGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	enemyGeneralArr := GetEnemyGeneralArr(tacticsParams)
	for _, general := range enemyGeneralArr {
		if general.IsMaster {
			return general
		}
	}
	panic("can't find master general , data error")
	return nil
}

// 找到当前执行战法武将的队伍副将
func GetPairViceGenerals(tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
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

// 找到当前执行战法武将的队伍一个副将
func GetPairViceGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	viceGenerals := GetPairViceGenerals(tacticsParams)
	if len(viceGenerals) == 1 {
		return viceGenerals[0]
	} else {
		if GenerateRate(0.5) {
			return viceGenerals[0]
		} else {
			return viceGenerals[1]
		}
	}
}

// 找到当前执行战法武将的队伍除自己之外的副将
func GetPairViceGeneralNotSelf(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	pairGeneralArr := GetPairGeneralArr(tacticsParams)
	for _, general := range pairGeneralArr {
		//不是主将，也不是当前执行战法武将自己的副将
		if !general.IsMaster &&
			(tacticsParams.CurrentGeneral.BaseInfo.UniqueId != general.BaseInfo.UniqueId) {
			return general
		}
	}
	return nil
}

// 找到当前传入武将之外的两个队友
func GetPairGeneralsNotSelf(tacticsParams *model.TacticsParams, targetGeneral *vo.BattleGeneral) []*vo.BattleGeneral {
	//找到队友
	pairGeneralArr := GetPairGeneralArr(tacticsParams)
	pairGenerals := make([]*vo.BattleGeneral, 0)
	for _, general := range pairGeneralArr {
		//除传入武将之外的队友
		if general.BaseInfo.UniqueId != targetGeneral.BaseInfo.UniqueId {
			pairGenerals = append(pairGenerals, general)
		}
	}
	if len(pairGenerals) != 2 {
		panic("general not self num is err , data error")
	}
	return pairGenerals
}

// 找到当前友军两到三个队友
func GetPairGeneralsTwoOrThreeMap(tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	//找到队友
	pairGeneralArr := GetPairGeneralArr(tacticsParams)
	pairGenerals := make([]*vo.BattleGeneral, 0)
	//两到三个
	totalNum := len(pairGeneralArr)
	hitIdxMap := GenerateHitTwoOrThreeIdxMap(totalNum)
	for idx, general := range pairGeneralArr {
		if _, ok := hitIdxMap[int64(idx)]; ok {
			pairGenerals = append(pairGenerals, general)
		}
	}
	return pairGenerals
}

// 找到当前敌军一个人
func GetEnemyOneGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	//找到敌军
	enemyGeneralArr := GetEnemyGeneralArr(tacticsParams)
	//随机1个人
	totalNum := len(enemyGeneralArr)
	hitIdx := GenerateHitOneIdx(totalNum)
	if enemyGeneralArr[hitIdx] != nil {
		return enemyGeneralArr[hitIdx]
	}
	panic("can't find any one general")
	return nil
}

// 找到当前敌军两个人
func GetEnemyGeneralsTwoArr(tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	//找到敌军
	enemyGeneralArr := GetEnemyGeneralArr(tacticsParams)
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	//随机两个人
	hitNum := 2
	totalNum := len(enemyGeneralArr)
	if totalNum < 2 {
		hitNum = 1
	}
	hitIdxMap := GenerateHitIdxMap(hitNum, totalNum)
	for idx, general := range enemyGeneralArr {
		if _, ok := hitIdxMap[int64(idx)]; ok {
			enemyGenerals = append(enemyGenerals, general)
		}
	}
	return enemyGenerals
}

// 找到当前友军两个队友
func GetPairGeneralsTwoArrByGeneral(currentGeneral *vo.BattleGeneral, tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	pairGeneralArr := make([]*vo.BattleGeneral, 0)
	currentGeneralId := currentGeneral.BaseInfo.UniqueId
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; ok {
		for _, general := range tacticsParams.FightingGeneralMap {
			if currentGeneralId != general.BaseInfo.UniqueId {
				pairGeneralArr = append(pairGeneralArr, general)
			}
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; ok {
		for _, general := range tacticsParams.EnemyGeneralMap {
			if currentGeneralId != general.BaseInfo.UniqueId {
				pairGeneralArr = append(pairGeneralArr, general)
			}
		}
	}
	return pairGeneralArr
}

// 找到当前友军两个队友
func GetPairGeneralsTwoArr(tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	//找到队友
	pairGeneralArr := GetPairGeneralArr(tacticsParams)
	pairGenerals := make([]*vo.BattleGeneral, 0)
	//随机两个队友
	hitNum := 2
	totalNum := len(pairGeneralArr)
	if totalNum < 2 {
		hitNum = 1
	}
	hitIdxMap := GenerateHitIdxMap(hitNum, totalNum)
	for idx, general := range pairGeneralArr {
		if _, ok := hitIdxMap[int64(idx)]; ok {
			pairGenerals = append(pairGenerals, general)
		}
	}
	return pairGenerals
}

// 找到当前敌军两到三个敌人
func GetEnemyGeneralsTwoOrThreeMap(tacticsParams *model.TacticsParams) map[int64]*vo.BattleGeneral {
	//找到敌人
	enemyGeneralArr := GetEnemyGeneralArr(tacticsParams)
	enemyGeneralMap := make(map[int64]*vo.BattleGeneral, 0)
	//两到三个
	totalNum := len(enemyGeneralArr)
	hitIdxMap := GenerateHitTwoOrThreeIdxMap(totalNum)
	for idx, general := range enemyGeneralArr {
		if _, ok := hitIdxMap[int64(idx)]; ok {
			enemyGeneralMap[general.BaseInfo.UniqueId] = general
		}
	}
	return enemyGeneralMap
}

// 找到当前敌军1到2个敌人
func GetEnemyGeneralsOneOrTwoMap(tacticsParams *model.TacticsParams) map[int64]*vo.BattleGeneral {
	//找到敌人
	enemyGeneralArr := make([]*vo.BattleGeneral, 0)
	enemyGeneralMap := make(map[int64]*vo.BattleGeneral, 0)

	if GenerateRate(0.5) {
		enemyGeneralArr = append(enemyGeneralArr, GetEnemyOneGeneral(tacticsParams))
	} else {
		enemyGeneralArr = append(enemyGeneralArr, GetEnemyGeneralsTwoArr(tacticsParams)...)
	}
	for _, general := range enemyGeneralArr {
		enemyGeneralMap[general.BaseInfo.UniqueId] = general
	}
	return enemyGeneralMap
}

// 找到我军损失兵力最多的武将
func GetPairMaxLossSoldierNumGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	pairGenerals := GetPairGeneralArr(tacticsParams)
	if pairGenerals == nil || len(pairGenerals) == 0 {
		return tacticsParams.CurrentGeneral
	}

	//找到我方损失兵力最多的我军单体
	maxLossSoldierNum := pairGenerals[0].LossSoldierNum
	maxLossSoldierNumGeneral := pairGenerals[0]
	for _, general := range pairGenerals {
		if maxLossSoldierNum < general.LossSoldierNum && general.LossSoldierNum > 0 {
			maxLossSoldierNum = general.LossSoldierNum
			maxLossSoldierNumGeneral = general
		}
	}
	return maxLossSoldierNumGeneral
}

// 移除兵力为0武将退场
func RemoveGeneralWhenSoldierNumIsEmpty(tacticsParams *model.TacticsParams) {
	newAllGenerals := make([]*vo.BattleGeneral, 0)

	for _, general := range tacticsParams.AllGeneralArr {
		if general.SoldierNum == 0 {
			delete(tacticsParams.AllGeneralMap, general.BaseInfo.UniqueId)
			delete(tacticsParams.FightingGeneralMap, general.BaseInfo.UniqueId)
			delete(tacticsParams.EnemyGeneralMap, general.BaseInfo.UniqueId)
		} else {
			newAllGenerals = append(newAllGenerals, general)
		}
	}
}

// 按速度排序，从快到慢
func MakeGeneralsOrderBySpeed(allGenerals []*vo.BattleGeneral) []*vo.BattleGeneral {
	var allGeneralsOrderBySpeed vo.BattleGeneralsOrderBySpeed
	allGeneralsOrderBySpeed = append(allGeneralsOrderBySpeed, allGenerals...)
	sort.Sort(allGeneralsOrderBySpeed)
	return allGeneralsOrderBySpeed
}

//获取全部武将
func GetAllGenerals(tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	return tacticsParams.AllGeneralArr
}

//获取我军武力最高的武将
func GetMostForcePairGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	pairGenerals := GetPairGeneralArr(tacticsParams)
	mostForceGeneral := pairGenerals[0]
	for _, general := range pairGenerals {
		if general.BaseInfo.AbilityAttr.ForceBase > mostForceGeneral.BaseInfo.AbilityAttr.ForceBase {
			mostForceGeneral = general
		}
	}
	return mostForceGeneral
}

//获取我军智力最高的武将
func GetMostIntelligencePairGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	pairGenerals := GetPairGeneralArr(tacticsParams)
	mostIntelligenceGeneral := pairGenerals[0]
	for _, general := range pairGenerals {
		if general.BaseInfo.AbilityAttr.IntelligenceBase > mostIntelligenceGeneral.BaseInfo.AbilityAttr.IntelligenceBase {
			mostIntelligenceGeneral = general
		}
	}
	return mostIntelligenceGeneral
}

//获取敌军武力最高的武将
func GetMostForceEnemyGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	enemyGenerals := GetEnemyGeneralArr(tacticsParams)
	mostForceGeneral := enemyGenerals[0]
	for _, general := range enemyGenerals {
		if general.BaseInfo.AbilityAttr.ForceBase > mostForceGeneral.BaseInfo.AbilityAttr.ForceBase {
			mostForceGeneral = general
		}
	}
	return mostForceGeneral
}

//获取敌军智力最高的武将
func GetMostIntelligenceEnemyGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	enemyGenerals := GetEnemyGeneralArr(tacticsParams)
	mostIntelligenceGeneral := enemyGenerals[0]
	for _, general := range enemyGenerals {
		if general.BaseInfo.AbilityAttr.IntelligenceBase > mostIntelligenceGeneral.BaseInfo.AbilityAttr.IntelligenceBase {
			mostIntelligenceGeneral = general
		}
	}
	return mostIntelligenceGeneral
}
