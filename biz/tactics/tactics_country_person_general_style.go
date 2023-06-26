package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 国士将风
// 战斗前3回合，使自己及友军单体获得先攻和必中状态，造成兵刃伤害和谋略伤害提升20%（受速度影响）
// 指挥，100%
type CountryPersonGeneralStyleTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (c CountryPersonGeneralStyleTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	c.tacticsParams = tacticsParams
	c.triggerRate = 1.0
	return c
}

func (c CountryPersonGeneralStyleTactic) Prepare() {
	ctx := c.tacticsParams.Ctx
	currentGeneral := c.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		c.Name(),
	)
	// 战斗前3回合，使自己及友军单体获得先攻和必中状态，造成兵刃伤害和谋略伤害提升20%（受速度影响）
	effectGenerals := make([]*vo.BattleGeneral, 0)
	pairGeneral := util.GetPairOneGeneralNotSelf(c.tacticsParams, currentGeneral)
	effectGenerals = append(effectGenerals, pairGeneral)
	effectGenerals = append(effectGenerals, currentGeneral)

	for _, general := range effectGenerals {
		//先攻
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_FirstAttack, &vo.EffectHolderParams{
			EffectRound:    3,
			FromTactic:     c.Id(),
			ProduceGeneral: general,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_FirstAttack,
					TacticId:   c.Id(),
				})

				return revokeResp
			})
		}
		//必中
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_MustHit, &vo.EffectHolderParams{
			EffectRound:    3,
			FromTactic:     c.Id(),
			ProduceGeneral: general,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_MustHit,
					TacticId:   c.Id(),
				})

				return revokeResp
			})
		}
		//兵刃伤害
		effectRate := 0.2 + currentGeneral.BaseInfo.AbilityAttr.SpeedBase/100/100
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
			EffectRate:     effectRate,
			EffectRound:    3,
			FromTactic:     c.Id(),
			ProduceGeneral: general,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
					TacticId:   c.Id(),
				})

				return revokeResp
			})
		}
		//谋略伤害
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
			EffectRate:     effectRate,
			EffectRound:    3,
			FromTactic:     c.Id(),
			ProduceGeneral: general,
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral

				util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
					Ctx:        ctx,
					General:    revokeGeneral,
					EffectType: consts.BuffEffectType_LaunchStrategyDamageImprove,
					TacticId:   c.Id(),
				})

				return revokeResp
			})
		}
	}
}

func (c CountryPersonGeneralStyleTactic) Id() consts.TacticId {
	return consts.CountryPersonGeneralStyle
}

func (c CountryPersonGeneralStyleTactic) Name() string {
	return "国士将风"
}

func (c CountryPersonGeneralStyleTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (c CountryPersonGeneralStyleTactic) GetTriggerRate() float64 {
	return c.triggerRate
}

func (c CountryPersonGeneralStyleTactic) SetTriggerRate(rate float64) {
	c.triggerRate = rate
}

func (c CountryPersonGeneralStyleTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (c CountryPersonGeneralStyleTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (c CountryPersonGeneralStyleTactic) Execute() {
}

func (c CountryPersonGeneralStyleTactic) IsTriggerPrepare() bool {
	return false
}
