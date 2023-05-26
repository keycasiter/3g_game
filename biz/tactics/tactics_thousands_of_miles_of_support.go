package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 千里驰援
type ThousandsOfMilesOfSupportTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThousandsOfMilesOfSupportTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (t ThousandsOfMilesOfSupportTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (t ThousandsOfMilesOfSupportTactic) Id() consts.TacticId {
	return consts.ThousandsOfMilesOfSupport
}

func (t ThousandsOfMilesOfSupportTactic) Name() string {
	return "千里驰援"
}

func (t ThousandsOfMilesOfSupportTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (t ThousandsOfMilesOfSupportTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (t ThousandsOfMilesOfSupportTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (t ThousandsOfMilesOfSupportTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (t ThousandsOfMilesOfSupportTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (t ThousandsOfMilesOfSupportTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (t ThousandsOfMilesOfSupportTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
