package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 搦战群雄
// 准备1回合，对敌军群体（2人）造成一次兵刃攻击（伤害率200%），随后使自己造成兵刃伤害提高25%，受到兵刃伤害降低25%，（受武力影响），持续2回合
// 主动，35%
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
