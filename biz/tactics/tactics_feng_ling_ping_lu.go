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

// 奉令平虏
// 战斗中，友军获得功能性增益时有35%概率（受统率影响，自身为主将时基础概率提升到40%）延长1回合；
// 此效果每触发3次，治疗友军单体（治疗率40%，受已损兵力影响）且若友军智力高于武力则获得10%攻心，否则获得10%倒戈，持续1回合
type FengLingPingLuTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	triggerTimes  int64
}

func (a FengLingPingLuTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 1.0
	return a
}

func (a FengLingPingLuTactic) Prepare() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)

	// 战斗中，友军获得功能性增益时有35%概率（受统率影响，自身为主将时基础概率提升到40%）延长1回合；
	pairGenerals := util.GetPairGeneralArr(currentGeneral, a.tacticsParams)
	for _, pairGeneral := range pairGenerals {
		util.TacticsTriggerWrapRegister(pairGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
			triggerResp := &vo.TacticsTriggerResult{}
			triggerGeneral := params.CurrentGeneral

			triggerRate := 0.35
			if currentGeneral.IsMaster {
				triggerRate = 0.4
			}
			triggerRate += currentGeneral.BaseInfo.AbilityAttr.CommandBase / 100 / 100

			if util.GenerateRate(triggerRate) {
				a.triggerTimes++
				//延长1回合
				for _, holderParams := range pairGeneral.BuffEffectHolderMap {
					for _, param := range holderParams {
						param.EffectRound += 1
					}
				}
			}

			// 此效果每触发3次，治疗友军单体（治疗率40%，受已损兵力影响）且若友军智力高于武力则获得10%攻心，否则获得10%倒戈，持续1回合
			if a.triggerTimes%3 == 0 {
				a.triggerTimes -= 3

				general := util.GetPairOneGeneral(a.tacticsParams, triggerGeneral)
				resumeNum := cast.ToInt64(cast.ToFloat64(general.LossSoldierNum) * 0.28)
				util.ResumeSoldierNum(&util.ResumeParams{
					Ctx:            ctx,
					TacticsParams:  a.tacticsParams,
					ProduceGeneral: currentGeneral,
					SufferGeneral:  general,
					ResumeNum:      resumeNum,
					TacticId:       a.Id(),
				})
				//若友军智力高于武力则获得10%攻心，否则获得10%倒戈，持续1回合
				attr, _ := util.GetGeneralHighestBetweenForceOrIntelligence(general)

				//获得10%攻心
				effectType := consts.BuffEffectType_AttackHeart
				if attr == consts.AbilityAttr_Force {
					//获得10%倒戈
					effectType = consts.BuffEffectType_Defection
				}

				if util.BuffEffectWrapSet(ctx, general, effectType, &vo.EffectHolderParams{
					EffectRate:     0.1,
					EffectRound:    1,
					FromTactic:     a.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: effectType,
							TacticId:   a.Id(),
						})

						return revokeResp
					})
				}
			}

			return triggerResp
		})
	}
}

func (a FengLingPingLuTactic) Id() consts.TacticId {
	return consts.FengLingPingLu
}

func (a FengLingPingLuTactic) Name() string {
	return "奉令平虏"
}

func (a FengLingPingLuTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a FengLingPingLuTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a FengLingPingLuTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a FengLingPingLuTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Command
}

func (a FengLingPingLuTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Spearman,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Apparatus,
	}
}

func (a FengLingPingLuTactic) Execute() {
}

func (a FengLingPingLuTactic) IsTriggerPrepare() bool {
	return false
}
