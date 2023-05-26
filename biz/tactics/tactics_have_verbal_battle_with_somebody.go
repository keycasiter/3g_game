package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 唇枪舌战
type HaveVerbalBattleWithSomebodyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HaveVerbalBattleWithSomebodyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (h HaveVerbalBattleWithSomebodyTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (h HaveVerbalBattleWithSomebodyTactic) Id() consts.TacticId {
	return consts.HaveVerbalBattleWithSomebody
}

func (h HaveVerbalBattleWithSomebodyTactic) Name() string {
	return "唇枪舌战"
}

func (h HaveVerbalBattleWithSomebodyTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (h HaveVerbalBattleWithSomebodyTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (h HaveVerbalBattleWithSomebodyTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (h HaveVerbalBattleWithSomebodyTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (h HaveVerbalBattleWithSomebodyTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (h HaveVerbalBattleWithSomebodyTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (h HaveVerbalBattleWithSomebodyTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
