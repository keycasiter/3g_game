package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//横扫千军
//对敌军全体造成100%兵刃伤害，若目标处于缴械或者计穷状态则有30%概率使目标处于震慑状态（无法行动），持续1回合
type SweepAwayTheMillionsOfEnemyTroopsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.4
	return s
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Prepare() {

}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Id() consts.TacticId {
	return consts.SweepAwayTheMillionsOfEnemyTroops
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Name() string {
	return "横扫千军"
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) Execute() {
}

func (s SweepAwayTheMillionsOfEnemyTroopsTactic) IsTriggerPrepare() bool {
	return false
}
