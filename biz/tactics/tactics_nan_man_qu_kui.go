package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type NanManQuKuiTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (n NanManQuKuiTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (n NanManQuKuiTactic) Prepare() {
	panic("implement me")
}

func (n NanManQuKuiTactic) Id() consts.TacticId {
	return consts.NanManQuKui
}

func (n NanManQuKuiTactic) Name() string {
	return "南蛮渠魁"
}

func (n NanManQuKuiTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (n NanManQuKuiTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (n NanManQuKuiTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (n NanManQuKuiTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (n NanManQuKuiTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (n NanManQuKuiTactic) Execute() {
	panic("implement me")
}

func (n NanManQuKuiTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
