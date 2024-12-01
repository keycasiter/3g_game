package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 知己知彼
// 战斗首回合，使敌我全体造成伤害降低35%（受智力影响），
// 第二回合起，每回合有35%概率（受魅力影响）使友军全体造成伤害提升6%（受最高属性影响），可叠加，持续到战斗结束
type KnowYourselfAsWellAsTheEnemyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a KnowYourselfAsWellAsTheEnemyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.00
	return a
}

func (a KnowYourselfAsWellAsTheEnemyTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := a.tacticsParams.CurrentGeneral

		// 战斗首回合，使敌我全体造成伤害降低35%（受智力影响）
		if params.CurrentRound == consts.Battle_Round_First {
			allGenerals := util.GetAllGenerals(a.tacticsParams)
			for _, general := range allGenerals {
				//兵刃
				if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
					EffectRate:  0.35 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100,
					EffectRound: 1,
					FromTactic:  a.Id(),
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral
						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_LaunchWeaponDamageDeduce,
							TacticId:   a.Id(),
						})
						return revokeResp
					})
				}
				//谋略
				if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
					EffectRate:  0.35 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100,
					EffectRound: 1,
					FromTactic:  a.Id(),
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral
						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_LaunchStrategyDamageDeduce,
							TacticId:   a.Id(),
						})
						return revokeResp
					})
				}
			}
		}

		// 第二回合起，每回合有35%概率（受魅力影响）使友军全体造成伤害提升6%（受最高属性影响），可叠加，持续到战斗结束
		if params.CurrentRound >= consts.Battle_Round_Second {
			triggerRate := 0.35 + triggerGeneral.BaseInfo.AbilityAttr.CharmBase/100/100
			if !util.GenerateRate(triggerRate) {
				return triggerResp
			}

			pairGenerals := util.GetPairGeneralArr(triggerGeneral, a.tacticsParams)
			_, val := util.GetGeneralHighestAttr(triggerGeneral)
			rate := 0.06 + val/100/100
			for _, general := range pairGenerals {
				util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate:     rate,
					FromTactic:     a.Id(),
					ProduceGeneral: triggerGeneral,
				})
			}
		}

		return triggerResp
	})
}

func (a KnowYourselfAsWellAsTheEnemyTactic) Id() consts.TacticId {
	return consts.KnowYourselfAsWellAsTheEnemy
}

func (a KnowYourselfAsWellAsTheEnemyTactic) Name() string {
	return "知己知彼"
}

func (a KnowYourselfAsWellAsTheEnemyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (a KnowYourselfAsWellAsTheEnemyTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a KnowYourselfAsWellAsTheEnemyTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a KnowYourselfAsWellAsTheEnemyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a KnowYourselfAsWellAsTheEnemyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a KnowYourselfAsWellAsTheEnemyTactic) Execute() {
}

func (a KnowYourselfAsWellAsTheEnemyTactic) IsTriggerPrepare() bool {
	return false
}

func (a KnowYourselfAsWellAsTheEnemyTactic) SetTriggerPrepare(triggerPrepare bool) {
}
