package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 暗潮涌动
// 准备1回合，对敌军主将造成一次兵刃攻击（伤害率272%）
// 主动，35%
type UndercurrentSurgeTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (u UndercurrentSurgeTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (u UndercurrentSurgeTactic) Prepare() {
	panic("implement me")
}

func (u UndercurrentSurgeTactic) Id() consts.TacticId {
	return consts.UndercurrentSurge
}

func (u UndercurrentSurgeTactic) Name() string {
	return "暗潮涌动"
}

func (u UndercurrentSurgeTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (u UndercurrentSurgeTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (u UndercurrentSurgeTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (u UndercurrentSurgeTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (u UndercurrentSurgeTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (u UndercurrentSurgeTactic) Execute() {
	panic("implement me")
}

func (u UndercurrentSurgeTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
