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

// 掣刀斫敌
// 使敌军单体受到兵刃伤害提高15%，然后对其造成兵刃伤害（伤害率208%）及震慑状态，持续1回合
// 主动，35%
type PullingSwordsAndChoppingEnemiesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (p PullingSwordsAndChoppingEnemiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 0.35
	return p
}

func (p PullingSwordsAndChoppingEnemiesTactic) Prepare() {

}

func (p PullingSwordsAndChoppingEnemiesTactic) Id() consts.TacticId {
	return consts.PullingSwordsAndChoppingEnemies
}

func (p PullingSwordsAndChoppingEnemiesTactic) Name() string {
	return "掣刀斫敌"
}

func (p PullingSwordsAndChoppingEnemiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (p PullingSwordsAndChoppingEnemiesTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PullingSwordsAndChoppingEnemiesTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PullingSwordsAndChoppingEnemiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (p PullingSwordsAndChoppingEnemiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PullingSwordsAndChoppingEnemiesTactic) Execute() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)

	// 使敌军单体受到兵刃伤害提高15%
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, p.tacticsParams)
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.15,
		EffectRound:    1,
		FromTactic:     p.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_SufferWeaponDamageImprove,
				TacticId:   p.Id(),
			})

			return revokeResp
		})
	}
	//然后对其造成兵刃伤害（伤害率208%）及震慑状态，持续1回合
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.08)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: p.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemyGeneral,
		DamageType:    consts.DamageType_Weapon,
		Damage:        dmg,
		TacticName:    p.Name(),
	})
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     p.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_Awe,
				TacticId:   p.Id(),
			})

			return revokeResp
		})
	}
}

func (p PullingSwordsAndChoppingEnemiesTactic) IsTriggerPrepare() bool {
	return false
}
