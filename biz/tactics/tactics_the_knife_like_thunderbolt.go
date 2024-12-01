package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 刀出如霆
// 准备1回合。自身及友军单体获得30%倒戈，持续2回合，并对敌军造成兵刃伤害（伤害率300%，由敌军全部武将平分，敌军每有1名副将总伤害率提高120%）及掠阵状态：
// 掠阵状态叠加两次时，移除掠阵状态并使目标受到兵刃伤害提高30%，可叠加；
// 若与张苞同时出战，则友军单体必定选择张苞
// 主动，40%
type TheKnifeLikeThunderboltTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (t TheKnifeLikeThunderboltTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.45
	return t
}

func (t TheKnifeLikeThunderboltTactic) Prepare() {
}

func (t TheKnifeLikeThunderboltTactic) Id() consts.TacticId {
	return consts.TheKnifeLikeThunderbolt
}

func (t TheKnifeLikeThunderboltTactic) Name() string {
	return "刀出如霆"
}

func (t TheKnifeLikeThunderboltTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t TheKnifeLikeThunderboltTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheKnifeLikeThunderboltTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheKnifeLikeThunderboltTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TheKnifeLikeThunderboltTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t TheKnifeLikeThunderboltTactic) Execute() {
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
				triggerGeneral.BaseInfo.Name,
				t.Name(),
			)
			// 准备1回合。自身及友军单体获得30%倒戈，持续2回合
			pairGenerals := make([]*vo.BattleGeneral, 0)
			// 若与张苞同时出战，则友军单体必定选择张苞
			allPairGenerals := util.GetPairGeneralArr(currentGeneral, t.tacticsParams)
			pairGeneral := util.GetPairOneGeneralNotSelf(t.tacticsParams, triggerGeneral)
			for _, general := range allPairGenerals {
				if consts.General_Id(general.BaseInfo.Id) == consts.ZhangBao {
					pairGeneral = general
				}
			}
			pairGenerals = append(pairGenerals, pairGeneral)
			pairGenerals = append(pairGenerals, triggerGeneral)
			for _, general := range pairGenerals {
				if util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_Defection, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     t.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(triggerGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.BuffEffectOfTacticCostRound(&util.BuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.BuffEffectType_Defection,
							TacticId:   t.Id(),
						})

						return revokeResp
					})
				}
			}
			// 并对敌军造成兵刃伤害（伤害率300%，由敌军全部武将平分，敌军每有1名副将总伤害率提高120%）及掠阵状态：
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, t.tacticsParams)
			dmgRate := 3.0
			for _, general := range enemyGenerals {
				if !general.IsMaster {
					dmgRate += 1.2
				}
			}
			for _, general := range enemyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     t.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     general,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: dmgRate,
					TacticId:          t.Id(),
					TacticName:        t.Name(),
				})
				//施加掠阵状态
				util.BuffEffectWrapSet(ctx, general, consts.BuffEffectType_GrazingArray, &vo.EffectHolderParams{
					EffectTimes:    1,
					FromTactic:     t.Id(),
					ProduceGeneral: triggerGeneral,
				})
				// 掠阵状态叠加两次时，移除掠阵状态并使目标受到兵刃伤害提高30%，可叠加；
				if util.BuffEffectGetCount(general, consts.BuffEffectType_GrazingArray) == 2 {
					if util.BuffEffectWrapRemove(ctx, general, consts.BuffEffectType_GrazingArray, 0) {
						util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_SufferWeaponDamageImprove, &vo.EffectHolderParams{
							EffectRate:     0.3,
							EffectTimes:    1,
							MaxEffectTimes: consts.INT64_MAX,
							FromTactic:     t.Id(),
							ProduceGeneral: triggerGeneral,
						})
					}
				}
			}
		}

		return triggerResp
	})

}

func (t TheKnifeLikeThunderboltTactic) IsTriggerPrepare() bool {
	return t.isTriggerPrepare
}

func (a TheKnifeLikeThunderboltTactic) SetTriggerPrepare(triggerPrepare bool) {
}
