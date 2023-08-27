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

// 击其惰归
// 自身在下回合行动前，若受到超过最大兵力20%的伤害，则恢复自身兵力（治疗率296%，受统率影响）并降低25%受到谋略伤害（受统率影响），持续1回合，
// 否则对敌军全体造成兵刃伤害（伤害率154%）
type StrikeItsLazyReturnTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (s StrikeItsLazyReturnTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	s.tacticsParams = tacticsParams
	s.triggerRate = 0.4
	return s
}

func (s StrikeItsLazyReturnTactic) Prepare() {
}

func (s StrikeItsLazyReturnTactic) Id() consts.TacticId {
	return consts.StrikeItsLazyReturn
}

func (s StrikeItsLazyReturnTactic) Name() string {
	return "击其惰归"
}

func (s StrikeItsLazyReturnTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Event
}

func (s StrikeItsLazyReturnTactic) GetTriggerRate() float64 {
	return s.triggerRate
}

func (s StrikeItsLazyReturnTactic) SetTriggerRate(rate float64) {
	s.triggerRate = rate
}

func (s StrikeItsLazyReturnTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (s StrikeItsLazyReturnTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (s StrikeItsLazyReturnTactic) Execute() {
	ctx := s.tacticsParams.Ctx
	currentGeneral := s.tacticsParams.CurrentGeneral
	currentRound := s.tacticsParams.CurrentRound
	lastLostSoliderNum := currentGeneral.LossSoldierNum

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		s.Name(),
	)

	// 自身在下回合行动前，若受到超过最大兵力20%的伤害，则恢复自身兵力（治疗率296%，受统率影响）并降低25%受到谋略伤害（受统率影响），持续1回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound
		triggerLostSoliderNum := triggerGeneral.LossSoldierNum

		if currentRound+1 == triggerRound {
			diffLostSoliderNum := triggerLostSoliderNum - lastLostSoliderNum
			if cast.ToFloat64(diffLostSoliderNum/triggerGeneral.SoldierNum) > 0.2 {
				//恢复
				resumeNum := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.CommandBase * 2.96)
				util.ResumeSoldierNum(&util.ResumeParams{
					Ctx:            ctx,
					TacticsParams:  s.tacticsParams,
					ProduceGeneral: triggerGeneral,
					SufferGeneral:  currentGeneral,
					ResumeNum:      resumeNum,
					TacticId:       s.Id(),
				})
				//效果
				if util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_SufferStrategyDamageDeduce, &vo.EffectHolderParams{
					EffectRate:     0.25,
					EffectRound:    1,
					FromTactic:     s.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_SufferStrategyDamageDeduce,
							TacticId:   s.Id(),
						})

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})

	// 否则对敌军全体造成兵刃伤害（伤害率154%）

}

func (s StrikeItsLazyReturnTactic) IsTriggerPrepare() bool {
	return false
}
