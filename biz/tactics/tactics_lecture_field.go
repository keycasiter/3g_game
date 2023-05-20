package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

type LectureFieldTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (l LectureFieldTactic) IsTriggerPrepare() bool {
	return false
}

//舌战群儒
//敌军尝试发动主动战法时，有25%几率令其发动几率降低5%（受智力影响），
//并提高自己及随机友军主动战法4%（受智力影响）发动几率，持续1回合
func (l LectureFieldTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 1.0
	return l
}

func (l LectureFieldTactic) Prepare() {
	ctx := l.tacticsParams.Ctx
	currentGeneral := l.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		l.Name(),
	)

	//敌军尝试发动主动战法时，有25%几率令其发动几率降低5%（受智力影响）
	//并提高自己及随机友军主动战法4%（受智力影响）发动几率，持续1回合
	enemyGenerals := util.GetEnemyGeneralArr(l.tacticsParams)
	for _, sufferGeneral := range enemyGenerals {
		if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_LectureField, &vo.EffectHolderParams{
			EffectRate: 1.0,
			FromTactic: l.Id(),
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerResp := &vo.TacticsTriggerResult{}
				triggerGeneral := params.CurrentGeneral
				triggerRound := params.CurrentRound

				if util.GenerateRate(0.25) {
					hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
						triggerGeneral.BaseInfo.Name,
						l.Name(),
						consts.DebuffEffectType_LectureField,
					)
					//敌人
					DecrRate := 0.05 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
					if util.DebuffEffectWrapSet(ctx, triggerGeneral, consts.DebuffEffectType_TacticsActiveTriggerDecr, &vo.EffectHolderParams{
						EffectRate: DecrRate,
						FromTactic: l.Id(),
					}).IsSuccess {
						//注册消失
						util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeRound := params.CurrentRound
							revokeGeneral := params.CurrentGeneral

							if triggerRound+1 == revokeRound {
								util.DebuffEffectWrapRemove(ctx, revokeGeneral, consts.DebuffEffectType_TacticsActiveTriggerDecr, l.Id())
							}
							return revokeResp
						})

						//自己
						improveRate := 0.04 + currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100
						if util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
							EffectRate: improveRate,
							FromTactic: l.Id(),
						}).IsSuccess {
							//注册消失
							util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
								revokeResp := &vo.TacticsTriggerResult{}
								revokeRound := params.CurrentRound
								revokeGeneral := params.CurrentGeneral

								if triggerRound+1 == revokeRound {
									util.BuffEffectWrapRemove(ctx, revokeGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, l.Id())
								}
								return revokeResp
							})
						}
						//队友单体
						pairGeneral := util.GetPairOneGeneralNotSelf(l.tacticsParams, currentGeneral)
						if pairGeneral != nil {
							if util.BuffEffectWrapSet(ctx, pairGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, &vo.EffectHolderParams{
								EffectRate: improveRate,
								FromTactic: l.Id(),
							}).IsSuccess {
								//注册消失
								util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
									revokeResp := &vo.TacticsTriggerResult{}
									revokeRound := params.CurrentRound
									revokeGeneral := params.CurrentGeneral

									if triggerRound+1 == revokeRound {
										util.BuffEffectWrapRemove(ctx, revokeGeneral, consts.BuffEffectType_TacticsActiveTriggerImprove, l.Id())
									}
									return revokeResp
								})
							}
						}
					}
				} else {
					hlog.CtxInfof(ctx, "[%s]来自[%s]【%s】的「%v」效果因几率没有生效",
						triggerGeneral.BaseInfo.Name,
						currentGeneral.BaseInfo.Name,
						l.Name(),
						consts.DebuffEffectType_LectureField,
					)
				}
				return triggerResp
			})
		}
	}
}

func (l LectureFieldTactic) Id() consts.TacticId {
	return consts.LectureField
}

func (l LectureFieldTactic) Name() string {
	return "舌战群儒"
}

func (l LectureFieldTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (l LectureFieldTactic) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LectureFieldTactic) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LectureFieldTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (l LectureFieldTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LectureFieldTactic) Execute() {

}
