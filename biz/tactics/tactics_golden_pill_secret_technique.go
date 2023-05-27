package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//金丹秘术
type GoldenPillSecretTechniqueTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (g GoldenPillSecretTechniqueTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (g GoldenPillSecretTechniqueTactic) Prepare() {
	panic("implement me")
}

func (g GoldenPillSecretTechniqueTactic) Id() consts.TacticId {
	return consts.GoldenPillSecretTechnique
}

func (g GoldenPillSecretTechniqueTactic) Name() string {
	return "金丹秘术"
}

func (g GoldenPillSecretTechniqueTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (g GoldenPillSecretTechniqueTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (g GoldenPillSecretTechniqueTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (g GoldenPillSecretTechniqueTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (g GoldenPillSecretTechniqueTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (g GoldenPillSecretTechniqueTactic) Execute() {
	panic("implement me")
}

func (g GoldenPillSecretTechniqueTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
