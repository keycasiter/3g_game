package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 狮子奋迅
// 对敌军单体及额外1～2名敌人造成兵刃攻击（伤害率118%），并使自身主动战法发动几率提高10%，
// 自身为主将时，发动几率提高至15%，持续2回合，如果单体目标为敌军主将则使其陷入叛逃状态，每回合持续造成伤害（伤害率102%），持续2回合
// 主动，35%
type TheLionFliesFastTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheLionFliesFastTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t TheLionFliesFastTactic) Prepare() {
	panic("implement me")
}

func (t TheLionFliesFastTactic) Id() consts.TacticId {
	return consts.TheLionFliesFast
}

func (t TheLionFliesFastTactic) Name() string {
	return "狮子奋迅"
}

func (t TheLionFliesFastTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (t TheLionFliesFastTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (t TheLionFliesFastTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (t TheLionFliesFastTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (t TheLionFliesFastTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (t TheLionFliesFastTactic) Execute() {
	panic("implement me")
}

func (t TheLionFliesFastTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
