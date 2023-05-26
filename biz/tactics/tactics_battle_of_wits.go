package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 斗智
type BattleOfWitsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BattleOfWitsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (b BattleOfWitsTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (b BattleOfWitsTactic) Id() consts.TacticId {
	return consts.BattleOfWits
}

func (b BattleOfWitsTactic) Name() string {
	return "斗智"
}

func (b BattleOfWitsTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (b BattleOfWitsTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (b BattleOfWitsTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (b BattleOfWitsTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (b BattleOfWitsTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (b BattleOfWitsTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (b BattleOfWitsTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
