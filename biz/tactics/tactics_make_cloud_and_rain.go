package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//兴云布雨
type MakeCloudAndRainTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MakeCloudAndRainTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (m MakeCloudAndRainTactic) Prepare() {
	panic("implement me")
}

func (m MakeCloudAndRainTactic) Id() consts.TacticId {
	return consts.MakeCloudAndRain
}

func (m MakeCloudAndRainTactic) Name() string {
	return "兴云布雨"
}

func (m MakeCloudAndRainTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (m MakeCloudAndRainTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (m MakeCloudAndRainTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (m MakeCloudAndRainTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (m MakeCloudAndRainTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (m MakeCloudAndRainTactic) Execute() {
	panic("implement me")
}

func (m MakeCloudAndRainTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
