package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//藤甲兵
//将盾兵进阶为刀枪不如的藤甲兵：
//我军全体受到兵刃伤害降低24%（受统率影响），但处于灼烧状态时每回合额外损失兵力（伤害率300%）；
//若兀突骨统领，则处于灼烧状态时的损失兵力降低（伤害率250%）
type TengjiaSoldierTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TengjiaSoldierTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TengjiaSoldierTactic) Prepare() {
}

func (t TengjiaSoldierTactic) Id() consts.TacticId {
	return consts.TengjiaSoldier
}

func (t TengjiaSoldierTactic) Name() string {
	return "藤甲兵"
}

func (t TengjiaSoldierTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TengjiaSoldierTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TengjiaSoldierTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TengjiaSoldierTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t TengjiaSoldierTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Mauler,
	}
}

func (t TengjiaSoldierTactic) Execute() {

}

func (t TengjiaSoldierTactic) IsTriggerPrepare() bool {
	return false
}