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

// 五雷轰顶
// 准备1回合，对敌军随机单体造成谋略攻击（伤害率136%，受智力影响），
// 并有30%概率使其进入震慑状态，持续1回合
// 共触发5次，每次独立选择目标
// 自身为主将时，若目标处于水攻状态、沙暴状态时，每多一种提高20%震慑概率
type ThunderStruckTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (t ThunderStruckTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	panic("implement me")
}

func (t ThunderStruckTactic) Prepare() {
	panic("implement me")
}

func (t ThunderStruckTactic) Id() consts.TacticId {
	return consts.ThunderStruck
}

func (t ThunderStruckTactic) Name() string {
	return "五雷轰顶"
}

func (t ThunderStruckTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (t ThunderStruckTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t ThunderStruckTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t ThunderStruckTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t ThunderStruckTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (t ThunderStruckTactic) Execute() {
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

			// 共触发5次，每次独立选择目标
			for i := 0; i < 5; i++ {
				// 准备1回合，对敌军随机单体造成谋略攻击（伤害率136%，受智力影响），
				enemyGeneral := util.GetEnemyOneGeneralByGeneral(triggerGeneral, t.tacticsParams)
				dmg := cast.ToInt64(enemyGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 1.36)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: t.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Strategy,
					Damage:        dmg,
					TacticId:      t.Id(),
					TacticName:    t.Name(),
				})
				// 并有30%概率使其进入震慑状态，持续1回合
				// 自身为主将时，若目标处于水攻状态、沙暴状态时，每多一种提高20%震慑概率
				triggerRate := 0.3
				if triggerGeneral.IsMaster {
					if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_WaterAttack) {
						triggerRate += 0.2
					}
					if util.DeBuffEffectContains(enemyGeneral, consts.DebuffEffectType_Sandstorm) {
						triggerRate += 0.2
					}
				}
				if util.GenerateRate(triggerRate) {
					if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Awe, &vo.EffectHolderParams{
						EffectRound:    1,
						FromTactic:     t.Id(),
						ProduceGeneral: triggerGeneral,
					}).IsSuccess {
						util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
							revokeResp := &vo.TacticsTriggerResult{}
							revokeGeneral := params.CurrentGeneral

							util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
								Ctx:        ctx,
								General:    revokeGeneral,
								EffectType: consts.DebuffEffectType_Awe,
								TacticId:   t.Id(),
							})

							return revokeResp
						})
					}
				}
			}
		}
		return triggerResp
	})
}

func (t ThunderStruckTactic) IsTriggerPrepare() bool {
	return t.isTriggerPrepare
}
