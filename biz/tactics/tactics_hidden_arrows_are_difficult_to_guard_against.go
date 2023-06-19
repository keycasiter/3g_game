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

// 暗箭难防
// 准备1回合，对敌军全体（2人）发动一次兵刃攻击（伤害率260%）；并有60%概率（受速度影响）捕获敌军单体武将（捕获效果无法被净化），
// 使其无法行动和造成伤害、禁用指挥和被动战法、进入禁疗状态，且无法被其友方武将选中，持续2回合；
// 同时最多捕获一名武将，若释放时已有敌军武将被捕获，则转而对其造成兵刃伤害（伤害率260%）
type HiddenArrowsAreDifficultToGuardAgainstTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	h.tacticsParams = tacticsParams
	h.triggerRate = 0.35
	return h
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Prepare() {
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Id() consts.TacticId {
	return consts.HiddenArrowsAreDifficultToGuardAgainst
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Name() string {
	return "暗箭难防"
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) GetTriggerRate() float64 {
	return h.triggerRate
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) SetTriggerRate(rate float64) {
	h.triggerRate = rate
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) SupportArmTypes() []consts.ArmType {
	panic("implement me")
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) Execute() {
	ctx := h.tacticsParams.Ctx
	currentGeneral := h.tacticsParams.CurrentGeneral
	currentRound := h.tacticsParams.CurrentRound

	h.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		h.Name(),
	)

	// 准备1回合，对敌军全体（2人）发动一次兵刃攻击（伤害率260%）；并有60%概率（受速度影响）捕获敌军单体武将（捕获效果无法被净化），
	// 使其无法行动和造成伤害、禁用指挥和被动战法、进入禁疗状态，且无法被其友方武将选中，持续2回合；
	// 同时最多捕获一名武将，若释放时已有敌军武将被捕获，则转而对其造成兵刃伤害（伤害率260%）
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			h.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if h.isTriggered {
				return triggerResp
			} else {
				h.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				h.Name(),
			)

			//找到敌军2人
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, h.tacticsParams)
			dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.6)
			for _, enemyGeneral := range enemyGenerals {
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: h.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticName:    h.Name(),
				})
			}
			// 并有60%概率（受速度影响）捕获敌军单体武将（捕获效果无法被净化），
			// 使其无法行动和造成伤害、禁用指挥和被动战法、进入禁疗状态，且无法被其友方武将选中，持续2回合；
			// 同时最多捕获一名武将，若释放时已有敌军武将被捕获，则转而对其造成兵刃伤害（伤害率260%）

			//是否有敌军被捕获
			allEnemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, h.tacticsParams)
			for _, enemyGeneral := range allEnemyGenerals {
				if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_Capture) {
					dmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.ForceBase * 2.6)
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams: h.tacticsParams,
						AttackGeneral: currentGeneral,
						SufferGeneral: enemyGeneral,
						DamageType:    consts.DamageType_Weapon,
						Damage:        dmg,
						TacticName:    h.Name(),
					})
					return triggerResp
				}
			}

			//没有敌军被捕获
			triggerRate := 0.6 + (triggerGeneral.BaseInfo.AbilityAttr.SpeedBase / 100 / 100)
			if util.GenerateRate(triggerRate) {
				fetchEnemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, h.tacticsParams)
				//施加效果
				if util.DebuffEffectWrapSet(ctx, fetchEnemyGeneral, consts.DebuffEffectType_Capture, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     h.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					//注册消失效果
					util.TacticsTriggerWrapRegister(fetchEnemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Capture,
							TacticId:   h.Id(),
						})

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (h HiddenArrowsAreDifficultToGuardAgainstTactic) IsTriggerPrepare() bool {
	return h.isTriggerPrepare
}
