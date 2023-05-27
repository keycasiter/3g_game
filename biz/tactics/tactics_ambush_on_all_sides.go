package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//十面埋伏
type AmbushOnAllSidesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AmbushOnAllSidesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (a AmbushOnAllSidesTactic) Prepare() {
	panic("implement me")
}

func (a AmbushOnAllSidesTactic) Id() consts.TacticId {
	return consts.AmbushOnAllSides
}

func (a AmbushOnAllSidesTactic) Name() string {
	return "十面埋伏"
}

func (a AmbushOnAllSidesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (a AmbushOnAllSidesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (a AmbushOnAllSidesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (a AmbushOnAllSidesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (a AmbushOnAllSidesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (a AmbushOnAllSidesTactic) Execute() {
	panic("implement me")
}

func (a AmbushOnAllSidesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
