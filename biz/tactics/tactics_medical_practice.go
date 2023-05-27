package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//青囊
type MedicalPracticeTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (m MedicalPracticeTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (m MedicalPracticeTactic) Prepare() {
	panic("implement me")
}

func (m MedicalPracticeTactic) Id() consts.TacticId {
	return consts.MedicalPractice
}

func (m MedicalPracticeTactic) Name() string {
	return "青囊"
}

func (m MedicalPracticeTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (m MedicalPracticeTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (m MedicalPracticeTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (m MedicalPracticeTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (m MedicalPracticeTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (m MedicalPracticeTactic) Execute() {
	panic("implement me")
}

func (m MedicalPracticeTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
