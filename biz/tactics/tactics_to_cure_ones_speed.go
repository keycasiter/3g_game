package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 将行其疾
// 普通攻击之后，对随机敌军单体发动一次兵刃攻击（伤害率158%）；
// 若命中敌军主将，则使其进入计穷状态，持续2回合
// 突击 60%
type ToCureOnesSpeedTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ToCureOnesSpeedTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t ToCureOnesSpeedTactic) Prepare() {
	panic("implement me")
}

func (t ToCureOnesSpeedTactic) Id() consts.TacticId {
	return consts.ToCureOnesSpeed
}

func (t ToCureOnesSpeedTactic) Name() string {
	return "将行其疾"
}

func (t ToCureOnesSpeedTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t ToCureOnesSpeedTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t ToCureOnesSpeedTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t ToCureOnesSpeedTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t ToCureOnesSpeedTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t ToCureOnesSpeedTactic) Execute() {
	panic("implement me")
}

func (t ToCureOnesSpeedTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
