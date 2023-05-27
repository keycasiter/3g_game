package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//傲睨王侯
type ProudPrinceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p ProudPrinceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (p ProudPrinceTactic) Prepare() {
	panic("implement me")
}

func (p ProudPrinceTactic) Id() consts.TacticId {
	return consts.ProudPrince
}

func (p ProudPrinceTactic) Name() string {
	return "傲睨王侯"
}

func (p ProudPrinceTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (p ProudPrinceTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (p ProudPrinceTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (p ProudPrinceTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (p ProudPrinceTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (p ProudPrinceTactic) Execute() {
	panic("implement me")
}

func (p ProudPrinceTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
