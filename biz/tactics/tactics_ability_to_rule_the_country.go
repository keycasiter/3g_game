package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//经天纬地
type AbilityToRuleTheCountryTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AbilityToRuleTheCountryTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (a AbilityToRuleTheCountryTactic) Prepare() {
	panic("implement me")
}

func (a AbilityToRuleTheCountryTactic) Id() consts.TacticId {
	return consts.AbilityToRuleTheCountry
}

func (a AbilityToRuleTheCountryTactic) Name() string {
	return "经天纬地"
}

func (a AbilityToRuleTheCountryTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (a AbilityToRuleTheCountryTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (a AbilityToRuleTheCountryTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (a AbilityToRuleTheCountryTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (a AbilityToRuleTheCountryTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (a AbilityToRuleTheCountryTactic) Execute() {
	panic("implement me")
}

func (a AbilityToRuleTheCountryTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
