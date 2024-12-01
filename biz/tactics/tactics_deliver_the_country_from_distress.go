package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 扶危定倾
// 指挥 100%
// 战斗首回合，使我军全体受到伤害降低39%（受自身最高属性影响，自身为主将时，基础值提升至42%）；
// 自身每回合40%概率（受统率影响）清除我军主将的负面状态并使其武力、智力、统率提升3%（受智力影响，可叠加，持续2回合）
type DeliverTheCountryFromDistressTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (d DeliverTheCountryFromDistressTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	d.tacticsParams = tacticsParams
	d.triggerRate = 1.0
	return d
}

func (d DeliverTheCountryFromDistressTactic) Prepare() {
	currentGeneral := d.tacticsParams.CurrentGeneral
	ctx := d.tacticsParams.Ctx

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		d.Name(),
	)

	//战斗首回合，使我军全体受到伤害降低39%（受自身最高属性影响，自身为主将时，基础值提升至42%）；
	//找到我军全体
	pairGenerals := util.GetPairGeneralArr(currentGeneral, d.tacticsParams)
	for _, general := range pairGenerals {
		effectRate := 0.39
		if currentGeneral.IsMaster {
			effectRate = 0.42
		} else {
			_, highestVal := util.GetGeneralHighestAttr(currentGeneral)
			effectRate += highestVal / 100 / 100
		}

		//受到谋略攻击减少
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
			EffectTimes: 1,
			EffectRate:  effectRate,
			FromTactic:  d.Id(),
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral
				currentRound := params.CurrentRound
				if currentRound == consts.Battle_Round_Second {
					util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
						TacticId:   d.Id(),
						CostTimes:  1,
					})
				}

				return revokeResp
			})
		}
		//受到兵刃攻击减少
		if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
			EffectTimes: 1,
			EffectRate:  effectRate,
			FromTactic:  d.Id(),
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				revokeResp := &vo.TacticsTriggerResult{}
				revokeGeneral := params.CurrentGeneral
				currentRound := params.CurrentRound
				if currentRound == consts.Battle_Round_Second {
					util.BuffEffectOfTacticCostTime(&util.BuffEffectOfTacticCostTimeParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
						TacticId:   d.Id(),
						CostTimes:  1,
					})
				}

				return revokeResp
			})
		}
	}

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRate := 0.4 + currentGeneral.BaseInfo.AbilityAttr.CommandBase/100/100

		//自身每回合40%概率（受统率影响）清除我军主将的负面状态并使其武力、智力、统率提升3%（受智力影响，可叠加，持续2回合）
		if util.GenerateRate(triggerRate) {
			//找到我军主将
			masterGeneral := util.GetPairMasterGeneral(currentGeneral, d.tacticsParams)

			//清除负面效果
			util.DebuffEffectClean(ctx, masterGeneral)

			//施加效果
			improveRate := 0.03 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
			//武力
			if util.BuffEffectWrapSet(ctx, masterGeneral, consts.BuffEffectType_IncrForce, &vo.EffectHolderParams{
				EffectRate:  improveRate,
				EffectRound: 2,
				FromTactic:  d.Id(),
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(masterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_IncrForce,
						TacticId:   d.Id(),
					})

					return revokeResp
				})
			}
			//智力
			if util.BuffEffectWrapSet(ctx, masterGeneral, consts.BuffEffectType_IncrIntelligence, &vo.EffectHolderParams{
				EffectRate:  improveRate,
				EffectRound: 2,
				FromTactic:  d.Id(),
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(masterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_IncrIntelligence,
						TacticId:   d.Id(),
					})

					return revokeResp
				})
			}
			//统率
			if util.BuffEffectWrapSet(ctx, masterGeneral, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
				EffectRate:  improveRate,
				EffectRound: 2,
				FromTactic:  d.Id(),
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(masterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_IncrCommand,
						TacticId:   d.Id(),
					})

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (d DeliverTheCountryFromDistressTactic) Id() consts.TacticId {
	return consts.DeliverTheCountryFromDistress
}

func (d DeliverTheCountryFromDistressTactic) Name() string {
	return "扶危定倾"
}

func (d DeliverTheCountryFromDistressTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (d DeliverTheCountryFromDistressTactic) GetTriggerRate() float64 {
	return d.triggerRate
}

func (d DeliverTheCountryFromDistressTactic) SetTriggerRate(rate float64) {
	d.triggerRate = rate
}

func (d DeliverTheCountryFromDistressTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (d DeliverTheCountryFromDistressTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (d DeliverTheCountryFromDistressTactic) Execute() {

}

func (d DeliverTheCountryFromDistressTactic) IsTriggerPrepare() bool {
	return false
}

func (a DeliverTheCountryFromDistressTactic) SetTriggerPrepare(triggerPrepare bool) {
}
