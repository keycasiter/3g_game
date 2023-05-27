package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//江东猛虎
type JiangdongTigerTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (j JiangdongTigerTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (j JiangdongTigerTactic) Prepare() {
	panic("implement me")
}

func (j JiangdongTigerTactic) Id() consts.TacticId {
	return consts.JiangdongTiger
}

func (j JiangdongTigerTactic) Name() string {
	return "江东猛虎"
}

func (j JiangdongTigerTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (j JiangdongTigerTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (j JiangdongTigerTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (j JiangdongTigerTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (j JiangdongTigerTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (j JiangdongTigerTactic) Execute() {
	panic("implement me")
}

func (j JiangdongTigerTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
