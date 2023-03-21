package model

import "github.com/keycasiter/3g_game/biz/consts"

type TacticsParams struct {
	//武将唯一ID
	GeneralId string
	//是否为主将
	IsMaster bool
	//当前回合
	CurrentRound consts.BattleRound
	//上一次触发战法的回合
	LastTriggerRound consts.BattleRound
}
