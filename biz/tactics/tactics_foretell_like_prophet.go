package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 料事如神
// 对敌军群体（2人）造成谋略伤害（伤害率106%，受智力影响），并使其造成伤害降低16%，持续2回合
type ForetellLikeProphetTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f ForetellLikeProphetTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.35
	return f
}

func (f ForetellLikeProphetTactic) Prepare() {
}

func (f ForetellLikeProphetTactic) Id() consts.TacticId {
	return consts.ForetellLikeProphet
}

func (f ForetellLikeProphetTactic) Name() string {
	return "料事如神"
}

func (f ForetellLikeProphetTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f ForetellLikeProphetTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f ForetellLikeProphetTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f ForetellLikeProphetTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f ForetellLikeProphetTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f ForetellLikeProphetTactic) Execute() {

}

func (f ForetellLikeProphetTactic) IsTriggerPrepare() bool {
	return false
}
