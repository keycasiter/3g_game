package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 虎痴
// 战斗中，每回合选择一名敌军单体，自身发动的所有攻击都会锁定该目标且对其造成伤害提高33%（受武力影响）
// 如果击败目标会使自身获得破阵状态，直到战斗结束
// 被动，100%
type TigerIdiotTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TigerIdiotTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t TigerIdiotTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	// 战斗中，每回合选择一名敌军单体，自身发动的所有攻击都会锁定该目标且对其造成伤害提高33%（受武力影响）,如果击败目标会使自身获得破阵状态，直到战斗结束
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerGeneral := params.CurrentGeneral
		triggerResp := &vo.TacticsTriggerResult{}

		enemyGeneral := util.GetEnemyOneGeneralByGeneral(currentGeneral, t.tacticsParams)
		effectRate := 0.33 + currentGeneral.BaseInfo.AbilityAttr.ForceBase/100/100
		if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_TigerIdiot_Locked, &vo.EffectHolderParams{
			EffectRate:     effectRate,
			EffectRound:    1,
			FromTactic:     t.Id(),
			ProduceGeneral: triggerGeneral,
			LockingTarget:  enemyGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_TigerIdiot_Locked,
					TacticId:   t.Id(),
				})

				return revokeResp
			})
		}

		return triggerResp
	})
}

func (t TigerIdiotTactic) Id() consts.TacticId {
	return consts.TigerIdiot
}

func (t TigerIdiotTactic) Name() string {
	return "虎痴"
}

func (t TigerIdiotTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t TigerIdiotTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TigerIdiotTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TigerIdiotTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t TigerIdiotTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TigerIdiotTactic) Execute() {

}

func (t TigerIdiotTactic) IsTriggerPrepare() bool {
	return false
}

func (a TigerIdiotTactic) SetTriggerPrepare(triggerPrepare bool) {
}
