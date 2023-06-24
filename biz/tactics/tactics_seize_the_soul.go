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

// 夺魂挟魄
// 发动概率55%
// 偷取敌军单体38点武力、智力、速度、统率（受智力影响），
// 持续2回合，可叠加2次
type SeizeTheSoulTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SeizeTheSoulTactic) IsTriggerPrepare() bool {
	return false
}

func (s SeizeTheSoulTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SeizeTheSoulTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (s SeizeTheSoulTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SeizeTheSoulTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.55
	return s
}

func (s SeizeTheSoulTactic) Prepare() {
	return
}

func (s SeizeTheSoulTactic) Id() consts.TacticId {
	return consts.SeizeTheSoul
}

func (s SeizeTheSoulTactic) Name() string {
	return "夺魂挟魄"
}

func (s SeizeTheSoulTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s SeizeTheSoulTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s SeizeTheSoulTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, s.tacticsParams)

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	// 偷取敌军单体38点武力、智力、速度、统率（受智力影响），
	// 持续2回合，可叠加2次

	effectValue := cast.ToInt64(38 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100)
	//降低武力
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrForce, &vo.EffectHolderParams{
		EffectValue:    effectValue,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrForce,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}

	//降低智力
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrIntelligence, &vo.EffectHolderParams{
		EffectValue:    effectValue,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrIntelligence,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}
	//降低统率
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
		EffectValue:    effectValue,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrCommand,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}
	//降低速度
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrSpeed, &vo.EffectHolderParams{
		EffectValue:    effectValue,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrSpeed,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}

	//提升武力
	if util.BuffEffectWrapSet(ctx, enemyGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
		EffectValue:    effectValue,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrForce,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}

	//提升智力
	if util.BuffEffectWrapSet(ctx, enemyGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
		EffectValue:    effectValue,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrIntelligence,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}
	//提升统率
	if util.BuffEffectWrapSet(ctx, enemyGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
		EffectValue:    effectValue,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrCommand,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}
	//提升速度
	if util.BuffEffectWrapSet(ctx, enemyGeneral, consts.BuffEffectType_IncrSpeed, &vo.EffectHolderParams{
		EffectValue:    effectValue,
		EffectRound:    2,
		EffectTimes:    1,
		MaxEffectTimes: 2,
		FromTactic:     s.Id(),
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrSpeed,
				TacticId:   s.Id(),
			})

			return revokeResp
		})
	}
}
