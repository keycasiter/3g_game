package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 校胜帷幄
// 战斗中，提高己方主将14%奇谋几率，及20%奇谋伤害，同时为己方主将分担30%的伤害（自身为主将时无效）
// 指挥，100%
type WinsTentTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WinsTentTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (w WinsTentTactic) Prepare() {
	panic("implement me")
}

func (w WinsTentTactic) Id() consts.TacticId {
	return consts.WinsTent
}

func (w WinsTentTactic) Name() string {
	return "校胜帷幄"
}

func (w WinsTentTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (w WinsTentTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (w WinsTentTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (w WinsTentTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (w WinsTentTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (w WinsTentTactic) Execute() {
	panic("implement me")
}

func (w WinsTentTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
