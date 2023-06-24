package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 刀出如霆
// 准备1回合。自身及友军单体获得30%倒戈，持续2回合，并对敌军造成兵刃伤害（伤害率300%，由敌军全部武将评分，敌军每有1名副将总伤害率提高120%）及掠阵状态：
// 掠阵状态叠加两次时，移除掠阵状态并使目标受到兵刃伤害提高30%，可叠加；
// 若与张苞同时出战，则友军单体必定选择张苞
// 主动，40%
type TheKnifeLikeThunderboltTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheKnifeLikeThunderboltTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) Prepare() {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) Id() consts.TacticId {
	return consts.TheKnifeLikeThunderbolt
}

func (t TheKnifeLikeThunderboltTactic) Name() string {
	return "刀出如霆"
}

func (t TheKnifeLikeThunderboltTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) Execute() {
	panic("implement me")
}

func (t TheKnifeLikeThunderboltTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
