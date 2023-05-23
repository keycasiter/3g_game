package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 净化
// 提升我军群体（2人）24点武力、智力、速度，持续2回合，并移除负面效果
type PurifyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PurifyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 0.45
	return p
}

func (p PurifyTactic) Prepare() {

}

func (p PurifyTactic) Id() consts.TacticId {
	return consts.Purify
}

func (p PurifyTactic) Name() string {
	return "净化"
}

func (p PurifyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (p PurifyTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PurifyTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PurifyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (p PurifyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PurifyTactic) Execute() {

}

func (p PurifyTactic) IsTriggerPrepare() bool {
	return false
}
