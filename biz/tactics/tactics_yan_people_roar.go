package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type YanPeopleRoarTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (y YanPeopleRoarTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (y YanPeopleRoarTactic) Prepare() {
	panic("implement me")
}

func (y YanPeopleRoarTactic) Id() consts.TacticId {
	return consts.YanPeopleRoar
}

func (y YanPeopleRoarTactic) Name() string {
	return "燕人咆哮"
}

func (y YanPeopleRoarTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (y YanPeopleRoarTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (y YanPeopleRoarTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (y YanPeopleRoarTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (y YanPeopleRoarTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (y YanPeopleRoarTactic) Execute() {
	panic("implement me")
}

func (y YanPeopleRoarTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
