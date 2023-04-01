package _interface

import "github.com/keycasiter/3g_game/biz/consts"

//战法负面效果
type TacticsDebuff interface {
	//负面效果
	DebuffEffect() consts.DebuffEffectType
	//负面率
	DebuffEffectRate() float64
	//受到伤害增加数值
	IncrSufferDamageNum() int64
	//受到伤害增加比率
	IncrSufferDamageRate() float64
	//造成伤害减少数值
	DecrLaunchDamageNum() int64
	//造成伤害减少比率
	DecrLaunchDamageRate() float64
	//降低属性
	DecrAttr() map[consts.AbilityAttr]float64
}
