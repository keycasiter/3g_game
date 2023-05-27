package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//骁健神行
type VigorousAndWalkTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (v VigorousAndWalkTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (v VigorousAndWalkTactic) Prepare() {
	panic("implement me")
}

func (v VigorousAndWalkTactic) Id() consts.TacticId {
	return consts.VigorousAndWalk
}

func (v VigorousAndWalkTactic) Name() string {
	return "骁健神行"
}

func (v VigorousAndWalkTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (v VigorousAndWalkTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (v VigorousAndWalkTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (v VigorousAndWalkTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (v VigorousAndWalkTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (v VigorousAndWalkTactic) Execute() {
	panic("implement me")
}

func (v VigorousAndWalkTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
