package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 义胆雄心
// 战斗中，奇数回合会对敌军单体造成184%兵刃伤害兵降低武力64点，持续2回合
// 偶数回合会对敌军群体(2人)造成76%谋略伤害（受智力影响）并降低智力34点，持续2回合；
// 自身为主将时，降低属性效果受自身对应属性影响
type BraveAmbitionTactic struct {
	tacticsParams *model.TacticsParams
}

func (b BraveAmbitionTactic) TriggerRate() float64 {
	return 1.0
}

func (b BraveAmbitionTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	return b
}

func (b BraveAmbitionTactic) Prepare() {
	return
}

func (b BraveAmbitionTactic) Id() consts.TacticId {
	return consts.BraveAmbition
}

func (b BraveAmbitionTactic) Name() string {
	return "义胆雄心"
}

func (b BraveAmbitionTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BraveAmbitionTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BraveAmbitionTactic) Execute() {
	return
}
