package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 破阵摧坚
// 准备1回合，使敌军群体（2人）统率、智力降低80点（受武力影响），持续2回合，并对其发动一次兵刃攻击（伤害率158%）
type BreakingThroughTheFormationAndDestroyingTheFirmTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.35
	return b
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Prepare() {
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Id() consts.TacticId {
	return consts.BreakingThroughTheFormationAndDestroyingTheFirm
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Name() string {
	return "破阵摧坚"
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) Execute() {

}

func (b BreakingThroughTheFormationAndDestroyingTheFirmTactic) IsTriggerPrepare() bool {
	return b.isTriggerPrepare
}
