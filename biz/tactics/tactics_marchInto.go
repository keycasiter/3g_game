package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//长驱直入
type MarchIntoTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MarchIntoTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (m MarchIntoTactic) Prepare() {
	panic("implement me")
}

func (m MarchIntoTactic) Id() consts.TacticId {
	return consts.MarchInto
}

func (m MarchIntoTactic) Name() string {
	return "长驱直入"
}

func (m MarchIntoTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (m MarchIntoTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (m MarchIntoTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (m MarchIntoTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (m MarchIntoTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (m MarchIntoTactic) Execute() {
	panic("implement me")
}

func (m MarchIntoTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
