package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//合军聚众
//战斗中，使自己获得休整状态，每回合恢复一定兵力（回复率124%）
type GatheringOfTroopsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GatheringOfTroopsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 1.0
	return g
}

func (g GatheringOfTroopsTactic) Prepare() {
}

func (g GatheringOfTroopsTactic) Id() consts.TacticId {
	return consts.GatheringOfTroops
}

func (g GatheringOfTroopsTactic) Name() string {
	return "合军聚众"
}

func (g GatheringOfTroopsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (g GatheringOfTroopsTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GatheringOfTroopsTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GatheringOfTroopsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (g GatheringOfTroopsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GatheringOfTroopsTactic) Execute() {

}

func (g GatheringOfTroopsTactic) IsTriggerPrepare() bool {
	return false
}