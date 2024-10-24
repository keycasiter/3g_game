package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 挟势弄权
// 对随机敌军单体造成谋略攻击（伤害率186%，受智力影响），并混乱（攻击和战法无差别选择目标）1回合
type TakingAdvantageOfTheSituationToGainPowerTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.35
	return t
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) Prepare() {
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) Id() consts.TacticId {
	return consts.TakingAdvantageOfTheSituationToGainPower
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) Name() string {
	return "挟势弄权"
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	// 对随机敌军单体造成谋略攻击（伤害率186%，受智力影响），
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, t.tacticsParams)
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.86)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams: t.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemyGeneral,
		DamageType:    consts.DamageType_Strategy,
		Damage:        dmg,
		TacticId:      t.Id(),
		TacticName:    t.Name(),
	})

	//并混乱（攻击和战法无差别选择目标）1回合
	if util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_Chaos,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
}

func (t TakingAdvantageOfTheSituationToGainPowerTactic) IsTriggerPrepare() bool {
	return false
}
