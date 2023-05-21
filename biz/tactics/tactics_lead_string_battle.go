package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//引弦力战
//普通攻击之后，有45%概率获得群攻（普通攻击时对目标同部队其他武将造成伤害）状态（伤害率52%），
//若已处于群攻状态，则提高16武力，持续3回合，最多可叠加6次
type LeadStringBattle struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LeadStringBattle) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 1.0
	return l
}

func (l LeadStringBattle) Prepare() {
}

func (l LeadStringBattle) Id() consts.TacticId {
	return consts.LeadStringBattle
}

func (l LeadStringBattle) Name() string {
	return "引弦力战"
}

func (l LeadStringBattle) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (l LeadStringBattle) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LeadStringBattle) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LeadStringBattle) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (l LeadStringBattle) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (l LeadStringBattle) Execute() {

}

func (l LeadStringBattle) IsTriggerPrepare() bool {
	return false
}
