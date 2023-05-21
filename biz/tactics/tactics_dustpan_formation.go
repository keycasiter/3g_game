package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//箕形阵
//战斗前3回合，使敌军主将造成伤害降低40%（受武力影响），并使我军随机副将受到兵刃伤害降低18%，另一名副将受到谋略伤害降低18%
type DustpanFormationTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DustpanFormationTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 1.0
	return d
}

func (d DustpanFormationTactic) Prepare() {
}

func (d DustpanFormationTactic) Id() consts.TacticId {
	return consts.DustpanFormation
}

func (d DustpanFormationTactic) Name() string {
	return "箕形阵"
}

func (d DustpanFormationTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (d DustpanFormationTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DustpanFormationTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DustpanFormationTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (d DustpanFormationTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DustpanFormationTactic) Execute() {
}

func (d DustpanFormationTactic) IsTriggerPrepare() bool {
	return false
}
