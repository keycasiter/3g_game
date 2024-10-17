package tactics

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/keycasiter/3g_game/biz/consts"
	"github.com/keycasiter/3g_game/biz/damage"
	"github.com/keycasiter/3g_game/biz/model/vo"
	_interface "github.com/keycasiter/3g_game/biz/tactics/interface"
	"github.com/keycasiter/3g_game/biz/tactics/model"
	"github.com/keycasiter/3g_game/biz/util"
	"github.com/spf13/cast"
)

// 精练策数
// 准备1回合，对敌军群体（2-3人）造成谋略攻击（伤害率210%，受智力影响），并缴械，持续2回合
// 主动，45%
type RefinedStrategiesTactic struct {
	tacticsParams    *model.TacticsParams
	triggerRate      float64
	isTriggerPrepare bool
	isTriggered      bool
}

func (r RefinedStrategiesTactic) Init(tacticsParams *model.TacticsParams) _interface.Tactics {
	r.tacticsParams = tacticsParams
	r.triggerRate = 0.45
	return r
}

func (r RefinedStrategiesTactic) Prepare() {
	ctx := r.tacticsParams.Ctx
	currentGeneral := r.tacticsParams.CurrentGeneral
	currentRound := r.tacticsParams.CurrentRound

	r.isTriggerPrepare = true
	hlog.CtxInfof(ctx, "[%s]准备发动战法【%s】",
		currentGeneral.BaseInfo.Name,
		r.Name(),
	)

	util.TacticsTriggerWrapRegister(currentGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
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
			//准备1回合，对敌军群体（2-3人）造成谋略攻击（伤害率210%，受智力影响），并缴械，持续2回合
			enemyGenerals := util.GetEnemyGeneralsTwoOrThreeMap(r.tacticsParams)
			dmg := cast.ToInt64(triggerGeneral.BaseInfo.AbilityAttr.IntelligenceBase * 2.10)
			for _, enemyGeneral := range enemyGenerals {
				//伤害
				damage.TacticDamage(&damage.TacticDamageParam{
					TacticsParams: r.tacticsParams,
					AttackGeneral: triggerGeneral,
					SufferGeneral: enemyGeneral,
					Damage:        dmg,
					DamageType:    consts.DamageType_Strategy,
					TacticId:      r.Id(),
					TacticName:    r.Name(),
				})
				//缴械
				if util.DebuffEffectWrapSet(ctx, enemyGeneral, consts.DebuffEffectType_CancelWeapon, &vo.EffectHolderParams{
					EffectRound:    2,
					FromTactic:     r.Id(),
					ProduceGeneral: currentGeneral,
				}).IsSuccess {
					util.TacticsTriggerWrapRegister(enemyGeneral, consts.BattleAction_BeginAction, func(params *vo.TacticsTriggerParams) *vo.TacticsTriggerResult {
						revokeResp := &vo.TacticsTriggerResult{}
						revokeGeneral := params.CurrentGeneral

						util.DeBuffEffectOfTacticCostRound(&util.DebuffEffectOfTacticCostRoundParams{
							Ctx:        ctx,
							General:    revokeGeneral,
							EffectType: consts.DebuffEffectType_CancelWeapon,
							TacticId:   r.Id(),
						})

						return revokeResp
					})
				}
			}
		}

		return triggerResp
	})
}

func (r RefinedStrategiesTactic) Id() consts.TacticId {
	return consts.RefinedStrategies
}

func (r RefinedStrategiesTactic) Name() string {
	return "精练策数"
}

func (r RefinedStrategiesTactic) TacticsSource() consts.TacticsSource {
	return consts.TacticsSource_SelfContained
}

func (r RefinedStrategiesTactic) GetTriggerRate() float64 {
	return r.triggerRate
}

func (r RefinedStrategiesTactic) SetTriggerRate(rate float64) {
	r.triggerRate = rate
}

func (r RefinedStrategiesTactic) TacticsType() consts.TacticsType {
	return consts.TacticsType_Active
}

func (r RefinedStrategiesTactic) SupportArmTypes() []consts.ArmType {
	return []consts.ArmType{
		consts.ArmType_Cavalry,
		consts.ArmType_Mauler,
		consts.ArmType_Archers,
		consts.ArmType_Spearman,
		consts.ArmType_Apparatus,
	}
}

func (r RefinedStrategiesTactic) Execute() {

}

func (r RefinedStrategiesTactic) IsTriggerPrepare() bool {
	return false
}
