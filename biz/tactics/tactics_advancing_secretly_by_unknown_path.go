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

//暗渡陈仓
//准备1回合，对敌军单体造成谋略攻击（伤害率260%，受智力影响）并使其进入震慑状态（无法行动），持续2回合（发动几率50%）
type AdvancingSecretlyByUnknownPathTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	//是否已经触发准备战法
	isTriggerPrepare bool
	//是否已经触发
	isTriggered bool
}

func (a AdvancingSecretlyByUnknownPathTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	a.tacticsParams = tacticsParams
	a.triggerRate = 0.5
	return a
}

func (a AdvancingSecretlyByUnknownPathTactic) Prepare() {

}

func (a AdvancingSecretlyByUnknownPathTactic) Id() consts.TacticId {
	return consts.AdvancingSecretlyByUnknownPath
}

func (a AdvancingSecretlyByUnknownPathTactic) Name() string {
	return "暗渡陈仓"
}

func (a AdvancingSecretlyByUnknownPathTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (a AdvancingSecretlyByUnknownPathTactic) GetTriggerRate() float64 {
	return a.triggerRate
}

func (a AdvancingSecretlyByUnknownPathTactic) SetTriggerRate(rate float64) {
	a.triggerRate = rate
}

func (a AdvancingSecretlyByUnknownPathTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (a AdvancingSecretlyByUnknownPathTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (a AdvancingSecretlyByUnknownPathTactic) Execute() {
	ctx := a.tacticsParams.Ctx
	currentGeneral := a.tacticsParams.CurrentGeneral
	currentRound := a.tacticsParams.CurrentRound

	//准备1回合，对敌军单体造成谋略攻击（伤害率260%，受智力影响）并使其进入震慑状态（无法行动），持续2回合（发动几率50%）
	a.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		a.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			a.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if a.isTriggered {
				return triggerResp
			} else {
				a.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				a.Name(),
			)

			//找到敌军单体
			enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, a.tacticsParams)
			//谋略伤害
			dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 2.6)
			_, _, _, isEffect := util.TacticDamage(&util.TacticDamageParam{
				TacticsParams: a.tacticsParams,
				AttackGeneral: triggerGeneral,
				SufferGeneral: enemyGeneral,
				Damage:        dmg,
				TacticName:    a.Name(),
			})
			if isEffect {
				//震慑效果施加
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
					EffectRate:  1.0,
					EffectRound: 2,
					FromTactic:  a.Id(),
				}).IsSuccess {
					//震慑效果消失
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Awe,
							TacticId:   a.Id(),
						})
						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (a AdvancingSecretlyByUnknownPathTactic) IsTriggerPrepare() bool {
	return a.isTriggerPrepare
}
