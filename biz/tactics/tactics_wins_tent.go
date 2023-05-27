package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//校胜帷幄
type WinsTentTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WinsTentTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (w WinsTentTactic) Prepare() {
	panic("implement me")
}

func (w WinsTentTactic) Id() consts.TacticId {
	return consts.WinsTent
}

func (w WinsTentTactic) Name() string {
	return "校胜帷幄"
}

func (w WinsTentTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (w WinsTentTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (w WinsTentTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (w WinsTentTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (w WinsTentTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (w WinsTentTactic) Execute() {
	panic("implement me")
}

func (w WinsTentTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
