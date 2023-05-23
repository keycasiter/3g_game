package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 风声鹤唳
// 准备1回合，对敌军群体（2人）造成谋略攻击（伤害率105%，受智力影响）并施加沙暴状态，每回合持续造成伤害（伤害率120%，受智力影响），持续1回合
type TheSoundOfTheWindAndTheCryOfTheStorkTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.45
	return t
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) Prepare() {
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) Id() consts.TacticId {
	return consts.TheSoundOfTheWindAndTheCryOfTheStork
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) Name() string {
	return "风声鹤唳"
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) Execute() {
}

func (t TheSoundOfTheWindAndTheCryOfTheStorkTactic) IsTriggerPrepare() bool {
	return t.isTriggerPrepare
}
