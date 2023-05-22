package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 千里走单骑
// 战斗中，自身准备发动自带准备战法时，有70%几率（受武力影响）获得洞察状态（免疫所有控制效果）并提高50武力，持续2回合，
// 在此期间，自身受到普通攻击时，对攻击者进行一次反击（伤害率238%），每回合最多触发1次
type ThousandMileWalkingSingleRiderTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThousandMileWalkingSingleRiderTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t ThousandMileWalkingSingleRiderTactic) Prepare() {
}

func (t ThousandMileWalkingSingleRiderTactic) Id() consts.TacticId {
	return consts.ThousandMileWalkingSingleRider
}

func (t ThousandMileWalkingSingleRiderTactic) Name() string {
	return "千里走单骑"
}

func (t ThousandMileWalkingSingleRiderTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t ThousandMileWalkingSingleRiderTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThousandMileWalkingSingleRiderTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThousandMileWalkingSingleRiderTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t ThousandMileWalkingSingleRiderTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThousandMileWalkingSingleRiderTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (t ThousandMileWalkingSingleRiderTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
