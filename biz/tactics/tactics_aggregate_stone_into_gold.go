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

// 聚石成金
// 主动 50%
// 奇数回合发动时，使我军全体统率提升68点（对黄巾武将生效时，受主将魅力影响），并使敌军魅力低于我军主将的武将进入禁疗状态，持续2回合
// 偶数回合发动时，使我军受到的主动战法和突击战法伤害降低24%（对黄巾武将生效时，受统率影响），持续2回合
type AggregateStoneIntoGoldTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (a AggregateStoneIntoGoldTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.5
	return a
}

func (a AggregateStoneIntoGoldTactic) Prepare() {

}

func (a AggregateStoneIntoGoldTactic) Id() consts.TacticId {
	return consts.AggregateStoneIntoGold
}

func (a AggregateStoneIntoGoldTactic) Name() string {
	return "聚石成金"
}

func (a AggregateStoneIntoGoldTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a AggregateStoneIntoGoldTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AggregateStoneIntoGoldTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AggregateStoneIntoGoldTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a AggregateStoneIntoGoldTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AggregateStoneIntoGoldTactic) Execute() {
	currentGeneral := a.tacticsParams.CurrentGeneral
	ctx := a.tacticsParams.Ctx
	currentRound := a.tacticsParams.CurrentRound

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)
	//奇数回合发动时，使我军全体统率提升68点（对黄巾武将生效时，受主将魅力影响），并使敌军魅力低于我军主将的武将进入禁疗状态，持续2回合
	if currentRound%2 != 0 {
		//找到我军全体
		pairGenerals := util.GetPairGeneralArr(currentGeneral, a.tacticsParams)
		for _, general := range pairGenerals {
			effectValue := int64(68)
			if util.IsContainsGeneralTag(general.BaseInfo.GeneralTag, consts.GeneralTag_YellowTurbans) {
				masterGeneral := util.GetPairMasterGeneral(currentGeneral, a.tacticsParams)
				effectValue += cast.ToInt64(masterGeneral.BaseInfo.AbilityAttr.CharmBase / 100)
			}

			if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_IncrCommand, &vo.EffectHolderParams{
				EffectValue: effectValue,
				EffectRound: 2,
				FromTactic:  a.Id(),
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					//效果消耗
					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_IncrCommand,
						TacticId:   a.Id(),
					})

					return revokeResp
				})
			}
		}
		//判断敌军魅力低于我军主将
		masterEnemyGeneral := util.GetEnemyMasterGeneral(currentGeneral, a.tacticsParams)
		enemyGenerals := util.GetEnemyGeneralArr(currentGeneral, a.tacticsParams)
		for _, general := range enemyGenerals {
			if masterEnemyGeneral.BaseInfo.AbilityAttr.CharmBase > general.BaseInfo.AbilityAttr.CharmBase {
				if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
					EffectRound: 2,
					FromTactic:  a.Id(),
				}).IsSuccess {
					//注册消失效果
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						//效果消耗
						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_ProhibitionTreatment,
							TacticId:   a.Id(),
						})

						return revokeResp
					})
				}
			}
		}
	}
	//偶数回合发动时，使我军受到的主动战法和突击战法伤害降低24%（对黄巾武将生效时，受统率影响），持续2回合
	if currentRound%2 == 0 {
		pairGenerals := util.GetPairGeneralArr(currentGeneral, a.tacticsParams)
		for _, general := range pairGenerals {
			effectRate := float64(0.24)
			if util.IsContainsGeneralTag(general.BaseInfo.GeneralTag, consts.GeneralTag_YellowTurbans) {
				effectRate += general.BaseInfo.AbilityAttr.CommandBase / 100 / 100
			}

			//主动战法伤害降低
			if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferActiveTacticDamageDeduce, &vo.EffectHolderParams{
				EffectRound: 2,
				EffectRate:  effectRate,
				FromTactic:  a.Id(),
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					//效果消耗
					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_SufferActiveTacticDamageDeduce,
						TacticId:   a.Id(),
					})

					return revokeResp
				})
			}
			//突击战法伤害降低
			if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferAssaultTacticDamageDeduce, &vo.EffectHolderParams{
				EffectRound: 2,
				EffectRate:  effectRate,
				FromTactic:  a.Id(),
			}).IsSuccess {
				//注册消失效果
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					//效果消耗
					util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_SufferAssaultTacticDamageDeduce,
						TacticId:   a.Id(),
					})

					return revokeResp
				})
			}
		}
	}
}

func (a AggregateStoneIntoGoldTactic) IsTriggerPrepare() bool {
	return false
}
