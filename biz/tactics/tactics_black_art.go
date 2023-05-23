package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 妖术
// 准备1回合，对敌军全体施加沙暴状态，每回合持续造成伤害（伤害率72%，受智力影响）并使自己获得1次抵御，可以免疫伤害，持续2回合
type BlackArtTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (b BlackArtTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 0.4
	return b
}

func (b BlackArtTactic) Prepare() {

}

func (b BlackArtTactic) Id() consts.TacticId {
	return consts.BlackArt
}

func (b BlackArtTactic) Name() string {
	return "妖术"
}

func (b BlackArtTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BlackArtTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BlackArtTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BlackArtTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (b BlackArtTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BlackArtTactic) Execute() {
}

func (b BlackArtTactic) IsTriggerPrepare() bool {
	return false
}
