package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 出其不意
// 发动概率35%
// 准备1回合，对敌军群体（2人）随机造成计穷（无法发动主动战法）或缴械（无法进行普通攻击），
// 持续1回合（有30%概率持续2回合）
type TakeBySurpriseTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (t TakeBySurpriseTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.35
	return t
}

func (t TakeBySurpriseTactic) Prepare() {
}

func (t TakeBySurpriseTactic) Id() consts.TacticId {
	return consts.TakeBySurprise
}

func (t TakeBySurpriseTactic) Name() string {
	return "出其不意"
}

func (t TakeBySurpriseTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TakeBySurpriseTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TakeBySurpriseTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TakeBySurpriseTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TakeBySurpriseTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TakeBySurpriseTactic) Execute() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral
	currentRound := t.tacticsParams.CurrentRound

	t.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			t.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if t.isTriggered {
				return triggerResp
			} else {
				t.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				t.Name(),
			)
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, t.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				debuffs := []consts.DebuffEffectType{
					consts.DebuffEffectType_CancelWeapon,
					consts.DebuffEffectType_NoStrategy,
				}
				hitIdx := util.GenerateHitOneIdx(2)
				debuff := debuffs[hitIdx]
				//施加效果
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, debuff, &vo.EffectHolderParams{
					EffectRate: 1.0,
					FromTactic: t.Id(),
				}).IsSuccess {
					//回合数设置
					effectRound := int64(1)
					if util.GenerateRate(0.3) {
						effectRound = 2
					}

					//注册消失效果
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral
						revokeRound := params.CurrentRound

						if revokeRound == consts.BattleRound(int64(currentRound)+effectRound) {
							util.DebuffEffectWrapRemove(ctx, revokeGeneral, debuff, t.Id())
						}
						return revokeResp
					})
				}
			}
		}
		return triggerResp
	})
}

func (t TakeBySurpriseTactic) IsTriggerPrepare() bool {
	return t.isTriggerPrepare
}
