package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//勇烈持重
type BraveAndResoluteTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BraveAndResoluteTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (b BraveAndResoluteTactic) Prepare() {
	panic("implement me")
}

func (b BraveAndResoluteTactic) Id() consts.TacticId {
	return consts.BraveAndResolute
}

func (b BraveAndResoluteTactic) Name() string {
	return "勇烈持重"
}

func (b BraveAndResoluteTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BraveAndResoluteTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BraveAndResoluteTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BraveAndResoluteTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BraveAndResoluteTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BraveAndResoluteTactic) Execute() {
	panic("implement me")
}

func (b BraveAndResoluteTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
