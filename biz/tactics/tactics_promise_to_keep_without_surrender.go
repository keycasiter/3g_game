package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 誓守无降
// 准备1回合，使我军群体（2人）进入洞察状态，持续2回合，并使自身2回合内受到下一次谋略伤害时，计穷敌军主将，持续2回合
// 主动，35%
type PromiseToKeepWithoutSurrenderTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (p PromiseToKeepWithoutSurrenderTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	p.tacticsParams = tacticsParams
	p.triggerRate = 0.35
	return p
}

func (p PromiseToKeepWithoutSurrenderTactic) Prepare() {

}

func (p PromiseToKeepWithoutSurrenderTactic) Id() consts.TacticId {
	return consts.PromiseToKeepWithoutSurrender
}

func (p PromiseToKeepWithoutSurrenderTactic) Name() string {
	return "誓守无降"
}

func (p PromiseToKeepWithoutSurrenderTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (p PromiseToKeepWithoutSurrenderTactic) GetTriggerRate() float64 {
	return p.triggerRate
}

func (p PromiseToKeepWithoutSurrenderTactic) SetTriggerRate(rate float64) {
	p.triggerRate = rate
}

func (p PromiseToKeepWithoutSurrenderTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (p PromiseToKeepWithoutSurrenderTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (p PromiseToKeepWithoutSurrenderTactic) Execute() {
	ctx := p.tacticsParams.Ctx
	currentGeneral := p.tacticsParams.CurrentGeneral
	currentRound := p.tacticsParams.CurrentRound

	p.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		p.Name(),
	)
	// 准备1回合，使我军群体（2人）进入洞察状态，持续2回合，并使自身2回合内受到下一次谋略伤害时，计穷敌军主将，持续2回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			p.isTriggerPrepare = false
		}
		if currentRound+1 == triggerRound {
			if p.isTriggered {
				return triggerResp
			} else {
				p.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				triggerGeneral.BaseInfo.Name,
				p.Name(),
			)

			pairGenerals := util.GetPairGeneralsTwoArr(currentGeneral, p.tacticsParams)
			for _, pairGeneral := range pairGenerals {
				if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Insight, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     p.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_Insight,
							TacticId:   p.Id(),
						})

						return revokeResp
					})
				}
			}
			//并使自身2回合内受到下一次谋略伤害时，计穷敌军主将，持续2回合
			util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferStrategyDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				round := params.CurrentRound
				resp := &vo.TacticsTriggerResult{}

				if round <= triggerRound+1 {
					enemyMasterGeneral := util.GetEnemyMasterGeneral(currentGeneral, p.tacticsParams)
					if util.DebuffEffectWrapSet(ctx, enemyMasterGeneral, consts.DebuffEffectType_NoStrategy, &vo.EffectHolderParams{
						EffectRound:    2,
						FromTactic:     p.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(enemyMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_NoStrategy,
								TacticId:   p.Id(),
							})

							return revokeResp
						})
					}
				}

				return resp
			})
		}

		return triggerResp
	})
}

func (p PromiseToKeepWithoutSurrenderTactic) IsTriggerPrepare() bool {
	return p.isTriggerPrepare
}
