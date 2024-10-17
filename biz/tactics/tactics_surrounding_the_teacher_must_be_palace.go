package tactics

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 围师必阙
// 战斗中，当敌军群体（2人）处于溃逃或叛逃状态时压制敌军，使敌军造成的谋略伤害降低39%（受统率影响，触发后每回合递减1/3），
// 战斗第2回合，使敌军全体进入叛逃状态（伤害率120%），持续2回合
// 自身为主将时，谋略伤害降低至提升至45%
// 指挥，100%
type SurroundingTheTeacherMustBePalaceTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SurroundingTheTeacherMustBePalaceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s SurroundingTheTeacherMustBePalaceTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	// 战斗中，当敌军群体（2人）处于溃逃或叛逃状态时压制敌军，使敌军造成的谋略伤害降低39%（受统率影响，触发后每回合递减1/3），
	enemyGenerals := util.GetEnemyTwoGeneralByGeneral(currentGeneral, s.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_SufferDebuffEffectEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.SufferDebuffEffectGeneral

			if params.DebuffEffect == consts.DebuffEffectType_Defect ||
				params.DebuffEffect == consts.DebuffEffectType_Escape {

				effectRate := 0.39 + currentGeneral.BaseInfo.AbilityAttr.CommandBase/100/100
				// 自身为主将时，谋略伤害降低至提升至45%
				if currentGeneral.IsMaster {
					effectRate = 0.45 + currentGeneral.BaseInfo.AbilityAttr.CommandBase/100/100
				}
				if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
					EffectRate:     effectRate,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRate(&util.DebuffEffectOfTacticCostRateParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_LaunchStrategyDamageDeduce,
							TacticId:   s.Id(),
							EffectRate: effectRate / 3,
						})

						return revokeResp
					})
				}
			}

			return triggerResp
		})
	}
	// 战斗第2回合，使敌军全体进入叛逃状态（伤害率120%），持续2回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound

		if triggerRound == consts.Battle_Round_Second {
			allEnemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, s.tacticsParams)
			for _, enemyGeneral := range allEnemyGenerals {
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Defect, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Defect,
							TacticId:   s.Id(),
						}) {
							attType, val := util.GetGeneralHighestBetweenForceOrIntelligence(currentGeneral)
							dmg := cast.ToInt64(val * 1.2)
							dmgType := consts.DamageType_Weapon
							if attType == consts.AbilityAttr_Intelligence {
								dmgType = consts.DamageType_Strategy
							}
							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams:  s.tacticsParams,
								AttackGeneral:  currentGeneral,
								SufferGeneral:  revokeGeneral,
								DamageType:     dmgType,
								Damage:         dmg,
								TacticId:       s.Id(),
								TacticName:     s.Name(),
								EffectName:     fmt.Sprintf("%v", consts.DebuffEffectType_Defect),
								IsIgnoreDefend: true,
							})
						}

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (s SurroundingTheTeacherMustBePalaceTactic) Id() consts.TacticId {
	return consts.SurroundingTheTeacherMustBePalace
}

func (s SurroundingTheTeacherMustBePalaceTactic) Name() string {
	return "围师必阙"
}

func (s SurroundingTheTeacherMustBePalaceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s SurroundingTheTeacherMustBePalaceTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SurroundingTheTeacherMustBePalaceTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SurroundingTheTeacherMustBePalaceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s SurroundingTheTeacherMustBePalaceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SurroundingTheTeacherMustBePalaceTactic) Execute() {
}

func (s SurroundingTheTeacherMustBePalaceTactic) IsTriggerPrepare() bool {
	return false
}
