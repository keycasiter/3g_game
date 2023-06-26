package tactics

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 云聚影从
// 战斗中，自身即将受到普通攻击时，有50%概率（受武力影响）使武力最高的友军单体获得反击效果（伤害率100%）和急救状态，
// 受到伤害时有30%概率恢复自身兵力（治疗率100%，受武力影响），持续1回合，随后使该友军为自己承担此次普通攻击
// 指挥，100%
type CloudGatheringShadowFromTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CloudGatheringShadowFromTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c CloudGatheringShadowFromTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)

	// 战斗中，自身即将受到普通攻击时，有50%概率（受武力影响）使武力最高的友军单体获得反击效果（伤害率100%）和急救状态受到伤害时有30%概率恢复自身兵力（治疗率100%，受武力影响），持续1回合，随后使该友军为自己承担此次普通攻击
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferGeneralAttack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral

		triggerRate := 0.5
		triggerRate += triggerGeneral.BaseInfo.AbilityAttr.ForceBase / 100 / 100
		if util.GenerateRate(triggerRate) {
			highPairGeneral := util.GetPairGeneralWhoIsHighestForce(c.tacticsParams)
			//反击效果，使武力最高的友军单体获得反击效果（伤害率100%）
			if util.BuffEffectWrapSet(ctx, highPairGeneral, consts.BuffEffectType_StrikeBack, &vo.EffectHolderParams{
				EffectRate:     1.0,
				EffectRound:    1,
				FromTactic:     c.Id(),
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(highPairGeneral, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral
					attackGeneral := params.AttackGeneral

					if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_StrikeBack,
						TacticId:   c.Id(),
					}) {
						dmg := cast.ToInt64(revokeGeneral.BaseInfo.AbilityAttr.ForceBase * 1.0)
						util.TacticDamage(&util.TacticDamageParam{
							TacticsParams: c.tacticsParams,
							AttackGeneral: revokeGeneral,
							SufferGeneral: attackGeneral,
							DamageType:    consts.DamageType_Weapon,
							Damage:        dmg,
							TacticId:      c.Id(),
							TacticName:    c.Name(),
							EffectName:    fmt.Sprintf("%v", consts.BuffEffectType_StrikeBack),
						})
					}

					return revokeResp
				})
			}
			//急救状态，受到伤害时有30%概率恢复自身兵力（治疗率100%，受武力影响），持续1回合
			if util.BuffEffectWrapSet(ctx, highPairGeneral, consts.BuffEffectType_EmergencyTreatment, &vo.EffectHolderParams{
				EffectRate:     1.0,
				EffectRound:    1,
				FromTactic:     c.Id(),
				ProduceGeneral: triggerGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(highPairGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_EmergencyTreatment,
						TacticId:   c.Id(),
					}) {
						if util.GenerateRate(0.3) {
							resumeNum := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 1.0)
							util.ResumeSoldierNum(&util.ResumeParams{
								Ctx:            ctx,
								TacticsParams:  c.tacticsParams,
								ProduceGeneral: triggerGeneral,
								SufferGeneral:  revokeGeneral,
								ResumeNum:      resumeNum,
							})
						}
					}

					return revokeResp
				})
			}
			//随后使该友军为自己承担此次普通攻击
			params.CurrentGeneral = highPairGeneral
		}

		return triggerResp
	})
}

func (c CloudGatheringShadowFromTactic) Id() consts.TacticId {
	return consts.CloudGatheringShadowFrom
}

func (c CloudGatheringShadowFromTactic) Name() string {
	return "云聚影从"
}

func (c CloudGatheringShadowFromTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c CloudGatheringShadowFromTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CloudGatheringShadowFromTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CloudGatheringShadowFromTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (c CloudGatheringShadowFromTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CloudGatheringShadowFromTactic) Execute() {

}

func (c CloudGatheringShadowFromTactic) IsTriggerPrepare() bool {
	return false
}
