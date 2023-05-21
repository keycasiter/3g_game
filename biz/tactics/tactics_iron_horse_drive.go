package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//铁骑驱驰
//战斗前2回合，使敌军全体处于遇袭状态（行动滞后），我军全体发动突击战法后，降低普通攻击目标15%统率，持续3回合，可叠加
type IronHorseDrive struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i IronHorseDrive) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 1.0
	return i
}

func (i IronHorseDrive) Prepare() {

}

func (i IronHorseDrive) Id() consts.TacticId {
	return consts.IronHorseDrive
}

func (i IronHorseDrive) Name() string {
	return "铁骑驱驰"
}

func (i IronHorseDrive) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (i IronHorseDrive) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i IronHorseDrive) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i IronHorseDrive) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (i IronHorseDrive) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
	}
}

func (i IronHorseDrive) Execute() {
}

func (i IronHorseDrive) IsTriggerPrepare() bool {
	return false
}
