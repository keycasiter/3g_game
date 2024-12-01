package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 太平道法
// 获得28%奇谋并提高自带主动战法发动率(6%，若为准备战法则提高12%，受智力影响)，
// 自身为黄巾军主将时，使黄巾军副将同样获得自带战法发动率提升
type TaipingLawTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TaipingLawTactic) IsTriggerPrepare() bool {
	return false
}

func (t TaipingLawTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TaipingLawTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t TaipingLawTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TaipingLawTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TaipingLawTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//获得28%奇谋并提高自带主动战法发动率(6%，若为准备战法则提高12%，受智力影响)，
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_EnhanceStrategy, &vo.EffectHolderParams{
		EffectRate: 0.28,
		FromTactic: t.Id(),
	}).IsSuccess {
		hlog.CtxInfof(ctx, "[%s]的奇谋提高了28%%",
			currentGeneral.BaseInfo.Name,
		)
	}

	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_HuangTianDangLi, &vo.EffectHolderParams{
		EffectRate: 1.0,
		FromTactic: t.Id(),
	}).IsSuccess {
		//TODO 自身为黄巾军主将时，使黄巾军副将同样获得自带战法发动率提升

		//注册效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral
			currentTactic := (params.CurrentTactic).(_interface.Tactics)

			if currentTactic.TacticsSource() == consts.TacticsSource_SelfContained {
				improveRate := 0.06 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
				if currentTactic.IsTriggerPrepare() {
					improveRate = 0.12 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
				}
				hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
					triggerGeneral.BaseInfo.Name,
					t.Name(),
					consts.BuffEffectType_HuangTianDangLi,
				)
				hlog.CtxInfof(ctx, "[%s]的【%s】发动几率提升了%.2f%%(%.2f%%↗%.2f%%)",
					triggerGeneral.BaseInfo.Name,
					currentTactic.Name(),
					improveRate*100,
					currentTactic.GetTriggerRate()*100,
					(currentTactic.GetTriggerRate()+improveRate)*100,
				)

				triggerRate := currentTactic.GetTriggerRate() + improveRate
				currentTactic.SetTriggerRate(triggerRate)

				//注册消失效果
				util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeTactic := params.CurrentTactic.(_interface.Tactics)

					if currentTactic.Id() == revokeTactic.Id() {
						revokeRate := revokeTactic.GetTriggerRate() - improveRate
						revokeTactic.SetTriggerRate(revokeRate)
					}

					return revokeResp
				})
			}

			return triggerResp
		})
	}
}

func (t TaipingLawTactic) Id() consts.TacticId {
	return consts.TaipingLaw
}

func (t TaipingLawTactic) Name() string {
	return "太平道法"
}

func (t TaipingLawTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t TaipingLawTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TaipingLawTactic) Execute() {
	return
}

func (a TaipingLawTactic) SetTriggerPrepare(triggerPrepare bool) {
}
