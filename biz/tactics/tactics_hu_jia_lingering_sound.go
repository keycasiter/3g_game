package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//胡笳余音
type HuJiaLingeringSoundTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HuJiaLingeringSoundTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (h HuJiaLingeringSoundTactic) Prepare() {
	panic("implement me")
}

func (h HuJiaLingeringSoundTactic) Id() consts.TacticId {
	return consts.HuJiaLingeringSound
}

func (h HuJiaLingeringSoundTactic) Name() string {
	return "胡笳余音"
}

func (h HuJiaLingeringSoundTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (h HuJiaLingeringSoundTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (h HuJiaLingeringSoundTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (h HuJiaLingeringSoundTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (h HuJiaLingeringSoundTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (h HuJiaLingeringSoundTactic) Execute() {
	panic("implement me")
}

func (h HuJiaLingeringSoundTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
