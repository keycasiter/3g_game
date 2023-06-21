package tactics

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 诱敌深入
// 准备1回合，对敌军群体（2人）施加沙暴状态，每回合持续造成伤害（伤害率126%，受智力影响），并使其受到的兵刃伤害提升25%，持续2回合
type LureTheEnemyInDeepTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	//是否已经触发准备战法
	isTriggerPrepare bool
	isTriggered      bool
}

func (l LureTheEnemyInDeepTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	l.tacticsParams = tacticsParams
	l.triggerRate = 0.4
	return l
}

func (l LureTheEnemyInDeepTactic) Prepare() {
}

func (l LureTheEnemyInDeepTactic) Id() consts.TacticId {
	return consts.LureTheEnemyInDeep
}

func (l LureTheEnemyInDeepTactic) Name() string {
	return "诱敌深入"
}

func (l LureTheEnemyInDeepTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (l LureTheEnemyInDeepTactic) GetTriggerRate() float64 {
	return l.triggerRate
}

func (l LureTheEnemyInDeepTactic) SetTriggerRate(rate float64) {
	l.triggerRate = rate
}

func (l LureTheEnemyInDeepTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (l LureTheEnemyInDeepTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (l LureTheEnemyInDeepTactic) Execute() {
	ctx := l.tacticsParams.Ctx
	currentGeneral := l.tacticsParams.CurrentGeneral
	currentRound := l.tacticsParams.CurrentRound

	//准备1回合，对敌军群体（2人）施加沙暴状态，每回合持续造成伤害（伤害率126%，受智力影响），并使其受到的兵刃伤害提升25%，持续2回合
	l.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		l.Name(),
	)
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
			//对敌军群体（2人）施加沙暴状态，每回合持续造成伤害（伤害率126%，受智力影响），并使其受到的兵刃伤害提升25%，持续2回合
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, l.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//沙暴状态
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Sandstorm, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     l.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Sandstorm,
							TacticId:   l.Id(),
						}) {
							dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.26)
							util.TacticDamage(&util.TacticDamageParam{
								TacticsParams: l.tacticsParams,
								AttackGeneral: triggerGeneral,
								SufferGeneral: enemyGeneral,
								DamageType:    consts.DamageType_Strategy,
								Damage:        dmg,
								TacticName:    l.Name(),
								EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Sandstorm),
							})
						}

						return revokeResp
					})
				}
				//并使其受到的兵刃伤害提升25%
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
					EffectRate:     0.25,
					EffectRound:    2,
					FromTactic:     l.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_SufferWeaponDamageImprove,
							TacticId:   l.Id(),
						})

						return revokeResp
					})
				}
			}

		}

		return triggerResp
	})
}

func (l LureTheEnemyInDeepTactic) IsTriggerPrepare() bool {
	return l.isTriggerPrepare
}
