package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//锦帆百翎
type JinfanArmyHundredFeathersTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JinfanArmyHundredFeathersTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (j JinfanArmyHundredFeathersTactic) Prepare() {
	panic("implement me")
}

func (j JinfanArmyHundredFeathersTactic) Id() consts.TacticId {
	return consts.JinfanArmyHundredFeathers
}

func (j JinfanArmyHundredFeathersTactic) Name() string {
	return "锦帆百翎"
}

func (j JinfanArmyHundredFeathersTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (j JinfanArmyHundredFeathersTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (j JinfanArmyHundredFeathersTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (j JinfanArmyHundredFeathersTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (j JinfanArmyHundredFeathersTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (j JinfanArmyHundredFeathersTactic) Execute() {
	panic("implement me")
}

func (j JinfanArmyHundredFeathersTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
