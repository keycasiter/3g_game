package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//固若金汤
type ImpregnableTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i ImpregnableTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (i ImpregnableTactic) Prepare() {
	panic("implement me")
}

func (i ImpregnableTactic) Id() consts.TacticId {
	return consts.Impregnable
}

func (i ImpregnableTactic) Name() string {
	return "固若金汤"
}

func (i ImpregnableTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (i ImpregnableTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (i ImpregnableTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (i ImpregnableTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (i ImpregnableTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (i ImpregnableTactic) Execute() {
	panic("implement me")
}

func (i ImpregnableTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
