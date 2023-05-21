package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//勇者得前
//普通攻击之后使自己获得一次抵御，可免疫伤害，并使下一个主动战法的伤害率提升80%
type TheBraveLeadTheWayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheBraveLeadTheWayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.45
	return t
}

func (t TheBraveLeadTheWayTactic) Prepare() {
}

func (t TheBraveLeadTheWayTactic) Id() consts.TacticId {
	return consts.TheBraveLeadTheWay
}

func (t TheBraveLeadTheWayTactic) Name() string {
	return "勇者得前"
}

func (t TheBraveLeadTheWayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheBraveLeadTheWayTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheBraveLeadTheWayTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheBraveLeadTheWayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (t TheBraveLeadTheWayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheBraveLeadTheWayTactic) Execute() {
}

func (t TheBraveLeadTheWayTactic) IsTriggerPrepare() bool {
	return false
}
