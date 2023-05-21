package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//威谋靡亢
//准备1回合，对敌军群体（2人）施加虚拟（无法造成伤害）状态，持续2回合；如果目标已处于虚弱状态则使其陷入叛逃状态，
//每回合持续造成伤害（伤害率158%，受武力或智力最高一项影响，无视防御），持续2回合
type IntenseAndPowerfulTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (i IntenseAndPowerfulTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.4
	return i
}

func (i IntenseAndPowerfulTactic) Prepare() {
	panic("implement me")
}

func (i IntenseAndPowerfulTactic) Id() consts.TacticId {
	return consts.IntenseAndPowerful
}

func (i IntenseAndPowerfulTactic) Name() string {
	return "威谋靡亢"
}

func (i IntenseAndPowerfulTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (i IntenseAndPowerfulTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i IntenseAndPowerfulTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i IntenseAndPowerfulTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i IntenseAndPowerfulTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i IntenseAndPowerfulTactic) Execute() {

}

func (i IntenseAndPowerfulTactic) IsTriggerPrepare() bool {
	return i.isTriggerPrepare
}
