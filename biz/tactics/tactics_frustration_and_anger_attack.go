package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 挫志怒袭
// 准备1回合，对敌军群体（2人）施加虚弱（无法造成伤害）状态，持续1回合；
// 如果目标已处于虚弱状态则改为造成一次猛击（伤害率188%）
type FrustrationAndAngerAttackTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (f FrustrationAndAngerAttackTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.35
	return f
}

func (f FrustrationAndAngerAttackTactic) Prepare() {
}

func (f FrustrationAndAngerAttackTactic) Id() consts.TacticId {
	return consts.FrustrationAndAngerAttack
}

func (f FrustrationAndAngerAttackTactic) Name() string {
	return "挫志怒袭"
}

func (f FrustrationAndAngerAttackTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FrustrationAndAngerAttackTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FrustrationAndAngerAttackTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FrustrationAndAngerAttackTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FrustrationAndAngerAttackTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FrustrationAndAngerAttackTactic) Execute() {
}

func (f FrustrationAndAngerAttackTactic) IsTriggerPrepare() bool {
	return f.isTriggerPrepare
}
