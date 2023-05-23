package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 敛众而击
// 对敌军群体(1-2人)造成兵刃伤害（伤害率164%），并有45%概率治疗自身（治疗率88%，受武力影响）
type GatherTheCrowdAndStrike struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GatherTheCrowdAndStrike) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 0.35
	return g
}

func (g GatherTheCrowdAndStrike) Prepare() {

}

func (g GatherTheCrowdAndStrike) Id() consts.TacticId {
	return consts.GatherTheCrowdAndStrike
}

func (g GatherTheCrowdAndStrike) Name() string {
	return "敛众而击"
}

func (g GatherTheCrowdAndStrike) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (g GatherTheCrowdAndStrike) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GatherTheCrowdAndStrike) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GatherTheCrowdAndStrike) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (g GatherTheCrowdAndStrike) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GatherTheCrowdAndStrike) Execute() {
}

func (g GatherTheCrowdAndStrike) IsTriggerPrepare() bool {
	return false
}
