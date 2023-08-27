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

// 焰逐风飞
// 对敌军单体熬成谋略攻击(伤害率226%，受智力影响)及震慑状态（无法行动）并有40%概率使其受到谋略伤害提高12%（受智力影响）， 持续1回合
type FlamesFlyingInTheWindTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f FlamesFlyingInTheWindTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.35
	return f
}

func (f FlamesFlyingInTheWindTactic) Prepare() {

}

func (f FlamesFlyingInTheWindTactic) Id() consts.TacticId {
	return consts.FlamesFlyingInTheWind
}

func (f FlamesFlyingInTheWindTactic) Name() string {
	return "焰逐风飞"
}

func (f FlamesFlyingInTheWindTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f FlamesFlyingInTheWindTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f FlamesFlyingInTheWindTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f FlamesFlyingInTheWindTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f FlamesFlyingInTheWindTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f FlamesFlyingInTheWindTactic) Execute() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)
	//对敌军单体熬成谋略攻击(伤害率226%，受智力影响)及震慑状态（无法行动）并有40%概率使其受到谋略伤害提高12%（受智力影响）， 持续1回合
	//找到敌军单体
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, f.tacticsParams)
	//谋略伤害
	dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 2.26)
	util.TacticDamage(&util.TacticDamageParam{
		TacticsParams: f.tacticsParams,
		AttackGeneral: currentGeneral,
		SufferGeneral: enemyGeneral,
		DamageType:    consts.DamageType_Strategy,
		Damage:        dmg,
		TacticId:      f.Id(),
		TacticName:    f.Name(),
	})
	//震慑效果
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
		EffectRound:    1,
		FromTactic:     f.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		//消失效果
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_Awe,
				TacticId:   f.Id(),
			})

			return revokeResp
		})
	}
	//受到谋略伤害提高
	if util.GenerateRate(0.4) {
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
			EffectRound:    1,
			FromTactic:     f.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//注册消失效果
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_SufferStrategyDamageImprove,
					TacticId:   f.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (f FlamesFlyingInTheWindTactic) IsTriggerPrepare() bool {
	return false
}
