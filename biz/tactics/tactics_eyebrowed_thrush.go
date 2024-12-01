package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 白眉
// 战斗中，自身主动战法的发动几率提高12%
// 被动 100%
type EyebrowedThrushTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (e EyebrowedThrushTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	e.tacticsParams = tacticsParams
	e.triggerRate = 1.0
	return e
}

func (e EyebrowedThrushTactic) Prepare() {
	ctx := e.tacticsParams.Ctx
	currentGeneral := e.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		e.Name(),
	)

	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
		EffectRate:     0.12,
		FromTactic:     e.Id(),
		ProduceGeneral: currentGeneral,
	})
}

func (e EyebrowedThrushTactic) Id() consts.TacticId {
	return consts.EyebrowedThrush
}

func (e EyebrowedThrushTactic) Name() string {
	return "白眉"
}

func (e EyebrowedThrushTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (e EyebrowedThrushTactic) GetTriggerRate() float64 {
	return e.triggerRate
}

func (e EyebrowedThrushTactic) SetTriggerRate(rate float64) {
	e.triggerRate = rate
}

func (e EyebrowedThrushTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (e EyebrowedThrushTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (e EyebrowedThrushTactic) Execute() {

}

func (e EyebrowedThrushTactic) IsTriggerPrepare() bool {
	return false
}

func (a EyebrowedThrushTactic) SetTriggerPrepare(triggerPrepare bool) {
}
