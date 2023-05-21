package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//大戟士
//将枪兵进阶为横冲直撞的大戟士：
//我军全体武力提升14点，进行普通攻击时，有35%几率对敌军单体造成兵刃伤害（伤害率122%）
//若张合统领，则发动几率提高为40%
type GreatHalberdWarriorTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GreatHalberdWarriorTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 1.0
	return g
}

func (g GreatHalberdWarriorTactic) Prepare() {

}

func (g GreatHalberdWarriorTactic) Id() consts.TacticId {
	return consts.GreatHalberdWarrior
}

func (g GreatHalberdWarriorTactic) Name() string {
	return "大戟士"
}

func (g GreatHalberdWarriorTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (g GreatHalberdWarriorTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GreatHalberdWarriorTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GreatHalberdWarriorTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (g GreatHalberdWarriorTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Spearman,
	}
}

func (g GreatHalberdWarriorTactic) Execute() {
}

func (g GreatHalberdWarriorTactic) IsTriggerPrepare() bool {
	return false
}
