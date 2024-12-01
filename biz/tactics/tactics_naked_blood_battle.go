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

// 裸衣血战
// 战斗中无法发动主动战法，战斗前3回合，获得90%连击及10%倒戈，并使自身及敌军单体统率降低40%
type NakedBloodBattleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (n NakedBloodBattleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	n.tacticsParams = tacticsParams
	n.triggerRate = 1.0
	return n
}

func (n NakedBloodBattleTactic) Prepare() {
	ctx := n.tacticsParams.Ctx
	currentGeneral := n.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		n.Name(),
	)
	//战斗中无法发动主动战法，战斗前3回合，获得90%连击及10%倒戈，并使自身及敌军单体统率降低40%
	util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_CanNotActiveTactic, &vo.EffectHolderParams{
		FromTactic:     n.Id(),
		ProduceGeneral: currentGeneral,
	})
	//连击
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ContinuousAttack, &vo.EffectHolderParams{
		EffectRate:     0.9,
		EffectRound:    3,
		FromTactic:     n.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_ContinuousAttack,
				TacticId:   n.Id(),
			})

			return revokeResp
		})
	}
	//倒戈
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Defection, &vo.EffectHolderParams{
		EffectRate:     0.1,
		EffectRound:    3,
		FromTactic:     n.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_Defection,
				TacticId:   n.Id(),
			})

			return revokeResp
		})
	}
	//自身统率降低
	if util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
		EffectValue:    cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.CommandBase * 0.4),
		EffectRound:    3,
		FromTactic:     n.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrCommand,
				TacticId:   n.Id(),
			})

			return revokeResp
		})
	}
	//敌军单体统率降低
	enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, n.tacticsParams)
	if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_DecrCommand, &vo.EffectHolderParams{
		EffectValue:    cast.ToInt64(enemyGeneral.BaseInfo.AbilityAttr.CommandBase * 0.4),
		EffectRound:    3,
		FromTactic:     n.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.DebuffEffectType_DecrCommand,
				TacticId:   n.Id(),
			})

			return revokeResp
		})
	}
}

func (n NakedBloodBattleTactic) Id() consts.TacticId {
	return consts.NakedBloodBattle
}

func (n NakedBloodBattleTactic) Name() string {
	return "裸衣血战"
}

func (n NakedBloodBattleTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (n NakedBloodBattleTactic) GetTriggerRate() float64 {
	return n.triggerRate
}

func (n NakedBloodBattleTactic) SetTriggerRate(rate float64) {
	n.triggerRate = rate
}

func (n NakedBloodBattleTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (n NakedBloodBattleTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (n NakedBloodBattleTactic) Execute() {

}

func (n NakedBloodBattleTactic) IsTriggerPrepare() bool {
	return false
}

func (a NakedBloodBattleTactic) SetTriggerPrepare(triggerPrepare bool) {
}
