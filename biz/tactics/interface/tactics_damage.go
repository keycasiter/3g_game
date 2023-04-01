package _interface

import "github.com/keycasiter/3g_game/biz/consts"

//战法伤害
type TacticsDamage interface {
	//伤害类型
	DamageType() consts.DamageType
	//伤害率
	DamageRate() float64
	//伤害值
	DamageNum() float64
	//伤害范围
	DamageRange() consts.GeneralNum
}
