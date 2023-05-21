package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//文武双全
//战斗中，自己每次造成谋略伤害时，增加30点智力，最多叠加5次，每次造成兵刃伤害时，增加30点武力，最多叠加5次
type BeAdeptWithBothPenAndSword struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BeAdeptWithBothPenAndSword) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BeAdeptWithBothPenAndSword) Prepare() {

}

func (b BeAdeptWithBothPenAndSword) Id() consts.TacticId {
	return consts.BeAdeptWithBothPenAndSword
}

func (b BeAdeptWithBothPenAndSword) Name() string {
	return "文武双全"
}

func (b BeAdeptWithBothPenAndSword) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BeAdeptWithBothPenAndSword) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BeAdeptWithBothPenAndSword) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BeAdeptWithBothPenAndSword) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BeAdeptWithBothPenAndSword) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BeAdeptWithBothPenAndSword) Execute() {
}

func (b BeAdeptWithBothPenAndSword) IsTriggerPrepare() bool {
	return false
}
