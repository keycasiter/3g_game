package tactics

import (
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

//奋矛英姿
type BraveSpearHeroicPoseTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BraveSpearHeroicPoseTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (b BraveSpearHeroicPoseTactic) Prepare() {
	panic("implement me")
}

func (b BraveSpearHeroicPoseTactic) Id() consts.TacticId {
	return consts.BraveSpearHeroicPose
}

func (b BraveSpearHeroicPoseTactic) Name() string {
	return "奋矛英姿"
}

func (b BraveSpearHeroicPoseTactic) TacticsSource() consts.TacticsSource {
	panic("implement me")
}

func (b BraveSpearHeroicPoseTactic) GetTriggerRate() float64 {
	panic("implement me")
}

func (b BraveSpearHeroicPoseTactic) SetTriggerRate(rate float64) {
	panic("implement me")
}

func (b BraveSpearHeroicPoseTactic) TacticsType() consts.TacticsType {
	panic("implement me")
}

func (b BraveSpearHeroicPoseTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (b BraveSpearHeroicPoseTactic) Execute() {
	panic("implement me")
}

func (b BraveSpearHeroicPoseTactic) IsTriggerPrepare() bool {
	panic("implement me")
}
