package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//搦战群雄
type ToSeizeThePowerOfGroupOfHeroesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Prepare() {
	panic("implement me")
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Id() consts.TacticId {
	return consts.ToSeizeThePowerOfGroupOfHeroes
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Name() string {
	return "搦战群雄"
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) Execute() {
	panic("implement me")
}

func (t ToSeizeThePowerOfGroupOfHeroesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
