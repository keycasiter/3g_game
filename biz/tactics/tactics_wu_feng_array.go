package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 武锋阵
// 我军三名武将兵种适性均不相同时，战斗前5回合，我军主将有80%几率优先成为敌军战法的目标，且使该战法对其造成的伤害降低30%（受主将统率影响），
// 战斗中，奇数回合使兵种适性较低的副将恢复我军单体兵力（治疗率184%）
// 偶数回合使兵种适性较高的副将造成伤害提高15%（可叠加）
type WuFengArrayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WuFengArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 1.0
	return w
}

func (w WuFengArrayTactic) Prepare() {

}

func (w WuFengArrayTactic) Id() consts.TacticId {
	return consts.WuFengArray
}

func (w WuFengArrayTactic) Name() string {
	return "武锋阵"
}

func (w WuFengArrayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (w WuFengArrayTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WuFengArrayTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WuFengArrayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (w WuFengArrayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (w WuFengArrayTactic) Execute() {

}

func (w WuFengArrayTactic) IsTriggerPrepare() bool {
	return false
}