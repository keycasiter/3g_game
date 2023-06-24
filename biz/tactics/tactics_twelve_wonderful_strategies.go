package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 十二奇策
// 移除敌军群体（1～2人）增益状态，提高我军全体1回合6%主动战法发动率（受智力影响）并使其瑕疵发动主动战法后，对敌军单体造成谋略攻击（伤害率102%，受智力影响）
// 主动，45%
type TwelveWonderfulStrategiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TwelveWonderfulStrategiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TwelveWonderfulStrategiesTactic) Prepare() {
	panic("implement me")
}

func (t TwelveWonderfulStrategiesTactic) Id() consts.TacticId {
	return consts.TwelveWonderfulStrategies
}

func (t TwelveWonderfulStrategiesTactic) Name() string {
	return "十二奇策"
}

func (t TwelveWonderfulStrategiesTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t TwelveWonderfulStrategiesTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t TwelveWonderfulStrategiesTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t TwelveWonderfulStrategiesTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t TwelveWonderfulStrategiesTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t TwelveWonderfulStrategiesTactic) Execute() {
	panic("implement me")
}

func (t TwelveWonderfulStrategiesTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
