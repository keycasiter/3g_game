package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
)

// 窃幸乘宠
// 我方主将恢复兵力且自身不为主将时，降低其20%治疗量，自身会恢复降低的治疗量，奇数回合对敌军群体（2人）造成谋略伤害（伤害率90%，受智力影响），
// 额外对其中智力低于自身的单位造成谋略伤害（伤害率120%，受智力影响）
// 指挥，100%
type StealingLuckAndRidingPetsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StealingLuckAndRidingPetsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s StealingLuckAndRidingPetsTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	// 我方主将恢复兵力且自身不为主将时，降低其20%治疗量，自身会恢复降低的治疗量，奇数回合对敌军群体（2人）造成谋略伤害（伤害率90%，受智力影响），
	// 额外对其中智力低于自身的单位造成谋略伤害（伤害率120%，受智力影响）

}

func (s StealingLuckAndRidingPetsTactic) Id() consts.TacticId {
	return consts.StealingLuckAndRidingPets
}

func (s StealingLuckAndRidingPetsTactic) Name() string {
	return "窃幸乘宠"
}

func (s StealingLuckAndRidingPetsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s StealingLuckAndRidingPetsTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s StealingLuckAndRidingPetsTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s StealingLuckAndRidingPetsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s StealingLuckAndRidingPetsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s StealingLuckAndRidingPetsTactic) Execute() {

}

func (s StealingLuckAndRidingPetsTactic) IsTriggerPrepare() bool {
	return false
}
