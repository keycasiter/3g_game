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

// 乘敌不虞
// 准备1回合，使敌军主将进入虚弱（无法造成伤害）状态，持续2回合，
// 并使我军主将进入休整状态（每回合恢复一次兵力，回复率105%，受智力影响），持续2回合
type RidingTheEnemyWithoutFearTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	//是否已经触发准备战法
	isTriggerPrepare bool
	isTriggered      bool
}

func (r RidingTheEnemyWithoutFearTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 0.35
	return r
}

func (r RidingTheEnemyWithoutFearTactic) Prepare() {

}

func (r RidingTheEnemyWithoutFearTactic) Id() consts.TacticId {
	return consts.RidingTheEnemyWithoutFear
}

func (r RidingTheEnemyWithoutFearTactic) Name() string {
	return "乘敌不虞"
}

func (r RidingTheEnemyWithoutFearTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RidingTheEnemyWithoutFearTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RidingTheEnemyWithoutFearTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RidingTheEnemyWithoutFearTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (r RidingTheEnemyWithoutFearTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RidingTheEnemyWithoutFearTactic) Execute() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral
	currentRound := r.tacticsParams.CurrentRound

	r.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound

		//准备回合释放
		if currentRound+2 == triggerRound {
			r.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if r.isTriggered {
				return triggerResp
			} else {
				r.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				r.Name(),
			)

			//准备1回合，使敌军主将进入虚弱（无法造成伤害）状态，持续2回合，
			enemyMasterGeneral := util.GetEnemyMasterGeneral(r.tacticsParams)
			if util.DebuffEffectWrapSet(ctx, enemyMasterGeneral, consts.DebuffEffectType_PoorHealth, &vo.EffectHolderParams{
				EffectRound:    2,
				FromTactic:     r.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(enemyMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.DebuffEffectType_PoorHealth,
						TacticId:   r.Id(),
					})

					return revokeResp
				})
			}
			//并使我军主将进入休整状态（每回合恢复一次兵力，回复率105%，受智力影响），持续2回合
			pairMasterGeneral := util.GetPairMasterGeneral(r.tacticsParams)
			if util.BuffEffectWrapSet(ctx, pairMasterGeneral, consts.BuffEffectType_Rest, &vo.EffectHolderParams{
				EffectRound:    2,
				FromTactic:     r.Id(),
				ProduceGeneral: currentGeneral,
			}).IsSuccess {
				util.TacticsTriggerWrapRegister(pairMasterGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					revokeResp := &vo.TacticsTriggerResult{}
					revokeGeneral := params.CurrentGeneral

					if util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
						Ctx:        ctx,
						General:    revokeGeneral,
						EffectType: consts.BuffEffectType_Rest,
						TacticId:   r.Id(),
					}) {
						resumeNum := cast.ToInt64(revokeGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.05)
						util.ResumeSoldierNum(&util.ResumeParams{
							Ctx:            ctx,
							TacticsParams:  r.tacticsParams,
							ProduceGeneral: revokeGeneral,
							SufferGeneral:  revokeGeneral,
							ResumeNum:      resumeNum,
						})
					}

					return revokeResp
				})
			}
		}

		return triggerResp
	})
}

func (r RidingTheEnemyWithoutFearTactic) IsTriggerPrepare() bool {
	return r.isTriggerPrepare
}
