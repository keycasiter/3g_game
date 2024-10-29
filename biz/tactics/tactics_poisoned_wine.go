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

// 鸩毒
// 对敌军单体施加鸩酒，降低其60%武力（受智力影响），持续1回合，1回合后毒发造成谋略攻击（伤害率226%，受智力影响）
// 同时降低其60点统率（受智力影响），可无限叠加，直到战斗结束
// 主动 70%
type PoisonedWineTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PoisonedWineTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 0.7
	return p
}

func (p PoisonedWineTactic) Prepare() {
}

func (p PoisonedWineTactic) Id() consts.TacticId {
	return consts.PoisonedWine
}

func (p PoisonedWineTactic) Name() string {
	return "鸩毒"
}

func (p PoisonedWineTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (p PoisonedWineTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PoisonedWineTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PoisonedWineTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (p PoisonedWineTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PoisonedWineTactic) Execute() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)
	// 对敌军单体施加鸩酒，降低其60%武力（受智力影响），持续1回合，1回合后毒发造成谋略攻击（伤害率226%，受智力影响）
	// 同时降低其60点统率（受智力影响），可无限叠加，直到战斗结束
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, p.tacticsParams)
	decrVal := cast.ToInt64(enemyGeneral.BaseInfo.AbilityAttr.ForceBase * (0.6 + enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100))
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrForce, &vo.EffectHolderParams{
		EffectValue:    decrVal,
		EffectRound:    1,
		FromTactic:     p.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrForce,
				TacticId:   p.Id(),
			}) {
				dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 2.26
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     p.tacticsParams,
					AttackGeneral:     currentGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: dmgRate,
					TacticName:        p.Name(),
				})
				//降低统率
				decrCommandVal := cast.ToInt64(60 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100)
				util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
					EffectValue:    decrCommandVal,
					FromTactic:     p.Id(),
					ProduceGeneral: currentGeneral,
				})
			}

			return revokeResp
		})
	}
}

func (p PoisonedWineTactic) IsTriggerPrepare() bool {
	return false
}
