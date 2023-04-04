package util

import "github.com/keycasiter/3g_game/biz/consts"

// 增益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
func BuffEffectWrapSet(holder map[consts.BuffEffectType]map[consts.BattleRound]float64,
	effectType consts.BuffEffectType,
	round consts.BattleRound,
	v float64) {
	mm := make(map[consts.BattleRound]float64, 0)
	mm[round] = v
	holder[effectType] = mm
}

// 减益效果容器处理
// @holder 效果容器
// @effectType 效果类型
// @v 效果值
func DebuffEffectWrapSet(holder map[consts.DebuffEffectType]map[consts.BattleRound]float64,
	effectType consts.DebuffEffectType,
	round consts.BattleRound,
	v float64) {
	mm := make(map[consts.BattleRound]float64, 0)
	mm[round] = v
	holder[effectType] = mm
}
