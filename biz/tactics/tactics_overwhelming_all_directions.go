package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 暴敛四方
type OverwhelmingAllDirectionsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (o OverwhelmingAllDirectionsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (o OverwhelmingAllDirectionsTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (o OverwhelmingAllDirectionsTactic) Id() consts.TacticId {
	return consts.OverwhelmingAllDirections
}

func (o OverwhelmingAllDirectionsTactic) Name() string {
	return "暴敛四方"
}

func (o OverwhelmingAllDirectionsTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (o OverwhelmingAllDirectionsTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (o OverwhelmingAllDirectionsTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (o OverwhelmingAllDirectionsTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (o OverwhelmingAllDirectionsTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (o OverwhelmingAllDirectionsTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (o OverwhelmingAllDirectionsTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
