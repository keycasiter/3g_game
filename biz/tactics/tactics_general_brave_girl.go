package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//将门虎女
type GeneralBraveGirlTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GeneralBraveGirlTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (g GeneralBraveGirlTactic) Prepare() {
	panic("implement me")
}

func (g GeneralBraveGirlTactic) Id() consts.TacticId {
	return consts.GeneralBraveGirl
}

func (g GeneralBraveGirlTactic) Name() string {
	return "将门虎女"
}

func (g GeneralBraveGirlTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (g GeneralBraveGirlTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (g GeneralBraveGirlTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (g GeneralBraveGirlTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (g GeneralBraveGirlTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (g GeneralBraveGirlTactic) Execute() {
	panic("implement me")
}

func (g GeneralBraveGirlTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
