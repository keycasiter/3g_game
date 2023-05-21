package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//智计
//使敌军群体（2人）的武力、智力降低38（受智力影响），持续2回合，最多叠加2次
type IntelligentStrategyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i IntelligentStrategyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.55
	return i
}

func (i IntelligentStrategyTactic) Prepare() {
}

func (i IntelligentStrategyTactic) Id() consts.TacticId {
	return consts.IntelligentStrategy
}

func (i IntelligentStrategyTactic) Name() string {
	return "智计"
}

func (i IntelligentStrategyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (i IntelligentStrategyTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i IntelligentStrategyTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i IntelligentStrategyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i IntelligentStrategyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i IntelligentStrategyTactic) Execute() {
}

func (i IntelligentStrategyTactic) IsTriggerPrepare() bool {
	return false
}
