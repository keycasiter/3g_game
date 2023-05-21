package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//陷阵营
//将盾兵进阶为无往不利的陷阵营：
//我军全体武力/统率提高22点，战斗前3回合获得急救状态，受到伤害时有30%概率获得治疗（治疗率60%，受智力影响）
//若高顺统领，则治疗率将额外受统率影响
type FallIntoCampTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FallIntoCampTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 1.0
	return f
}

func (f FallIntoCampTactic) Prepare() {
}

func (f FallIntoCampTactic) Id() consts.TacticId {
	return consts.FallIntoCamp
}

func (f FallIntoCampTactic) Name() string {
	return "陷阵营"
}

func (f FallIntoCampTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FallIntoCampTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FallIntoCampTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FallIntoCampTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Arm
}

func (f FallIntoCampTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Mauler,
	}
}

func (f FallIntoCampTactic) Execute() {
}

func (f FallIntoCampTactic) IsTriggerPrepare() bool {
	return false
}
