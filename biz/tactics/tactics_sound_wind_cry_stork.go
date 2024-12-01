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

// 风声鹤唳
// 准备1回合，对敌军群体（2人）造成谋略攻击（伤害率105%，受智力影响）并施加沙暴状态，每回合持续造成伤害（伤害率120%，受智力影响），持续1回合
type SoundOfTheWindAndTheCryOfTheStorkTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.45
	return s
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) Prepare() {
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) Id() consts.TacticId {
	return consts.SoundOfTheWindAndTheCryOfTheStork
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) Name() string {
	return "风声鹤唳"
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral
	currentRound := s.tacticsParams.CurrentRound

	s.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			s.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if s.isTriggered {
				return triggerResp
			} else {
				s.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				triggerGeneral.BaseInfo.Name,
				s.Name(),
			)

			//准备1回合，对敌军群体（2人）造成谋略攻击（伤害率105%，受智力影响）并施加沙暴状态，
			//每回合持续造成伤害（伤害率120%，受智力影响），持续1回合
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, s.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//攻击
				dmgRate := triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.05
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     s.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: dmgRate,
					TacticId:          s.Id(),
					TacticName:        s.Name(),
				})
				//施加状态
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Sandstorm, &vo.EffectHolderParams{
					EffectRound:    1,
					FromTactic:     s.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Sandstorm,
							TacticId:   s.Id(),
						}) {
							roundDmgRate := triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.2
							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams:     s.tacticsParams,
								AttackGeneral:     triggerGeneral,
								SufferGeneral:     revokeGeneral,
								DamageType:        consts.DamageType_Strategy,
								DamageImproveRate: roundDmgRate,
								TacticId:          s.Id(),
								TacticName:        s.Name(),
								EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Sandstorm),
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

func (s SoundOfTheWindAndTheCryOfTheStorkTactic) IsTriggerPrepare() bool {
	return s.isTriggerPrepare
}

func (a SoundOfTheWindAndTheCryOfTheStorkTactic) SetTriggerPrepare(triggerPrepare bool) {
}
