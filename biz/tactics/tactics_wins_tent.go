package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 校胜帷幄
// 战斗中，提高己方主将14%奇谋几率，及20%奇谋伤害，同时为己方主将分担30%的伤害（自身为主将时无效）
// 指挥，100%
type WinsTentTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WinsTentTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 1.0
	return w
}

func (w WinsTentTactic) Prepare() {
	ctx := w.tacticsParams.Ctx
	currentGeneral := w.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		w.Name(),
	)

	// 战斗中，提高己方主将14%奇谋几率，及20%奇谋伤害，同时为己方主将分担30%的伤害（自身为主将时无效）
	pairMasterGeneral := util.GetPairMasterGeneral(currentGeneral, w.tacticsParams)
	//奇谋几率及伤害
	util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_EnhanceStrategy, &vo.EffectHolderParams{
		TriggerRate:    0.14,
		EffectRate:     0.2,
		FromTactic:     w.Id(),
		ProduceGeneral: currentGeneral,
	})
	//分担伤害
	if !currentGeneral.IsMaster {
		util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_ShareResponsibilityFor, &vo.EffectHolderParams{
			EffectRate:                      0.3,
			FromTactic:                      w.Id(),
			ProduceGeneral:                  currentGeneral,
			ShareResponsibilityForByGeneral: currentGeneral,
		})
	}
}

func (w WinsTentTactic) Id() consts.TacticId {
	return consts.WinsTent
}

func (w WinsTentTactic) Name() string {
	return "校胜帷幄"
}

func (w WinsTentTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (w WinsTentTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WinsTentTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WinsTentTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (w WinsTentTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (w WinsTentTactic) Execute() {

}

func (w WinsTentTactic) IsTriggerPrepare() bool {
	return false
}

func (a WinsTentTactic) SetTriggerPrepare(triggerPrepare bool) {
}
