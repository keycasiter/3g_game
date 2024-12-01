package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 料事如神
// 对敌军群体（2人）造成谋略伤害（伤害率106%，受智力影响），并使其造成伤害降低16%，持续2回合
type ForetellLikeProphetTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (f ForetellLikeProphetTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	f.tacticsParams = tacticsParams
	f.triggerRate = 0.35
	return f
}

func (f ForetellLikeProphetTactic) Prepare() {
}

func (f ForetellLikeProphetTactic) Id() consts.TacticId {
	return consts.ForetellLikeProphet
}

func (f ForetellLikeProphetTactic) Name() string {
	return "料事如神"
}

func (f ForetellLikeProphetTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (f ForetellLikeProphetTactic) GetTriggerRate() float64 {
	return f.triggerRate
}

func (f ForetellLikeProphetTactic) SetTriggerRate(rate float64) {
	f.triggerRate = rate
}

func (f ForetellLikeProphetTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (f ForetellLikeProphetTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (f ForetellLikeProphetTactic) Execute() {
	ctx := f.tacticsParams.Ctx
	currentGeneral := f.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		f.Name(),
	)

	//对敌军群体（2人）造成谋略伤害（伤害率106%，受智力影响），并使其造成伤害降低16%，持续2回合
	//找到敌军群体2人
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, f.tacticsParams)
	for _, general := range enemyGenerals {
		//谋略伤害
		dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.06
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     f.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     general,
			DamageType:        consts.DamageType_Strategy,
			DamageImproveRate: dmgRate,
			TacticId:          f.Id(),
			TacticName:        f.Name(),
		})
		//施加效果
		if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.16,
			EffectRound:    2,
			FromTactic:     f.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//消失效果
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_LaunchWeaponDamageDeduce,
					TacticId:   f.Id(),
				})
				return revokeResp
			})
		}

		if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectRate:     0.16,
			EffectRound:    2,
			FromTactic:     f.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//消失效果
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_LaunchStrategyDamageDeduce,
					TacticId:   f.Id(),
				})
				return revokeResp
			})
		}
	}
}

func (f ForetellLikeProphetTactic) IsTriggerPrepare() bool {
	return false
}

func (a ForetellLikeProphetTactic) SetTriggerPrepare(triggerPrepare bool) {
}
