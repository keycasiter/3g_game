package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 声东击西
// 准备1回合，对敌军群体（2人）造成谋略攻击（伤害率175%，受智力影响），并降低30点速度，持续2回合
type MakeFeintToTheEastButAttackInTheWest struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (m MakeFeintToTheEastButAttackInTheWest) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	m.tacticsParams = tacticsParams
	m.triggerRate = 0.4
	return m
}

func (m MakeFeintToTheEastButAttackInTheWest) Prepare() {
}

func (m MakeFeintToTheEastButAttackInTheWest) Id() consts.TacticId {
	return consts.MakeFeintToTheEastButAttackInTheWest
}

func (m MakeFeintToTheEastButAttackInTheWest) Name() string {
	return "声东击西"
}

func (m MakeFeintToTheEastButAttackInTheWest) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (m MakeFeintToTheEastButAttackInTheWest) GetTriggerRate() float64 {
	return m.triggerRate
}

func (m MakeFeintToTheEastButAttackInTheWest) SetTriggerRate(rate float64) {
	m.triggerRate = rate
}

func (m MakeFeintToTheEastButAttackInTheWest) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (m MakeFeintToTheEastButAttackInTheWest) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (m MakeFeintToTheEastButAttackInTheWest) Execute() {
}

func (m MakeFeintToTheEastButAttackInTheWest) IsTriggerPrepare() bool {
	return m.isTriggerPrepare
}
