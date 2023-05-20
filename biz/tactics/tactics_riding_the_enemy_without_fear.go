package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//乘敌不虞
//准备1回合，使敌军主将进入虚弱（无法造成伤害）状态，持续2回合，
//并使我军主将进入休整状态（每回合恢复一次兵力，回复率105%，受智力影响），持续2回合
type RidingTheEnemyWithoutFearTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	//是否已经触发准备战法
	isTriggerPrepare bool
}

func (r RidingTheEnemyWithoutFearTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 0.35
	return r
}

func (r RidingTheEnemyWithoutFearTactic) Prepare() {

}

func (r RidingTheEnemyWithoutFearTactic) Id() consts.TacticId {
	return consts.RidingTheEnemyWithoutFear
}

func (r RidingTheEnemyWithoutFearTactic) Name() string {
	return "乘敌不虞"
}

func (r RidingTheEnemyWithoutFearTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RidingTheEnemyWithoutFearTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RidingTheEnemyWithoutFearTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RidingTheEnemyWithoutFearTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (r RidingTheEnemyWithoutFearTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RidingTheEnemyWithoutFearTactic) Execute() {

}

func (r RidingTheEnemyWithoutFearTactic) IsTriggerPrepare() bool {
	return r.isTriggerPrepare
}
