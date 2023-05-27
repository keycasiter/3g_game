package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type TheLionFliesFastTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheLionFliesFastTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TheLionFliesFastTactic) Prepare() {
	panic("implement me")
}

func (t TheLionFliesFastTactic) Id() consts.TacticId {
	return consts.TheLionFliesFast
}

func (t TheLionFliesFastTactic) Name() string {
	return "狮子奋迅"
}

func (t TheLionFliesFastTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t TheLionFliesFastTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t TheLionFliesFastTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t TheLionFliesFastTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t TheLionFliesFastTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t TheLionFliesFastTactic) Execute() {
	panic("implement me")
}

func (t TheLionFliesFastTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
