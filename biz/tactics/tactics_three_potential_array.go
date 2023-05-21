package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//三势阵
//我军三名武将阵营均不相同，且我方主将自带战法为主动战法或突击战法时，战斗前5回合，主将提高16%自带主动、
//突击战法发动几率，每回合行动前，使损失兵力较多的副将受到伤害降低30%，另一面副将造成伤害提高25%，持续1回合
type ThreePotentialArray struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThreePotentialArray) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t ThreePotentialArray) Prepare() {
}

func (t ThreePotentialArray) Id() consts.TacticId {
	return consts.ThreePotentialArray
}

func (t ThreePotentialArray) Name() string {
	return "三势阵"
}

func (t ThreePotentialArray) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t ThreePotentialArray) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThreePotentialArray) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThreePotentialArray) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (t ThreePotentialArray) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThreePotentialArray) Execute() {
}

func (t ThreePotentialArray) IsTriggerPrepare() bool {
	return false
}
