package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 天下无双
// 对敌军单体发起决斗，决斗双方轮流向对方普通攻击3次，自己率先出售。
// 决斗途中，双手不受缴械和震慑状态影响，并且可以触发群攻和突击战法，
// 自身为主将时，决斗后自身受到兵刃伤害降低7%（受武力影响），持续2回合
// 主动35%
type UniqueTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (u UniqueTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (u UniqueTactic) Prepare() {
	panic("implement me")
}

func (u UniqueTactic) Id() consts.TacticId {
	return consts.Unique
}

func (u UniqueTactic) Name() string {
	return "天下无双"
}

func (u UniqueTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (u UniqueTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (u UniqueTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (u UniqueTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (u UniqueTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (u UniqueTactic) Execute() {
	panic("implement me")
}

func (u UniqueTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
