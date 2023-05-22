package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 形一阵
// 我军三名武将自带战法类型相同时，战斗中，自身最高属性提高60点，友军群体（2人）造成及受到伤害降低30%
// 此效果每回合降低10%，该效果结束后，每回合使其造成伤害提高16%，受到伤害提高4%，可叠加
type ShapelyArrayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s ShapelyArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s ShapelyArrayTactic) Prepare() {

}

func (s ShapelyArrayTactic) Id() consts.TacticId {
	return consts.ShapelyArray
}

func (s ShapelyArrayTactic) Name() string {
	return "形一阵"
}

func (s ShapelyArrayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (s ShapelyArrayTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s ShapelyArrayTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s ShapelyArrayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (s ShapelyArrayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s ShapelyArrayTactic) Execute() {

}

func (s ShapelyArrayTactic) IsTriggerPrepare() bool {
	return false
}
