package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//沉沙决水
//准备1回合，对敌军群体(2人)施加水攻状态，每回合持续造成伤害(伤害率126%,受智力影响)，并使其受到的谋略伤害提升25%
//持续2回合
type SinkingSandAndBreakingWaterTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SinkingSandAndBreakingWaterTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.4
	return s
}

func (s SinkingSandAndBreakingWaterTactic) Prepare() {
}

func (s SinkingSandAndBreakingWaterTactic) Id() consts.TacticId {
	return consts.SinkingSandAndBreakingWater
}

func (s SinkingSandAndBreakingWaterTactic) Name() string {
	return "沉沙决水"
}

func (s SinkingSandAndBreakingWaterTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SinkingSandAndBreakingWaterTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SinkingSandAndBreakingWaterTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SinkingSandAndBreakingWaterTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SinkingSandAndBreakingWaterTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SinkingSandAndBreakingWaterTactic) Execute() {
	panic("implement me")
}

func (s SinkingSandAndBreakingWaterTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
