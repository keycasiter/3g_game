package tactics

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
)

// 万箭齐发
// 准备1回合,对敌军全体造成兵刃攻击（伤害率140%），并有50%概率造成溃逃状态，每回合持续造成伤害（伤害率120%，受武力影响），持续1回合
// 主动，40%
type TenThousandArrowsShotAtOnceTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (t TenThousandArrowsShotAtOnceTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.4
	return t
}

func (t TenThousandArrowsShotAtOnceTactic) Prepare() {
}

func (t TenThousandArrowsShotAtOnceTactic) Id() consts.TacticId {
	return consts.TenThousandArrowsShotAtOnce
}

func (t TenThousandArrowsShotAtOnceTactic) Name() string {
	return "万箭齐发"
}

func (t TenThousandArrowsShotAtOnceTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TenThousandArrowsShotAtOnceTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TenThousandArrowsShotAtOnceTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TenThousandArrowsShotAtOnceTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TenThousandArrowsShotAtOnceTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Archers,
	}
}

func (t TenThousandArrowsShotAtOnceTactic) Execute() {
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
			// 准备1回合,对敌军全体造成兵刃攻击（伤害率140%）
			enemyGenerals := util.GetEnemyGeneralsByGeneral(triggerGeneral, t.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     t.tacticsParams,
					AttackGeneral:     triggerGeneral,
					SufferGeneral:     enemyGeneral,
					DamageType:        consts.DamageType_Weapon,
					DamageImproveRate: 1.4,
					TacticId:          t.Id(),
					TacticName:        t.Name(),
				})
				//并有50%概率造成溃逃状态，每回合持续造成伤害（伤害率120%，受武力影响），持续1回合
				if util.GenerateRate(0.5) {
					if util.DebuffEffectWrapSet(ctx, currentGeneral, consts.DebuffEffectType_Escape, &vo.EffectHolderParams{
						EffectRound:    1,
						FromTactic:     t.Id(),
						ProduceGeneral: currentGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_Escape,
								TacticId:   t.Id(),
							}) {
								effectDmgRate := currentGeneral.BaseInfo.AbilityAttr.ForceBase/100/100 + 1.2
								damage.TacticDamage(&damage.TacticDamageParam{
									TacticsParams:     t.tacticsParams,
									AttackGeneral:     currentGeneral,
									SufferGeneral:     revokeGeneral,
									DamageType:        consts.DamageType_Weapon,
									DamageImproveRate: effectDmgRate,
									TacticId:          t.Id(),
									TacticName:        t.Name(),
									EffectName:        fmt.Sprintf("%v", consts.DebuffEffectType_Escape),
									IsIgnoreDefend:    true,
								})
							}

							return revokeResp
						})
					}
				}
			}
		}

		return triggerResp
	})
}

func (t TenThousandArrowsShotAtOnceTactic) IsTriggerPrepare() bool {
	return t.isTriggerPrepare
}
