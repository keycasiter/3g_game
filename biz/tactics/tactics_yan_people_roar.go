package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 燕人咆哮
// 战斗第2、4回合，对敌军全体造成兵刃攻击（伤害率104%）
// 若目标处于缴械状态，则额外使目标统率降低50%，持续2回合
// 自身为主将时，降低统率效果额外对计穷状态的目标生效
// 被动，100%
type YanPeopleRoarTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (y YanPeopleRoarTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (y YanPeopleRoarTactic) Prepare() {
	panic("implement me")
}

func (y YanPeopleRoarTactic) Id() consts.TacticId {
	return consts.YanPeopleRoar
}

func (y YanPeopleRoarTactic) Name() string {
	return "燕人咆哮"
}

func (y YanPeopleRoarTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (y YanPeopleRoarTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (y YanPeopleRoarTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (y YanPeopleRoarTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (y YanPeopleRoarTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (y YanPeopleRoarTactic) Execute() {
	panic("implement me")
}

func (y YanPeopleRoarTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
