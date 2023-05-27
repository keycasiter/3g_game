package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type SittingIntheSoutheastTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SittingIntheSoutheastTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s SittingIntheSoutheastTactic) Prepare() {
	panic("implement me")
}

func (s SittingIntheSoutheastTactic) Id() consts.TacticId {
	return consts.SittingIntheSoutheast
}

func (s SittingIntheSoutheastTactic) Name() string {
	return "坐断东南"
}

func (s SittingIntheSoutheastTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s SittingIntheSoutheastTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s SittingIntheSoutheastTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s SittingIntheSoutheastTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s SittingIntheSoutheastTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s SittingIntheSoutheastTactic) Execute() {
	panic("implement me")
}

func (s SittingIntheSoutheastTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
