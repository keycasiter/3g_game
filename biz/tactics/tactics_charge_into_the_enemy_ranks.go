package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//登锋陷阵
type ChargeIntoTheEnemyRanksTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c ChargeIntoTheEnemyRanksTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (c ChargeIntoTheEnemyRanksTactic) Prepare() {
	panic("implement me")
}

func (c ChargeIntoTheEnemyRanksTactic) Id() consts.TacticId {
	return consts.ChargeIntoTheEnemyRanks
}

func (c ChargeIntoTheEnemyRanksTactic) Name() string {
	return "登锋陷阵"
}

func (c ChargeIntoTheEnemyRanksTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (c ChargeIntoTheEnemyRanksTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (c ChargeIntoTheEnemyRanksTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (c ChargeIntoTheEnemyRanksTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (c ChargeIntoTheEnemyRanksTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (c ChargeIntoTheEnemyRanksTactic) Execute() {
	panic("implement me")
}

func (c ChargeIntoTheEnemyRanksTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
