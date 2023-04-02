package util

import (
	"github.com/spf13/cast"
	"math/rand"
	"time"
)

// 生成随机数，min为下限，max为上限
func Random(min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + cast.ToFloat64(min)
}

//根据概率设定返回是否成功
func GenerateRate(rate float64) bool {
	rand := Random(0, 1)
	if rand >= 0 && rand <= rate {
		return true
	}
	return false
}

//生成totalNum数中命中hitNum数量的索引数组
func GenerateHitIdxArr(hitNum, totalNum int) []int64 {
	idxMap := make(map[int64]bool, 0)
	for {
		if len(idxMap) == hitNum {
			break
		}

		randNum := Random(0, float64(totalNum))

		if _, ok := idxMap[int64(randNum)]; ok {
			continue
		} else {
			idxMap[int64(randNum)] = true
		}
	}

	idxArr := make([]int64, 0)
	for k, _ := range idxMap {
		idxArr = append(idxArr, k)
	}
	return idxArr
}

func GenerateHitIdxMap(hitNum, totalNum int) map[int64]bool {
	idxMap := make(map[int64]bool, 0)
	for {
		if len(idxMap) == hitNum {
			break
		}

		randNum := Random(0, float64(totalNum))

		if _, ok := idxMap[int64(randNum)]; ok {
			continue
		} else {
			idxMap[int64(randNum)] = true
		}
	}
	return idxMap
}

//生成total数中的索引返回
func GenerateHitOneIdx(total int) int64 {
	rand := Random(0, float64(total))
	return int64(rand)
}
