package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 毒泉拒蜀
// 对敌军单体（80%概率选择敌军统率最高的武将）施加猛毒，2回合后消失
// 拥有猛毒的敌军每次受到普通攻击会叠加一层猛毒，最多叠加3层，层数叠满时会立即消失并清除自身所有战法冷却时间，
// 猛毒消失时会对敌军全体造成谋略伤害（伤害率150%，每有一层猛毒伤害率提高40%，受智力影响），
// 我方蛮族造成的非普攻伤害也会增加猛毒层数，该战法发动后会进入1回合冷却
// 主动 60%
type PoisonousSpringRefusesShuTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PoisonousSpringRefusesShuTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 1.0
	return p
}

func (p PoisonousSpringRefusesShuTactic) Prepare() {

}

func (p PoisonousSpringRefusesShuTactic) Id() consts.TacticId {
	return consts.PoisonousSpringRefusesShu
}

func (p PoisonousSpringRefusesShuTactic) Name() string {
	return "毒泉拒蜀"
}

func (p PoisonousSpringRefusesShuTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (p PoisonousSpringRefusesShuTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) Execute() {
	panic("implement me")
}

func (p PoisonousSpringRefusesShuTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
