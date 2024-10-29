package util

import (
	"math/rand"
	"time"

	"github.com/spf13/cast"
)

// 生成随机数，min为下限，max为上限
func Random(min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + cast.ToFloat64(min)
}

// 根据概率设定返回是否成功
func GenerateRate(rate float64) bool {
	rand := Random(0, 1)
	if rand >= 0 && rand <= rate {
		return true
	}
	return false
}

// 生成totalNum数中命中hitNum数量的索引数组
func GenerateHitIdxArr(hitNum, totalNum int) []int64 {
	//hlog.CtxInfof(context.Background(), "hitNum:%v, totalNum:%v", hitNum, totalNum)
	if hitNum > totalNum {
		hitNum = totalNum
	}
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

	if hitNum > totalNum {
		hitNum = totalNum
	}

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

// 生成total数中的索引返回
func GenerateHitOneIdx(total int) int64 {
	rand := Random(0, float64(total))
	return int64(rand)
}

// 生成total数中的索引返回
func GenerateHitTwoOrThreeIdxArr() []int64 {
	if GenerateRate(0.5) {
		return GenerateHitIdxArr(2, 3)
	}
	return GenerateHitIdxArr(3, 3)
}

// 生成total数中的索引返回
func GenerateHitTwoOrThreeIdxMap(totalNum int) map[int64]bool {
	arr := make([]int64, 0)
	m := make(map[int64]bool, 0)

	if totalNum > 3 && totalNum < 0 {
		panic(any("totalNum is err"))
	}

	if totalNum == 1 { //1人
		arr = GenerateHitIdxArr(1, 1)
	} else if totalNum == 2 { //2人
		arr = GenerateHitIdxArr(2, 2)
	} else { //3人
		if GenerateRate(0.5) {
			arr = GenerateHitIdxArr(2, 3)
		} else {
			arr = GenerateHitIdxArr(3, 3)
		}
	}
	for _, idx := range arr {
		m[idx] = true
	}
	return m
}
