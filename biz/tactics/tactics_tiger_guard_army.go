package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//虎卫军
//将盾兵进阶为善固疆场的虎卫军：
//战斗中，我军主将即将受到普通攻击时，副将提高12武力，最多提高5次
//并会分别对攻击者造成兵刃伤害（伤害率72%，受鸽子损失兵力影响，最多提高40%），每回合最多触发1次
//若典韦或许褚统领，自身统率提高50
type TigerGuardArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TigerGuardArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TigerGuardArmyTactic) Prepare() {
	panic("implement me")
}

func (t TigerGuardArmyTactic) Id() consts.TacticId {
	return consts.TigerGuardArmy
}

func (t TigerGuardArmyTactic) Name() string {
	return "虎卫军"
}

func (t TigerGuardArmyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t TigerGuardArmyTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TigerGuardArmyTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TigerGuardArmyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (t TigerGuardArmyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TigerGuardArmyTactic) Execute() {
}

func (t TigerGuardArmyTactic) IsTriggerPrepare() bool {
	return false
}
