package tactics

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 死战不退
// 战斗中，使自己免疫混乱；自身受到伤害时，有80%概率获得一层蓄威效果，可积攒20层；
// 普攻后，有50%概率（受武力影响）消耗一层蓄威对敌军单体造成一次兵刃伤害（伤害率130%）
// 触发后可继续判定，每次触发后几率降低8%，每次普攻后最多触发5次
// 被动 100%
type NeverRetreatFromDeadBattleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (n NeverRetreatFromDeadBattleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	n.tacticsParams = tacticsParams
	n.triggerRate = 1.0
	return n
}

func (n NeverRetreatFromDeadBattleTactic) Prepare() {
	ctx := n.tacticsParams.Ctx
	currentGeneral := n.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		n.Name(),
	)
	// 战斗中，使自己免疫混乱；
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_ImmunityChaos, &vo.EffectHolderParams{
		FromTactic:     n.Id(),
		ProduceGeneral: currentGeneral,
	})
	//自身受到伤害时，有80%概率获得一层蓄威效果，可积攒20层；
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.SufferAttackGeneral

		if util.GenerateRate(0.8) {
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_AccumulatePower, &vo.EffectHolderParams{
				EffectTimes:    1,
				MaxEffectTimes: 20,
				FromTactic:     n.Id(),
				ProduceGeneral: triggerGeneral,
			})
		}

		return triggerResp
	})
	// 普攻后，有50%概率（受武力影响）消耗一层蓄威对敌军单体造成一次兵刃伤害（伤害率130%）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		sufferGeneral := params.SufferAttackGeneral
		triggerRate := 0.5 + triggerGeneral.BaseInfo.AbilityAttr.ForceBase/100/100
		triggerMaxCnt := 5

		for {
			if util.GenerateRate(triggerRate) {
				n.costAccumulatePowerHandler(ctx, triggerGeneral, sufferGeneral)
				// 触发后可继续判定，每次触发后几率降低8%，每次普攻后最多触发5次
				triggerRate -= 0.08
				triggerMaxCnt--
			} else {
				break
			}
			if triggerMaxCnt == 0 {
				break
			}
		}

		return triggerResp
	})
}

func (n NeverRetreatFromDeadBattleTactic) costAccumulatePowerHandler(ctx context.Context, currentGeneral *vo.BattleGeneral, sufferAttackGeneral *vo.BattleGeneral) {
	if util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
		Ctx:        ctx,
		General:    currentGeneral,
		EffectType: consts.BuffEffectType_AccumulatePower,
		TacticId:   n.Id(),
		CostTimes:  1,
	}) {
		enemyGeneral := util.GetEnemyOneGeneralByGeneral(sufferAttackGeneral, n.tacticsParams)
		damage.TacticDamage(&damage.TacticDamageParam{
			TacticsParams:     n.tacticsParams,
			AttackGeneral:     currentGeneral,
			SufferGeneral:     enemyGeneral,
			DamageType:        consts.DamageType_Weapon,
			DamageImproveRate: 1.3,
			TacticId:          n.Id(),
			TacticName:        n.Name(),
			EffectName:        fmt.Sprintf("%v", consts.BuffEffectType_AccumulatePower),
		})
	}
}

func (n NeverRetreatFromDeadBattleTactic) Id() consts.TacticId {
	return consts.NeverRetreatFromDeadBattle
}

func (n NeverRetreatFromDeadBattleTactic) Name() string {
	return "死战不退"
}

func (n NeverRetreatFromDeadBattleTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (n NeverRetreatFromDeadBattleTactic) GetTriggerRate() float64 {
	return n.triggerRate
}

func (n NeverRetreatFromDeadBattleTactic) SetTriggerRate(rate float64) {
	n.triggerRate = rate
}

func (n NeverRetreatFromDeadBattleTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (n NeverRetreatFromDeadBattleTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (n NeverRetreatFromDeadBattleTactic) Execute() {
}

func (n NeverRetreatFromDeadBattleTactic) IsTriggerPrepare() bool {
	return false
}

func (a NeverRetreatFromDeadBattleTactic) SetTriggerPrepare(triggerPrepare bool) {
}
