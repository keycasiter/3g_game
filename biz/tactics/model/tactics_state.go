package model

import "github.com/keycasiter/3g_game/biz/consts"

// 战法状态
type TacticsState struct {
	//战法触发回合
	tacticsTriggerMap map[string]consts.BattleRound
	//战法生效回合
	tacticsEffectMap map[string]consts.BattleRound
}

func NewTacticsState() *TacticsState {
	return &TacticsState{
		tacticsTriggerMap: make(map[string]consts.BattleRound, 0),
		tacticsEffectMap:  make(map[string]consts.BattleRound, 0),
	}
}

// 设置战法触发回合
func (t TacticsState) SetTriggerRound(uniqueId string, round consts.BattleRound) {
	t.tacticsTriggerMap[uniqueId] = round
}

// 获取最近触发战法回合
func (t TacticsState) GetLastTriggerRound(uniqueId string) consts.BattleRound {
	if round, ok := t.tacticsTriggerMap[uniqueId]; ok {
		return round
	}
	return consts.Battle_Round_Unknow
}
