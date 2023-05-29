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

// 战法名称：镇扼防拒
// 战法描述：每回合有50%概率（受智力影响）使我军单体(优先选除自己之外的副将)援护所有友军并获得休整状态（每回合恢复一次兵力，治疗率192%，受智力影响），
// 持续1回合，同时使其在1回合内受到普通攻击时，有55%概率（受智力影响）移除攻击者的增益状态
type SuppressChokesAndPreventRefusalsTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s SuppressChokesAndPreventRefusalsTactic) IsTriggerPrepare() bool {
	return false
}

func (s SuppressChokesAndPreventRefusalsTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s SuppressChokesAndPreventRefusalsTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (s SuppressChokesAndPreventRefusalsTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s SuppressChokesAndPreventRefusalsTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 1.00
	return s
}

func (s SuppressChokesAndPreventRefusalsTactic) Prepare() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	hlog.CtxInfof(ctx, "[%s]的「%v」效果已施加",
		currentGeneral.BaseInfo.Name,
		consts.BuffEffectType_SuppressChokesAndPreventRefusals_Prepare,
	)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound

		//每回合有50%概率
		//TODO（受智力影响）
		if !util.GenerateRate(0.5) {
			hlog.CtxInfof(ctx, "[%s]来自[%s]【%s】的「%v」效果因几率没有生效",
				currentGeneral.BaseInfo.Name,
				currentGeneral.BaseInfo.Name,
				s.Name(),
				consts.BuffEffectType_SuppressChokesAndPreventRefusals_Prepare,
			)
			return triggerResp
		}

		hlog.CtxInfof(ctx, "[%s]执行来自[%s]【%s】的「%v」效果",
			currentGeneral.BaseInfo.Name,
			currentGeneral.BaseInfo.Name,
			s.Name(),
			consts.BuffEffectType_SuppressChokesAndPreventRefusals_Prepare,
		)
		//使我军单体(优先选除自己之外的副将)援护所有友军并获得休整状态（每回合恢复一次兵力，治疗率192%，受智力影响）
		//找到除当前战法执行外的副将
		viceGeneral := util.GetPairViceGeneralNotSelf(s.tacticsParams)
		//如果不存在除自己的副将则选择为自己
		if viceGeneral == nil {
			viceGeneral = currentGeneral
		}
		//让这个副将援护友军
		generals := util.GetPairGeneralsNotSelf(s.tacticsParams, viceGeneral)
		//援护效果
		for _, general := range generals {
			//注册
			if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_Intervene, &vo.EffectHolderParams{
				EffectRate: 1.0,
				FromTactic: s.Id(),
			}).IsSuccess {
				general.HelpByGeneral = viceGeneral

				util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}

					//消失
					if triggerRound+1 == params.CurrentRound {
						if !util.BuffEffectWrapRemove(ctx, general, consts.BuffEffectType_Intervene, s.Id()) {
							return revokeResp
						}
						general.HelpByGeneral = nil
					}

					return revokeResp
				})
			}
		}

		//休整效果
		if util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_Rest, &vo.EffectHolderParams{
			EffectRate: 1.0,
			FromTactic: s.Id(),
		}).IsSuccess {
			util.TacticsTriggerWrapRegister(viceGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				triggerGeneral := params.CurrentGeneral
				//恢复一次兵力
				if util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_Rest, &vo.EffectHolderParams{
					EffectRate: 1.0,
					FromTactic: s.Id(),
				}).IsSuccess {
					hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
						triggerGeneral.BaseInfo.Name,
						s.Name(),
						consts.BuffEffectType_Rest,
					)
					resumeNum := cast.ToInt64(1.92 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
					finalResumeNum, holdNum, finalNum := util.ResumeSoldierNum(ctx, viceGeneral, resumeNum)

					hlog.CtxInfof(ctx, "[%s]恢复了兵力%d(%d↗%d)",
						triggerGeneral.BaseInfo.Name,
						finalResumeNum,
						holdNum,
						finalNum,
					)
				}
				return triggerResp
			})

			//注册被攻击效果
			if util.BuffEffectWrapSet(ctx, viceGeneral, consts.BuffEffectType_SuppressChokesAndPreventRefusals_Prepare, &vo.EffectHolderParams{
				EffectRate: 1.0,
				FromTactic: s.Id(),
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(viceGeneral, consts.BattleAction_SufferGeneralAttack, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					triggerGeneral := params.CurrentGeneral
					attackGeneral := params.AttackGeneral

					if !util.BuffEffectContains(triggerGeneral, consts.BuffEffectType_SuppressChokesAndPreventRefusals_Prepare) {
						return triggerResp
					}
					//55%概率
					if util.GenerateRate(0.55) {
						hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
							triggerGeneral.BaseInfo.Name,
							s.Name(),
							consts.BuffEffectType_SuppressChokesAndPreventRefusals_Prepare,
						)
						//移除攻击者增益效果
						util.BuffEffectClean(ctx, attackGeneral)
					}
					util.BuffEffectWrapRemove(ctx, triggerGeneral, consts.BuffEffectType_SuppressChokesAndPreventRefusals_Prepare, s.Id())

					return triggerResp
				})
				return triggerResp
			}
		}
		return triggerResp
	})
}

func (s SuppressChokesAndPreventRefusalsTactic) Name() string {
	return "镇扼防拒"
}

func (s SuppressChokesAndPreventRefusalsTactic) Execute() {
	return
}

func (s SuppressChokesAndPreventRefusalsTactic) Trigger() {
	return
}

func (s SuppressChokesAndPreventRefusalsTactic) Id() consts.TacticId {
	return consts.SuppressChokesAndPreventRefusals
}

func (s SuppressChokesAndPreventRefusalsTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (s SuppressChokesAndPreventRefusalsTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}
