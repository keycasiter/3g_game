package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 顾盼生姿
// 偷取敌军一名男性武将35点智力和武力给自身和友军单体（受魅力影响）持续2回合，可叠加2次
// 主动 45%
type LookAroundCharminglyTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LookAroundCharminglyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 0.45
	return l
}

func (l LookAroundCharminglyTactic) Prepare() {
}

func (l LookAroundCharminglyTactic) Id() consts.TacticId {
	return consts.LookAroundCharmingly
}

func (l LookAroundCharminglyTactic) Name() string {
	return "顾盼生姿"
}

func (l LookAroundCharminglyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (l LookAroundCharminglyTactic) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LookAroundCharminglyTactic) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LookAroundCharminglyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (l LookAroundCharminglyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LookAroundCharminglyTactic) Execute() {
	ctx := l.tacticsParams.Ctx
	currentGeneral := l.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		l.Name(),
	)
	//偷取敌军一名男性武将35点智力和武力给自身和友军单体（受魅力影响）持续2回合，可叠加2次
	enemyGeneral := util.GetEnemyOneMaleGeneral(currentGeneral, l.tacticsParams)

	//偷取武力
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrForce, &vo.EffectHolderParams{
		EffectValue:    35,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     l.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrForce,
				TacticId:   l.Id(),
			})

			return revokeResp
		})
	}
	//偷取智力
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
		EffectValue:    35,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     l.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrIntelligence,
				TacticId:   l.Id(),
			})

			return revokeResp
		})
	}

	//找到我军单体和自己
	allGenerals := make([]*vo.BattleGeneral, 0)
	pairGeneral := util.GetPairOneGeneralNotSelf(l.tacticsParams, currentGeneral)
	allGenerals = append(allGenerals, pairGeneral)
	allGenerals = append(allGenerals, currentGeneral)
	for _, general := range allGenerals {
		//提高武力
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
			EffectValue:    35,
			EffectRound:    2,
			EffectTimes:    1,
			MaxEffectTimes: 2,
			FromTactic:     l.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrForce,
					TacticId:   l.Id(),
				})

				return revokeResp
			})
		}
		//提高智力
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
			EffectValue:    35,
			EffectRound:    2,
			EffectTimes:    1,
			MaxEffectTimes: 2,
			FromTactic:     l.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_IncrIntelligence,
					TacticId:   l.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (l LookAroundCharminglyTactic) IsTriggerPrepare() bool {
	return false
}

func (a LookAroundCharminglyTactic) SetTriggerPrepare(triggerPrepare bool) {
}
