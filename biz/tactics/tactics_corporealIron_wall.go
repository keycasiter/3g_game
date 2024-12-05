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

// 肉身铁壁
// 战斗中，为友军承担30%，友军主将60%的伤害效果（承担的兵刃伤害额外受自身统率影响降低），
// 当友军兵力高于70%时，使其造成兵刃伤害和谋略伤害提高18%（受统率影响），若主将为孙权，造成伤害提高的基础值增加至30%
// 被动，100%
type CorporealIronWallTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CorporealIronWallTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c CorporealIronWallTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)
	// 战斗中，为友军承担30%，友军主将60%的伤害效果（承担的兵刃伤害额外受自身统率影响降低），
	pairGenerals := util.GetPairGeneralsNotSelf(c.tacticsParams, currentGeneral)
	for _, pairGeneral := range pairGenerals {
		shareRate := 0.3
		if pairGeneral.IsMaster {
			shareRate = 0.6
		}
		//TODO （承担的兵刃伤害额外受自身统率影响降低）

		util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_ShareResponsibilityFor, &vo.EffectHolderParams{
			EffectRate:                      shareRate,
			FromTactic:                      c.Id(),
			ShareResponsibilityForByGeneral: currentGeneral,
			ProduceGeneral:                  currentGeneral,
		})

		// 当友军兵力高于70%时，使其造成兵刃伤害和谋略伤害提高18%（受统率影响），若主将为孙权，造成伤害提高的基础值增加至30%
		effectBase := 0.18
		if consts.General_Id(pairGeneral.BaseInfo.Id) == consts.SunQuan {
			effectBase = 0.3
		}
		effectRate := effectBase + currentGeneral.BaseInfo.AbilityAttr.CommandBase/100/100
		//兵刃伤害提升
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     effectRate,
			EffectTimes:    1,
			FromTactic:     c.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.SufferAttackGeneral
				if cast.ToFloat64(util.DivInt64(revokeGeneral.LossSoldierNum, revokeGeneral.SoldierNum)) < 0.7 {
					util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
						TacticId:   c.Id(),
						CostTimes:  1,
					})
				}

				return revokeResp
			})
		}

		//谋略伤害提升
		if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
			EffectRate:     effectRate,
			EffectTimes:    1,
			FromTactic:     c.Id(),
			ProduceGeneral: currentGeneral,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_SufferDamageEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.SufferAttackGeneral

				if cast.ToFloat64(util.DivInt64(revokeGeneral.LossSoldierNum, revokeGeneral.SoldierNum)) < 0.7 {
					util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_LaunchStrategyDamageImprove,
						TacticId:   c.Id(),
						CostTimes:  1,
					})
				}

				return revokeResp
			})
		}
	}
}

func (c CorporealIronWallTactic) Id() consts.TacticId {
	return consts.CorporealIronWall
}

func (c CorporealIronWallTactic) Name() string {
	return "肉身铁壁"
}

func (c CorporealIronWallTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c CorporealIronWallTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CorporealIronWallTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CorporealIronWallTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (c CorporealIronWallTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CorporealIronWallTactic) Execute() {

}

func (c CorporealIronWallTactic) IsTriggerPrepare() bool {
	return false
}

func (a CorporealIronWallTactic) SetTriggerPrepare(triggerPrepare bool) {
}
