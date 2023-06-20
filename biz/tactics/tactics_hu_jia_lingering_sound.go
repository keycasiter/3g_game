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

// 胡笳余音
// 治疗我军群体（2人，治疗率122%，受智力影响）并有50%概率使自身及友军单体造成的兵刃伤害和谋略伤害提高，受到的兵刃伤害和谋略伤害降低26%（受智力影响），
// 独立判断，持续2回合
// 主动，50%
type HuJiaLingeringSoundTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (h HuJiaLingeringSoundTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.5
	return h
}

func (h HuJiaLingeringSoundTactic) Prepare() {
}

func (h HuJiaLingeringSoundTactic) Id() consts.TacticId {
	return consts.HuJiaLingeringSound
}

func (h HuJiaLingeringSoundTactic) Name() string {
	return "胡笳余音"
}

func (h HuJiaLingeringSoundTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (h HuJiaLingeringSoundTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HuJiaLingeringSoundTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HuJiaLingeringSoundTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (h HuJiaLingeringSoundTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (h HuJiaLingeringSoundTactic) Execute() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)
	// 治疗我军群体（2人，治疗率122%，受智力影响）并有50%概率使自身及友军单体造成的兵刃伤害和谋略伤害提高，受到的兵刃伤害和谋略伤害降低26%（受智力影响），
	// 独立判断，持续2回合
	pairGenerals := util.GetPairGeneralsTwoArr(h.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		//治疗
		resumeNum := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.22)
		util.ResumeSoldierNum(ctx, pairGeneral, resumeNum)
		//效果施加
		if util.GenerateRate(0.5) {
			//找到自己和友军单体
			twoPairGenerals := make([]*vo.BattleGeneral, 0)
			twoPairGenerals = append(twoPairGenerals, currentGeneral)
			onePairGeneral := util.GetPairOneGeneralNotSelf(h.tacticsParams, currentGeneral)
			twoPairGenerals = append(twoPairGenerals, onePairGeneral)

			for _, general := range twoPairGenerals {
				//兵刃提高
				if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.26,
					EffectRound:    2,
					FromTactic:     h.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_LaunchWeaponDamageImprove,
							TacticId:   h.Id(),
						})

						return revokeResp
					})
				}

				//谋略提高
				if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.26,
					EffectRound:    2,
					FromTactic:     h.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_LaunchStrategyDamageImprove,
							TacticId:   h.Id(),
						})

						return revokeResp
					})
				}
				//受到兵刃降低
				if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferWeaponDamageDeduce, &vo.EffectHolderParams{
					EffectRate:     0.26,
					EffectRound:    2,
					FromTactic:     h.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_SufferWeaponDamageDeduce,
							TacticId:   h.Id(),
						})

						return revokeResp
					})
				}
				//受到谋略降低
				if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
					EffectRate:     0.26,
					EffectRound:    2,
					FromTactic:     h.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
							TacticId:   h.Id(),
						})

						return revokeResp
					})
				}
			}
		}
	}
}

func (h HuJiaLingeringSoundTactic) IsTriggerPrepare() bool {
	return false
}
