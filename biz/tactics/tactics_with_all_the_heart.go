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

// 竭忠尽智
// 准备1回合，使敌军群体（1～2人）速度降低15%（受智力影响）
// 并进入混乱状态，持续2回合
// 并使友军单体获得1次抵御，持续1回合
// 主动，50%
type WithAllTheHeartTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (w WithAllTheHeartTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	w.tacticsParams = tacticsParams
	w.triggerRate = 0.5
	return w
}

func (w WithAllTheHeartTactic) Prepare() {
}

func (w WithAllTheHeartTactic) Id() consts.TacticId {
	return consts.WithAllTheHeart
}

func (w WithAllTheHeartTactic) Name() string {
	return "竭忠尽智"
}

func (w WithAllTheHeartTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (w WithAllTheHeartTactic) GetTriggerRate() float64 {
	return w.triggerRate
}

func (w WithAllTheHeartTactic) SetTriggerRate(rate float64) {
	w.triggerRate = rate
}

func (w WithAllTheHeartTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (w WithAllTheHeartTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (w WithAllTheHeartTactic) Execute() {
	ctx := w.tacticsParams.Ctx
	currentGeneral := w.tacticsParams.CurrentGeneral
	currentRound := w.tacticsParams.CurrentRound

	w.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		w.Name(),
	)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			w.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if w.isTriggered {
				return triggerResp
			} else {
				w.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				w.Name(),
			)

			// 准备1回合，使敌军群体（1～2人）速度降低15%（受智力影响）,并进入混乱状态，持续2回合
			enemyGenerals := util.GetEnemyGeneralsOneOrTwoArr(currentGeneral, w.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//速度降低
				effectDecrRate := 0.15 + triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
				effectVal := cast.ToInt64(enemyGeneral.BaseInfo.AbilityAttr.SpeedBase * effectDecrRate)
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrSpeed, &vo.EffectHolderParams{
					EffectValue:    effectVal,
					EffectRound:    2,
					FromTactic:     w.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_DecrSpeed,
							TacticId:   w.Id(),
						})

						return revokeResp
					})
				}
				//混乱
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Chaos, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     w.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Chaos,
							TacticId:   w.Id(),
						})

						return revokeResp
					})
				}
			}
			//并使友军单体获得1次抵御，持续1回合
			pairGeneral := util.GetPairOneGeneralNotSelf(w.tacticsParams, triggerGeneral)
			if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_Defend, &vo.EffectHolderParams{
				EffectRound:    1,
				EffectTimes:    1,
				FromTactic:     w.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_Defend,
						TacticId:   w.Id(),
					})

					return revokeResp
				})
			}
		}
		return triggerResp
	})
}

func (w WithAllTheHeartTactic) IsTriggerPrepare() bool {
	return false
}

func (a WithAllTheHeartTactic) SetTriggerPrepare(triggerPrepare bool) {
}
