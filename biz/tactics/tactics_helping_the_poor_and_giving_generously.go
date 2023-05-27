package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//济贫好施
type HelpingThePoorAndGivingGenerouslyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Prepare() {
	panic("implement me")
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Id() consts.TacticId {
	return consts.HelpingThePoorAndGivingGenerously
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Name() string {
	return "济贫好施"
}

func (h HelpingThePoorAndGivingGenerouslyTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (h HelpingThePoorAndGivingGenerouslyTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (h HelpingThePoorAndGivingGenerouslyTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (h HelpingThePoorAndGivingGenerouslyTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (h HelpingThePoorAndGivingGenerouslyTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (h HelpingThePoorAndGivingGenerouslyTactic) Execute() {
	panic("implement me")
}

func (h HelpingThePoorAndGivingGenerouslyTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
