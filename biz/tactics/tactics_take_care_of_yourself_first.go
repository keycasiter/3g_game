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

// 先成其虑
// 普通攻击之后，对攻击目标再次造成一次谋略攻击（伤害率145%，受智力影响）并使自身主动战法的发动几率提高15%，持续1回合
type TakeCareOfYourselfFirstTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TakeCareOfYourselfFirstTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.45
	return t
}

func (t TakeCareOfYourselfFirstTactic) Prepare() {
}

func (t TakeCareOfYourselfFirstTactic) Id() consts.TacticId {
	return consts.TakeCareOfYourselfFirst
}

func (t TakeCareOfYourselfFirstTactic) Name() string {
	return "先成其虑"
}

func (t TakeCareOfYourselfFirstTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TakeCareOfYourselfFirstTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TakeCareOfYourselfFirstTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TakeCareOfYourselfFirstTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Assault
}

func (t TakeCareOfYourselfFirstTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TakeCareOfYourselfFirstTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//普通攻击之后，对攻击目标再次造成一次谋略攻击（伤害率145%，受智力影响）并使自身主动战法的发动几率提高15%，持续1回合
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.45)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: t.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: t.tacticsParams.CurrentSufferGeneral,
		DamageType:    consts.DamageType_Strategy,
		Damage:        dmg,
		TacticId:      t.Id(),
		TacticName:    t.Name(),
	})

	//使自身主动战法的发动几率提高15%，持续1回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_TacticsActiveTriggerWithSelfImprove, &vo.EffectHolderParams{
		EffectRound:    1,
		TriggerRate:    0.15,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_TacticsActiveTriggerWithSelfImprove,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
}

func (t TakeCareOfYourselfFirstTactic) IsTriggerPrepare() bool {
	return false
}
