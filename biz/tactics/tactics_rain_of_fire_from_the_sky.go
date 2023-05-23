package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 天降火雨
// 准备1回合，对敌军群体（2人）造成一次兵刃攻击（伤害率118%），并附加灼烧状态，每回合持续造成伤害（伤害率66%，受智力影响），持续1回合
type RainOfFireFromTheSkyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RainOfFireFromTheSkyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 0.5
	return r
}

func (r RainOfFireFromTheSkyTactic) Prepare() {
}

func (r RainOfFireFromTheSkyTactic) Id() consts.TacticId {
	return consts.RainOfFireFromTheSky
}

func (r RainOfFireFromTheSkyTactic) Name() string {
	return "天降火雨"
}

func (r RainOfFireFromTheSkyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RainOfFireFromTheSkyTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RainOfFireFromTheSkyTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RainOfFireFromTheSkyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (r RainOfFireFromTheSkyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RainOfFireFromTheSkyTactic) Execute() {
}

func (r RainOfFireFromTheSkyTactic) IsTriggerPrepare() bool {
	return false
}
