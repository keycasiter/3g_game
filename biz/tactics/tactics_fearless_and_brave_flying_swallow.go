package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 轻勇飞燕
type FearlessAndBraveFlyingSwallowTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FearlessAndBraveFlyingSwallowTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	//TODO implement me
	panic("implement me")
}

func (f FearlessAndBraveFlyingSwallowTactic) Prepare() {
	//TODO implement me
	panic("implement me")
}

func (f FearlessAndBraveFlyingSwallowTactic) Id() consts.TacticId {
	return consts.FearlessAndBraveFlyingSwallow
}

func (f FearlessAndBraveFlyingSwallowTactic) Name() string {
	return "轻勇飞燕"
}

func (f FearlessAndBraveFlyingSwallowTactic) TacticsSource() consts.TacticsSource {
	//TODO implement me
	panic("implement me")
}

func (f FearlessAndBraveFlyingSwallowTactic) GetTriggerRate() float64 {
	//TODO implement me
	panic("implement me")
}

func (f FearlessAndBraveFlyingSwallowTactic) SetTriggerRate(rate float64) {
	//TODO implement me
	panic("implement me")
}

func (f FearlessAndBraveFlyingSwallowTactic) TacticsType() consts.TacticsType {
	//TODO implement me
	panic("implement me")
}

func (f FearlessAndBraveFlyingSwallowTactic) SupportArmTypes() []consts.ArmType {
	//TODO implement me
	panic("implement me")
}

func (f FearlessAndBraveFlyingSwallowTactic) Execute() {
	//TODO implement me
	panic("implement me")
}

func (f FearlessAndBraveFlyingSwallowTactic) IsTriggerPrepare() bool {
	//TODO implement me
	panic("implement me")
}
