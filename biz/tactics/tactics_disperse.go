package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 驱散
// 解除敌军全体身上的增益效果，并提高自己28点智力，持续3回合
type DisperseTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DisperseTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 0.45
	return d
}

func (d DisperseTactic) Prepare() {
}

func (d DisperseTactic) Id() consts.TacticId {
	return consts.Disperse
}

func (d DisperseTactic) Name() string {
	return "驱散"
}

func (d DisperseTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DisperseTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DisperseTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DisperseTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (d DisperseTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DisperseTactic) Execute() {
}

func (d DisperseTactic) IsTriggerPrepare() bool {
	return false
}
