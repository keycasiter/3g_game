package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

type IkkiTousenTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

//一骑当千
//普通攻击之后，对敌军全体发动一次兵刃攻击（伤害率72%），自身为主将时，该次兵刃攻击更为强力（伤害率108%）
//发动概率30%
func (i IkkiTousenTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.3
	return i
}

func (i IkkiTousenTactic) Prepare() {

}

func (i IkkiTousenTactic) Id() consts.TacticId {
	return consts.IkkiTousen
}

func (i IkkiTousenTactic) Name() string {
	return "一骑当千"
}

func (i IkkiTousenTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (i IkkiTousenTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i IkkiTousenTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i IkkiTousenTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (i IkkiTousenTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i IkkiTousenTactic) Execute() {

}

func (i IkkiTousenTactic) IsTriggerPrepare() bool {
	return false
}