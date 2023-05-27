package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type RiverFireFlameTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RiverFireFlameTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (r RiverFireFlameTactic) Prepare() {
	panic("implement me")
}

func (r RiverFireFlameTactic) Id() consts.TacticId {
	return consts.RiverFireFlame
}

func (r RiverFireFlameTactic) Name() string {
	return "江天长焰"
}

func (r RiverFireFlameTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (r RiverFireFlameTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (r RiverFireFlameTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (r RiverFireFlameTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (r RiverFireFlameTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (r RiverFireFlameTactic) Execute() {
	panic("implement me")
}

func (r RiverFireFlameTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
