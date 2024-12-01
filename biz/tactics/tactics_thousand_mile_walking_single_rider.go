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

// 千里走单骑
// 战斗中，自身准备发动自带准备战法时，有70%几率（受武力影响）获得洞察状态（免疫所有控制效果）并提高50武力，持续2回合，
// 在此期间，自身受到普通攻击时，对攻击者进行一次反击（伤害率238%），每回合最多触发1次
type ThousandMileWalkingSingleRiderTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThousandMileWalkingSingleRiderTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t ThousandMileWalkingSingleRiderTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral
	StrikeTriggerRoundHolder := map[consts.BattleRound]bool{}

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	// 战斗中，自身准备发动自带准备战法时，有70%几率（受武力影响）获得洞察状态（免疫所有控制效果）并提高50武力，持续2回合，
	// 在此期间，自身受到普通攻击时，对攻击者进行一次反击（伤害率238%），每回合最多触发1次
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTacticEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		tactic := params.CurrentTactic.(_interface.Tactics)

		if currentGeneral.EquipTactics[0].Id == tactic.Id() &&
			consts.ActivePrepareTacticsMap[tactic.Id()] {
			triggerRate := 0.7 + triggerGeneral.BaseInfo.AbilityAttr.ForceBase/100/100
			if util.GenerateRate(triggerRate) {
				//效果
				if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_Insight, &vo.EffectHolderParams{
					EffectRound:    2,
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_Insight,
							TacticId:   t.Id(),
						})

						return revokeResp
					})
				}
				//属性
				if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
					EffectRound:    2,
					EffectValue:    50,
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_IncrForce,
							TacticId:   t.Id(),
						})

						return revokeResp
					})
				}
				//施加预备效果
				if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_ThousandMileWalkingSingleRider_Prepare, &vo.EffectHolderParams{
					EffectRound:    2,
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					//消耗注册
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_ThousandMileWalkingSingleRider_Prepare,
							TacticId:   t.Id(),
						})

						return revokeResp
					})
					//在此期间，自身受到普通攻击时，对攻击者进行一次反击（伤害率238%），每回合最多触发1次
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						tResp := &vo.TacticsTriggerResult{}
						tGeneral := params.CurrentGeneral
						tRound := params.CurrentRound
						attackGeneral := params.AttackGeneral

						if util.BuffEffectContains(tGeneral, consts.BuffEffectType_ThousandMileWalkingSingleRider_Prepare) &&
							!StrikeTriggerRoundHolder[tRound] {

							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams:     t.tacticsParams,
								AttackGeneral:     tGeneral,
								SufferGeneral:     attackGeneral,
								DamageType:        consts.DamageType_Weapon,
								DamageImproveRate: 2.38,
								TacticId:          t.Id(),
								TacticName:        t.Name(),
								EffectName:        fmt.Sprintf("%v", consts.BuffEffectType_StrikeBack),
							})

							//每回合触发一次
							StrikeTriggerRoundHolder[tRound] = true
						}

						return tResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (t ThousandMileWalkingSingleRiderTactic) Id() consts.TacticId {
	return consts.ThousandMileWalkingSingleRider
}

func (t ThousandMileWalkingSingleRiderTactic) Name() string {
	return "千里走单骑"
}

func (t ThousandMileWalkingSingleRiderTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (t ThousandMileWalkingSingleRiderTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThousandMileWalkingSingleRiderTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThousandMileWalkingSingleRiderTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t ThousandMileWalkingSingleRiderTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThousandMileWalkingSingleRiderTactic) Execute() {

}

func (t ThousandMileWalkingSingleRiderTactic) IsTriggerPrepare() bool {
	return false
}

func (a ThousandMileWalkingSingleRiderTactic) SetTriggerPrepare(triggerPrepare bool) {
}
