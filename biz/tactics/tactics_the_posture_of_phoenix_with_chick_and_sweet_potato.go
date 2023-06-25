package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 鸱苕凤姿
// 普通攻击伤害提高60%（受目标损失兵力影响），
// 战斗第5回合时，锁定敌方兵力最低单体直到战斗结束，并且普通攻击时有70%概率使目标进入禁疗状态，持续1回合
// 被动 ，100%
type ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 1.0
	return t
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Prepare() {
	ctx := t.tacticsParams.Ctx
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)
	// 普通攻击伤害提高60%（受目标损失兵力影响），
	util.BuffEffectWrapSet(ctx, currentGeneral, consts.BuffEffectType_LaunchWeaponDamageImprove, &vo.EffectHolderParams{
		EffectRate:     0.6,
		FromTactic:     t.Id(),
		ProduceGeneral: currentGeneral,
	})
	// 战斗第5回合时，锁定敌方兵力最低单体直到战斗结束，并且普通攻击时有70%概率使目标进入禁疗状态，持续1回合
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerGeneral := params.CurrentGeneral
		triggerRound := params.CurrentRound

		if triggerRound == consts.Battle_Round_Fifth {
			enemeyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, t.tacticsParams)
			enemyGeneral := util.GetLowestSoliderNumGeneral(enemeyGenerals)
			//锁定
			util.BuffEffectWrapSet(ctx, triggerGeneral, consts.BuffEffectType_ThePostureOfAPhoenixWithAChickAndASweetPotato_Locked, &vo.EffectHolderParams{
				FromTactic:     t.Id(),
				LockingTarget:  enemyGeneral,
				ProduceGeneral: triggerGeneral,
			})
			//并且普通攻击时有70%概率使目标进入禁疗状态，持续1回合
			util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_AttackEnd, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
				tResp := &vo.TacticsTriggerResult{}
				tGeneral := params.CurrentGeneral
				sufferGeneral := t.tacticsParams.CurrentSufferGeneral

				if util.GenerateRate(0.7) {
					if util.DebuffEffectWrapSet(ctx, sufferGeneral, consts.DebuffEffectType_ProhibitionTreatment, &vo.EffectHolderParams{
						EffectRound:    1,
						FromTactic:     t.Id(),
						ProduceGeneral: tGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(sufferGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_ProhibitionTreatment,
								TacticId:   t.Id(),
							})

							return revokeResp
						})
					}
				}

				return tResp
			})
		}

		return triggerResp
	})
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Id() consts.TacticId {
	return consts.ThePostureOfAPhoenixWithAChickAndASweetPotato
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Name() string {
	return "鸱苕凤姿"
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Passive
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) Execute() {
}

func (t ThePostureOfAPhoenixWithAChickAndASweetPotatoTactic) IsTriggerPrepare() bool {
	return false
}
