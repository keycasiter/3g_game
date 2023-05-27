package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type TenWinsAndTenLossesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TenWinsAndTenLossesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TenWinsAndTenLossesTactic) Prepare() {
	panic("implement me")
}

func (t TenWinsAndTenLossesTactic) Id() consts.TacticId {
	return consts.TenWinsAndTenLosses
}

func (t TenWinsAndTenLossesTactic) Name() string {
	return "十胜十败"
}

func (t TenWinsAndTenLossesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t TenWinsAndTenLossesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t TenWinsAndTenLossesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t TenWinsAndTenLossesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t TenWinsAndTenLossesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t TenWinsAndTenLossesTactic) Execute() {
	panic("implement me")
}

func (t TenWinsAndTenLossesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
