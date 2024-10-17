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

// 落雷
// 对随机其他单体（有5%几率对友军释放）造成谋略攻击（伤害率170%，受智力影响），并使其受到谋略伤害增加18%，持续2回合
type ThunderboltTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThunderboltTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.5
	return t
}

func (t ThunderboltTactic) Prepare() {

}

func (t ThunderboltTactic) Id() consts.TacticId {
	return consts.Thunderbolt
}

func (t ThunderboltTactic) Name() string {
	return "落雷"
}

func (t ThunderboltTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t ThunderboltTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThunderboltTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThunderboltTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t ThunderboltTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThunderboltTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	//对随机其他单体（有5%几率对友军释放）造成谋略攻击（伤害率170%，受智力影响），并使其受到谋略伤害增加18%，持续2回合
	targetGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, t.tacticsParams)
	if util.GenerateRate(0.05) {
		targetGeneral = util.GetPairOneGeneralNotSelf(t.tacticsParams, currentGeneral)
	}
	//伤害
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.7)
	damage.TacticDamage(&damage.TacticDamageParam{
		TacticsParams: t.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: targetGeneral,
		DamageType:    consts.DamageType_Strategy,
		Damage:        dmg,
		TacticId:      t.Id(),
		TacticName:    t.Name(),
	})
	//效果施加
	if util.DebuffEffectWrapSet(ctx, targetGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.18,
		EffectRound:    2,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(targetGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_SufferStrategyDamageImprove,
				TacticId:   t.Id(),
			})

			return revokeResp
		})
	}
}

func (t ThunderboltTactic) IsTriggerPrepare() bool {
	return false
}
