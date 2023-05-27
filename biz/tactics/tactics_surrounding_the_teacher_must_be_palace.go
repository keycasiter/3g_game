package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//围师必阙
type SurroundingTheTeacherMustBePalaceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SurroundingTheTeacherMustBePalaceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (s SurroundingTheTeacherMustBePalaceTactic) Prepare() {
	panic("implement me")
}

func (s SurroundingTheTeacherMustBePalaceTactic) Id() consts.TacticId {
	return consts.SurroundingTheTeacherMustBePalace
}

func (s SurroundingTheTeacherMustBePalaceTactic) Name() string {
	return "围师必阙"
}

func (s SurroundingTheTeacherMustBePalaceTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (s SurroundingTheTeacherMustBePalaceTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (s SurroundingTheTeacherMustBePalaceTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (s SurroundingTheTeacherMustBePalaceTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (s SurroundingTheTeacherMustBePalaceTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (s SurroundingTheTeacherMustBePalaceTactic) Execute() {
	panic("implement me")
}

func (s SurroundingTheTeacherMustBePalaceTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
