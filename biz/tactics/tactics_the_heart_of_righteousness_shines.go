package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//义心昭烈
//战斗中，自身造成治疗效果时，使目标受到主动战法伤害降低14%（受智力影响），持续1回合，
//当自身兵力首次低于40%时，使我军全体受到伤害时会分摊50%伤害，持续1回合
type TheHeartOfRighteousnessShinesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheHeartOfRighteousnessShinesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TheHeartOfRighteousnessShinesTactic) Prepare() {

}

func (t TheHeartOfRighteousnessShinesTactic) Id() consts.TacticId {
	return consts.TheHeartOfRighteousnessShines
}

func (t TheHeartOfRighteousnessShinesTactic) Name() string {
	return "义心昭烈"
}

func (t TheHeartOfRighteousnessShinesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheHeartOfRighteousnessShinesTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheHeartOfRighteousnessShinesTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheHeartOfRighteousnessShinesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (t TheHeartOfRighteousnessShinesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheHeartOfRighteousnessShinesTactic) Execute() {
}

func (t TheHeartOfRighteousnessShinesTactic) IsTriggerPrepare() bool {
	return false
}
