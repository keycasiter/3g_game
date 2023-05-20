package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//诱敌深入
//准备1回合，对敌军群体（2人）施加沙暴状态，每回合持续造成伤害（伤害率126%，受智力影响），并使其受到的兵刃伤害提升25%，持续2回合
type LureTheEnemyInDeepTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	//是否已经触发准备战法
	isTriggerPrepare bool
}

func (l LureTheEnemyInDeepTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 0.4
	return l
}

func (l LureTheEnemyInDeepTactic) Prepare() {
}

func (l LureTheEnemyInDeepTactic) Id() consts.TacticId {
	return consts.LureTheEnemyInDeep
}

func (l LureTheEnemyInDeepTactic) Name() string {
	return "诱敌深入"
}

func (l LureTheEnemyInDeepTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (l LureTheEnemyInDeepTactic) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LureTheEnemyInDeepTactic) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LureTheEnemyInDeepTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (l LureTheEnemyInDeepTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LureTheEnemyInDeepTactic) Execute() {
}

func (l LureTheEnemyInDeepTactic) IsTriggerPrepare() bool {
	return l.isTriggerPrepare
}
