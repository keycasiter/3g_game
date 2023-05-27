package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//火神英风
type FireGodHeroStyleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FireGodHeroStyleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (f FireGodHeroStyleTactic) Prepare() {
	panic("implement me")
}

func (f FireGodHeroStyleTactic) Id() consts.TacticId {
	return consts.FireGodHeroStyle
}

func (f FireGodHeroStyleTactic) Name() string {
	return "火神英风"
}

func (f FireGodHeroStyleTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (f FireGodHeroStyleTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (f FireGodHeroStyleTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (f FireGodHeroStyleTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (f FireGodHeroStyleTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (f FireGodHeroStyleTactic) Execute() {
	panic("implement me")
}

func (f FireGodHeroStyleTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
