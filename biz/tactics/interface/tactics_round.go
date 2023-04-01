package _interface

import "github.com/keycasiter/3g_game/biz/consts"

//战法回合
type TacticsRound interface {
	//获取当前回合
	GetCurrentRound() consts.BattleRound
	//最近一次触发回合
	LastTriggerRound() consts.BattleRound
}
