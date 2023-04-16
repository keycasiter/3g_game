package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//偃旗息鼓
type LowerBannersAndMuffleDrumsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LowerBannersAndMuffleDrumsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 1.0
	return l
}

func (l LowerBannersAndMuffleDrumsTactic) Prepare() {
	panic("implement me")
}

func (l LowerBannersAndMuffleDrumsTactic) Id() consts.TacticId {
	return consts.LowerBannersAndMuffleDrums
}

func (l LowerBannersAndMuffleDrumsTactic) Name() string {
	return "偃旗息鼓"
}

func (l LowerBannersAndMuffleDrumsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (l LowerBannersAndMuffleDrumsTactic) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LowerBannersAndMuffleDrumsTactic) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LowerBannersAndMuffleDrumsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (l LowerBannersAndMuffleDrumsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LowerBannersAndMuffleDrumsTactic) Execute() {
	return
}
