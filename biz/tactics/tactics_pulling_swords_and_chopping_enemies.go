package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type PullingSwordsAndChoppingEnemiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PullingSwordsAndChoppingEnemiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) Id() consts.TacticId {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) Name() string {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (p PullingSwordsAndChoppingEnemiesTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
