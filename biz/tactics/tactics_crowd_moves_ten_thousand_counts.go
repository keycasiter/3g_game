package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 众动万计
// 受到普通攻击时，有45%概率对攻击来源造成兵刃伤害（伤害率140%），并使其下一次造成的伤害伤减少40%
// 被动 100%
type CrowdMovesTenThousandCountsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CrowdMovesTenThousandCountsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c CrowdMovesTenThousandCountsTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)

	//受到普通攻击时，有45%概率对攻击来源造成兵刃伤害（伤害率140%），并使其下一次造成的伤害伤减少40%
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferGeneralAttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}

		//概率攻击
		if util.GenerateRate(0.45) {
			damage.TacticDamage(&damage.TacticDamageParam{
				TacticsParams:     c.tacticsParams,
				AttackGeneral:     params.SufferAttackGeneral,
				SufferGeneral:     params.AttackGeneral,
				DamageType:        consts.DamageType_Weapon,
				DamageImproveRate: 1.4,
				TacticId:          c.Id(),
				TacticName:        c.Name(),
			})

			//施加效果
			if util.DebuffEffectWrapSet(ctx, params.AttackGeneral, consts.DebuffEffectType_LaunchWeaponDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     0.4,
				EffectRound:    1,
				EffectTimes:    1,
				MaxEffectTimes: 1,
				FromTactic:     c.Id(),
			}).IsSuccess {
				//效果消失注册
				util.TacticsTriggerWrapRegister(params.AttackGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_LaunchWeaponDamageDeduce,
						TacticId:   c.Id(),
					})

					return revokeResp
				})
			}
			if util.DebuffEffectWrapSet(ctx, params.AttackGeneral, consts.DebuffEffectType_LaunchStrategyDamageDeduce, &vo.EffectHolderParams{
				EffectRate:     0.4,
				EffectRound:    1,
				EffectTimes:    1,
				MaxEffectTimes: 1,
				FromTactic:     c.Id(),
			}).IsSuccess {
				//效果消失注册
				util.TacticsTriggerWrapRegister(params.AttackGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_LaunchStrategyDamageDeduce,
						TacticId:   c.Id(),
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (c CrowdMovesTenThousandCountsTactic) Id() consts.TacticId {
	return consts.CrowdMovesTenThousandCounts
}

func (c CrowdMovesTenThousandCountsTactic) Name() string {
	return "众动万计"
}

func (c CrowdMovesTenThousandCountsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (c CrowdMovesTenThousandCountsTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CrowdMovesTenThousandCountsTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CrowdMovesTenThousandCountsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (c CrowdMovesTenThousandCountsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CrowdMovesTenThousandCountsTactic) Execute() {

}

func (c CrowdMovesTenThousandCountsTactic) IsTriggerPrepare() bool {
	return false
}

func (a CrowdMovesTenThousandCountsTactic) SetTriggerPrepare(triggerPrepare bool) {
}
