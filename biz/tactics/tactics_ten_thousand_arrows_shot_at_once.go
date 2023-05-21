package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//万箭齐发
//准备1回合
type TenThousandArrowsShotAtOnceTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
}

func (t TenThousandArrowsShotAtOnceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.4
	return t
}

func (t TenThousandArrowsShotAtOnceTactic) Prepare() {
}

func (t TenThousandArrowsShotAtOnceTactic) Id() consts.TacticId {
	return consts.TenThousandArrowsShotAtOnce
}

func (t TenThousandArrowsShotAtOnceTactic) Name() string {
	return "万箭齐发"
}

func (t TenThousandArrowsShotAtOnceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TenThousandArrowsShotAtOnceTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TenThousandArrowsShotAtOnceTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TenThousandArrowsShotAtOnceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TenThousandArrowsShotAtOnceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (t TenThousandArrowsShotAtOnceTactic) Execute() {
}

func (t TenThousandArrowsShotAtOnceTactic) IsTriggerPrepare() bool {
	return t.isTriggerPrepare
}
