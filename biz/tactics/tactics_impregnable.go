package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 固若金汤
// 使自己获得洞察（免疫所有控制效果）状态，并嘲讽敌军全体，同时提高自己150统率，持续2回合
// 主动，45%
type ImpregnableTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (i ImpregnableTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	i.tacticsParams = tacticsParams
	i.triggerRate = 0.45
	return i
}

func (i ImpregnableTactic) Prepare() {

}

func (i ImpregnableTactic) Id() consts.TacticId {
	return consts.Impregnable
}

func (i ImpregnableTactic) Name() string {
	return "固若金汤"
}

func (i ImpregnableTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (i ImpregnableTactic) GetTriggerRate() float64 {
	return i.triggerRate
}

func (i ImpregnableTactic) SetTriggerRate(rate float64) {
	i.triggerRate = rate
}

func (i ImpregnableTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (i ImpregnableTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (i ImpregnableTactic) Execute() {
	ctx := i.tacticsParams.Ctx
	currentGeneral := i.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		i.Name(),
	)
	//使自己获得洞察（免疫所有控制效果）状态，并嘲讽敌军全体，同时提高自己150统率，持续2回合
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_Insight, &vo.EffectHolderParams{
		EffectRound:    2,
		FromTactic:     i.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		//注册消失效果
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_Insight,
				TacticId:   i.Id(),
			})

			return revokeResp
		})
	}
	//嘲讽全体敌军
	enemyGenerals := util.GetEnemyGeneralsByGeneral(currentGeneral, i.tacticsParams)
	for _, enemyGeneral := range enemyGenerals {
		if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Taunt, &vo.EffectHolderParams{
			EffectRound:    2,
			FromTactic:     i.Id(),
			TauntByTarget:  currentGeneral,
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			//消失效果
			util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.DebuffEffectType_Taunt,
					TacticId:   i.Id(),
				})

				return revokeResp
			})
		}
	}
	//提高自身统率
	if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
		EffectValue:    150,
		EffectRound:    2,
		FromTactic:     i.Id(),
		ProduceGeneral: currentGeneral,
	}).IsSuccess {
		util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			revokeResp := &vo.TacticsTriggerResult{}
			revokeGeneral := params.CurrentGeneral

			util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
				Ctx:        ctx,
				General:    revokeGeneral,
				EffectType: consts.BuffEffectType_IncrCommand,
				TacticId:   i.Id(),
			})

			return revokeResp
		})
	}
}

func (i ImpregnableTactic) IsTriggerPrepare() bool {
	return false
}
