package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Id() consts.TacticId {
	return consts.GainMasteryByStrikingOnlyAfterTheEnemyHasStruck
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Name() string {
	return "后发制人"
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (g GainMasteryByStrikingOnlyAfterTheEnemyHasStruckTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
