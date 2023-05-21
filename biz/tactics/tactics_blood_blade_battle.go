package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//血刃争锋
//战斗中，提高75%普通攻击伤害，普通攻击之后对目标造成酣斗效果，每回合最多触发1次，
//若目标身上存在3次酣斗效果，消耗全部酣斗效果额外提高自身（11%x酣斗次数）普通攻击伤害，额外提升效果不可叠加，持续到战斗结束
type BloodBladeBattle struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BloodBladeBattle) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BloodBladeBattle) Prepare() {
}

func (b BloodBladeBattle) Id() consts.TacticId {
	return consts.BloodBladeBattle
}

func (b BloodBladeBattle) Name() string {
	return "血刃争锋"
}

func (b BloodBladeBattle) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (b BloodBladeBattle) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BloodBladeBattle) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BloodBladeBattle) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BloodBladeBattle) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BloodBladeBattle) Execute() {

}

func (b BloodBladeBattle) IsTriggerPrepare() bool {
	return false
}
