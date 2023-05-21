package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//虎踞鹰扬
//战斗中使自身免疫缴械（无法进行普通攻击）状态，普通攻击之后，使自己造成兵刃伤害提高7%，最多叠加4次，
//叠加4次后，使自身获得群攻状态（伤害率30%），持续1回合
type TigerCrouchingAndEagleSoaring struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TigerCrouchingAndEagleSoaring) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TigerCrouchingAndEagleSoaring) Prepare() {
}

func (t TigerCrouchingAndEagleSoaring) Id() consts.TacticId {
	return consts.TigerCrouchingAndEagleSoaring
}

func (t TigerCrouchingAndEagleSoaring) Name() string {
	return "虎踞鹰扬"
}

func (t TigerCrouchingAndEagleSoaring) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t TigerCrouchingAndEagleSoaring) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TigerCrouchingAndEagleSoaring) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TigerCrouchingAndEagleSoaring) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t TigerCrouchingAndEagleSoaring) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TigerCrouchingAndEagleSoaring) Execute() {
}

func (t TigerCrouchingAndEagleSoaring) IsTriggerPrepare() bool {
	return false
}
