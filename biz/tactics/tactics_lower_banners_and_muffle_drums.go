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

// 偃旗息鼓
// 准备1回合。使我军群体（2人）造成谋略伤害增加25%，持续一回合，
// 随后对敌军单体造成谋略攻击（伤害率210%，受智力影响）
type LowerBannersAndMuffleDrumsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	//是否已经触发准备战法
	isTriggerPrepare bool
	isTriggered      bool
}

func (l LowerBannersAndMuffleDrumsTactic) IsTriggerPrepare() bool {
	return l.isTriggerPrepare
}

func (l LowerBannersAndMuffleDrumsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 1.0
	return l
}

func (l LowerBannersAndMuffleDrumsTactic) Prepare() {
}

func (l LowerBannersAndMuffleDrumsTactic) Id() consts.TacticId {
	return consts.LowerBannersAndMuffleDrums
}

func (l LowerBannersAndMuffleDrumsTactic) Name() string {
	return "偃旗息鼓"
}

func (l LowerBannersAndMuffleDrumsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (l LowerBannersAndMuffleDrumsTactic) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LowerBannersAndMuffleDrumsTactic) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LowerBannersAndMuffleDrumsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (l LowerBannersAndMuffleDrumsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LowerBannersAndMuffleDrumsTactic) Execute() {
	ctx := l.tacticsParams.Ctx
	currentGeneral := l.tacticsParams.CurrentGeneral
	currentRound := l.tacticsParams.CurrentRound
	//准备1回合。使我军群体（2人）造成谋略伤害增加25%，持续一回合
	l.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		l.Name(),
	)
	//注册延迟效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			l.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if l.isTriggered {
				return triggerResp
			} else {
				l.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				l.Name(),
			)
			pairGenerals := util.GetPairGeneralsTwoArrByGeneral(triggerGeneral, l.tacticsParams)
			for _, general := range pairGenerals {
				rate := 0.25
				if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_LaunchStrategyDamageImprove, &vo.EffectHolderParams{
					EffectRate: rate,
					FromTactic: l.Id(),
				}).IsSuccess {
					hlog.CtxInfof(ctx, "[%s]造成的谋略伤害提高了%.2f%%",
						general.BaseInfo.Name,
						rate*100,
					)

					//注册消失效果
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral
						revokeRound := params.CurrentRound

						if triggerRound+1 == revokeRound {
							if util.BuffEffectWrapRemove(ctx, revokeGeneral, consts.BuffEffectType_LaunchStrategyDamageImprove, l.Id()) {
								hlog.CtxInfof(ctx, "[%s]造成的谋略伤害降低了%.2f%%",
									revokeGeneral.BaseInfo.Name,
									rate*100,
								)
							}
						}

						return revokeResp
					})
				}
			}
			//随后对敌军单体造成谋略攻击（伤害率210%，受智力影响）
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				dmgTriggerRound := params.CurrentRound
				dmgTriggerGeneral := params.CurrentGeneral
				resp := &vo.TacticsTriggerResult{}

				if triggerRound+1 == dmgTriggerRound {
					enemyGeneral := util.GetEnemyOneGeneralByGeneral(dmgTriggerGeneral, l.tacticsParams)
					dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 2.1
					damage.TacticDamage(&damage.TacticDamageParam{
						TacticsParams:     l.tacticsParams,
						AttackGeneral:     currentGeneral,
						SufferGeneral:     enemyGeneral,
						DamageType:        consts.DamageType_Strategy,
						DamageImproveRate: dmgRate,
						TacticId:          l.Id(),
						TacticName:        l.Name(),
						EffectName:        fmt.Sprintf("%v", consts.BuffEffectType_LowerBannersAndMuffleDrums_Prepare),
					})
				}
				return resp
			})
		}

		return triggerResp
	})
}

func (a LowerBannersAndMuffleDrumsTactic) SetTriggerPrepare(triggerPrepare bool) {
}
