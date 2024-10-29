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
)

// 镇压黄巾
// 战斗第2回合，使我军全体免疫沙暴状态，并且使敌军全体陷入叛逃状态，每回合持续造成伤害（伤害率88%，受武力或智力最高一项影响，无视防御），持续2回合
// 若目标为黄巾军，伤害率提高20%
// 指挥，100%
type SuppressYellowScarvesTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SuppressYellowScarvesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.0
	return s
}

func (s SuppressYellowScarvesTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	//战斗第2回合，使我军全体免疫沙暴状态，并且使敌军全体陷入叛逃状态，每回合持续造成伤害（伤害率88%，受武力或智力最高一项影响，无视防御），持续2回合
	// 若目标为黄巾军，伤害率提高20%
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		if triggerRound == consts.Battle_Round_Second {
			//我军
			pairGenerals := util.GetPairGeneralArr(s.tacticsParams)
			for _, pairGeneral := range pairGenerals {
				if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_ImmunitySandstorm, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     s.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_ImmunitySandstorm,
							TacticId:   s.Id(),
						})

						return revokeResp
					})
				}
			}
			//敌军
			allEnemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, s.tacticsParams)
			for _, enemyGeneral := range allEnemyGenerals {
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Defect, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     s.Id(),
					ProduceGeneral: triggerGeneral,
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
							// 每回合持续造成伤害（伤害率88%，受武力或智力最高一项影响，无视防御），持续2回合
							// 若目标为黄巾军，伤害率提高20%
							attType, _ := util.GetGeneralHighestBetweenForceOrIntelligence(revokeGeneral)
							dmgType := consts.DamageType_Weapon
							if attType == consts.AbilityAttr_Intelligence {
								dmgType = consts.DamageType_Strategy
							}

							dmgRate := 0.88
							if util.IsContainsGeneralTag(revokeGeneral.BaseInfo.GeneralTag, consts.GeneralTag_YellowTurbans) {
								dmgRate = 0.88 * (1 + 0.2)
							}
							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams:     s.tacticsParams,
								AttackGeneral:     currentGeneral,
								SufferGeneral:     revokeGeneral,
								DamageType:        dmgType,
								DamageImproveRate: dmgRate,
								TacticId:          s.Id(),
								TacticName:        s.Name(),
								EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Defect),
								IsIgnoreDefend:    true,
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

func (s SuppressYellowScarvesTactic) Id() consts.TacticId {
	return consts.SuppressYellowScarves
}

func (s SuppressYellowScarvesTactic) Name() string {
	return "镇压黄巾"
}

func (s SuppressYellowScarvesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s SuppressYellowScarvesTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SuppressYellowScarvesTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SuppressYellowScarvesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s SuppressYellowScarvesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SuppressYellowScarvesTactic) Execute() {
}

func (s SuppressYellowScarvesTactic) IsTriggerPrepare() bool {
	return false
}
