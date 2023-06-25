package util

import (
	"github.com/keycasiter/3g_game/biz/consts"
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

// 找到兵力最低武将
func GetLowestSoliderNumGeneral(generals []*vo.BattleGeneral) *vo.BattleGeneral {
	lowestGeneral := generals[0]
	for _, general := range generals {
		if general.SoldierNum < lowestGeneral.SoldierNum {
			lowestGeneral = general
		}
	}
	return lowestGeneral
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

// 找到敌军一名男性武将
func GetEnemyOneMaleGeneral(tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	enemyGenerals := GetEnemyGeneralArr(tacticsParams)

	maleGenerals := make([]*vo.BattleGeneral, 0)
	for _, enemyGeneral := range enemyGenerals {
		if enemyGeneral.BaseInfo.Gender == consts.Gender_Male {
			maleGenerals = append(maleGenerals, enemyGeneral)
		}
	}
	if len(maleGenerals) == 0 {
		return nil
	}
	hitIdx := GenerateHitOneIdx(len(maleGenerals))
	return maleGenerals[hitIdx]
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

// 找到当前我军任意一人
func GetPairOneGeneral(tacticsParams *model.TacticsParams, general *vo.BattleGeneral) *vo.BattleGeneral {
	//找到我军
	pairGenerals := make([]*vo.BattleGeneral, 0)
	if _, ok := tacticsParams.FightingGeneralMap[general.BaseInfo.UniqueId]; ok {
		for _, battleGeneral := range tacticsParams.FightingGeneralMap {
			pairGenerals = append(pairGenerals, battleGeneral)
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[general.BaseInfo.UniqueId]; ok {
		for _, battleGeneral := range tacticsParams.EnemyGeneralMap {
			pairGenerals = append(pairGenerals, battleGeneral)
		}
	}

	//随机1个人
	totalNum := len(pairGenerals)
	hitIdx := GenerateHitOneIdx(totalNum)
	if pairGenerals[hitIdx] != nil {
		return pairGenerals[hitIdx]
	}
	return nil
}

// 找到当前我军除自己之外一个人
func GetPairOneGeneralNotSelf(tacticsParams *model.TacticsParams, general *vo.BattleGeneral) *vo.BattleGeneral {
	//找到我军
	pairGeneralArrs := GetPairGeneralsNotSelf(tacticsParams, general)
	//随机1个人
	totalNum := len(pairGeneralArrs)
	hitIdx := GenerateHitOneIdx(totalNum)
	if pairGeneralArrs[hitIdx] != nil {
		return pairGeneralArrs[hitIdx]
	}
	return nil
}

// 找到兵力最低友军
func GetPairLowestSoldierNumGeneral(tacticsParams *model.TacticsParams, general *vo.BattleGeneral) *vo.BattleGeneral {
	pairGenerals := GetPairGeneralsNotSelf(tacticsParams, general)
	if len(pairGenerals) == 0 {
		return nil
	}
	lowestSoliderGeneral := pairGenerals[0]
	for _, pairGeneral := range pairGenerals {
		//找到兵力最低的友军单体
		if pairGeneral.SoldierNum < lowestSoliderGeneral.SoldierNum {
			lowestSoliderGeneral = pairGeneral
		}
	}
	return lowestSoliderGeneral
}

// 找到武将除自己之外的队友们
func GetPairGeneralsNotSelf(tacticsParams *model.TacticsParams, general *vo.BattleGeneral) []*vo.BattleGeneral {
	//找到我军
	pairGenerals := make([]*vo.BattleGeneral, 0)
	if _, ok := tacticsParams.FightingGeneralMap[general.BaseInfo.UniqueId]; ok {
		for uniqueId, perGeneral := range tacticsParams.FightingGeneralMap {
			if uniqueId != general.BaseInfo.UniqueId {
				pairGenerals = append(pairGenerals, perGeneral)
			}
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[general.BaseInfo.UniqueId]; ok {
		for uniqueId, perGeneral := range tacticsParams.EnemyGeneralMap {
			if uniqueId != general.BaseInfo.UniqueId {
				pairGenerals = append(pairGenerals, perGeneral)
			}
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

func GetEnemyGeneralsByGeneral(general *vo.BattleGeneral, tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	currentGeneralId := general.BaseInfo.UniqueId
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; !ok {
		for _, general := range tacticsParams.FightingGeneralMap {
			enemyGenerals = append(enemyGenerals, general)
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; !ok {
		for _, general := range tacticsParams.EnemyGeneralMap {
			enemyGenerals = append(enemyGenerals, general)
		}
	}
	return enemyGenerals
}

func GetEnemyOneGeneralNotSelfByGeneral(general *vo.BattleGeneral, tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	currentGeneralId := general.BaseInfo.UniqueId
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; !ok {
		for _, each := range tacticsParams.FightingGeneralMap {
			if each.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
				continue
			}
			enemyGenerals = append(enemyGenerals, each)
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; !ok {
		for _, each := range tacticsParams.EnemyGeneralMap {
			if each.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
				continue
			}
			enemyGenerals = append(enemyGenerals, each)
		}
	}

	hitIdx := GenerateHitOneIdx(len(enemyGenerals))

	return enemyGenerals[hitIdx]
}

func GetEnemyGeneralsNotSelfByGeneral(general *vo.BattleGeneral, tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	currentGeneralId := general.BaseInfo.UniqueId
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; !ok {
		for _, each := range tacticsParams.FightingGeneralMap {
			if each.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
				continue
			}
			enemyGenerals = append(enemyGenerals, each)
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; !ok {
		for _, each := range tacticsParams.EnemyGeneralMap {
			if each.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
				continue
			}
			enemyGenerals = append(enemyGenerals, each)
		}
	}
	return enemyGenerals
}

func GetAllGeneralsNotSelfByGeneral(general *vo.BattleGeneral, tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	allGenerals := make([]*vo.BattleGeneral, 0)
	for _, each := range tacticsParams.FightingGeneralMap {
		if each.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
			continue
		}
		allGenerals = append(allGenerals, each)
	}
	for _, each := range tacticsParams.EnemyGeneralMap {
		if each.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
			continue
		}
		allGenerals = append(allGenerals, each)
	}
	return allGenerals
}

func GetEnemyOneGeneralByGeneral(general *vo.BattleGeneral, tacticsParams *model.TacticsParams) *vo.BattleGeneral {
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	currentGeneralId := general.BaseInfo.UniqueId
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; !ok {
		for _, general := range tacticsParams.FightingGeneralMap {
			enemyGenerals = append(enemyGenerals, general)
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; !ok {
		for _, general := range tacticsParams.EnemyGeneralMap {
			enemyGenerals = append(enemyGenerals, general)
		}
	}
	hitIdx := GenerateHitOneIdx(len(enemyGenerals))
	return enemyGenerals[hitIdx]
}

func GetEnemyTwoGeneralByGeneral(general *vo.BattleGeneral, tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	enemyGenerals := make([]*vo.BattleGeneral, 0)
	resGenerals := make([]*vo.BattleGeneral, 0)
	currentGeneralId := general.BaseInfo.UniqueId
	if _, ok := tacticsParams.FightingGeneralMap[currentGeneralId]; !ok {
		for _, general := range tacticsParams.FightingGeneralMap {
			enemyGenerals = append(enemyGenerals, general)
		}
	}
	if _, ok := tacticsParams.EnemyGeneralMap[currentGeneralId]; !ok {
		for _, general := range tacticsParams.EnemyGeneralMap {
			enemyGenerals = append(enemyGenerals, general)
		}
	}
	hitIdxArr := GenerateHitIdxArr(2, len(enemyGenerals))
	for _, idx := range hitIdxArr {
		resGenerals = append(resGenerals, enemyGenerals[idx])
	}
	return resGenerals
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

// 找到当前敌军1到2个敌人
func GetEnemyGeneralsOneOrTwoArr(tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	//找到敌人
	enemyGeneralArr := make([]*vo.BattleGeneral, 0)

	if GenerateRate(0.5) {
		enemyGeneralArr = append(enemyGeneralArr, GetEnemyOneGeneral(tacticsParams))
	} else {
		enemyGeneralArr = append(enemyGeneralArr, GetEnemyGeneralsTwoArr(tacticsParams)...)
	}
	return enemyGeneralArr
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

// 获取全部武将
func GetAllGenerals(tacticsParams *model.TacticsParams) []*vo.BattleGeneral {
	return tacticsParams.AllGeneralArr
}

// 获取我军武力最高的武将
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

// 获取我军智力最高的武将
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

// 获取敌军武力最高的武将
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

// 获取敌军智力最高的武将
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

// 是否为友军
func IsPair(currentGeneral *vo.BattleGeneral, compareGeneral *vo.BattleGeneral, tacticsParams *model.TacticsParams) bool {
	currentFlag := false
	compareFlag := false
	for _, general := range tacticsParams.FightingGeneralMap {
		if currentGeneral.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
			currentFlag = true
		}
		if compareGeneral.BaseInfo.UniqueId == compareGeneral.BaseInfo.UniqueId {
			compareFlag = true
		}
	}
	if currentFlag && compareFlag {
		return true
	}
	return false
}

// 获取我军兵种
func GetPairArmType(currentGeneral *vo.BattleGeneral, tacticsParams *model.TacticsParams) consts.ArmType {
	for _, general := range tacticsParams.FightingTeam.BattleGenerals {
		if currentGeneral.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
			return tacticsParams.FightingTeam.ArmType
		}
	}
	for _, general := range tacticsParams.EnemyTeam.BattleGenerals {
		if currentGeneral.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
			return tacticsParams.EnemyTeam.ArmType
		}
	}
	return consts.ArmType_Unknow
}

// 获取敌军兵种
func GetEnemyArmType(currentGeneral *vo.BattleGeneral, tacticsParams *model.TacticsParams) consts.ArmType {
	isFighting := false
	for _, general := range tacticsParams.FightingTeam.BattleGenerals {
		if currentGeneral.BaseInfo.UniqueId == general.BaseInfo.UniqueId {
			isFighting = true
			break
		}
	}
	if isFighting {
		return tacticsParams.EnemyTeam.ArmType
	} else {
		return tacticsParams.FightingTeam.ArmType
	}
}

// 获取武将武力/智力最高一项
func GetGeneralHighestBetweenForceOrIntelligence(general *vo.BattleGeneral) (consts.AbilityAttr, float64) {
	if general.BaseInfo.AbilityAttr.ForceBase > general.BaseInfo.AbilityAttr.IntelligenceBase {
		return consts.AbilityAttr_Force, general.BaseInfo.AbilityAttr.ForceBase
	}
	return consts.AbilityAttr_Intelligence, general.BaseInfo.AbilityAttr.IntelligenceBase
}

// 获取敌军智力最高的武将
func GetEnemyGeneralWhoIsHighestIntelligence(params *model.TacticsParams) *vo.BattleGeneral {
	enemyGenerals := GetEnemyGeneralArr(params)
	highestGeneral := enemyGenerals[0]
	for _, general := range enemyGenerals {
		if general.BaseInfo.AbilityAttr.IntelligenceBase > highestGeneral.BaseInfo.AbilityAttr.IntelligenceBase {
			highestGeneral = general
		}
	}
	return highestGeneral
}

// 获取敌军武力最高的武将
func GetEnemyGeneralWhoIsHighestForce(params *model.TacticsParams) *vo.BattleGeneral {
	enemyGenerals := GetEnemyGeneralArr(params)
	highestGeneral := enemyGenerals[0]
	for _, general := range enemyGenerals {
		if general.BaseInfo.AbilityAttr.ForceBase > highestGeneral.BaseInfo.AbilityAttr.ForceBase {
			highestGeneral = general
		}
	}
	return highestGeneral
}

// 获取我军武力最高的武将
func GetPairGeneralWhoIsHighestForce(params *model.TacticsParams) *vo.BattleGeneral {
	pairGenerals := GetPairGeneralArr(params)
	highestGeneral := pairGenerals[0]
	for _, general := range pairGenerals {
		if general.BaseInfo.AbilityAttr.ForceBase > highestGeneral.BaseInfo.AbilityAttr.ForceBase {
			highestGeneral = general
		}
	}
	return highestGeneral
}

// 获取我军智力最高的武将
func GetPairGeneralWhoIsHighestIntelligence(params *model.TacticsParams) *vo.BattleGeneral {
	pairGenerals := GetPairGeneralArr(params)
	highestGeneral := pairGenerals[0]
	for _, general := range pairGenerals {
		if general.BaseInfo.AbilityAttr.IntelligenceBase > highestGeneral.BaseInfo.AbilityAttr.IntelligenceBase {
			highestGeneral = general
		}
	}
	return highestGeneral
}

// 获取统率最低的敌军单体
func GetEnemyGeneralWhoLowestCommand(currentGeneral *vo.BattleGeneral, params *model.TacticsParams) *vo.BattleGeneral {
	//找到敌军全体
	enemyGenerals := GetEnemyGeneralsByGeneral(currentGeneral, params)
	//找到统率最低的敌军单体
	lowestCommandGeneral := enemyGenerals[0]
	for _, general := range enemyGenerals {
		if general.BaseInfo.AbilityAttr.CommandBase < lowestCommandGeneral.BaseInfo.AbilityAttr.CommandBase {
			lowestCommandGeneral = general
		}
	}
	return lowestCommandGeneral
}

// 提高武将属性
func ImproveGeneralAttr(general *vo.BattleGeneral, attr consts.AbilityAttr, attrValue float64) {
	switch attr {
	//武力
	case consts.AbilityAttr_Force:
		general.BaseInfo.AbilityAttr.ForceBase += attrValue
	//智力
	case consts.AbilityAttr_Intelligence:
		general.BaseInfo.AbilityAttr.IntelligenceBase += attrValue
	//统率
	case consts.AbilityAttr_Command:
		general.BaseInfo.AbilityAttr.CommandBase += attrValue
	//速度
	case consts.AbilityAttr_Speed:
		general.BaseInfo.AbilityAttr.SpeedBase += attrValue
	}
}

// 降低武将属性
func DeduceGeneralAttr(general *vo.BattleGeneral, attr consts.AbilityAttr, attrValue float64) {
	switch attr {
	//武力
	case consts.AbilityAttr_Force:
		general.BaseInfo.AbilityAttr.ForceBase -= attrValue
	//智力
	case consts.AbilityAttr_Intelligence:
		general.BaseInfo.AbilityAttr.IntelligenceBase -= attrValue
	//统率
	case consts.AbilityAttr_Command:
		general.BaseInfo.AbilityAttr.CommandBase -= attrValue
	//速度
	case consts.AbilityAttr_Speed:
		general.BaseInfo.AbilityAttr.SpeedBase -= attrValue
	}

	//兜底
	if general.BaseInfo.AbilityAttr.ForceBase < 0 {
		general.BaseInfo.AbilityAttr.ForceBase = 0
	}
	if general.BaseInfo.AbilityAttr.IntelligenceBase < 0 {
		general.BaseInfo.AbilityAttr.IntelligenceBase = 0
	}
	if general.BaseInfo.AbilityAttr.CommandBase < 0 {
		general.BaseInfo.AbilityAttr.CommandBase = 0
	}
	if general.BaseInfo.AbilityAttr.SpeedBase < 0 {
		general.BaseInfo.AbilityAttr.SpeedBase = 0
	}
}

// 获取兵种适性较高武将
func GetHighestArmAbilityGeneral(generals []*vo.BattleGeneral, armType consts.ArmType) *vo.BattleGeneral {
	highestGeneral := generals[0]
	for _, general := range generals {
		switch armType {
		case consts.ArmType_Apparatus:
			if general.BaseInfo.ArmsAttr.Apparatus < highestGeneral.BaseInfo.ArmsAttr.Apparatus {
				highestGeneral = general
			}
		case consts.ArmType_Mauler:
			if general.BaseInfo.ArmsAttr.Mauler < highestGeneral.BaseInfo.ArmsAttr.Mauler {
				highestGeneral = general
			}
		case consts.ArmType_Cavalry:
			if general.BaseInfo.ArmsAttr.Cavalry < highestGeneral.BaseInfo.ArmsAttr.Cavalry {
				highestGeneral = general
			}
		case consts.ArmType_Spearman:
			if general.BaseInfo.ArmsAttr.Spearman < highestGeneral.BaseInfo.ArmsAttr.Spearman {
				highestGeneral = general
			}
		case consts.ArmType_Archers:
			if general.BaseInfo.ArmsAttr.Archers < highestGeneral.BaseInfo.ArmsAttr.Archers {
				highestGeneral = general
			}
		}
	}
	return highestGeneral
}

// 获取兵种适性较低武将
func GetLowestArmAbilityGeneral(generals []*vo.BattleGeneral, armType consts.ArmType) *vo.BattleGeneral {
	lowestGeneral := generals[0]
	for _, general := range generals {
		switch armType {
		case consts.ArmType_Apparatus:
			if general.BaseInfo.ArmsAttr.Apparatus > lowestGeneral.BaseInfo.ArmsAttr.Apparatus {
				lowestGeneral = general
			}
		case consts.ArmType_Mauler:
			if general.BaseInfo.ArmsAttr.Mauler > lowestGeneral.BaseInfo.ArmsAttr.Mauler {
				lowestGeneral = general
			}
		case consts.ArmType_Cavalry:
			if general.BaseInfo.ArmsAttr.Cavalry > lowestGeneral.BaseInfo.ArmsAttr.Cavalry {
				lowestGeneral = general
			}
		case consts.ArmType_Spearman:
			if general.BaseInfo.ArmsAttr.Spearman > lowestGeneral.BaseInfo.ArmsAttr.Spearman {
				lowestGeneral = general
			}
		case consts.ArmType_Archers:
			if general.BaseInfo.ArmsAttr.Archers > lowestGeneral.BaseInfo.ArmsAttr.Archers {
				lowestGeneral = general
			}
		}
	}
	return lowestGeneral
}

// 找到武将最高属性
func GetGeneralHighestAttr(general *vo.BattleGeneral) (attr consts.AbilityAttr, attrValue float64) {
	highAttr := general.BaseInfo.AbilityAttr.ForceBase
	att := consts.AbilityAttr_Force

	if general.BaseInfo.AbilityAttr.IntelligenceBase > highAttr {
		highAttr = general.BaseInfo.AbilityAttr.IntelligenceBase
		att = consts.AbilityAttr_Intelligence
	}
	if general.BaseInfo.AbilityAttr.CommandBase > highAttr {
		highAttr = general.BaseInfo.AbilityAttr.CommandBase
		att = consts.AbilityAttr_Command
	}
	if general.BaseInfo.AbilityAttr.SpeedBase > highAttr {
		highAttr = general.BaseInfo.AbilityAttr.SpeedBase
		att = consts.AbilityAttr_Speed
	}

	return att, highAttr
}

// 计算属性之差
func CalculateAttrDiff(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return 0
}

// 是否包含武将标签
func IsContainsGeneralTag(generalTags []consts.GeneralTag, generalTag consts.GeneralTag) bool {
	for _, tag := range generalTags {
		if tag == generalTag {
			return true
		}
	}
	return false
}
