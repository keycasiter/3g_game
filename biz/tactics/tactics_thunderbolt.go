package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 落雷
// 对随机其他单体（有5%几率对友军释放）造成谋略攻击（伤害率170%，受智力影响），并使其受到谋略伤害增加18%，持续2回合
type ThunderboltTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThunderboltTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.5
	return t
}

func (t ThunderboltTactic) Prepare() {

}

func (t ThunderboltTactic) Id() consts.TacticId {
	return consts.Thunderbolt
}

func (t ThunderboltTactic) Name() string {
	return "落雷"
}

func (t ThunderboltTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t ThunderboltTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThunderboltTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThunderboltTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t ThunderboltTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThunderboltTactic) Execute() {
}

func (t ThunderboltTactic) IsTriggerPrepare() bool {
	return false
}
