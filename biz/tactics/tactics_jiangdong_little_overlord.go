package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//江东小霸王
type JiangdongLittleOverlordTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JiangdongLittleOverlordTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (j JiangdongLittleOverlordTactic) Prepare() {
	panic("implement me")
}

func (j JiangdongLittleOverlordTactic) Id() consts.TacticId {
	return consts.JiangdongLittleOverlord
}

func (j JiangdongLittleOverlordTactic) Name() string {
	return "江东小霸王"
}

func (j JiangdongLittleOverlordTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (j JiangdongLittleOverlordTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (j JiangdongLittleOverlordTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (j JiangdongLittleOverlordTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (j JiangdongLittleOverlordTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (j JiangdongLittleOverlordTactic) Execute() {
	panic("implement me")
}

func (j JiangdongLittleOverlordTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
