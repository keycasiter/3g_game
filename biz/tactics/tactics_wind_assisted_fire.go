package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type WindAssistedFireTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

//风助火势
//对敌军单体造成谋略伤害（伤害率154%，受智力影响）
//若目标处于灼烧状态，额外对目标造成一次谋略伤害（伤害率228%，受智力影响）
func (w WindAssistedFireTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 0.5
	return w
}

func (w WindAssistedFireTactic) Prepare() {
}

func (w WindAssistedFireTactic) Id() consts.TacticId {
	return consts.WindAssistedFire
}

func (w WindAssistedFireTactic) Name() string {
	return "风助火势"
}

func (w WindAssistedFireTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (w WindAssistedFireTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WindAssistedFireTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WindAssistedFireTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (w WindAssistedFireTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (w WindAssistedFireTactic) Execute() {
}

func (w WindAssistedFireTactic) IsTriggerPrepare() bool {
	return false
}