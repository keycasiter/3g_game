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

// 战法名称：熯天炽地
// 战法描述：准备1回合，对敌军全体施放火攻（伤害率102%，受智力影响），并施加灼烧状态，
// 每回合持续造成伤害（伤害率72%，受智力影响），持续2回合。
// 主动战法 发动率35%
type TheSkyIsBlazingTactic struct {
	tacticsParams *model.TacticsParams
	triggerRate   float64
	//是否已经触发准备战法
	isTriggerPrepare bool
}

func (t TheSkyIsBlazingTactic) IsTriggerPrepare() bool {
	return t.isTriggerPrepare
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

	t.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		t.Name(),
	)

	//找到敌军全体
	enemyGenerals := util.GetEnemyGeneralArr(currentGeneral, t.tacticsParams)
	//注册效果
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerRound := params.CurrentRound
		triggerResp := &vo.TacticsTriggerResult{}
		//释放回合
		if currentRound+2 == triggerRound {
			t.isTriggerPrepare = false
		}
		//准备1回合
		if currentRound+1 == triggerRound {
			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				t.Name(),
			)
			for _, general := range enemyGenerals {
				dmgRate := currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 1.02
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams:     t.tacticsParams,
					AttackGeneral:     currentGeneral,
					SufferGeneral:     general,
					DamageType:        consts.DamageType_Strategy,
					DamageImproveRate: dmgRate,
					TacticId:          t.Id(),
					TacticName:        t.Name(),
				})

				//每回合持续造成伤害（伤害率72%，受智力影响），持续2回合
				if util.DebuffEffectWrapSet(ctx, general, consts.DebuffEffectType_Firing, &vo.EffectHolderParams{
					FromTactic:  t.Id(),
					EffectRound: 2,
				}).IsSuccess {
					//注册持续效果
					util.TacticsTriggerWrapRegister(general, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeGeneral := params.CurrentGeneral
						revokeRound := params.CurrentRound

						hlog.CtxInfof(ctx, "[%s]执行来自【%s】的「%v」效果",
							general.BaseInfo.Name,
							t.Name(),
							consts.DebuffEffectType_Firing,
						)
						firingDmgRate := revokeGeneral.BaseInfo.AbilityAttr.IntelligenceBase/100/100 + 0.72
						damage.TacticDamage(&damage.TacticDamageParam{
							TacticsParams:     t.tacticsParams,
							AttackGeneral:     currentGeneral,
							SufferGeneral:     general,
							DamageType:        consts.DamageType_Strategy,
							DamageImproveRate: firingDmgRate,
							TacticId:          t.Id(),
							TacticName:        t.Name(),
						})

						if currentRound+2 == revokeRound {
							util.DebuffEffectWrapRemove(ctx, revokeGeneral, consts.DebuffEffectType_Firing, t.Id())
							hlog.CtxInfof(ctx, "[%s]的【%s】「%v」效果已消失",
								revokeGeneral.BaseInfo.Name,
								t.Name(),
								consts.DebuffEffectType_Firing,
							)
						}
						return triggerResp
					})
				}
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
