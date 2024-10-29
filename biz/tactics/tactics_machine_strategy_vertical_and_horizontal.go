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

// 机略纵横
// 准备1回合，对敌军群体（2人）造成灼烧、中毒状态，每回合持续造成伤害（伤害率58%，受智力影响），持续2回合
// 主动 45%
type MachineStrategyVerticalAndHorizontalTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (m MachineStrategyVerticalAndHorizontalTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	m.tacticsParams = tacticsParams
	m.triggerRate = 0.45
	return m
}

func (m MachineStrategyVerticalAndHorizontalTactic) Prepare() {

}

func (m MachineStrategyVerticalAndHorizontalTactic) Id() consts.TacticId {
	return consts.MachineStrategyVerticalAndHorizontal
}

func (m MachineStrategyVerticalAndHorizontalTactic) Name() string {
	return "机略纵横"
}

func (m MachineStrategyVerticalAndHorizontalTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (m MachineStrategyVerticalAndHorizontalTactic) GetTriggerRate() float64 {
	return m.triggerRate
}

func (m MachineStrategyVerticalAndHorizontalTactic) SetTriggerRate(rate float64) {
	m.triggerRate = rate
}

func (m MachineStrategyVerticalAndHorizontalTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (m MachineStrategyVerticalAndHorizontalTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (m MachineStrategyVerticalAndHorizontalTactic) Execute() {
	ctx := m.tacticsParams.Ctx
	currentGeneral := m.tacticsParams.CurrentGeneral
	currentRound := m.tacticsParams.CurrentRound

	m.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		m.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			m.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if m.isTriggered {
				return triggerResp
			} else {
				m.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				m.Name(),
			)
			// 准备1回合，对敌军群体（2人）造成灼烧、中毒状态，每回合持续造成伤害（伤害率58%，受智力影响），持续2回合
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, m.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//灼烧
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Firing, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     m.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Firing,
							TacticId:   m.Id(),
						}) {
							dmgRate := triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 0.58
							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams:     m.tacticsParams,
								AttackGeneral:     triggerGeneral,
								SufferGeneral:     revokeGeneral,
								DamageType:        consts.DamageType_Strategy,
								DamageImproveRate: dmgRate,
								TacticId:          m.Id(),
								TacticName:        m.Name(),
								EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Firing),
							})
						}

						return revokeResp
					})
				}
				//中毒
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Methysis, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     m.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Methysis,
							TacticId:   m.Id(),
						}) {
							dmgRate := triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 0.58
							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams:     m.tacticsParams,
								AttackGeneral:     triggerGeneral,
								SufferGeneral:     revokeGeneral,
								DamageType:        consts.DamageType_Strategy,
								DamageImproveRate: dmgRate,
								TacticId:          m.Id(),
								TacticName:        m.Name(),
								EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Methysis),
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

func (m MachineStrategyVerticalAndHorizontalTactic) IsTriggerPrepare() bool {
	return false
}
