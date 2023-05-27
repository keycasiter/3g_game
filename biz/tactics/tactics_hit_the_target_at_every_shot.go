package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//百步穿杨
type HitTheTargetAtEveryShotTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HitTheTargetAtEveryShotTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (h HitTheTargetAtEveryShotTactic) Prepare() {
	panic("implement me")
}

func (h HitTheTargetAtEveryShotTactic) Id() consts.TacticId {
	return consts.HitTheTargetAtEveryShot
}

func (h HitTheTargetAtEveryShotTactic) Name() string {
	return "百步穿杨"
}

func (h HitTheTargetAtEveryShotTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (h HitTheTargetAtEveryShotTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (h HitTheTargetAtEveryShotTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (h HitTheTargetAtEveryShotTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (h HitTheTargetAtEveryShotTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (h HitTheTargetAtEveryShotTactic) Execute() {
	panic("implement me")
}

func (h HitTheTargetAtEveryShotTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
