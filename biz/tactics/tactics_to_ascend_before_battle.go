package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 临战先登
// 对敌军群体（2人）造成兵刃攻击（伤害率150%），之后自己进入虚弱状态，持续1回合
// 主动，100%
type ToAscendBeforeBattleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ToAscendBeforeBattleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t ToAscendBeforeBattleTactic) Prepare() {
	panic("implement me")
}

func (t ToAscendBeforeBattleTactic) Id() consts.TacticId {
	return consts.ToAscendBeforeBattle
}

func (t ToAscendBeforeBattleTactic) Name() string {
	return "临战先登"
}

func (t ToAscendBeforeBattleTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t ToAscendBeforeBattleTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t ToAscendBeforeBattleTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t ToAscendBeforeBattleTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t ToAscendBeforeBattleTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t ToAscendBeforeBattleTactic) Execute() {
	panic("implement me")
}

func (t ToAscendBeforeBattleTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
