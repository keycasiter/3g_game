package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//绝地反击
//战斗中，自己每次受到兵刃伤害后，武力提升6点，最大叠加10次；第5回合时，根据叠加次数对敌军全体造成兵刃伤害（伤害率120%，每次提高14%伤害率）
type JediCounterattackTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JediCounterattackTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	j.tacticsParams = tacticsParams
	j.triggerRate = 1.0
	return j
}

func (j JediCounterattackTactic) Prepare() {
}

func (j JediCounterattackTactic) Id() consts.TacticId {
	return consts.JediCounterattack
}

func (j JediCounterattackTactic) Name() string {
	return "绝地反击"
}

func (j JediCounterattackTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (j JediCounterattackTactic) GetTriggerRate() float64 {
	return j.triggerRate
}

func (j JediCounterattackTactic) SetTriggerRate(rate float64) {
	j.triggerRate = rate
}

func (j JediCounterattackTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (j JediCounterattackTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Spearman,
	}
}

func (j JediCounterattackTactic) Execute() {

}

func (j JediCounterattackTactic) IsTriggerPrepare() bool {
	return false
}
