package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//鬼神霆威
//普通攻击之后，对攻击目标再次发起一次兵刃攻击（伤害率204%）
//自身为主将且当目标兵力低于50%时，额外提高伤害（受目标损失兵力影响，最多提高50%）
type GhostGodThunderForceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GhostGodThunderForceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	g.tacticsParams = tacticsParams
	g.triggerRate = 0.35
	return g
}

func (g GhostGodThunderForceTactic) Prepare() {

}

func (g GhostGodThunderForceTactic) Id() consts.TacticId {
	return consts.GhostGodThunderForce
}

func (g GhostGodThunderForceTactic) Name() string {
	return "鬼神霆威"
}

func (g GhostGodThunderForceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (g GhostGodThunderForceTactic) GetTriggerRate() float64 {
	return g.triggerRate
}

func (g GhostGodThunderForceTactic) SetTriggerRate(rate float64) {
	g.triggerRate = rate
}

func (g GhostGodThunderForceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (g GhostGodThunderForceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (g GhostGodThunderForceTactic) Execute() {

}

func (g GhostGodThunderForceTactic) IsTriggerPrepare() bool {
	return false
}
