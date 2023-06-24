package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 竭忠尽智
// 准备1回合，使敌军群体（1～2人）速度降低15%（受智力影响）
// 并进入混乱状态，持续2回合
// 并使友军单体获得1次抵御，持续1回合
// 主动，50%
type WithAllTheHeartTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WithAllTheHeartTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (w WithAllTheHeartTactic) Prepare() {
	panic("implement me")
}

func (w WithAllTheHeartTactic) Id() consts.TacticId {
	return consts.WithAllTheHeart
}

func (w WithAllTheHeartTactic) Name() string {
	return "竭忠尽智"
}

func (w WithAllTheHeartTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (w WithAllTheHeartTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (w WithAllTheHeartTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (w WithAllTheHeartTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (w WithAllTheHeartTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (w WithAllTheHeartTactic) Execute() {
	panic("implement me")
}

func (w WithAllTheHeartTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
