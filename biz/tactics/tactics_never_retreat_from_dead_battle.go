package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//死战不退
type NeverRetreatFromDeadBattleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (n NeverRetreatFromDeadBattleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (n NeverRetreatFromDeadBattleTactic) Prepare() {
	panic("implement me")
}

func (n NeverRetreatFromDeadBattleTactic) Id() consts.TacticId {
	return consts.NeverRetreatFromDeadBattle
}

func (n NeverRetreatFromDeadBattleTactic) Name() string {
	return "死战不退"
}

func (n NeverRetreatFromDeadBattleTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (n NeverRetreatFromDeadBattleTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (n NeverRetreatFromDeadBattleTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (n NeverRetreatFromDeadBattleTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (n NeverRetreatFromDeadBattleTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (n NeverRetreatFromDeadBattleTactic) Execute() {
	panic("implement me")
}

func (n NeverRetreatFromDeadBattleTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
