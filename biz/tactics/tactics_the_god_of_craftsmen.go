package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 工神
// 战斗前3回合，使我军全体获得先攻状态，我军主将造成的兵刃伤害和谋略伤害提高30%，我军副将造成伤害提升15%；
// 第4回合开始，我军全体造成伤害降低15%，持续2回合
// 指挥，100%
type TheGodOfCraftsmenTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheGodOfCraftsmenTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TheGodOfCraftsmenTactic) Prepare() {
	panic("implement me")
}

func (t TheGodOfCraftsmenTactic) Id() consts.TacticId {
	return consts.TheGodOfCraftsmen
}

func (t TheGodOfCraftsmenTactic) Name() string {
	return "工神"
}

func (t TheGodOfCraftsmenTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t TheGodOfCraftsmenTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t TheGodOfCraftsmenTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t TheGodOfCraftsmenTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t TheGodOfCraftsmenTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t TheGodOfCraftsmenTactic) Execute() {
	panic("implement me")
}

func (t TheGodOfCraftsmenTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
