package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 黄天太平
// 准备1回合，以自己进入混乱（攻击和战法无差别选择目标）状态为代价，使敌军群体（2人）进入计穷（无法发动主动战法）状态，持续2回合
type HuangTianTaiPingTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (a HuangTianTaiPingTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.4
	return a
}

func (a HuangTianTaiPingTactic) Prepare() {

}

func (a HuangTianTaiPingTactic) Id() consts.TacticId {
	return consts.HuangTiantaiPing
}

func (a HuangTianTaiPingTactic) Name() string {
	return "黄天太平"
}

func (a HuangTianTaiPingTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (a HuangTianTaiPingTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a HuangTianTaiPingTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a HuangTianTaiPingTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a HuangTianTaiPingTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a HuangTianTaiPingTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral
	currentRound := a.tacticsParams.CurrentRound

	// 准备1回合，以自己进入混乱（攻击和战法无差别选择目标）状态为代价，使敌军群体（2人）进入计穷（无法发动主动战法）状态，持续2回合
	a.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			a.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if a.isTriggered {
				return triggerResp
			} else {
				a.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				a.Name(),
			)

			//自己进入混乱
			if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
				EffectRound:    1,
				FromTactic:     a.Id(),
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				//混乱效果消失
				util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_Chaos,
						TacticId:   a.Id(),
					})
					return revokeResp
				})
			}
			//使敌军群体（2人）进入计穷（无法发动主动战法）状态，持续2回合
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, a.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_NoStrategy, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     a.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					//混乱效果消失
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_NoStrategy,
							TacticId:   a.Id(),
						})
						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (a HuangTianTaiPingTactic) IsTriggerPrepare() bool {
	return a.isTriggerPrepare
}
