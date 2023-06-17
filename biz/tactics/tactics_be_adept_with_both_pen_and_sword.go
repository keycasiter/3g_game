package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 文武双全
// 战斗中，自己每次造成谋略伤害时，增加30点智力，最多叠加5次，每次造成兵刃伤害时，增加30点武力，最多叠加5次
type BeAdeptWithBothPenAndSwordTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (b BeAdeptWithBothPenAndSwordTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	b.tacticsParams = tacticsParams
	b.triggerRate = 1.0
	return b
}

func (b BeAdeptWithBothPenAndSwordTactic) Prepare() {
	ctx := b.tacticsParams.Ctx
	currentGeneral := b.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		b.Name(),
	)

	//战斗中，自己每次造成谋略伤害时，增加30点智力，最多叠加5次，每次造成兵刃伤害时，增加30点武力，最多叠加5次
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_StrategyDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
			EffectValue:    30,
			EffectTimes:    1,
			MaxEffectTimes: 5,
			FromTactic:     b.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase += 30
			hlog.CtxInfof(ctx, "[%s]的智力提高了%.2d",
				triggerGeneral.BaseInfo.Name,
				30)
		}

		return triggerResp
	})
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_WeaponDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
			EffectValue:    30,
			EffectTimes:    1,
			MaxEffectTimes: 5,
			FromTactic:     b.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			triggerGeneral.BaseInfo.AbilityAttr.ForceBase += 30
			hlog.CtxInfof(ctx, "[%s]的武力提高了%.2d",
				triggerGeneral.BaseInfo.Name,
				30)
		}

		return triggerResp
	})
}

func (b BeAdeptWithBothPenAndSwordTactic) Id() consts.TacticId {
	return consts.BeAdeptWithBothPenAndSword
}

func (b BeAdeptWithBothPenAndSwordTactic) Name() string {
	return "文武双全"
}

func (b BeAdeptWithBothPenAndSwordTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (b BeAdeptWithBothPenAndSwordTactic) GetTriggerRate() float64 {
	return b.triggerRate
}

func (b BeAdeptWithBothPenAndSwordTactic) SetTriggerRate(rate float64) {
	b.triggerRate = rate
}

func (b BeAdeptWithBothPenAndSwordTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (b BeAdeptWithBothPenAndSwordTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (b BeAdeptWithBothPenAndSwordTactic) Execute() {
}

func (b BeAdeptWithBothPenAndSwordTactic) IsTriggerPrepare() bool {
	return false
}
