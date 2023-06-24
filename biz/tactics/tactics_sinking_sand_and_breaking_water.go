package tactics

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 沉沙决水
// 准备1回合，对敌军群体(2人)施加水攻状态，每回合持续造成伤害(伤害率126%,受智力影响)，并使其受到的谋略伤害提升25% ,持续2回合
type SinkingSandAndBreakingWaterTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (s SinkingSandAndBreakingWaterTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.4
	return s
}

func (s SinkingSandAndBreakingWaterTactic) Prepare() {
}

func (s SinkingSandAndBreakingWaterTactic) Id() consts.TacticId {
	return consts.SinkingSandAndBreakingWater
}

func (s SinkingSandAndBreakingWaterTactic) Name() string {
	return "沉沙决水"
}

func (s SinkingSandAndBreakingWaterTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SinkingSandAndBreakingWaterTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SinkingSandAndBreakingWaterTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SinkingSandAndBreakingWaterTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SinkingSandAndBreakingWaterTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SinkingSandAndBreakingWaterTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral
	currentRound := s.tacticsParams.CurrentRound

	s.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
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
				currentGeneral.BaseInfo.Name,
				s.Name(),
			)
			// 准备1回合，对敌军群体(2人)施加水攻状态，每回合持续造成伤害(伤害率126%,受智力影响)，并使其受到的谋略伤害提升25% ,持续2回合
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, s.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//对敌军群体(2人)施加水攻状态，每回合持续造成伤害(伤害率126%,受智力影响)
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_WaterAttack, &vo.EffectHolderParams{
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
							EffectType: consts.DebuffEffectType_WaterAttack,
							TacticId:   s.Id(),
						}) {
							dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.26)
							util.TacticDamage(&util.TacticDamageParam{
								TacticsParams: s.tacticsParams,
								AttackGeneral: currentGeneral,
								SufferGeneral: enemyGeneral,
								DamageType:    consts.DamageType_Strategy,
								Damage:        dmg,
								TacticId:      s.Id(),
								TacticName:    s.Name(),
								EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_WaterAttack),
							})
						}

						return revokeResp
					})
				}
				//并使其受到的谋略伤害提升25%
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferStrategyDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.25,
					EffectRound:    2,
					FromTactic:     s.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_SufferStrategyDamageImprove,
							TacticId:   s.Id(),
						})

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (s SinkingSandAndBreakingWaterTactic) IsTriggerPrepare() bool {
	return false
}
