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

// 战法名称：熯天炽地
// 战法描述：准备1回合，对敌军全体施放火攻（伤害率102%，受智力影响），并施加灼烧状态，
// 每回合持续造成伤害（伤害率72%，受智力影响），持续2回合。
// 主动战法 发动率35%
type TheSkyIsBlazingTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
}

func (t TheSkyIsBlazingTactic) SetTriggerRate(rate float64) {
	t.triggerRate = rate
}

func (t TheSkyIsBlazingTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (t TheSkyIsBlazingTactic) GetTriggerRate() float64 {
	return t.triggerRate
}

func (t TheSkyIsBlazingTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	t.tacticsParams = tacticsParams
	t.triggerRate = 0.35
	return t
}

func (t TheSkyIsBlazingTactic) Prepare() {

}

func (t TheSkyIsBlazingTactic) Name() string {
	return "熯天炽地"
}

func (t TheSkyIsBlazingTactic) Execute() {
	//准备1回合，对敌军全体施放火攻（伤害率102%，受智力影响），并施加灼烧状态，
	// 每回合持续造成伤害（伤害率72%，受智力影响），持续2回合。
	ctx := t.tacticsParams.Ctx
	currentRound := t.tacticsParams.CurrentRound
	currentGeneral := t.tacticsParams.CurrentGeneral

	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//找到敌军全体
	enemyGenerals := util.GetEnemyGeneralArr(t.tacticsParams)
	//注册效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerRound := params.CurrentRound
		triggerResp := &vo.TacticsTriggerResult{}
		//准备1回合
		if currentRound+1 == triggerRound {
			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				t.Name(),
			)
			for _, general := range enemyGenerals {
				dmg := cast.ToInt64(1.02 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
				util.TacticDamage(&util.TacticDamageParam{
					TacticsParams: t.tacticsParams,
					AttackGeneral: currentGeneral,
					SufferGeneral: general,
					Damage:        dmg,
					TacticName:    t.Name(),
				})

				//每回合持续造成伤害（伤害率72%，受智力影响），持续2回合
				if !util.TacticsDebuffEffectCountWrapIncr(ctx, general, consts.DebuffEffectType_Firing, 2, 2, true) {
					return triggerResp
				}
				if !util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_Firing, 1.0) {
					return triggerResp
				}
				//注册持续效果
				util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
					if util.DeBuffEffectContains(general, consts.DebuffEffectType_Firing) &&
						!util.TacticsDebuffEffectCountWrapDecr(ctx, general, consts.DebuffEffectType_Firing, 1) {
						//次数不足移除效果
						util.DebuffEffectWrapRemove(ctx, general, consts.DebuffEffectType_Firing)
						hlog.CtxInfof(ctx, "[%s]的【%s】「%v」效果已消失",
							general.BaseInfo.Name,
							t.Name(),
							consts.DebuffEffectType_Firing,
						)

						return triggerResp
					}
					hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
						general.BaseInfo.Name,
						t.Name(),
						consts.DebuffEffectType_Firing,
					)
					firingDmg := cast.ToInt64(0.72 * currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase)
					util.TacticDamage(&util.TacticDamageParam{
						TacticsParams: t.tacticsParams,
						AttackGeneral: currentGeneral,
						SufferGeneral: general,
						Damage:        firingDmg,
						TacticName:    t.Name(),
					})
					return triggerResp
				})

			}
		}

		return triggerResp
	})
}

func (t TheSkyIsBlazingTactic) Trigger() {
	return
}

func (t TheSkyIsBlazingTactic) Id() consts.TacticId {
	return consts.TheSkyIsBlazing
}

func (t TheSkyIsBlazingTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (t TheSkyIsBlazingTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}
