package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//士争先赴
//提高自带主动战法20%伤害，成功发动自带主动战法前，50%概率对敌方群体（2～3人）造成兵刃伤害（伤害率120%）
type ScholarsStriveToGoFirstTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s ScholarsStriveToGoFirstTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s ScholarsStriveToGoFirstTactic) Prepare() {
}

func (s ScholarsStriveToGoFirstTactic) Id() consts.TacticId {
	return consts.ScholarsStriveToGoFirst
}

func (s ScholarsStriveToGoFirstTactic) Name() string {
	return "士争先赴"
}

func (s ScholarsStriveToGoFirstTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s ScholarsStriveToGoFirstTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s ScholarsStriveToGoFirstTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s ScholarsStriveToGoFirstTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (s ScholarsStriveToGoFirstTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s ScholarsStriveToGoFirstTactic) Execute() {

}

func (s ScholarsStriveToGoFirstTactic) IsTriggerPrepare() bool {
	return false
}
