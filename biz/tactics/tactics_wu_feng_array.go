package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 武锋阵
// 我军三名武将兵种适性均不相同时，战斗前5回合，我军主将有80%几率优先成为敌军战法的目标，且使该战法对其造成的伤害降低30%（受主将统率影响），
// 战斗中，奇数回合使兵种适性较低的副将恢复我军单体兵力（治疗率184%）
// 偶数回合使兵种适性较高的副将造成伤害提高15%（可叠加）
type WuFengArrayTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (w WuFengArrayTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 1.0
	return w
}

func (w WuFengArrayTactic) Prepare() {
	ctx := w.tacticsParams.Ctx
	currentGeneral := w.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		w.Name(),
	)
	// 我军三名武将兵种适性均不相同时，战斗前5回合，我军主将有80%几率优先成为敌军战法的目标，且使该战法对其造成的伤害降低30%（受主将统率影响），
	armsAbilityMap := make(map[consts.ArmsAbility]*vo.BattleGeneral, 0)
	pairGenerals := util.GetPairGeneralArr(w.tacticsParams)

	armType := util.GetPairArmType(currentGeneral, w.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		switch armType {
		case consts.ArmType_Mauler:
			armsAbilityMap[pairGeneral.BaseInfo.ArmsAttr.Mauler] = pairGeneral
		case consts.ArmType_Spearman:
			armsAbilityMap[pairGeneral.BaseInfo.ArmsAttr.Spearman] = pairGeneral
		case consts.ArmType_Cavalry:
			armsAbilityMap[pairGeneral.BaseInfo.ArmsAttr.Cavalry] = pairGeneral
		case consts.ArmType_Archers:
			armsAbilityMap[pairGeneral.BaseInfo.ArmsAttr.Archers] = pairGeneral
		case consts.ArmType_Apparatus:
			armsAbilityMap[pairGeneral.BaseInfo.ArmsAttr.Apparatus] = pairGeneral
		}
	}
	if len(armsAbilityMap) != 3 {
		hlog.CtxInfof(ctx, "武将兵种适性不符合均不相同，无法发动")
		return
	}

	viceGenerals := util.GetPairViceGenerals(w.tacticsParams)
	highGeneral := util.GetHighestArmAbilityGeneral(viceGenerals, armType)
	lowGeneral := util.GetLowestArmAbilityGeneral(viceGenerals, armType)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggeRound := params.CurrentRound

		// 战斗中，奇数回合使兵种适性较低的副将恢复我军单体兵力（治疗率184%）
		if triggeRound%2 != 0 {
			pairGeneral := util.GetPairOneGeneral(w.tacticsParams, triggerGeneral)
			resumeNum := cast.ToInt64(lowGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.84)
			util.ResumeSoldierNum(&util.ResumeParams{
				Ctx:            ctx,
				TacticsParams:  w.tacticsParams,
				ProduceGeneral: lowGeneral,
				SufferGeneral:  pairGeneral,
				ResumeNum:      resumeNum,
				TacticId:       w.Id(),
			})
		} else {
			// 偶数回合使兵种适性较高的副将造成伤害提高15%（可叠加）
			//兵刃伤害
			util.BuffEffectWrapSet(ctx, highGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.15,
				EffectTimes:    1,
				MaxEffectTimes: consts.INT64_MAX,
				FromTactic:     w.Id(),
				ProduceGeneral: highGeneral,
			})
			//谋略伤害
			util.BuffEffectWrapSet(ctx, highGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
				EffectRate:     0.15,
				EffectTimes:    1,
				MaxEffectTimes: consts.INT64_MAX,
				FromTactic:     w.Id(),
				ProduceGeneral: highGeneral,
			})
		}

		return triggerResp
	})
}

func (w WuFengArrayTactic) Id() consts.TacticId {
	return consts.WuFengArray
}

func (w WuFengArrayTactic) Name() string {
	return "武锋阵"
}

func (w WuFengArrayTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (w WuFengArrayTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WuFengArrayTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WuFengArrayTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_TroopsTactics
}

func (w WuFengArrayTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (w WuFengArrayTactic) Execute() {

}

func (w WuFengArrayTactic) IsTriggerPrepare() bool {
	return false
}
