package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//挥兵谋胜
type WieldTroopsToSeekVictoryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WieldTroopsToSeekVictoryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) Prepare() {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) Id() consts.TacticId {
	return consts.WieldTroopsToSeekVictory
}

func (w WieldTroopsToSeekVictoryTactic) Name() string {
	return "挥兵谋胜"
}

func (w WieldTroopsToSeekVictoryTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) Execute() {
	panic("implement me")
}

func (w WieldTroopsToSeekVictoryTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
