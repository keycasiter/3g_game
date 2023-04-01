package _interface

import "github.com/keycasiter/3g_game/biz/consts"

//战法增益效果
type TacticsBuff interface {
	//增益效果
	BuffEffect() consts.BuffEffectType
	//增益率
	BuffEffectRate() float64
	//造成伤害增加数值
	IncrLaunchDamageNum() int64
	//造成伤害增加比率
	IncrLaunchDamageRate() float64
	//受到伤害减少数值
	DecrSufferDamageNum() int64
	//受到伤害减少比率
	DecrSufferDamageRate() float64
	//增加属性
	IncrAttr() map[consts.AbilityAttr]float64
}
