package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//古之恶来
type AncientEvilComesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AncientEvilComesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (a AncientEvilComesTactic) Prepare() {
	panic("implement me")
}

func (a AncientEvilComesTactic) Id() consts.TacticId {
	return consts.AncientEvilComes
}

func (a AncientEvilComesTactic) Name() string {
	return "古之恶来"
}

func (a AncientEvilComesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (a AncientEvilComesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (a AncientEvilComesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (a AncientEvilComesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (a AncientEvilComesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (a AncientEvilComesTactic) Execute() {
	panic("implement me")
}

func (a AncientEvilComesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
