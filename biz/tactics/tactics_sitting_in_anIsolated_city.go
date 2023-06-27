package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 坐守孤城
// 恢复我军群体（2人）兵力（治疗率116%，受智力影响）
// 主动，45%
type SittingInAnIsolatedCityTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SittingInAnIsolatedCityTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.45
	return s
}

func (s SittingInAnIsolatedCityTactic) Prepare() {
}

func (s SittingInAnIsolatedCityTactic) Id() consts.TacticId {
	return consts.SittingInAnIsolatedCity
}

func (s SittingInAnIsolatedCityTactic) Name() string {
	return "坐守孤城"
}

func (s SittingInAnIsolatedCityTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SittingInAnIsolatedCityTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SittingInAnIsolatedCityTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SittingInAnIsolatedCityTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SittingInAnIsolatedCityTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SittingInAnIsolatedCityTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	// 恢复我军群体（2人）兵力（治疗率116%，受智力影响）
	pairGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, s.tacticsParams)
	resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.16)
	for _, general := range pairGenerals {
		util.ResumeSoldierNum(&util.ResumeParams{
			Ctx:            ctx,
			TacticsParams:  s.tacticsParams,
			ProduceGeneral: currentGeneral,
			SufferGeneral:  general,
			ResumeNum:      resumeNum,
		})
	}
}

func (s SittingInAnIsolatedCityTactic) IsTriggerPrepare() bool {
	return false
}
