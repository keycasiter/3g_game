package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 国士将风
type CountryPersonGeneralStyleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CountryPersonGeneralStyleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (c CountryPersonGeneralStyleTactic) Prepare() {
	panic("implement me")
}

func (c CountryPersonGeneralStyleTactic) Id() consts.TacticId {
	return consts.CountryPersonGeneralStyle
}

func (c CountryPersonGeneralStyleTactic) Name() string {
	return "国士将风"
}

func (c CountryPersonGeneralStyleTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (c CountryPersonGeneralStyleTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (c CountryPersonGeneralStyleTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (c CountryPersonGeneralStyleTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (c CountryPersonGeneralStyleTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (c CountryPersonGeneralStyleTactic) Execute() {
	panic("implement me")
}

func (c CountryPersonGeneralStyleTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
