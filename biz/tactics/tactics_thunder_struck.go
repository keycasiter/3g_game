package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 五雷轰顶
// 准备1回合，对敌军随机单体造成谋略攻击（伤害率136%，受智力影响），
// 并有30%概率使其进入震慑状态，持续1回合
// 共触发5次，每次独立选择目标
// 自身为主将时，若目标处于水攻状态、沙暴状态时，每多一种提高20%震慑概率
type ThunderStruckTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThunderStruckTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t ThunderStruckTactic) Prepare() {
	panic("implement me")
}

func (t ThunderStruckTactic) Id() consts.TacticId {
	return consts.ThunderStruck
}

func (t ThunderStruckTactic) Name() string {
	return "五雷轰顶"
}

func (t ThunderStruckTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t ThunderStruckTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t ThunderStruckTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t ThunderStruckTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t ThunderStruckTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t ThunderStruckTactic) Execute() {
	panic("implement me")
}

func (t ThunderStruckTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
