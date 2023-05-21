package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//白毦兵
//将枪兵进阶为攻无不破的白毦兵：
//我军全体战斗重普通攻击后有45%概率对攻击目标再次发起一次谋略攻击（伤害率110%，受智力影响）
//若陈到统领，则谋略攻击更为强力（伤害率130%，受智力影响）
type WhiteArmyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WhiteArmyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 1.0
	return w
}

func (w WhiteArmyTactic) Prepare() {

}

func (w WhiteArmyTactic) Id() consts.TacticId {
	return consts.WhiteArmy
}

func (w WhiteArmyTactic) Name() string {
	return "白毦兵"
}

func (w WhiteArmyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (w WhiteArmyTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WhiteArmyTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WhiteArmyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (w WhiteArmyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Spearman,
	}
}

func (w WhiteArmyTactic) Execute() {
}

func (w WhiteArmyTactic) IsTriggerPrepare() bool {
	return false
}
