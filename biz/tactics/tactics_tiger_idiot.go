package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//虎痴
type TigerIdiotTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TigerIdiotTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TigerIdiotTactic) Prepare() {
	panic("implement me")
}

func (t TigerIdiotTactic) Id() consts.TacticId {
	return consts.TigerIdiot
}

func (t TigerIdiotTactic) Name() string {
	return "虎痴"
}

func (t TigerIdiotTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t TigerIdiotTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t TigerIdiotTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t TigerIdiotTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t TigerIdiotTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t TigerIdiotTactic) Execute() {
	panic("implement me")
}

func (t TigerIdiotTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
