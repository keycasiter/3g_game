package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//枪舞如风
type GunDanceLikeTheWindTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GunDanceLikeTheWindTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (g GunDanceLikeTheWindTactic) Prepare() {
	panic("implement me")
}

func (g GunDanceLikeTheWindTactic) Id() consts.TacticId {
	return consts.GunDanceLikeTheWind
}

func (g GunDanceLikeTheWindTactic) Name() string {
	return "枪舞如风"
}

func (g GunDanceLikeTheWindTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (g GunDanceLikeTheWindTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (g GunDanceLikeTheWindTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (g GunDanceLikeTheWindTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (g GunDanceLikeTheWindTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (g GunDanceLikeTheWindTactic) Execute() {
	panic("implement me")
}

func (g GunDanceLikeTheWindTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
