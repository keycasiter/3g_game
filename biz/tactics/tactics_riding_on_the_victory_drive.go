package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//乘胜长驱
//战斗中，每回合使自己造成伤害提高11%，可叠加，直到战斗结束
type RidingOnTheVictoryDriveTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (r RidingOnTheVictoryDriveTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 1.0
	return r
}

func (r RidingOnTheVictoryDriveTactic) Prepare() {
}

func (r RidingOnTheVictoryDriveTactic) Id() consts.TacticId {
	return consts.RidingOnTheVictoryDrive
}

func (r RidingOnTheVictoryDriveTactic) Name() string {
	return "乘胜长驱"
}

func (r RidingOnTheVictoryDriveTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RidingOnTheVictoryDriveTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RidingOnTheVictoryDriveTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RidingOnTheVictoryDriveTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (r RidingOnTheVictoryDriveTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RidingOnTheVictoryDriveTactic) Execute() {
}

func (r RidingOnTheVictoryDriveTactic) IsTriggerPrepare() bool {
	return false
}
