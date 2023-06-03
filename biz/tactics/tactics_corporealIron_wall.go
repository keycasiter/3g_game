package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 肉身铁壁
type CorporealIronWallTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CorporealIronWallTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (c CorporealIronWallTactic) Prepare() {
	panic("implement me")
}

func (c CorporealIronWallTactic) Id() consts.TacticId {
	return consts.CorporealIronWall
}

func (c CorporealIronWallTactic) Name() string {
	return "肉身铁壁"
}

func (c CorporealIronWallTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (c CorporealIronWallTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (c CorporealIronWallTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (c CorporealIronWallTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (c CorporealIronWallTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (c CorporealIronWallTactic) Execute() {
	panic("implement me")
}

func (c CorporealIronWallTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
