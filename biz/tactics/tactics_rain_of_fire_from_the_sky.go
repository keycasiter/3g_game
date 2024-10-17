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
	"github.com/spf13/cast"
)

// 天降火雨
// 准备1回合，对敌军群体（2人）造成一次兵刃攻击（伤害率118%），并附加灼烧状态，每回合持续造成伤害（伤害率66%，受智力影响），持续1回合
type RainOfFireFromTheSkyTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (r RainOfFireFromTheSkyTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 0.5
	return r
}

func (r RainOfFireFromTheSkyTactic) Prepare() {
}

func (r RainOfFireFromTheSkyTactic) Id() consts.TacticId {
	return consts.RainOfFireFromTheSky
}

func (r RainOfFireFromTheSkyTactic) Name() string {
	return "天降火雨"
}

func (r RainOfFireFromTheSkyTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_Inherit
}

func (r RainOfFireFromTheSkyTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RainOfFireFromTheSkyTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RainOfFireFromTheSkyTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (r RainOfFireFromTheSkyTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RainOfFireFromTheSkyTactic) Execute() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral
	currentRound := r.tacticsParams.CurrentRound

	// 准备1回合，对敌军群体（2人）造成一次兵刃攻击（伤害率118%），并附加灼烧状态，每回合持续造成伤害（伤害率66%，受智力影响），持续1回合
	r.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)
	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_ActiveTactic, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
		triggerResp := &vo.TacticsTriggerResult{}
		triggerRound := params.CurrentRound
		triggerGeneral := params.CurrentGeneral

		//准备回合释放
		if currentRound+2 == triggerRound {
			r.isTriggerPrepare = false
		}

		if currentRound+1 == triggerRound {
			if r.isTriggered {
				return triggerResp
			} else {
				r.isTriggered = true
			}

			hlog.CtxInfof(ctx, "[%s]发动战法【%s】",
				currentGeneral.BaseInfo.Name,
				r.Name(),
			)
			//对敌军群体（2人）造成一次兵刃攻击（伤害率118%），并附加灼烧状态，每回合持续造成伤害（伤害率66%，受智力影响），持续1回合
			enemyGenerals := util.GetEnemyTwoGeneralByGeneral(triggerGeneral, r.tacticsParams)
			for _, enemyGeneral := range enemyGenerals {
				//伤害
				dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.ForceBase * 1.18)
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams: r.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					DamageType:    consts.DamageType_Weapon,
					Damage:        dmg,
					TacticId:      r.Id(),
					TacticName:    r.Name(),
				})
				//灼烧状态
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_Firing, &vo.EffectHolderParams{
					EffectRound:    1,
					FromTactic:     r.Id(),
					ProduceGeneral: triggerGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						if util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_Firing,
							TacticId:   r.Id(),
						}) {
							fireDmg := cast.ToInt64(currentGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 0.66)
							damage.TacticDamage(&damage.TacticDamageParam{
								TacticsParams: r.tacticsParams,
								AttackGeneral: currentGeneral,
								SufferGeneral: triggerGeneral,
								DamageType:    consts.DamageType_Strategy,
								Damage:        fireDmg,
								TacticId:      r.Id(),
								TacticName:    r.Name(),
								EffectName:    fmt.Sprintf("%v", consts.DebuffEffectType_Firing),
							})
						}

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (r RainOfFireFromTheSkyTactic) IsTriggerPrepare() bool {
	return false
}
