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
